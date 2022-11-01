package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var NewDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 170.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.PointExclusiveSet{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 220.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointExclusiveSet{}.Points,
					Middlevertice: &uml.Vertice{
						X: 331.500000,
						Y: 339.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
		},
	},
	Notes: []*uml.Note{
	},
}
