package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongStructShapeAdapter struct {
	gongStructShape *gongdoc_models.GongStructShape

	element *gong_models.GongStruct
}

// GetElement implements diagrammer.Shape.
func (gongStructShapeAdapter *GongStructShapeAdapter) GetElement() diagrammer.ModelElement {
	return gongStructShapeAdapter.element
}
