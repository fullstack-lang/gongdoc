package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
	// insertion points for import of the illustrated modelgithub.com/fullstack-lang/gongdoc/go/tests/geometry/models
)

var UmlscDiagram1 uml.Umlsc = uml.Umlsc{
	Activestate: string(models.AFTER),
	States: []*uml.UmlState{
		{
			X:    20.000000,
			Y:    90.000000,
			Name: string(models.AFTER),
		},
		{
			X:    20.000000,
			Y:    30.000000,
			Name: string(models.BEFORE),
		},
	},
}