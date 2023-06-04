package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ClassdiagramDrawButtonImpl struct {
	nodeCB *NodeCB
}

func NewClassdiagramDrawButtonImpl(nodeCB *NodeCB) (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) {

	classdiagramDrawButtonImpl = new(ClassdiagramDrawButtonImpl)
	classdiagramDrawButtonImpl.nodeCB = nodeCB
	return
}

func (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) ButtonUpdated(
	stage *gongdoc_models.StageStruct,
	front *gongdoc_models.Button) {

	log.Println("ButtonImpl, ButtonUpdated", front.Name)

	selectedClassdiagram := classdiagramDrawButtonImpl.nodeCB.GetSelectedClassdiagram()
	selectedClassdiagram.IsInDrawMode = true
	classdiagramDrawButtonImpl.nodeCB.computeDiagramNodesConfigurations(stage)
}
