package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ClassdiagramDrawButtonImpl struct {
}

func NewClassdiagramDrawButtonImpl(nodeCB *NodeCB) (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) {

	classdiagramDrawButtonImpl = new(ClassdiagramDrawButtonImpl)
	return
}

func (classdiagramDrawButtonImpl *ClassdiagramDrawButtonImpl) ButtonUpdated(
	stage *gongdoc_models.StageStruct,
	front *gongdoc_models.Button) {

	log.Println("ClassdiagramDrawButtonImpl, ButtonUpdated", front.Name)
}
