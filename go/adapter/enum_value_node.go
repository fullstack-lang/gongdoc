package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewEnumValueNode(
	stage *gong_models.StageStruct,
	enumvalue *gong_models.GongEnumValue) (enumvalueNode *EnumValueNode) {
	enumvalueNode = new(EnumValueNode)
	enumvalueNode.stage = stage
	enumvalueNode.EnumValue = enumvalue
	return
}

type EnumValueNode struct {
	stage     *gong_models.StageStruct
	EnumValue *gong_models.GongEnumValue
}

func (enumvalueNode *EnumValueNode) IsExpanded() bool {
	return true
}

func (enumvalueNode *EnumValueNode) HasCheckboxButton() bool {
	return true
}

// GetChildren implements bridge.Node.
func (enumvalueNode *EnumValueNode) GetChildren() (children []bridge.ModelNode) {
	return
}

// GetName implements bridge.Node.
func (enumvalueNode *EnumValueNode) GetName() string {
	return enumvalueNode.EnumValue.GetName()
}

// IsNameEditable implements bridge.Node.
func (enumvalueNode *EnumValueNode) IsNameEditable() bool {
	return false
}
