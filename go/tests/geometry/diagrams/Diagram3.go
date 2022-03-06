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
				X: 159.000000,
				Y: 134.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 559.500000,
						Y: 238.500000,
					},
					Multiplicity: "0..1",
				},
			},
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
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 80.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
			},
		},
	},
}
