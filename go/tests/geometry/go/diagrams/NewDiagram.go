package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

var NewDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: new(models.LineType),
			Position: &uml.Position{
				X: 78.000000,
				Y: 16.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
	Notes: []*uml.Note{
	},
}
