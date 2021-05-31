package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/models"
)

var Diagram6 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 130.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Line2{}),
			Position: &uml.Position{
				X: 430.000000,
				Y: 270.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 330.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
