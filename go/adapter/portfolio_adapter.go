package adapter

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	stage                 *gongdoc_models.StageStruct
	selectedPortfolioNode diagrammer.PortfolioNode
}

// GetSelectedPortfolioNode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedPortfolioNode() diagrammer.PortfolioNode {
	return portfolioAdapter.selectedPortfolioNode
}

// SetSelectedPortfolioNode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) SetSelectedPortfolioNode(portfolioNode diagrammer.PortfolioNode) {
	portfolioAdapter.selectedPortfolioNode = portfolioNode
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (*PortfolioAdapter) IsInSelectionMode() bool {
	return true // temp, will be false when a diagram is selected
}

// GetSelectedDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedDiagram() (diagram diagrammer.PortfolioNode) {
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
