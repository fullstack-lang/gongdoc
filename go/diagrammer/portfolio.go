package diagrammer

type Portfolio interface {
	GetChildren() []PortfolioNode
	GetSelectedPortfolioNode() PortfolioNode
	SetSelectedPortfolioNode(portfolioNode PortfolioNode)
	IsInSelectionMode() bool // the end user can select a diagram to display
}
