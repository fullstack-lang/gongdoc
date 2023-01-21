package models

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// LoadDiagramPackage fill up the stage with the diagrams elements
func LoadDiagramPackage(pkgPath string, modelPkg *gong_models.ModelPkg, editable bool) (diagramPackage *DiagramPackage, err error) {

	gongdocStage := Stage
	_ = gongdocStage

	diagramPackage = (&DiagramPackage{}).Stage()
	diagramPackage.IsEditable = editable
	diagramPackage.ModelPkg = modelPkg
	diagramPackage.Name = modelPkg.Name + "_diagrams"

	// parse the diagram package
	diagramPkgPath := filepath.Join(pkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage, _ = filepath.Abs(diagramPkgPath)
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	// diagram package, when marshalled, will reference identifiers in the
	// model package. Both of the variable need to be set up for the
	// generic marshalling/unmarshalling to work
	Stage.MetaPackageImportAlias = "ref_" + filepath.Base(diagramPackage.GongModelPath)
	Stage.MetaPackageImportPath = `"` + diagramPackage.GongModelPath + `"`
	if Stage.Map_DocLink_Renaming == nil {
		Stage.Map_DocLink_Renaming = make(map[string]string)
	}

	// if diagrams directory does not exist create it
	_, errd := os.Stat(diagramPkgPath)
	if os.IsNotExist(errd) {
		log.Printf(diagramPkgPath, " does not exist, gongdoc creates it")

		errd := os.MkdirAll(diagramPkgPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPkgPath)
		}
		if os.IsExist(errd) {
			log.Println("directory " + diagramPkgPath + " allready exists")
		}
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset, diagramPkgPath, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ")
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		diagramPackage.IsEditable = editable
		FillUpNodeTree(diagramPackage)
		Stage.Commit()
		return diagramPackage, nil
	}
	diagramPackage.ast = diagramPackageAst
	diagramPackage.fset = fset

	// load all diagram files
	for diagramName, inFile := range diagramPackageAst.Files {
		diagramName := strings.TrimSuffix(filepath.Base(diagramName), ".go")
		diagramPackage.UnmarshallOneDiagram(diagramName, inFile, fset)
	}

	diagramPackage.IsEditable = editable
	FillUpNodeTree(diagramPackage)
	Stage.Commit()
	return diagramPackage, nil
}
