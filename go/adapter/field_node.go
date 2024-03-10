package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type FieldNode struct {
	ElementNodeBase
	Field gong_models.FieldInterface
}

// AddToDiagram implements diagrammer.ElementNode.
func (fieldNode *FieldNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &FieldNode{}

func NewFieldNode(
	stage *gong_models.StageStruct,
	field gong_models.FieldInterface) (fieldNode *FieldNode) {
	fieldNode = &FieldNode{ElementNodeBase: ElementNodeBase{stage: stage}}
	fieldNode.stage = stage
	fieldNode.Field = field
	return
}

// GenerateChildren implements diagrammer.Node.
func (fieldNode *FieldNode) GenerateChildren() (children []diagrammer.ModelNode) {
	return
}

// GetName implements diagrammer.Node.
func (fieldNode *FieldNode) GetName() string {
	return fieldNode.Field.GetName()
}

func (fieldNode *FieldNode) GetElement() any {
	return fieldNode.Field
}
