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
				X: 150.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.DOTTED,
				},
				{
					Field: models.CONTINUOUS,
				},
			},
		},
		{
			ReferencedGong: new(models.SimulationStage),
			Position: &uml.Position{
				X: 240.000000,
				Y: 130.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.AFTER,
				},
				{
					Field: models.BEFORE,
				},
			},
		},
	},
	Notes: []*uml.NoteShape{
	},
}
