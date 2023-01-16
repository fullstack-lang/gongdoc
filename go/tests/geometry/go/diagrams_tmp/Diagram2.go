package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Diagram2 models.StageStruct
var ___dummy__Time_Diagram2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Diagram2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Diagram2 map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Diagram2"] = Diagram2Injection
// }

// Diagram2Injection will stage objects of database "Diagram2"
func Diagram2Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Diagram2 := (&models.Classdiagram{Name: `Diagram2`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0009 := (&models.Classshape{Name: `Classshape0009`}).Stage()
	__Classshape__000001_Classshape0010 := (&models.Classshape{Name: `Classshape0010`}).Stage()
	__Classshape__000002_Classshape0011 := (&models.Classshape{Name: `Classshape0011`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()
	__Link__000001_Points := (&models.Link{Name: `Points`}).Stage()
	__Link__000002_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Position_0009 := (&models.Position{Name: `Position-0009`}).Stage()
	__Position__000001_Position_0010 := (&models.Position{Name: `Position-0010`}).Stage()
	__Position__000002_Position_0011 := (&models.Position{Name: `Position-0011`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()
	__Reference__000002_PointExclusiveSet := (&models.Reference{Name: `PointExclusiveSet`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0006 := (&models.Vertice{Name: `Vertice-0006`}).Stage()
	__Vertice__000001_Vertice_0007 := (&models.Vertice{Name: `Vertice-0007`}).Stage()
	__Vertice__000002_Vertice_0008 := (&models.Vertice{Name: `Vertice-0008`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0009.Name = `Classshape0009`
	__Classshape__000000_Classshape0009.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line]
	__Classshape__000000_Classshape0009.Identifier = `Ref_models.Line`
	__Classshape__000000_Classshape0009.ShowNbInstances = true
	__Classshape__000000_Classshape0009.NbInstances = 47
	__Classshape__000000_Classshape0009.Width = 240.000000
	__Classshape__000000_Classshape0009.Heigth = 78.000000
	__Classshape__000000_Classshape0009.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0010.Name = `Classshape0010`
	__Classshape__000001_Classshape0010.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000001_Classshape0010.Identifier = `Ref_models.Point`
	__Classshape__000001_Classshape0010.ShowNbInstances = true
	__Classshape__000001_Classshape0010.NbInstances = 59
	__Classshape__000001_Classshape0010.Width = 240.000000
	__Classshape__000001_Classshape0010.Heigth = 63.000000
	__Classshape__000001_Classshape0010.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0011.Name = `Classshape0011`
	__Classshape__000002_Classshape0011.ReferenceName = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet]
	__Classshape__000002_Classshape0011.Identifier = `Ref_models.PointExclusiveSet`
	__Classshape__000002_Classshape0011.ShowNbInstances = true
	__Classshape__000002_Classshape0011.NbInstances = 81
	__Classshape__000002_Classshape0011.Width = 240.000000
	__Classshape__000002_Classshape0011.Heigth = 63.000000
	__Classshape__000002_Classshape0011.IsSelected = false

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`
	__Field__000000_CreationDate.Fieldname = `CreationDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.CreationDate]
	__Field__000000_CreationDate.Identifier = `Ref_models.Line.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Fieldname = `End`
	__Link__000000_End.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.End]
	__Link__000000_End.Identifier = `Ref_models.Line.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Points.Name = `Points`
	__Link__000001_Points.Fieldname = `Points`
	__Link__000001_Points.Structname = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet.Points]
	__Link__000001_Points.Identifier = `Ref_models.PointExclusiveSet.Points`
	__Link__000001_Points.Fieldtypename = `Point`
	__Link__000001_Points.TargetMultiplicity = models.MANY
	__Link__000001_Points.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_Start.Name = `Start`
	__Link__000002_Start.Fieldname = `Start`
	__Link__000002_Start.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Start]
	__Link__000002_Start.Identifier = `Ref_models.Line.Start`
	__Link__000002_Start.Fieldtypename = `Point`
	__Link__000002_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_Start.SourceMultiplicity = models.MANY

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 90.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 60.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = true

	// Position values setup
	__Position__000000_Position_0009.X = 50.000000
	__Position__000000_Position_0009.Y = 210.000000
	__Position__000000_Position_0009.Name = `Position-0009`

	// Position values setup
	__Position__000001_Position_0010.X = 680.000000
	__Position__000001_Position_0010.Y = 220.000000
	__Position__000001_Position_0010.Name = `Position-0010`

	// Position values setup
	__Position__000002_Position_0011.X = 30.000000
	__Position__000002_Position_0011.Y = 350.000000
	__Position__000002_Position_0011.Name = `Position-0011`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 47
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 59
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_PointExclusiveSet.Name = `PointExclusiveSet`
	__Reference__000002_PointExclusiveSet.NbInstances = 81
	__Reference__000002_PointExclusiveSet.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0006.X = 449.000000
	__Vertice__000000_Vertice_0006.Y = 329.500000
	__Vertice__000000_Vertice_0006.Name = `Vertice-0006`

	// Vertice values setup
	__Vertice__000001_Vertice_0007.X = 485.000000
	__Vertice__000001_Vertice_0007.Y = 201.500000
	__Vertice__000001_Vertice_0007.Name = `Vertice-0007`

	// Vertice values setup
	__Vertice__000002_Vertice_0008.X = 699.000000
	__Vertice__000002_Vertice_0008.Y = 389.500000
	__Vertice__000002_Vertice_0008.Name = `Vertice-0008`

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Classshape0009)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000001_Classshape0010)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000002_Classshape0011)
	__Classdiagram__000000_Diagram2.Notes = append(__Classdiagram__000000_Diagram2.Notes, __NoteShape__000000_ShortNodeOnModels)
	__Classshape__000000_Classshape0009.Position = __Position__000000_Position_0009
	__Classshape__000000_Classshape0009.Reference = __Reference__000000_Line
	__Classshape__000000_Classshape0009.Fields = append(__Classshape__000000_Classshape0009.Fields, __Field__000000_CreationDate)
	__Classshape__000000_Classshape0009.Links = append(__Classshape__000000_Classshape0009.Links, __Link__000000_End)
	__Classshape__000000_Classshape0009.Links = append(__Classshape__000000_Classshape0009.Links, __Link__000002_Start)
	__Classshape__000001_Classshape0010.Position = __Position__000001_Position_0010
	__Classshape__000001_Classshape0010.Reference = __Reference__000001_Point
	__Classshape__000002_Classshape0011.Position = __Position__000002_Position_0011
	__Classshape__000002_Classshape0011.Reference = __Reference__000002_PointExclusiveSet
	__Classshape__000002_Classshape0011.Links = append(__Classshape__000002_Classshape0011.Links, __Link__000001_Points)
	__Link__000000_End.Middlevertice = __Vertice__000000_Vertice_0006
	__Link__000001_Points.Middlevertice = __Vertice__000002_Vertice_0008
	__Link__000002_Start.Middlevertice = __Vertice__000001_Vertice_0007
}


