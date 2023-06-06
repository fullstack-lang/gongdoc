package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ClassdiagramDrawButtonImpl struct {
	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongdoc_models.Node

	legacyDiagramPackageNode *gongdoc_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongdoc_models.Tree
}

func NewClassdiagramDrawButtonImpl(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	legacyDiagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
) (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) {

	classdiagramDrawButtonImpl = new(ClassdiagramDrawButtonImpl)

	classdiagramDrawButtonImpl.diagramPackage = diagramPackage
	classdiagramDrawButtonImpl.classdiagram = classdiagram
	classdiagramDrawButtonImpl.diagramPackageNode = diagramPackageNode
	classdiagramDrawButtonImpl.legacyDiagramPackageNode = legacyDiagramPackageNode
	classdiagramDrawButtonImpl.treeOfGongObjects = treeOfGongObjects

	return
}

func (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) ButtonUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	front *gongdoc_models.Button) {

	log.Println("ClassdiagramDrawButtonImpl, ButtonUpdated", front.Name)

	// set the classdiagram in draw mode
	classdiagramDrawButtonImpl.classdiagram.IsInDrawMode = true

	computeNodeConfs(gongdocStage,
		classdiagramDrawButtonImpl.diagramPackageNode,
		classdiagramDrawButtonImpl.diagramPackage,
		classdiagramDrawButtonImpl.treeOfGongObjects)
}
