package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongdoc/go/models"
)

var defaultDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Classdiagram{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Classdiagram{}.Classshapes,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 200.000000,
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
			Struct: &(models.Classshape{}),
			Position: &uml.Position{
				X: 340.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Links: []*uml.Link{
				{
					Field: models.Classshape{}.Fields,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 200.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.Classshape{}.Links,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 250.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.Classshape{}.Position,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 300.000000,
					},
					TargetMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Classshape{}.ClassshapeTargetType,
				},
				{
					Field: models.Classshape{}.Heigth,
				},
				{
					Field: models.Classshape{}.Name,
				},
				{
					Field: models.Classshape{}.Struct,
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
			Struct: new(models.ClassshapeTargetType),
			Position: &uml.Position{
				X: 40.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.ENUM,
				},
				{
					Field: models.STRUCT,
				},
			},
		},
		{
			Struct: &(models.Field{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Fields: []*uml.Field{
				{
					Field: models.Field{}.Field,
				},
				{
					Field: models.Field{}.Fieldname,
				},
				{
					Field: models.Field{}.Fieldtypename,
				},
				{
					Field: models.Field{}.Name,
				},
				{
					Field: models.Field{}.Structname,
				},
			},
		},
		{
			Struct: &(models.Link{}),
			Position: &uml.Position{
				X: 1240.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Links: []*uml.Link{
				{
					Field: models.Link{}.Middlevertice,
					Middlevertice: &uml.Vertice{
						X: 1490.000000,
						Y: 200.000000,
					},
					TargetMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Link{}.Field,
				},
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
					Field: models.Link{}.Name,
				},
				{
					Field: models.Link{}.Structname,
				},
			},
		},
		{
			Struct: new(models.MultiplicityType),
			Position: &uml.Position{
				X: 340.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.MANY,
				},
				{
					Field: models.ONE,
				},
				{
					Field: models.ZERO_ONE,
				},
			},
		},
		{
			Struct: &(models.DiagramPackage{}),
			Position: &uml.Position{
				X: 1540.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.DiagramPackage{}.Classdiagrams,
					Middlevertice: &uml.Vertice{
						X: 1790.000000,
						Y: 200.000000,
					},
					TargetMultiplicity: "*",
				},
				{
					Field: models.DiagramPackage{}.Umlscs,
					Middlevertice: &uml.Vertice{
						X: 1790.000000,
						Y: 250.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.DiagramPackage{}.Name,
				},
			},
		},
		{
			Struct: &(models.Position{}),
			Position: &uml.Position{
				X: 1840.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.Position{}.Name,
				},
				{
					Field: models.Position{}.X,
				},
				{
					Field: models.Position{}.Y,
				},
			},
		},
		{
			Struct: &(models.UmlState{}),
			Position: &uml.Position{
				X: 2140.000000,
				Y: 40.000000,
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
			Struct: &(models.Umlsc{}),
			Position: &uml.Position{
				X: 2440.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Umlsc{}.States,
					Middlevertice: &uml.Vertice{
						X: 2690.000000,
						Y: 200.000000,
					},
					TargetMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Umlsc{}.Activestate,
				},
				{
					Field: models.Umlsc{}.Name,
				},
			},
		},
		{
			Struct: &(models.Vertice{}),
			Position: &uml.Position{
				X: 2740.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.Vertice{}.Name,
				},
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
