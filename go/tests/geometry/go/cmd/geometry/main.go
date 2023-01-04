package main

import (
	"path/filepath"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	// for diagrams
	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	"github.com/fullstack-lang/gongdoc/go/models"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	geometry "github.com/fullstack-lang/gongdoc/go/tests/geometry"
)

func main() {

	var embeddedDiagrams *bool

	r := gin.Default()
	r.Use(cors.Default())

	// Analyse package
	gong_fullstack.Init(r)
	gongdoc_fullstack.Init(r)
	modelPackage, _ := gong_models.LoadEmbedded(geometry.GoDir)

	// create the diagrams
	// prepare the model views
	var diagramPackage *gongdoc_models.DiagramPackage

	// first, get all gong struct in the model
	for gongStruct := range gong_models.Stage.GongStructs {

		// let create the gong struct in the gongdoc models
		// and put the numbre of instances
		reference := (&gongdoc_models.Reference{Name: gongStruct.Name}).Stage()
		reference.Type = gongdoc_models.REFERENCE_GONG_STRUCT
		nbInstances, ok := models.Stage.Map_GongStructName_InstancesNb[gongStruct.Name]
		if ok {
			reference.NbInstances = nbInstances
		}
	}

	if *embeddedDiagrams {
		diagramPackage, _ = gongdoc_models.LoadEmbedded(geometry.GoDir, modelPackage)
	} else {
		diagramPackage, _ = gongdoc_models.Load(filepath.Join("../../diagrams"), modelPackage, true)
	}

	diagramPackage.GongModelPath = "geometry/go/models"

}
