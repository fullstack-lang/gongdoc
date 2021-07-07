package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/models"
)

var UmlscDiagram1 uml.Umlsc = uml.Umlsc{
	Activestate: string(models.APRES_CALCUL),
	States: []*uml.UmlState{
		{
			X:    20.000000,
			Y:    90.000000,
			Name: string(models.APRES_CALCUL),
		},
		{
			X:    20.000000,
			Y:    30.000000,
			Name: string(models.AVANT_CALCUL),
		},
	},
}
