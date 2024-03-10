package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type FieldNode struct {
	ElementNodeBase
	gongStructNode *GongStructNode // parent node
	Field          gong_models.FieldInterface
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (fieldNode *FieldNode) RemoveFromDiagram() {
	panic("unimplemented")
}

// AddToDiagram implements diagrammer.ElementNode.
func (fieldNode *FieldNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &FieldNode{}

func NewFieldNode(
	portfolioAdapter *PortfolioAdapter,
	gongStructNode *GongStructNode,
	field gong_models.FieldInterface) (fieldNode *FieldNode) {
	fieldNode = &FieldNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	fieldNode.gongStructNode = gongStructNode
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

func (fieldNode *FieldNode) CanBeDisplayed() (ok bool) {

	// the parent node must already be displayed in order to be able to display the node
	ok = fieldNode.portfolioAdapter.diagrammer.IsElementNodeDisplayed(fieldNode.gongStructNode)
	return
}
