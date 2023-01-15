package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var Diagram2 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: &(models.Line{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 210.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 449.000000,
						Y: 329.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 485.000000,
						Y: 201.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
		},
		{
			ReferencedGong: &(models.Point{}),
			Position: &uml.Position{
				X: 680.000000,
				Y: 220.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			ReferencedGong: &(models.PointExclusiveSet{}),
			Position: &uml.Position{
				X: 30.000000,
				Y: 350.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointExclusiveSet{}.Points,
					Middlevertice: &uml.Vertice{
						X: 699.000000,
						Y: 389.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
		},
	},
	Notes: []*uml.NoteShape{
		{
			Name: `ShortNodeOnModels`,
			Body: `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`,
			X:      90.000000,
			Y:      60.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
