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
				X: 40.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 535.000000,
						Y: 184.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 85.000000,
						Y: 274.000000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Line2{}),
			Position: &uml.Position{
				X: 834.000000,
				Y: 79.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line2{}.End,
					Middlevertice: &uml.Vertice{
						X: 932.000000,
						Y: 206.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line2{}.Start,
					Middlevertice: &uml.Vertice{
						X: 762.000000,
						Y: 156.000000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 310.000000,
				Y: 270.000000,
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
			Struct: &(models.PointWithName{}),
			Position: &uml.Position{
				X: 150.000000,
				Y: 572.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Polyline{}),
			Position: &uml.Position{
				X: 701.000000,
				Y: 606.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Polyline{}.Follower,
					Middlevertice: &uml.Vertice{
						X: 1061.000000,
						Y: 637.500000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Polyline{}.Followers,
					Middlevertice: &uml.Vertice{
						X: 1061.000000,
						Y: 637.500000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Polyline{}.Points,
					Middlevertice: &uml.Vertice{
						X: 865.500000,
						Y: 469.500000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Polyline{}.Start_Point,
					Middlevertice: &uml.Vertice{
						X: 525.500000,
						Y: 509.500000,
					},
					Multiplicity: "0..1",
				},
			},
		},
	},
}
