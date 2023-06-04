package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_static "github.com/fullstack-lang/gongdoc/go/static"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongdoc/go/doc2svg"
	"github.com/fullstack-lang/gongdoc/go/load"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

var (
	logBBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	setUpRandomNumberOfInstances = flag.Bool("setUpRandomNumberOfInstances", false,
		"set up random number of instance (between 0 and 100)")

	editable = flag.Bool("editable", true, "have the diagram editable")

	marshallOnCommit = flag.String("marshallOnCommit", "",
		"on all commits, marshall staged data to a go file with the marshall name and '.go' "+
			"(must be lowercased without spaces). If marshall arg is '', no marshalling")

	port = flag.Int("port", 8080, "port server")

	// transport secure layer
	tls = flag.Bool("tls", false, "serve on https//localhost:443/")

	selectedDiagramOnLoad = flag.String("selectedDiagramOnLoad", "",
		"load diagram at startup, if empty, no diagram is loaded")
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	// in case one wants a marshalling of the gongdoc stack
	marshallOnCommit string

	// for generating SVG
	docSVGMapper *doc2svg.DocSVGMapper
	gongsvgStage *gongsvg_models.StageStruct
}

func (beforeCommitImplementation *BeforeCommitImplementation) BeforeCommit(gongdocStage *gongdoc_models.StageStruct) {

	if beforeCommitImplementation.marshallOnCommit != "" {
		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		gongdoc_models.GetDefaultStage().Checkout()
		gongdoc_models.GetDefaultStage().Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")
	}
	beforeCommitImplementation.docSVGMapper.GenerateSvg(gongdocStage, beforeCommitImplementation.gongsvgStage)
}

type StackConfigs struct {
	Stacks []string
}

func main() {
	log.SetPrefix("gongdoc: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("usage : gongdoc <flags> <path to models 1> <path to models 2> ...")
	}

	r := gongdoc_static.ServeStaticFiles(*logGINFlag)

	// setup one gongdoc stack per model to analyse
	// list of stack configs
	stacksConfig := new(StackConfigs)
	for _, modelPath := range flag.Args() {

		// compute absolute path
		absPkgPath, _ := filepath.Abs(modelPath)

		_, fullPkgPath := gong_models.ComputePkgPathFromGoModFile(absPkgPath)
		gongdocStage := gongdoc_fullstack.NewStackInstance(r, fullPkgPath)
		gongStage := gong_fullstack.NewStackInstance(r, fullPkgPath)
		gongsvgStage := gongsvg_fullstack.NewStackInstance(r, fullPkgPath)
		stacksConfig.Stacks = append(stacksConfig.Stacks, fullPkgPath)

		// gongdocStage := gongdoc_fullstack.NewStackInstance(r, "")
		// gongStage := gong_fullstack.NewStackInstance(r, "")
		// stacksConfig.Stacks = append(stacksConfig.Stacks, "")

		// load package to analyse
		modelPkg, _ := gong_models.LoadSource(gongStage, absPkgPath)
		diagramPackage, _ := load.LoadDiagramPackage(gongdocStage, absPkgPath, modelPkg, *editable)

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		gongdoc_models.SetupMapDocLinkRenaming(gongStage, diagramPackage.Stage_)
		// end of the be removed

		// set up gong structs for diagram package
		if *setUpRandomNumberOfInstances {
			for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](gongStage) {
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

		diagramPackageCallbackSingloton := new(load.DiagramPackageCallbacksSingloton)
		gongdocStage.OnAfterDiagramPackageUpdateCallback = diagramPackageCallbackSingloton

		gongStage.Commit()
		gongdocStage.Commit()
		gongsvgStage.Commit()

		beforeCommitImplementation := new(BeforeCommitImplementation)
		beforeCommitImplementation.marshallOnCommit = *marshallOnCommit

		docSVGMapper := new(doc2svg.DocSVGMapper)
		beforeCommitImplementation.docSVGMapper = docSVGMapper
		beforeCommitImplementation.gongsvgStage = gongsvgStage

		gongdocStage.OnInitCommitFromFrontCallback = beforeCommitImplementation
		gongdocStage.OnInitCommitFromBackCallback = beforeCommitImplementation

		// enable the inversion control mechanism on the gongsvg backend
		gongsvg_models.SetOrchestratorOnAfterUpdate[gongsvg_models.Rect](gongsvgStage)
		gongsvg_models.SetOrchestratorOnAfterUpdate[gongsvg_models.Link](gongsvgStage)
		gongsvg_models.SetOrchestratorOnAfterUpdate[gongsvg_models.LinkAnchoredText](gongsvgStage)

		// load diagram at startup if requested
		if *selectedDiagramOnLoad != "" {
			for _, classDiagram := range diagramPackage.Classdiagrams {
				if classDiagram.Name == *selectedDiagramOnLoad {
					diagramPackage.SelectedClassdiagram = classDiagram
					gongdocStage.Commit()
					// docSVGMapper.GenerateSvg(gongdocStage, gongsvgStage)
				}
			}
		}

	}

	log.Printf("Server ready to serve on http://localhost:" + strconv.Itoa(*port) + "/")

	// the client will fetch the list of stacks
	r.GET("/api/stacks", func(c *gin.Context) {
		c.JSON(http.StatusOK, stacksConfig.Stacks)
	})

	if *tls {
		err := r.RunTLS(":443", "cert.pem", "key-no-pass.pem")
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		err := r.Run(":" + strconv.Itoa(*port))
		if err != nil {
			log.Fatalln(err.Error())
		}

	}
}
