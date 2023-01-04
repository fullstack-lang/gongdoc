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
	var embeddedDiagrams *bool
	var map_StructName_InstanceNb *map[string]int

	r := gin.Default()
	r.Use(cors.Default())
	var goSourceDirectories embed.FS
	goSourceDirectories = geometry.GoDir

	stackName := "geometry"

	gongdoc_load.Load(
		stackName,
		goSourceDirectories,
		r,
		*embeddedDiagrams,
		map_StructName_InstanceNb)

}
