package diagrammer

// Portfolio interface to be implemented by the Portfolio adapter
type Portfolio interface {
	GenerateTree()

	// GetChildren returns the root nodes of the tree for navigating the portfolio
	GetChildren() []PortfolioNode

	// IsInSelectionMode is true is at least one diagram has been selected
	IsInSelectionMode() bool // the end user can select a diagram to display
	GetSelectedPortfolioDiagramNode() PortfolioDiagramNode

	IsInDrawingMode() bool
	AddElement(ModelNode) map[ModelNode]Shape
	RemoveElement(ModelNode)
}
