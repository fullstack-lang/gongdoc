package models

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/walk"
	"golang.org/x/tools/go/packages"
)

// Pkgelt stores all diagrams related to a gong package
// swagger:model Pkgelt
type Pkgelt struct {
	Name string

	// Path to the "diagrams" directory
	Path string

	// GongModelPath is the package path, e.g. "fullstack-lang/gongxlsx/go/models"
	GongModelPath string

	// Classdiagrams store UML Classdiagrams
	Classdiagrams []*Classdiagram

	// Umlscs stores UML State charts diagrams
	Umlscs []*Umlsc
}

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

func closeFile(f *os.File) {
	fmt.Println("Closing file ", f.Name())

	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

const preludeRef string = `package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model{{Imports}}
)

`

// Marshall translates all elements of a Pkgelt into a go file
// it recusively call Marshall function into the elements
func (pkgelt *Pkgelt) Marshall(pkgPath string) error {

	// sort Classdiagrams
	sort.Slice(pkgelt.Classdiagrams[:], func(i, j int) bool {
		return pkgelt.Classdiagrams[i].Name < pkgelt.Classdiagrams[j].Name
	})
	for _, classdiagram := range pkgelt.Classdiagrams {
		// open file
		file, err := os.Create(filepath.Join(pkgPath, classdiagram.Name) + ".go")
		defer closeFile(file)

		log.SetFlags(log.Lshortfile)
		filename := walk.CaptureOutput(func() { log.Printf("") })
		log.SetFlags(log.LstdFlags)
		prelude := strings.ReplaceAll(preludeRef, "{{filename}}", filename)
		prelude = strings.ReplaceAll(prelude, "{{ClassdiagramName}}", classdiagram.Name)
		if len(classdiagram.Classshapes) > 0 {
			prelude = strings.ReplaceAll(prelude, "{{Imports}}", "\n\t\""+
				pkgelt.GongModelPath+"\"")
		} else {
			prelude = strings.ReplaceAll(prelude, "{{Imports}}", "")
		}
		prelude = strings.ReplaceAll(prelude, "docs", "models")

		if err == nil {
			fmt.Fprintf(file, prelude)
		} else {
			log.Fatal(err)
		}
		if err2 := classdiagram.MarshallAsVariable(file); err != nil {
			return err2
		}
	}

	for _, umlsc := range pkgelt.Umlscs {
		// open file
		file, err := os.Create(filepath.Join(pkgPath, umlsc.Name) + ".go")
		defer closeFile(file)

		prelude := strings.ReplaceAll(preludeRef, "{{Imports}}", strings.ReplaceAll(pkgelt.Name, "diagrams", "models"))
		prelude = strings.ReplaceAll(prelude, "docs", "models")

		if err == nil {
			fmt.Fprintf(file, prelude)
		} else {
			log.Fatal(err)
		}

		if err2 := umlsc.MarshallAsVariable(file); err != nil {
			return err2
		}
	}

	return nil
}

// PkgeltMap is a Map of all Classdiagrams via their Name
type PkgeltMap map[string]*Pkgelt

// PkgeltStore is a handy ClassdiagramsMap
var PkgeltStore PkgeltMap = make(map[string]*Pkgelt, 0)

// Unmarshall parse the diagram package to get diagrams
// diagramPackagePath is "../diagrams" relative to the "models"
// gongModelPackagePath is the model package path, e.g. "github.com/fullstack-lang/gongxlsx/go/models"
func (pkgelt *Pkgelt) Unmarshall(modelPkg *gong_models.ModelPkg, astPkg *ast.Package, fset2 *token.FileSet, diagramPackagePath string) {

	pkgelt.Path = diagramPackagePath
	pkgelt.GongModelPath = modelPkg.PkgPath

	ast.Inspect(astPkg, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if len(x.Specs) > 0 {
				log.Println("Found declaration ")
				switch vs := x.Specs[0].(type) {
				case *ast.ValueSpec:
					log.Println("Found value spec ", vs.Names[0])

					switch se := vs.Type.(type) {
					case *ast.SelectorExpr:
						switch se.Sel.Name {
						case "Classdiagram":
							var classdiagram Classdiagram
							classdiagram.Name = vs.Names[0].Name
							_ = astPkg
							log.Println("nb files ", len(astPkg.Files))
							astNode := vs.Values[0]
							classdiagram.Unmarshall(modelPkg, astNode, fset2)

							// pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, &classdiagram)
						case "Umlsc":
							var umlsc Umlsc
							umlsc.Name = vs.Names[0].Name
							astNode := vs.Values[0]
							umlsc.Unmarshall(modelPkg, astNode, fset2)

							// pkgelt.Umlscs = append(pkgelt.Umlscs, &umlsc)
						}

					}
				}
			}
		}
		return true
	})

	var directory string
	var err error
	if directory, err = filepath.Abs(diagramPackagePath); err != nil {
		log.Panic("Diagram package path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

	var fset token.FileSet
	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,

		Fset: &fset,
	}

	// if "diagrams" directory does not exists, creates it
	// check existance of path
	_, err = os.Stat(diagramPackagePath)

	// if directory does not exist, creates it
	if os.IsNotExist(err) {
		errd := os.Mkdir(diagramPackagePath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPackagePath)
		}
	}

	var pkgs []*packages.Package
	if pkgs, err = packages.Load(cfg, "./..."); err != nil {
		s := fmt.Sprintf("cannot process package at path %s, err %s", diagramPackagePath, err.Error())
		log.Panic(s)
	}

	if len(pkgs) != 1 {
		// empty package
		return
	}
	pkg := pkgs[0]
	pkgelt.Name = pkg.ID

	// fetch all gd (generic declaration node)
	for _, f := range pkg.Syntax {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			// we are interested in gd with "var" lexical token
			if gd.Tok != token.VAR {
				continue
			}

			// we should find only one ValueSpec (constant or variable declaration)
			if len(gd.Specs) != 1 {
				log.Panicf("One variable declaration should be present instead of %d, %s",
					len(gd.Specs), fset.Position(gd.Pos()))
			}

			// value spec is the top node for
			// "var position Position = position{ X : 10, Y : 20 }"
			vs, ok := gd.Specs[0].(*ast.ValueSpec)
			if !ok {
				log.Panicf("Expected a variable declaration at %s", fset.Position(vs.Pos()))
			}

			// analyse name
			// produce error is vs.Names is no a single element array
			// of type ast.Ident. We expect the name of the diagram
			if len(vs.Names) != 1 {
				log.Panicf("bad variable declaration: %s", fset.Position(vs.Pos()))
			}
			variableName := vs.Names[0]

			// analyse the type of the variable declaration
			// X   Expr   // expression
			// Sel *Ident // field selector
			var se *ast.SelectorExpr
			if se, ok = vs.Type.(*ast.SelectorExpr); !ok {
				log.Panicf("bad type variable declaration, if should something like uml.Classshape or uml.<...>: %s",
					fset.Position(vs.Pos()))
			}

			//  fetch the X field which should an Ident node
			// var X *ast.Ident
			// if X, ok = se.X.(*ast.Ident); !ok {
			// 	log.Panicf("bad type variable declaration, selector should something like uml: %s",
			// 		fset.Position(se.Pos()))
			// }
			// log.Printf("expression is %s, selector is %s", X.Name, se.Sel.Name)

			// produce error if vs.Values is not a single element array
			// of type ast.Ident. We expect the value of the diagram
			if len(vs.Values) != 1 {
				log.Panicf("variable declaration with more than one variable at %s", fset.Position(vs.Pos()))
			}

			switch se.Sel.Name {
			case "Classdiagram":
				var classdiagram Classdiagram
				classdiagram.Name = variableName.Name
				_ = astPkg
				log.Println("nb files ", len(astPkg.Files))
				astNode := vs.Values[0]
				classdiagram.Unmarshall(modelPkg, astNode, &fset)

				pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, &classdiagram)

			case "Umlsc":
				var umlsc Umlsc
				umlsc.Name = variableName.Name
				astNode := vs.Values[0]
				umlsc.Unmarshall(modelPkg, astNode, &fset)

				pkgelt.Umlscs = append(pkgelt.Umlscs, &umlsc)

			case "Document":
			default:
				log.Panicf("Unexpected type of variable: %s at pos %s", se.Sel.Name, fset.Position(se.Pos()))
			}
		}
	}
}

func exprString(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, expr)
	return buf.String()
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (pkgelt *Pkgelt) SerializeToStage() {

	pkgelt.Stage()

	for _, classdiagram := range pkgelt.Classdiagrams {
		classdiagram.SerializeToStage()
	}

	for _, umlsc := range pkgelt.Umlscs {
		umlsc.SerializeToStage()
	}

}
