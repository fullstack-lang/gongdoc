package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewFieldNode(
	stage *gong_models.StageStruct,
	field gong_models.FieldInterface) (fieldNode *FieldNode) {
	fieldNode = new(FieldNode)
	fieldNode.stage = stage
	fieldNode.Field = field
	return
}

type FieldNode struct {
	stage *gong_models.StageStruct
	Field gong_models.FieldInterface
}

func (fieldNode *FieldNode) IsExpanded() bool {
	return true
}

func (fieldNode *FieldNode) HasCheckboxButton() bool {
	return true
}

// GetChildren implements bridge.Node.
func (fieldNode *FieldNode) GetChildren() (children []bridge.Node) {

	// for field := range *gong_models.GetFieldInstancesSet[gong_models.Field](fieldNode.stage) {

	// 	fieldRootNode.Children = append(fieldRootNode.Children, nodeField)
	// }

	return
}

// GetName implements bridge.Node.
func (fieldNode *FieldNode) GetName() string {
	return fieldNode.Field.GetName()
}

// IsNameEditable implements bridge.Node.
func (fieldNode *FieldNode) IsNameEditable() bool {
	return false
}
