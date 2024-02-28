package diagrammer

type ElementNode interface {
	ModelNode

	GetElement() any
}
