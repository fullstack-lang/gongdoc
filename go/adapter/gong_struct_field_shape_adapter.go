package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongStructFieldShapeAdapter struct {
	fieldShape *gongdoc_models.Field

	element gong_models.FieldInterface
}

// GetElement implements diagrammer.Shape.
func (gongStructFieldShapeAdapter *GongStructFieldShapeAdapter) GetElement() diagrammer.ModelElement {
	return gongStructFieldShapeAdapter.element
}
