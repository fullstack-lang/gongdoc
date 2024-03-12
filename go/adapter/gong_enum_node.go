package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumNode struct {
	ElementNodeBase
	gongEnum *gong_models.GongEnum
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (gongEnumNode *GongEnumNode) RemoveFromDiagram() {
	panic("unimplemented")
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongEnumNode *GongEnumNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &GongEnumNode{}

func NewGongEnumNode(
	portfolioAdapter *PortfolioAdapter,
	gongEnum *gong_models.GongEnum) (gongEnumNode *GongEnumNode) {
	gongEnumNode = &GongEnumNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	gongEnumNode.gongEnum = gongEnum
	return
}

// GenerateProgeny implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GenerateProgeny() []diagrammer.ModelNode {

	for _, gongEnumValue := range gongEnumNode.gongEnum.GongEnumValues {
		enumValueNode := NewEnumValueNode(gongEnumNode.portfolioAdapter, gongEnumValue)
		gongEnumNode.children = append(gongEnumNode.children, enumValueNode)
	}

	return gongEnumNode.children
}

// GetName implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GetName() string {
	return gongEnumNode.gongEnum.GetName()
}

func (gongEnumNode *GongEnumNode) GetElement() any {
	return gongEnumNode.gongEnum
}
