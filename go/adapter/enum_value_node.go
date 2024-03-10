package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type EnumValueNode struct {
	ElementNodeBase
	EnumValue *gong_models.GongEnumValue
}

// AddToDiagram implements diagrammer.ElementNode.
func (enumvalueNode *EnumValueNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &EnumValueNode{}

func NewEnumValueNode(
	stage *gong_models.StageStruct,
	enumvalue *gong_models.GongEnumValue) (enumvalueNode *EnumValueNode) {
	enumvalueNode = &EnumValueNode{ElementNodeBase: ElementNodeBase{stage: stage}}
	enumvalueNode.stage = stage
	enumvalueNode.EnumValue = enumvalue
	return
}

// GenerateChildren implements diagrammer.Node.
func (enumvalueNode *EnumValueNode) GenerateChildren() (children []diagrammer.ModelNode) {
	return
}

// GetName implements diagrammer.Node.
func (enumvalueNode *EnumValueNode) GetName() string {
	return enumvalueNode.EnumValue.GetName()
}

// GetElement implements diagrammer.ModelNode.
func (enumvalueNode *EnumValueNode) GetElement() any {
	return enumvalueNode.EnumValue
}
