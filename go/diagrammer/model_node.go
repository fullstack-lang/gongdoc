package diagrammer

type ModelNode interface {
	GetChildren() []ModelNode
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool

	// GetElement return the pointer to the model element
	GetElement() any
}
