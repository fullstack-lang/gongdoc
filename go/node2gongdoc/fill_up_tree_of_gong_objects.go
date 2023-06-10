package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpTreeOfGongObjects(
	gongdocStage *gongdoc_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) (gongTree *gongdoc_models.Tree) {

	// set up the gongTree to display elements
	gongTree = (&gongdoc_models.Tree{Name: "gong"}).Stage(gongdocStage)

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage(gongdocStage)
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](diagramPackage.ModelPkg.GetStage()) {

		nodeGongstruct := (&gongdoc_models.Node{Name: gongStruct.Name}).Stage(gongdocStage)
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)

		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up inversion control
		nodeGongstruct.Impl = NewNodeImplGongstruct(gongStruct, diagramPackage)
	}

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage(gongdocStage)
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage(gongdocStage)
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)

	return
}
