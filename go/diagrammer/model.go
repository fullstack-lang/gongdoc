package diagrammer

type Model interface {
	GetChildren() []ModelNode
}
