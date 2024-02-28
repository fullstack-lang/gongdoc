package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongEnumValueShapeAdapter struct {
	valueShape *gongdoc_models.GongEnumValueEntry

	element *gong_models.GongEnumValue
}

// GetElement implements diagrammer.Shape.
func (gongEnumValueShapeAdapter *GongEnumValueShapeAdapter) GetElement() diagrammer.ModelElement {
	return gongEnumValueShapeAdapter.element
}
