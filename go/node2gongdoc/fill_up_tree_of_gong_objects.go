package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpTreeOfGongObjects(
	gongdocStage *gongdoc_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) (treeOfGongObjects *gongdoc_models.Tree) {

	// set up the gongTree to display elements
	treeOfGongObjects = (&gongdoc_models.Tree{Name: "gong"}).Stage(gongdocStage)

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage(gongdocStage)
	gongstructRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](diagramPackage.ModelPkg.GetStage()) {

		nodeGongstruct := (&gongdoc_models.Node{Name: gongStruct.Name}).Stage(gongdocStage)
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)

		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up inversion control
		nodeGongstruct.Impl = NewNodeImplGongstruct(gongStruct, diagramPackage, treeOfGongObjects)

		for _, field := range gongStruct.Fields {
			nodeGongField := (&gongdoc_models.Node{Name: field.GetName()}).Stage(gongdocStage)
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)

			nodeGongField.HasCheckboxButton = true
			nodeGongField.Impl = NewNodeImplField(
				gongStruct,
				field,
				diagramPackage,
				nodeGongstruct,
				nodeGongField,
				treeOfGongObjects)
		}
	}

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage(gongdocStage)
	gongenumRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongenumRootNode)

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage(gongdocStage)
	gongNotesRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongNotesRootNode)

	return
}
