package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var Diagram2 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 120.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 450.000000,
				Y: 170.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
	Notes: []*uml.Note{
	},
}