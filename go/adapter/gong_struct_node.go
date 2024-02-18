package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongStructNode(
	stage *gong_models.StageStruct,
	gongStruct *gong_models.GongStruct) (gongStructNode *GongStructNode) {
	gongStructNode = &GongStructNode{ElementNodeBase: ElementNodeBase{stage: stage}}

	gongStructNode.stage = stage
	gongStructNode.gongStruct = gongStruct
	return
}

type GongStructNode struct {
	ElementNodeBase
	gongStruct *gong_models.GongStruct
}

// GetChildren implements bridge.Node.
func (gongStructNode *GongStructNode) GetChildren() (children []bridge.ModelNode) {

	for _, field := range gongStructNode.gongStruct.Fields {
		fieldNode := NewFieldNode(gongStructNode.stage, field)
		children = append(children, fieldNode)
	}

	return
}

// GetName implements bridge.Node.
func (gongStructNode *GongStructNode) GetName() string {
	return gongStructNode.gongStruct.GetName()
}
