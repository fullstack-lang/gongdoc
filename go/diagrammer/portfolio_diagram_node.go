package diagrammer

type PortfolioDiagramNode interface {
	PortfolioNode

	GetDiagram() Diagram

	// DisplayDiagram request the portfolio to display the diagram
	DisplayDiagram() map[ModelNode]Shape

	HasEditButton() bool
	EditDiagram() map[ModelNode]Shape
	CancelEdit() map[ModelNode]Shape
	SaveDiagram() map[ModelNode]Shape
	IsInDrawingMode() bool

	HasDiagramRenameButton() bool
	RenameDiagram(newName string) error
	IsInRenameMode() bool
	SetIsInRenameMode(isInRenameMode bool)

	HasDuplicateButton() bool
	DuplicateDiagram() PortfolioDiagramNode // returns the new diagram

	HasDeleteButton() bool
	DeleteDiagram()
}
