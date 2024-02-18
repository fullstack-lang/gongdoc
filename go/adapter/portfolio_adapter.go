package adapter

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	stage *gongdoc_models.StageStruct
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (*PortfolioAdapter) IsInSelectionMode() bool {
	return true // temp, will be false when a diagram is selected
}

// GetSelectedDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedDiagram() (diagram diagrammer.Diagram) {
	return
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
