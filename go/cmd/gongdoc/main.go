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

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gongdoc "github.com/fullstack-lang/gongdoc"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	// for import of the embed ng directory of the `github.com/fullstack-lang/gong/ng/projects`
	// directory.
	// with vendoring, this will import the gong code of the github.com/fullstack-lang/gong/ng/projects/gong library
	_ "github.com/fullstack-lang/gong/ng"
)

var (
	logBBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	genDefaultDiagramFlag = flag.Bool("genDefaultDiagram", false, "generate default diagram")
	svg                   = flag.Bool("svg", false, "generate svg output and exits")

	pkgPath = flag.String("pkgPath", ".", "path to the models package")

	setUpRandomNumberOfInstances = flag.Bool("setUpRandomNumberOfInstances", false,
		"set up random number of instance (between 0 and 100)")

	editable = flag.Bool("editable", true, "have the diagram editable")
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

	if *genDefaultDiagramFlag {
		log.Printf("Generating default diagram")

		// generates default UML diagrams
		gongdoc_models.GenGoDefaultDiagram(modelPkg, *pkgPath)
		return
	}

	r.Use(static.Serve("/", EmbedFolder(gongdoc.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	// set up gong structs for diagram package
	if *setUpRandomNumberOfInstances {
		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

			// let create the gong struct in the gongdoc models
			var reference *gongdoc_models.Reference
			var ok bool
			reference, ok = (*gongdoc_models.GetGongstructInstancesMap[gongdoc_models.Reference]())[gongStruct.Name]
			if !ok {
				reference = (&gongdoc_models.Reference{Name: gongStruct.Name}).Stage()
				reference.Type = gongdoc_models.REFERENCE_GONG_STRUCT
			}
			reference.NbInstances = rand.Intn(100)
		}
	}

	diagramPackage, _ := gongdoc_models.Load(*pkgPath, modelPkg, *editable)

	if *svg {
		for _, classDiagram := range diagramPackage.Classdiagrams {
			classDiagram.OutputSVG(filepath.Join(*pkgPath, "../diagrams"))
		}
		os.Exit(0)
	}

	if false {
		for _, classDiagram := range diagramPackage.Classdiagrams {
			for _, classShape := range classDiagram.Classshapes {
				classShape.ShowNbInstances = true
				classShape.NbInstances = rand.Intn(100)
			}
		}
	}
	classshapeCallbackSingloton := new(gongdoc_models.ClassshapeCallbacksSingloton)
	gongdoc_models.Stage.OnAfterClassshapeUpdateCallback = classshapeCallbackSingloton

	diagramPackageCallbackSingloton := new(gongdoc_models.DiagramPackageCallbacksSingloton)
	gongdoc_models.Stage.OnAfterDiagramPackageUpdateCallback = diagramPackageCallbackSingloton

	r.Run(":8080")
}
