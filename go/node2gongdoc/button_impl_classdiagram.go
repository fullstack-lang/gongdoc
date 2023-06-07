package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImplClassdiagram struct {
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

	// type of button
	Icon ButtonType
}

func NewButtonImplClassdiagram(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	legacyDiagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
	classdiagramNode *gongdoc_models.Node,
	nodeImplClassdiagram *NodeImplClasssiagram,
	icon ButtonType,
) (buttonImplClassdiagram *ButtonImplClassdiagram) {

	buttonImplClassdiagram = new(ButtonImplClassdiagram)

	buttonImplClassdiagram.diagramPackage = diagramPackage
	buttonImplClassdiagram.classdiagram = classdiagram
	buttonImplClassdiagram.diagramPackageNode = diagramPackageNode
	buttonImplClassdiagram.legacyDiagramPackageNode = legacyDiagramPackageNode
	buttonImplClassdiagram.treeOfGongObjects = treeOfGongObjects
	buttonImplClassdiagram.classdiagramNode = classdiagramNode
	buttonImplClassdiagram.nodeImplClassdiagram = nodeImplClassdiagram
	buttonImplClassdiagram.Icon = icon

	return
}

func (buttonImplClassdiagram *ButtonImplClassdiagram) ButtonUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	stageButton, front *gongdoc_models.Button) {

	log.Println("ButtonImplClassdiagramDraw, ButtonUpdated", front.Name)

	switch buttonImplClassdiagram.Icon {
	case BUTTON_draw:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = true
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = true
	case BUTTON_edit_off:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = false
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = false
	}

	computeNodeConfs(gongdocStage,
		buttonImplClassdiagram.diagramPackageNode,
		buttonImplClassdiagram.diagramPackage,
		buttonImplClassdiagram.treeOfGongObjects)
}
