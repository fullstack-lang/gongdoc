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
				X: 50.000000,
				Y: 220.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 452.500000,
						Y: 346.000000,
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
			Name: `Short note on the models`,
			Body: `this is an example of a short note
`,
			X:      320.000000,
			Y:      120.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
