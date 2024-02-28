package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type NoteLinkShapeAdapter struct {
	nodeShapeLink *gongdoc_models.NoteShapeLink

	element *gong_models.GongLink
}

// GetElement implements diagrammer.Shape.
func (noteLinkShapeAdapter *NoteLinkShapeAdapter) GetElement() diagrammer.ModelElement {
	return noteLinkShapeAdapter.element
}
