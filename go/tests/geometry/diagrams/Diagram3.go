package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/models"
)

var Diagram3 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 310.000000,
				Y: 120.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 598.000000,
						Y: 280.500000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 166.000000,
				Y: 233.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
				{
					Field: models.Point{}.Y,
				},
			},
		},
		{
			Struct: &(models.PointWithName{}),
			Position: &uml.Position{
				X: 89.000000,
				Y: 320.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
	},
}
