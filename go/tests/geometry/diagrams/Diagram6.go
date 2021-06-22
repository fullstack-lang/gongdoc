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
				X: 280.000000,
				Y: 119.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 380.000000,
						Y: 556.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 540.000000,
						Y: 248.500000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Line2{}),
			Position: &uml.Position{
				X: 620.000000,
				Y: 232.000000,
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
	},
}
