package main

import (
	"embed"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	geometry "github.com/fullstack-lang/gongdoc/go/tests/geometry"
)

func main() {

	// paramters
	var embeddedDiagrams bool

	map_StructName_InstanceNb := make(map[string]int)
	r := gin.Default()
	r.Use(cors.Default())
	var goSourceDirectories embed.FS
	goSourceDirectories = geometry.GoDir

	stackName := "geometry"
	embeddedDiagrams = true

	gongdoc_load.Load(
		stackName,
		goSourceDirectories,
		r,
		embeddedDiagrams,
		&map_StructName_InstanceNb)

}
