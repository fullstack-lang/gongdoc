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

	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	"github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"

	gongdoc "github.com/fullstack-lang/gongdoc"

	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"

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

	pkgPath = flag.String("pkgPath", "go/models", "path to the models package in order to reveal gong elements in the package")

	setUpRandomNumberOfInstances = flag.Bool("setUpRandomNumberOfInstances", false,
		"set up random number of instance (between 0 and 100)")
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

	// setup GORM
	gongdoc_orm.SetupModels(*logBBFlag, "file:memdb1?mode=memory&cache=shared")
	gong_orm.SetupModels(*logBBFlag, "file:memdb2?mode=memory&cache=shared")

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// load package to analyse
	modelPkg := &gong_models.ModelPkg{}

	// compute absolute path
	absPkgPath, _ := filepath.Abs(*pkgPath)
	*pkgPath = absPkgPath
	diagramPkgPath := filepath.Join(*pkgPath, "../diagrams")

	gong_models.Walk(*pkgPath, modelPkg)
	modelPkg.SerializeToStage()
	gong_models.Stage.Commit()

	if *genDefaultDiagramFlag {
		log.Printf("Generating default diagram")

		// generates default UML diagrams
		gongdoc_models.GenGoDefaultDiagram(modelPkg, *pkgPath)
		return
	}

	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)

	r.Use(static.Serve("/", EmbedFolder(gongdoc.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	var pkgelt gongdoc_models.Pkgelt

	// set up gong structs for pkgelet

	if *setUpRandomNumberOfInstances {
		for gongStruct, _ := range gong_models.Stage.GongStructs {

			// let create the gong struct in the gongdoc models
			gongStruct_ := (&models.GongStruct{Name: gongStruct.Name}).Stage()
			gongStruct_.NbInstances = rand.Intn(100)
		}
	}

	// parse the diagram package
	pkgelt.Unmarshall(modelPkg.PkgPath, diagramPkgPath)

	if *svg {
		for _, classDiagram := range pkgelt.Classdiagrams {
			classDiagram.OutputSVG(diagramPkgPath)
		}
		os.Exit(0)
	}

	if false {
		for _, classDiagram := range pkgelt.Classdiagrams {
			for _, classShape := range classDiagram.Classshapes {
				classShape.ShowNbInstances = true
				classShape.NbInstances = rand.Intn(100)
			}
		}
	}
	pkgelt.SerializeToStage()
	gongdoc_models.Stage.Commit()
	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")

	r.Run(":8080")
}
