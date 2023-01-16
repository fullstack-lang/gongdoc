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

func Load(pkgPath string, modelPkg *gong_models.ModelPkg, editable bool) (diagramPackage *DiagramPackage, err error) {
	diagramPackage = (&DiagramPackage{})
	diagramPackage.IsEditable = editable
	diagramPackage.ModelPkg = modelPkg

	// parse the diagram package
	diagramPkgPath := filepath.Join(pkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage, _ = filepath.Abs(diagramPkgPath)

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
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		return diagramPackage, nil
	}
	diagramPackage.ast = diagramPackageAst
	diagramPackage.fset = fset

	for fileName := range diagramPackageAst.Files {
		diagramPackage.Files = append(diagramPackage.Files, strings.TrimSuffix(filepath.Base(fileName), ".go"))
	}

	// diagramPackage.Unmarshall(modelPkg, pkgsParser["diagrams"], fset, diagramPkgPath)
	diagramPackage.IsEditable = editable

	diagramPackage.SerializeToStage()
	FillUpNodeTree(diagramPackage)
	return diagramPackage, nil
}
