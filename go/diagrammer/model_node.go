package diagrammer

type ModelNode interface {
	GenerateProgeny() []ModelNode
	GetChildren() []ModelNode

	GetName() string
	IsNameEditable() bool

	IsExpanded() bool
	SetIsExpanded(isExpanded bool)
}
