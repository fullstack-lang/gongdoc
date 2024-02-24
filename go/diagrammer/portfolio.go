package diagrammer

// Portfolio interface to be implemented by the Portfolio adapter
type Portfolio interface {

	// GetChildren returns the root nodes of the tree for navigating the portfolio
	GetChildren() []PortfolioNode

	// IsInSelectionMode is true is at least diagram has been selected
	IsInSelectionMode() bool // the end user can select a diagram to display

	// GenerateSVG is called from the adapter and implemented by the adapter
	// Its use is to allow for the clean switch off of the function
	GenerateSVG()

	// AddDiagram allows the end user to request the creation a new Diagram/PortfolioNode
	AddDiagram() PortfolioNode
}
