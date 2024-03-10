package diagrammer

type Model interface {
	GenerateChildren() []ModelNode
}
