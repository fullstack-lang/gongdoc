package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type EnumValueNode struct {
	ElementNodeBase
	EnumValue *gong_models.GongEnumValue
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (enumvalueNode *EnumValueNode) RemoveFromDiagram() {
	panic("unimplemented")
}

// AddToDiagram implements diagrammer.ElementNode.
func (enumvalueNode *EnumValueNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &EnumValueNode{}

func NewEnumValueNode(
	portfolioAdapter *PortfolioAdapter,
	enumvalue *gong_models.GongEnumValue) (enumvalueNode *EnumValueNode) {
	enumvalueNode = &EnumValueNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}
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
