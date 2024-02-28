package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongEnumShapeAdapter struct {
	gongEnumShape *gongdoc_models.GongEnumShape

	element *gong_models.GongEnum
}

// GetElement implements diagrammer.Shape.
func (gongEnumShapeAdapter *GongEnumShapeAdapter) GetElement() diagrammer.ModelElement {
	return gongEnumShapeAdapter.element
}
