package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_static "github.com/fullstack-lang/gongdoc/go/static"

	"github.com/fullstack-lang/gongdoc/go/load"
	"github.com/fullstack-lang/gongdoc/go/node2gongdoc"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"
)

var (
	logBBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	svg = flag.Bool("svg", false, "generate svg output and exits")

	pkgPath = flag.String("pkgPath", ".", "path to the models package")

	setUpRandomNumberOfInstances = flag.Bool("setUpRandomNumberOfInstances", false,
		"set up random number of instance (between 0 and 100)")

	editable = flag.Bool("editable", true, "have the diagram editable")

	marshallOnCommit = flag.String("marshallOnCommit", "",
		"on all commits, marshall staged data to a go file with the marshall name and '.go' "+
			"(must be lowercased without spaces). If marshall arg is '', no marshalling")

	port = flag.Int("port", 8080, "port server")
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongdoc_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	gongdoc_models.GetDefaultStage().Checkout()
	gongdoc_models.GetDefaultStage().Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")
}

func main() {
	log.SetPrefix("gongdoc: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatal("too many non-flag arguments to the gongdoc command")
	}
	if len(flag.Args()) == 1 {
		*pkgPath = flag.Arg(0)
	}

	r := gongdoc_static.ServeStaticFiles(*logGINFlag)

	// setup stacks
	gongdoc_fullstack.NewStackInstance(r, "")
	gongStage := gong_fullstack.NewStackInstance(r, "")
	_ = gongStage

	// compute absolute path
	absPkgPath, _ := filepath.Abs(*pkgPath)
	*pkgPath = absPkgPath

	// load package to analyse
	modelPkg, _ := gong_models.LoadSource(*pkgPath)

	diagramPackage, _ := load.LoadDiagramPackage(*pkgPath, modelPkg, *editable)

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.Stage_)
	// end of the be removed

	// set up gong structs for diagram package
	if *setUpRandomNumberOfInstances {
		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
			diagramPackage.Map_Identifier_NbInstances[gongStruct.Name] = rand.Intn(100)
		}
		// parse all classdiagrams and all classshape and put the number
		// of instances
		for _, classdiagram := range diagramPackage.Classdiagrams {
			for _, classshape := range classdiagram.GongStructShapes {

				gongStructName := gongdoc_models.IdentifierToGongObjectName(classshape.Identifier)

				nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongStructName]

				if ok {
					classshape.ShowNbInstances = true
					classshape.NbInstances = nbInstances
				}
			}
		}

	}

	if *svg {
		for _, classDiagram := range diagramPackage.Classdiagrams {
			classDiagram.OutputSVG(filepath.Join(*pkgPath, "../diagrams"))
		}
		os.Exit(0)
	}

	gongdocStage := gongdoc_models.GetDefaultStage()

	gongStructShapeCallbackSingloton := new(node2gongdoc.GongStructShapeCallbacksSingloton)

	// very strangely,
	// gongdoc_models.GetDefaultStage().OnAfterClassshapeUpdateCallback = classshapeCallbackSingloton
	// does not seem to be executed
	gongdocStage.OnAfterGongStructShapeUpdateCallback = gongStructShapeCallbackSingloton

	diagramPackageCallbackSingloton := new(load.DiagramPackageCallbacksSingloton)
	gongdocStage.OnAfterDiagramPackageUpdateCallback = diagramPackageCallbackSingloton

	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		gongdoc_models.GetDefaultStage().OnInitCommitFromFrontCallback = hook
		gongdoc_models.GetDefaultStage().OnInitCommitFromBackCallback = hook
	}

	log.Printf("Server ready to serve on http://localhost:" + strconv.Itoa(*port) + "/")
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
