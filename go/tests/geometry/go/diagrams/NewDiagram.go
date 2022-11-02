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
		{
			Name: `Note on the models`,
			Body: `This is an example of a note that
could be displayed on a diagram.

It could explain one aspect of the model
for intance, describing relations between structs

The text of a UML note refers a comment with the GONGNOTE keyword which is
a special case of go Note convention. See example
for details in the go code of the models.

This follows the go code convention described in https://pkg.go.dev/go/doc#Note

"A Note represents a marked comment starting with "MARKER(uid): note body".
Any note with a marker of 2 or more upper case [A-Z] letters and a uid of at least one character is recognized.
The ":" following the uid is optional. Notes are collected in the Package.Notes map indexed by the notes marker."

In the UML diagram, the size of the note is automaticaly computed from the note
number of lines (for the width) and the number of characters per line (for the height)
in the go code
`,
			X:      30.000000,
			Y:      30.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
