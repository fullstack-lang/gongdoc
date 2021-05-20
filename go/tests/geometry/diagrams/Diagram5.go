package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/models"
)

var Diagram5 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 90.000000,
				Y: 80.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 300.000000,
				Y: 220.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
