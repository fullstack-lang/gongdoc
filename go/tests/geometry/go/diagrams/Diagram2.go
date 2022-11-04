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
				X: 120.000000,
				Y: 190.000000,
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
	},
	Notes: []*uml.NoteShape{
		{
			Name: `Short note on the models`,
			Body: `this is an of a short note
`,
			X:      430.000000,
			Y:      100.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
