package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumNode struct {
	ElementNodeBase
	gongEnum *gong_models.GongEnum
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongEnumNode *GongEnumNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &GongEnumNode{}

func NewGongEnumNode(
	stage *gong_models.StageStruct,
	gongEnum *gong_models.GongEnum) (gongEnumNode *GongEnumNode) {
	gongEnumNode = &GongEnumNode{ElementNodeBase: ElementNodeBase{stage: stage}}

	gongEnumNode.stage = stage
	gongEnumNode.gongEnum = gongEnum
	return
}

// GenerateChildren implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for _, gongEnumValue := range gongEnumNode.gongEnum.GongEnumValues {
		enumValueNode := NewEnumValueNode(gongEnumNode.stage, gongEnumValue)
		children = append(children, enumValueNode)
	}

	return
}

// GetName implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GetName() string {
	return gongEnumNode.gongEnum.GetName()
}

func (gongEnumNode *GongEnumNode) GetElement() any {
	return gongEnumNode.gongEnum
}
