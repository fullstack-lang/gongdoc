package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var Diagram3 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 130.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 699.000000,
						Y: 95.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Line{}.CreationDate,
				},
				{
					Field: models.Line{}.JourneyTime,
				},
				{
					Field: models.Line{}.LineType,
				},
				{
					Field: models.Line{}.Name,
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.Name,
				},
				{
					Field: models.Point{}.X,
				},
				{
					Field: models.Point{}.Y,
				},
			},
		},
		{
			Struct: &(models.PointPartition{}),
			Position: &uml.Position{
				X: 140.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointPartition{}.Points,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 321.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointPartition{}.Name,
				},
			},
		},
	},
	Notes: []*uml.Note{
		{
			Name:    "First note",
			Content: "This the is displayed content of the note",
			X:       200,
			Y:       200,
			Width:   300,
			Heigth:  150,
		},
		{
			Name:    "Second note",
			Content: "This the is displayed content of the note\nThe is the second line",
			X:       400,
			Y:       500,
			Width:   300,
			Heigth:  150,
		},
	},
}
