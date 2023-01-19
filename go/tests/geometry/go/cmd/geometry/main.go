package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongdoc"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	geometry "github.com/fullstack-lang/gongdoc/go/tests/geometry"
)

var (
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
	r := gin.Default()
	r.Use(cors.Default())

	gongdoc_load.Load(
		"geometry",
		"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models",
		geometry.GoDir,
		r,
		*embeddedDiagrams,
		&map_StructName_InstanceNb)

	r.Use(static.Serve("/", EmbedFolder(gongdoc.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready to serve on http://localhost:" + strconv.Itoa(*port) + "/")
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
