package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

func NewGongStructNode(
	stage *gong_models.StageStruct,
	gongStruct *gong_models.GongStruct) (gongStructNode *GongStructNode) {
	gongStructNode = &GongStructNode{ElementNodeBase: ElementNodeBase{stage: stage}}

	gongStructNode.stage = stage
	gongStructNode.gongStruct = gongStruct
	return
}

var _ diagrammer.ModelNode = &GongStructNode{}

type GongStructNode struct {
	ElementNodeBase
	gongStruct *gong_models.GongStruct
}

// GenerateChildren implements diagrammer.Node.
func (gongStructNode *GongStructNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for _, field := range gongStructNode.gongStruct.Fields {
		fieldNode := NewFieldNode(gongStructNode.stage, field)
		children = append(children, fieldNode)
	}

	return
}

// GetName implements diagrammer.Node.
func (gongStructNode *GongStructNode) GetName() string {
	return gongStructNode.gongStruct.GetName()
}

func (gongStructNode *GongStructNode) GetElement() any {
	return gongStructNode.gongStruct
}
