package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImplClassdiagramEditOff struct {
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

func NewButtonImplClassdiagramEditOff(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	legacyDiagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
	classdiagramNode *gongdoc_models.Node,
	nodeImplClassdiagram *NodeImplClasssiagram,
) (buttonImplClassdiagramEditOff *ButtonImplClassdiagramEditOff) {

	buttonImplClassdiagramEditOff = new(ButtonImplClassdiagramEditOff)

	buttonImplClassdiagramEditOff.diagramPackage = diagramPackage
	buttonImplClassdiagramEditOff.classdiagram = classdiagram
	buttonImplClassdiagramEditOff.diagramPackageNode = diagramPackageNode
	buttonImplClassdiagramEditOff.legacyDiagramPackageNode = legacyDiagramPackageNode
	buttonImplClassdiagramEditOff.treeOfGongObjects = treeOfGongObjects
	buttonImplClassdiagramEditOff.classdiagramNode = classdiagramNode
	buttonImplClassdiagramEditOff.nodeImplClassdiagram = nodeImplClassdiagram

	return
}

func (buttonImplClassdiagramEditOff *ButtonImplClassdiagramEditOff) ButtonUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	stageButton, front *gongdoc_models.Button) {

	log.Println("ButtonImplClassdiagramEditOff, ButtonUpdated", front.Name)

	// set the classdiagram in editoff mode
	buttonImplClassdiagramEditOff.classdiagram.IsInDrawMode = false

	// set the node in editoff mode
	buttonImplClassdiagramEditOff.classdiagramNode.IsInDrawMode = false

	// SetButtonDiaplayState(buttonImplClassdiagramEditOff.classdiagramNode, BUTTON_draw, true)
	// SetButtonDiaplayState(buttonImplClassdiagramEditOff.classdiagramNode, BUTTON_edit_off, false)
	buttonImplClassdiagramEditOff.nodeImplClassdiagram.IsInDrawMode = false

	computeNodeConfs(gongdocStage,
		buttonImplClassdiagramEditOff.diagramPackageNode,
		buttonImplClassdiagramEditOff.diagramPackage,
		buttonImplClassdiagramEditOff.treeOfGongObjects)
}
