package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/models"
)

var DiagramUMLduUML uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: &(models.Classdiagram{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Classdiagram{}.Classshapes,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 420.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Classdiagram{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Classshape{}),
			Position: &uml.Position{
				X: 70.000000,
				Y: 490.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Links: []*uml.Link{
				{
					Field: models.Classshape{}.Fields,
					Middlevertice: &uml.Vertice{
						X: 470.000000,
						Y: 610.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.Classshape{}.Links,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 620.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.Classshape{}.Position,
					Middlevertice: &uml.Vertice{
						X: 725.000000,
						Y: 609.000000,
					},
					TargetMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Classshape{}.Heigth,
				},
				{
					Field: models.Classshape{}.ReferenceName,
				},
				{
					Field: models.Classshape{}.Width,
				},
			},
		},
		{
			ReferencedGong: &(models.Field{}),
			Position: &uml.Position{
				X: 350.000000,
				Y: 680.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Field{}.FieldTypeAsString,
				},
				{
					Field: models.Field{}.Fieldname,
				},
				{
					Field: models.Field{}.Fieldtypename,
				},
				{
					Field: models.Field{}.Structname,
				},
			},
		},
		{
			ReferencedGong: &(models.Link{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 680.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Links: []*uml.Link{
				{
					Field: models.Link{}.Middlevertice,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 790.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Link{}.Fieldname,
				},
				{
					Field: models.Link{}.Fieldtypename,
				},
				{
					Field: models.Link{}.TargetMultiplicity,
				},
				{
					Field: models.Link{}.Structname,
				},
			},
		},
		{
			ReferencedGong: &(models.DiagramPackage{}),
			Position: &uml.Position{
				X: 350.000000,
				Y: 90.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.DiagramPackage{}.Classdiagrams,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 170.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.DiagramPackage{}.Umlscs,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 180.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.DiagramPackage{}.Name,
				},
				{
					Field: models.DiagramPackage{}.Path,
				},
			},
		},
		{
			ReferencedGong: &(models.Position{}),
			Position: &uml.Position{
				X: 600.000000,
				Y: 680.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Position{}.X,
				},
				{
					Field: models.Position{}.Y,
				},
			},
		},
		{
			ReferencedGong: &(models.UmlState{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 480.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.UmlState{}.Name,
				},
				{
					Field: models.UmlState{}.X,
				},
				{
					Field: models.UmlState{}.Y,
				},
			},
		},
		{
			ReferencedGong: &(models.Umlsc{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Umlsc{}.States,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 420.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Umlsc{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Vertice{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 850.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Vertice{}.X,
				},
				{
					Field: models.Vertice{}.Y,
				},
			},
		},
	},
}
