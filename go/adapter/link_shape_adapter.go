package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type LinkShapeAdapter struct {
	linkShape *gongdoc_models.Link

	element gong_models.FieldInterface
}

// GetElement implements diagrammer.Shape.
func (linkShapeAdapter *LinkShapeAdapter) GetElement() diagrammer.ModelElement {
	return linkShapeAdapter.element
}
