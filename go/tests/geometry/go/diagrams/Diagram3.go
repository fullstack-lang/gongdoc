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
						X: 695.000000,
						Y: 99.000000,
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
				Y: 170.000000,
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
			Struct: &(models.PointExclusiveSet{}),
			Position: &uml.Position{
				X: 140.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointExclusiveSet{}.Points,
					Middlevertice: &uml.Vertice{
						X: 620.000000,
						Y: 321.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointExclusiveSet{}.Name,
				},
			},
		},
		{
			Struct: &(models.PointNonExclusiveSet{}),
			Position: &uml.Position{
				X: 140.000000,
				Y: 420.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.PointNonExclusiveSet{}.Points,
					Middlevertice: &uml.Vertice{
						X: 515.000000,
						Y: 459.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointNonExclusiveSet{}.Name,
				},
			},
		},
		{
			Struct: &(models.PointUse{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 420.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.PointUse{}.Points,
					Middlevertice: &uml.Vertice{
						X: 770.000000,
						Y: 364.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointUse{}.Name,
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
			X:      70.000000,
			Y:      550.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Name: `Short note on the models`,
			Body: `this is an of a short note
`,
			X:      50.000000,
			Y:      190.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
