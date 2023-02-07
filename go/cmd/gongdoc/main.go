package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	"github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/node2gongdoc"

	"github.com/fullstack-lang/gongdoc/go/load"

	gongdoc "github.com/fullstack-lang/gongdoc"

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

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongdoc_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	gongdoc_models.Stage.Checkout()
	gongdoc_models.Stage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")
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

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// setup stacks
	gongdoc_fullstack.Init(r)
	gong_fullstack.Init(r)

	// compute absolute path
	absPkgPath, _ := filepath.Abs(*pkgPath)
	*pkgPath = absPkgPath

	// load package to analyse
	modelPkg, _ := gong_models.LoadSource(*pkgPath)

	r.Use(static.Serve("/", EmbedFolder(gongdoc.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	diagramPackage, _ := load.LoadDiagramPackage(*pkgPath, modelPkg, *editable)

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming()
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

				gongStructName := gongdoc_models.IdentifierToGongStructName(classshape.Identifier)

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

	gongdocStage := &gongdoc_models.Stage
	_ = gongdocStage

	gongStructShapeCallbackSingloton := new(node2gongdoc.GongStructShapeCallbacksSingloton)

	// very strangely,
	// gongdoc_models.Stage.OnAfterClassshapeUpdateCallback = classshapeCallbackSingloton
	// does not seem to be executed
	gongdocStage.OnAfterGongStructShapeUpdateCallback = gongStructShapeCallbackSingloton

	diagramPackageCallbackSingloton := new(load.DiagramPackageCallbacksSingloton)
	gongdocStage.OnAfterDiagramPackageUpdateCallback = diagramPackageCallbackSingloton

	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		models.Stage.OnInitCommitFromFrontCallback = hook
		models.Stage.OnInitCommitFromBackCallback = hook
	}

	log.Printf("Server ready to serve on http://localhost:" + strconv.Itoa(*port) + "/")
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
