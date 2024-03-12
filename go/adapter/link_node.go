package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type LinkNode struct {
	ElementNodeBase
	stage *gong_models.StageStruct
	Link  *gong_models.GongLink
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (linkNode *LinkNode) RemoveFromDiagram() {
	panic("unimplemented")
}

// AddToDiagram implements diagrammer.ElementNode.
func (linkNode *LinkNode) AddToDiagram() {
	panic("unimplemented")
}

// GetChildren implements diagrammer.ModelElementNode.
func (*LinkNode) GetChildren() []diagrammer.ModelNode {
	return nil
}

func NewLinkNode(
	portfolioAdapter *PortfolioAdapter,
	link *gong_models.GongLink) (linkNode *LinkNode) {
	linkNode = new(LinkNode)

	linkNode.Link = link
	return
}

var _ diagrammer.ModelElementNode = &LinkNode{}

func (linkNode *LinkNode) IsExpanded() bool {
	return true
}

// GenerateProgeny implements diagrammer.Node.
func (linkNode *LinkNode) GenerateProgeny() (children []diagrammer.ModelNode) {

	// for link := range *gong_models.GetLinkInstancesSet[gong_models.Link](linkNode.stage) {

	// 	linkRootNode.Children = append(linkRootNode.Children, nodeLink)
	// }

	return
}

// GetName implements diagrammer.Node.
func (linkNode *LinkNode) GetName() string {
	return linkNode.Link.GetName()
}

// IsNameEditable implements diagrammer.Node.
func (linkNode *LinkNode) IsNameEditable() bool {
	return false
}

func (linkNode *LinkNode) GetElement() any {
	return linkNode.Link
}
