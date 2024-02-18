package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumNode struct {
	ElementNodeBase
	gongEnum *gong_models.GongEnum
}

func NewGongEnumNode(
	stage *gong_models.StageStruct,
	gongEnum *gong_models.GongEnum) (gongEnumNode *GongEnumNode) {
	gongEnumNode = &GongEnumNode{ElementNodeBase: ElementNodeBase{stage: stage}}

	gongEnumNode.stage = stage
	gongEnumNode.gongEnum = gongEnum
	return
}

// GetChildren implements bridge.Node.
func (gongEnumNode *GongEnumNode) GetChildren() (children []diagrammer.ModelNode) {

	for _, gongEnumValue := range gongEnumNode.gongEnum.GongEnumValues {
		enumValueNode := NewEnumValueNode(gongEnumNode.stage, gongEnumValue)
		children = append(children, enumValueNode)
	}

	return
}

// GetName implements bridge.Node.
func (gongEnumNode *GongEnumNode) GetName() string {
	return gongEnumNode.gongEnum.GetName()
}
