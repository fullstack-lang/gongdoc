package diagrammer

type Portfolio interface {
	GetChildren() []PortfolioNode
	IsInSelectionMode() bool // the end user can select a diagram to display
	GenerateSVG()
}
