package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

type FieldNode struct {
	ElementNodeBase
	Field gong_models.FieldInterface
}

func NewFieldNode(
	stage *gong_models.StageStruct,
	field gong_models.FieldInterface) (fieldNode *FieldNode) {
	fieldNode = &FieldNode{ElementNodeBase: ElementNodeBase{stage: stage}}
	fieldNode.stage = stage
	fieldNode.Field = field
	return
}

// GetChildren implements bridge.Node.
func (fieldNode *FieldNode) GetChildren() (children []bridge.ModelNode) {
	return
}

// GetName implements bridge.Node.
func (fieldNode *FieldNode) GetName() string {
	return fieldNode.Field.GetName()
}
