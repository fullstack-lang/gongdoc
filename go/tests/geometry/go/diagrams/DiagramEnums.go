package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var DiagramEnums uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: new(models.LineType),
			Position: &uml.Position{
				X: 40.000000,
				Y: 61.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.CONTINUOUS,
				},
			},
		},
	},
	Notes: []*uml.NoteShape{},
}
