package diagrammer

type ModelNode interface {
	GenerateChildren() []ModelNode

	GetName() string
	IsNameEditable() bool

	IsExpanded() bool
	SetIsExpanded(isExpanded bool)
}
