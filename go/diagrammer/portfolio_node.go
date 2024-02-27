package diagrammer

// PortfolioNode has to be implemented by the Adapter Nodes
type PortfolioNode interface {

	// GetName() returns the Name that is displayed on the node
	GetName() string

	// GetChildren returns the nodes below the node
	GetChildren() []PortfolioNode

	// IsExpanded is true if the the node is visualy expanded
	IsExpanded() bool

	// HasAddButton is true if a "Add" button has to be displayed
	HasAddButton() bool
}
