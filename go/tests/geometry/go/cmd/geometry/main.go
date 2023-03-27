package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
	gongdoc_static "github.com/fullstack-lang/gongdoc/go/static"

	geometry "github.com/fullstack-lang/gongdoc/go/tests/geometry"
)

var (
	logGINFlag       = flag.Bool("logGIN", false, "log mode for gin")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	port             = flag.Int("port", 8080, "port server")
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

	map_StructName_InstanceNb := make(map[string]int)

	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}

	r := gongdoc_static.ServeStaticFiles(*logGINFlag)

	gongdoc_load.Load(
		"geometry",
		"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models",
		geometry.GoDir,
		r,
		*embeddedDiagrams,
		&map_StructName_InstanceNb)

	myArray := []string{"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"}
	r.GET("/api/stacks", func(c *gin.Context) {

		log.Println("answering the stacks request")

		c.JSON(http.StatusOK, myArray)
	})

	log.Printf("Server ready to serve on http://localhost:" + strconv.Itoa(*port) + "/")
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
