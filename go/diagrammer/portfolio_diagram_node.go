package diagrammer

type PortfolioDiagramNode interface {
	PortfolioNode

	GetDiagram() Diagram

	// DisplayDiagram request the portfolio to display the diagram
	DisplayDiagram() map[ModelNode]Shape

	HasEditButton() bool
	CancelEdit()
	SaveDiagram()
	IsInEditMode() bool
	SetIsInEditMode(isInEditMode bool)

	HasDiagramRenameButton() bool
	RenameDiagram(newName string) error
	IsInRenameMode() bool
	SetIsInRenameMode(isInRenameMode bool)

	HasDuplicateButton() bool
	DuplicateDiagram() PortfolioDiagramNode // returns the new diagram

	HasDeleteButton() bool
	DeleteDiagram()
}