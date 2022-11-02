package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var Diagram4 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 130.000000,
				Y: 150.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 505.000000,
						Y: 309.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 525.000000,
						Y: 99.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Line{}.JourneyTime,
				},
				{
					Field: models.Line{}.Name,
				},
				{
					Field: models.Line{}.Type,
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 520.000000,
				Y: 150.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.Name,
				},
			},
		},
	},
	Notes: []*uml.Note{
	},
}
