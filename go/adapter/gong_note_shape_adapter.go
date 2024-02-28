package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongNoteShapeAdapter struct {
	gongNoteShape *gongdoc_models.NoteShape

	element *gong_models.GongNote
}

// GetElement implements diagrammer.Shape.
func (gongNoteShapeAdapter *GongNoteShapeAdapter) GetElement() diagrammer.ModelElement {
	return gongNoteShapeAdapter.element
}
