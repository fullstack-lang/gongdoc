package load

import (
	"embed"
	"path/filepath"
	"sort"
	"strings"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	// for diagrams
	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/gin-gonic/gin"
)

// Load have gongdoc init itself and the gong stack as well
// then parse the model source code in [goSourceDirectories]
// [stackName], for instance "gongsvg"
// of the gong stack [pathPath], for instance "github.com/fullstack-lang/gongsvg/go/models"
// then parse the diagram.
// the diagram  can be embedded if [embeddedDiagrams] is true (possible if
// no edit is wished and if the binary need to be shipped as a standalone item)
// Number of instances if the working models can be computed to be
// displayed. This is stored in the [map_StructName_InstanceNb] parameter
func Load(
	stackName string,
	pkgPath string,
	goSourceDirectories embed.FS,
	r *gin.Engine,
	embeddedDiagrams bool,
	map_StructName_InstanceNb *map[string]int) {

	gong_fullstack.Init(r)
	gongdoc_fullstack.Init(r)
	modelPackage, _ := gong_models.LoadEmbedded(goSourceDirectories)
	modelPackage.Name = stackName
	modelPackage.PkgPath = pkgPath

	// create the diagrams
	// prepare the model views
	var diagramPackage *gongdoc_models.DiagramPackage

	gongStage := gong_models.Stage
	_ = gongStage

	gongdoc_models.Stage.MetaPackageImportAlias = stackName
	gongdoc_models.Stage.MetaPackageImportPath = pkgPath

	// first, get all gong struct in the model
	for gongStruct := range gong_models.Stage.GongStructs {
		nbInstances, ok := (*map_StructName_InstanceNb)[gongStruct.Name]
		if ok {
			gongdoc_models.Map_Identifier_NbInstances[gongStruct.Name] = nbInstances
		}
	}
	if embeddedDiagrams {
		diagramPackage, _ = gongdoc_models.LoadEmbeddedDiagramPackage(goSourceDirectories, modelPackage)
	} else {
		diagramPackage, _ = gongdoc_models.LoadDiagramPackage(filepath.Join("../../diagrams"), modelPackage, true)
	}
	diagramPackage.GongModelPath = pkgPath

	// set up Map_DocLink_Renaming
	//  TO BE REMOVED
	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.Stage.Map_DocLink_Renaming = make(map[string]string)

	gongstructOrdered := []*gong_models.GongStruct{}
	for gongstruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		gongstructOrdered = append(gongstructOrdered, gongstruct)
	}
	sort.Slice(gongstructOrdered[:], func(i, j int) bool {
		return gongstructOrdered[i].Name < gongstructOrdered[j].Name
	})
	for _, gongStruct := range gongstructOrdered {

		ident := gongdoc_models.RefPrefixReferencedPackage + "models." + gongStruct.Name

		gongdoc_models.Stage.Map_DocLink_Renaming[ident] = ident

		for _, field := range gongStruct.Fields {
			ident := gongdoc_models.RefPrefixReferencedPackage + "models." + gongStruct.Name + "." + field.GetName()
			gongdoc_models.Stage.Map_DocLink_Renaming[ident] = ident
		}
	}
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		ident := gongdoc_models.RefPrefixReferencedPackage + "models." + gongEnum.Name
		_ = ident

		// to do after fix of https://github.com/fullstack-lang/gongdoc/issues/100
		// gongdoc_models.Stage.Map_DocLink_Renaming[ident] = ident
	}
	// end of TO BE REMOVED

	// set up the number of instance per classshape
	if map_StructName_InstanceNb != nil {

		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
			gongdoc_models.Map_Identifier_NbInstances[gongStruct.Name] =
				(*map_StructName_InstanceNb)[gongStruct.Name]

		}

		for _, classdiagram := range diagramPackage.Classdiagrams {
			for _, classshape := range classdiagram.Classshapes {

				gongStructName := strings.TrimPrefix(
					classshape.Identifier,
					gongdoc_models.RefPrefixReferencedPackage+"models.")

				nbInstances, ok := gongdoc_models.Map_Identifier_NbInstances[gongStructName]

				if ok {
					classshape.ShowNbInstances = true
					classshape.NbInstances = nbInstances
				}
			}
		}
	}
}
