package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	gongStage    *gong_models.StageStruct
	gongdocStage *gongdoc_models.StageStruct
	gongsvgStage *gongsvg_models.StageStruct
	diagrammer   *diagrammer.Diagrammer
}

func NewPortfolioAdapter(
	gongStage *gong_models.StageStruct,
	gongdocStage *gongdoc_models.StageStruct,
	svgStage *gongsvg_models.StageStruct,
) (portfolioAdapter *PortfolioAdapter) {
	portfolioAdapter = new(PortfolioAdapter)

	portfolioAdapter.gongStage = gongStage
	portfolioAdapter.gongdocStage = gongdocStage
	portfolioAdapter.gongsvgStage = svgStage

	return
}

func (portfolioAdapter *PortfolioAdapter) SetDiagrammer(diagrammer *diagrammer.Diagrammer) {
	portfolioAdapter.diagrammer = diagrammer
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (*PortfolioAdapter) IsInSelectionMode() bool {
	return true // temp, will be false when a diagram is selected
}

// GetSelectedDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedDiagram() (diagram diagrammer.PortfolioNode) {
	return
}

// GetRootNodes implements bridge.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetChildren() (rootNodes []diagrammer.PortfolioNode) {
	classDiagramCategoryNode := NewClassDiagramCategoryNode(portfolioAdapter, "class diagrams", portfolioAdapter.diagrammer)
	rootNodes = append(rootNodes, classDiagramCategoryNode)

	return
}
