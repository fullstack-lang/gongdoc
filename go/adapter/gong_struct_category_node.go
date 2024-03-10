package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongStructCategoryNode struct {
	ModelCategoryNodeBase
}

func NewGongStructCategoryNode(portfolioAdapter *PortfolioAdapter, name string) *GongStructCategoryNode {
	return &GongStructCategoryNode{ModelCategoryNodeBase: ModelCategoryNodeBase{portfolioAdapter: portfolioAdapter, Name: name}}
}

// GenerateChildren implements diagrammer.Node.
func (categoryNode *GongStructCategoryNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](categoryNode.portfolioAdapter.gongStage) {

		gongStructNode := NewGongStructNode(categoryNode.portfolioAdapter, gongStruct)
		children = append(children, gongStructNode)
	}

	slices.SortFunc(children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
