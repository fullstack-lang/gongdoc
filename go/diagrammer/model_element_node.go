package diagrammer

type ModelElementNode interface {
	ModelNode

	GetElement() any
	AddToDiagram()
	RemoveFromDiagram()

	// CanBeDisplayed returns true if the model element can be displayed
	CanBeDisplayed() bool
}
