package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongEnumNode(
	stage *gong_models.StageStruct,
	gongEnum *gong_models.GongEnum) (gongEnumNode *GongEnumNode) {
	gongEnumNode = new(GongEnumNode)

	gongEnumNode.stage = stage
	gongEnumNode.gongEnum = gongEnum
	return
}

type GongEnumNode struct {
	stage    *gong_models.StageStruct
	gongEnum *gong_models.GongEnum
}

func (gongEnumNode *GongEnumNode) IsExpanded() bool {
	return false
}

func (gongEnumNode *GongEnumNode) HasCheckboxButton() bool {
	return true
}

// GetChildren implements bridge.Node.
func (gongEnumNode *GongEnumNode) GetChildren() (children []bridge.ModelNode) {

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

// IsNameEditable implements bridge.Node.
func (gongEnumNode *GongEnumNode) IsNameEditable() bool {
	return false
}
