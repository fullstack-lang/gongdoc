package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImplClassdiagramDraw struct {
	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongdoc_models.Node

	legacyDiagramPackageNode *gongdoc_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongdoc_models.Tree

	classdiagramNode     *gongdoc_models.Node
	nodeImplClassdiagram *NodeImplClasssiagram
}

func NewButtonImplClassdiagramDraw(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	legacyDiagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
	classdiagramNode *gongdoc_models.Node,
	nodeImplClassdiagram *NodeImplClasssiagram,
) (buttonImplClassdiagramDraw *ButtonImplClassdiagramDraw) {

	buttonImplClassdiagramDraw = new(ButtonImplClassdiagramDraw)

	buttonImplClassdiagramDraw.diagramPackage = diagramPackage
	buttonImplClassdiagramDraw.classdiagram = classdiagram
	buttonImplClassdiagramDraw.diagramPackageNode = diagramPackageNode
	buttonImplClassdiagramDraw.legacyDiagramPackageNode = legacyDiagramPackageNode
	buttonImplClassdiagramDraw.treeOfGongObjects = treeOfGongObjects
	buttonImplClassdiagramDraw.classdiagramNode = classdiagramNode
	buttonImplClassdiagramDraw.nodeImplClassdiagram = nodeImplClassdiagram

	return
}

func (buttonImplClassdiagramDraw *ButtonImplClassdiagramDraw) ButtonUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	stageButton, front *gongdoc_models.Button) {

	log.Println("ButtonImplClassdiagramDraw, ButtonUpdated", front.Name)

	// LEGACY
	// set the classdiagram in draw mode
	// set the node in draw mode
	buttonImplClassdiagramDraw.classdiagram.IsInDrawMode = true
	buttonImplClassdiagramDraw.classdiagramNode.IsInDrawMode = true

	// NEW
	buttonImplClassdiagramDraw.nodeImplClassdiagram.IsInDrawMode = true

	computeNodeConfs(gongdocStage,
		buttonImplClassdiagramDraw.diagramPackageNode,
		buttonImplClassdiagramDraw.diagramPackage,
		buttonImplClassdiagramDraw.treeOfGongObjects)
}
