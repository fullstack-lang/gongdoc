package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumCategoryNode struct {
	ModelCategoryNodeBase
}

func NewGongEnumCategoryNode(portfolioAdapter *PortfolioAdapter, name string) *GongEnumCategoryNode {
	return &GongEnumCategoryNode{ModelCategoryNodeBase: ModelCategoryNodeBase{portfolioAdapter: portfolioAdapter, Name: name}}
}

// GenerateChildren implements diagrammer.Node.
func (categoryNode *GongEnumCategoryNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](categoryNode.portfolioAdapter.gongStage) {

		gongEnumNode := NewGongEnumNode(categoryNode.portfolioAdapter, gongEnum)
		children = append(children, gongEnumNode)
	}

	slices.SortFunc(children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
