package diagrammer

// Portfolio interface to be implemented by the Portfolio adapter
type Portfolio interface {

	// GetChildren returns the root nodes of the tree for navigating the portfolio
	GetChildren() []PortfolioNode

	// IsInSelectionMode is true is at least one diagram has been selected
	IsInSelectionMode() bool // the end user can select a diagram to display

	// DisplayDiagram
	DisplayDiagram(diagramNode DiagramNode) map[ModelNode]Shape
}
