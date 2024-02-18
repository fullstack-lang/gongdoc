package adapter

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	stage *gongdoc_models.StageStruct
}

func NewPortfolioAdapter(stage *gongdoc_models.StageStruct) (adapter *PortfolioAdapter) {
	adapter = new(PortfolioAdapter)
	adapter.stage = stage
	return
}

// GetRootNodes implements bridge.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetChildren() (rootNodes []diagrammer.PortfolioNode) {
	classDiagramCategoryNode := NewClassDiagramCategoryNode(portfolioAdapter.stage, "class diagrams")
	rootNodes = append(rootNodes, classDiagramCategoryNode)

	return
}
