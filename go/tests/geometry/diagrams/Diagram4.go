package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/models"
)

var Diagram4 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 70.000000,
				Y: 100.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 563.500000,
						Y: 237.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 563.500000,
						Y: 237.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 563.500000,
						Y: 237.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 98.500000,
						Y: 237.000000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 227.000000,
				Y: 266.000000,
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
