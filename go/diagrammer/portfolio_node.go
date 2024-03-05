package diagrammer

// PortfolioNode has to be implemented by the Adapter Nodes
type PortfolioNode interface {

	// GetName() returns the Name that is displayed on the node
	GetName() string

	// GetParent returns the node above
	GetParent() PortfolioNode

	// GetChildren returns the nodes below the node
	GetChildren() []PortfolioNode

	// AppendChildren
	AppendChildren(PortfolioNode)

	// RemoveChildren
	RemoveChildren(PortfolioNode)

	// IsExpanded is true if the the node is visualy expanded
	IsExpanded() bool
}
