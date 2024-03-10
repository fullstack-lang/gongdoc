package diagrammer

type ModelElementNode interface {
	ModelNode

	GetElement() any
}
