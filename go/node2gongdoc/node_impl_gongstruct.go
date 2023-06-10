package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type NodeImplGongstruct struct {
	gongStruct        *gong_models.GongStruct
	diagramPackage    *gongdoc_models.DiagramPackage
	treeOfGongObjects *gongdoc_models.Tree
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
	treeOfGongObjects *gongdoc_models.Tree,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.diagramPackage = diagramPackage
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.treeOfGongObjects = treeOfGongObjects

	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongdocStage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		// remove the gongstructshape from the selected diagram
		classDiagram := nodeImplGongstruct.diagramPackage.SelectedClassdiagram
		classDiagram.RemoveGongStructShape(gongdocStage, stagedNode.Name)

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		classDiagram := nodeImplGongstruct.diagramPackage.SelectedClassdiagram
		classDiagram.AddGongStructShape(gongdocStage, nodeImplGongstruct.diagramPackage, frontNode.Name)
	}

	computeGongNodesConfigurations(
		gongdocStage,
		nodeImplGongstruct.diagramPackage.SelectedClassdiagram,
		nodeImplGongstruct.treeOfGongObjects)
}
