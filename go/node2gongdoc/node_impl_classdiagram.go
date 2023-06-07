package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type NodeImplClasssiagram struct {

	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongdoc_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongdoc_models.Tree

	// peculiar
	IsInDrawMode bool
}

func NewNodeImplClasssiagram(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
) (nodeImplClasssiagram *NodeImplClasssiagram) {

	nodeImplClasssiagram = new(NodeImplClasssiagram)
	nodeImplClasssiagram.diagramPackage = diagramPackage
	nodeImplClasssiagram.classdiagram = classdiagram
	nodeImplClasssiagram.diagramPackageNode = diagramPackageNode
	nodeImplClasssiagram.treeOfGongObjects = treeOfGongObjects

	return
}

func (nodeImplClasssiagram *NodeImplClasssiagram) NodeUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	stagedNode,
	frontNode *gongdoc_models.Node) {

	log.Println("NodeImplClasssiagram: NodeUpdated")

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(gongdocStage)
		nodeImplClasssiagram.diagramPackage.SelectedClassdiagram = nodeImplClasssiagram.classdiagram

		for _, otherDiagramNode := range nodeImplClasssiagram.diagramPackageNode.Children {
			if otherDiagramNode == stagedNode {
				continue
			}

			// uncheck the other node
			if otherDiagramNode.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should be unchecked")
				otherDiagramNode.IsChecked = false
				otherDiagramNode.Commit(gongdocStage)
			}
		}
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(gongdocStage)
	}

	computeNodeConfs(gongdocStage,
		nodeImplClasssiagram.diagramPackageNode,
		nodeImplClasssiagram.diagramPackage,
		nodeImplClasssiagram.treeOfGongObjects)
}
