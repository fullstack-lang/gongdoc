package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type EnumValueNode struct {
	ElementNodeBase
	EnumValue *gong_models.GongEnumValue
}

func NewEnumValueNode(
	stage *gong_models.StageStruct,
	enumvalue *gong_models.GongEnumValue) (enumvalueNode *EnumValueNode) {
	enumvalueNode = &EnumValueNode{ElementNodeBase: ElementNodeBase{stage: stage}}
	enumvalueNode.stage = stage
	enumvalueNode.EnumValue = enumvalue
	return
}

// GetChildren implements bridge.Node.
func (enumvalueNode *EnumValueNode) GetChildren() (children []diagrammer.ModelNode) {
	return
}

// GetName implements bridge.Node.
func (enumvalueNode *EnumValueNode) GetName() string {
	return enumvalueNode.EnumValue.GetName()
}
