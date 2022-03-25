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
				X: 118.000000,
				Y: 153.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 699.000000,
						Y: 65.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 150.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
			},
		},
		{
			Struct: &(models.PointPartition{}),
			Position: &uml.Position{
				X: 447.000000,
				Y: 352.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.PointPartition{}.Points,
					Middlevertice: &uml.Vertice{
						X: 703.500000,
						Y: 302.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
		},
	},
}
