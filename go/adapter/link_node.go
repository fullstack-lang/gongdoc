package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

func NewLinkNode(
	stage *gong_models.StageStruct,
	link *gong_models.GongLink) (linkNode *LinkNode) {
	linkNode = new(LinkNode)
	linkNode.stage = stage
	linkNode.Link = link
	return
}

type LinkNode struct {
	ElementNodeBase
	stage *gong_models.StageStruct
	Link  *gong_models.GongLink
}

func (linkNode *LinkNode) IsExpanded() bool {
	return true
}

// GetChildren implements bridge.Node.
func (linkNode *LinkNode) GetChildren() (children []diagrammer.ModelNode) {

	// for link := range *gong_models.GetLinkInstancesSet[gong_models.Link](linkNode.stage) {

	// 	linkRootNode.Children = append(linkRootNode.Children, nodeLink)
	// }

	return
}

// GetName implements bridge.Node.
func (linkNode *LinkNode) GetName() string {
	return linkNode.Link.GetName()
}

// IsNameEditable implements bridge.Node.
func (linkNode *LinkNode) IsNameEditable() bool {
	return false
}

func (linkNode *LinkNode) GetElement() any {
	return linkNode.Link
}
