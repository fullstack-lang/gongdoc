package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ModelAdapter struct {
	portfolioAdapter *PortfolioAdapter
}

func NewModelAdapter(portfolioAdapter *PortfolioAdapter) (adapter *ModelAdapter) {
	adapter = new(ModelAdapter)
	adapter.portfolioAdapter = portfolioAdapter
	return
}

// GetRootNodes implements bridge.Model.
func (modelAdapter *ModelAdapter) GenerateChildren() (rootNodes []diagrammer.ModelNode) {
	gongStructCategoryNode := NewGongStructCategoryNode(modelAdapter.portfolioAdapter, "gongstructs")
	rootNodes = append(rootNodes, gongStructCategoryNode)

	gongEnumCategoryNode := NewGongEnumCategoryNode(modelAdapter.portfolioAdapter, "gongenums")
	rootNodes = append(rootNodes, gongEnumCategoryNode)

	gongNoteCategoryNode := NewGongNoteCategoryNode(modelAdapter.portfolioAdapter, "gongnotes")
	rootNodes = append(rootNodes, gongNoteCategoryNode)

	return
}
