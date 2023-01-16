package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Diagram4 models.StageStruct
var ___dummy__Time_Diagram4 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Diagram4 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Diagram4 map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Diagram4"] = Diagram4Injection
// }

// Diagram4Injection will stage objects of database "Diagram4"
func Diagram4Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Diagram4 := (&models.Classdiagram{Name: `Diagram4`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0008 := (&models.Classshape{Name: `Classshape0008`}).Stage()
	__Classshape__000001_Classshape0009 := (&models.Classshape{Name: `Classshape0009`}).Stage()
	__Classshape__000002_Classshape0010 := (&models.Classshape{Name: `Classshape0010`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_JourneyTime := (&models.Field{Name: `JourneyTime`}).Stage()
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000003_Type := (&models.Field{Name: `Type`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()
	__Link__000001_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Position_0008 := (&models.Position{Name: `Position-0008`}).Stage()
	__Position__000001_Position_0009 := (&models.Position{Name: `Position-0009`}).Stage()
	__Position__000002_Position_0010 := (&models.Position{Name: `Position-0010`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()
	__Reference__000002_PointExclusiveSet := (&models.Reference{Name: `PointExclusiveSet`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0007 := (&models.Vertice{Name: `Vertice-0007`}).Stage()
	__Vertice__000001_Vertice_0008 := (&models.Vertice{Name: `Vertice-0008`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram4.Name = `Diagram4`
	__Classdiagram__000000_Diagram4.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0008.Name = `Classshape0008`
	__Classshape__000000_Classshape0008.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line]
	__Classshape__000000_Classshape0008.Identifier = `Ref_models.Line`
	__Classshape__000000_Classshape0008.ShowNbInstances = true
	__Classshape__000000_Classshape0008.NbInstances = 87
	__Classshape__000000_Classshape0008.Width = 240.000000
	__Classshape__000000_Classshape0008.Heigth = 108.000000
	__Classshape__000000_Classshape0008.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0009.Name = `Classshape0009`
	__Classshape__000001_Classshape0009.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000001_Classshape0009.Identifier = `Ref_models.Point`
	__Classshape__000001_Classshape0009.ShowNbInstances = true
	__Classshape__000001_Classshape0009.NbInstances = 47
	__Classshape__000001_Classshape0009.Width = 240.000000
	__Classshape__000001_Classshape0009.Heigth = 78.000000
	__Classshape__000001_Classshape0009.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0010.Name = `Classshape0010`
	__Classshape__000002_Classshape0010.ReferenceName = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet]
	__Classshape__000002_Classshape0010.Identifier = `Ref_models.PointExclusiveSet`
	__Classshape__000002_Classshape0010.ShowNbInstances = true
	__Classshape__000002_Classshape0010.NbInstances = 59
	__Classshape__000002_Classshape0010.Width = 240.000000
	__Classshape__000002_Classshape0010.Heigth = 63.000000
	__Classshape__000002_Classshape0010.IsSelected = false

	// Field values setup
	__Field__000000_JourneyTime.Name = `JourneyTime`
	__Field__000000_JourneyTime.Fieldname = `JourneyTime`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.JourneyTime]
	__Field__000000_JourneyTime.Identifier = `Ref_models.Line.JourneyTime`
	__Field__000000_JourneyTime.FieldTypeAsString = ``
	__Field__000000_JourneyTime.Structname = `Line`
	__Field__000000_JourneyTime.Fieldtypename = `Duration`

	// Field values setup
	__Field__000001_Name.Name = `Name`
	__Field__000001_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Name]
	__Field__000001_Name.Identifier = `Ref_models.Line.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Line`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`
	__Field__000002_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point.Name]
	__Field__000002_Name.Identifier = `Ref_models.Point.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Point`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Type.Name = `Type`
	__Field__000003_Type.Fieldname = `Type`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Type]
	__Field__000003_Type.Identifier = `Ref_models.Line.Type`
	__Field__000003_Type.FieldTypeAsString = ``
	__Field__000003_Type.Structname = `Line`
	__Field__000003_Type.Fieldtypename = `LineType`

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
	__Link__000001_Start.Name = `Start`
	__Link__000001_Start.Fieldname = `Start`
	__Link__000001_Start.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Start]
	__Link__000001_Start.Identifier = `Ref_models.Line.Start`
	__Link__000001_Start.Fieldtypename = `Point`
	__Link__000001_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Start.SourceMultiplicity = models.MANY

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 40.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 310.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = true

	// Position values setup
	__Position__000000_Position_0008.X = 130.000000
	__Position__000000_Position_0008.Y = 150.000000
	__Position__000000_Position_0008.Name = `Position-0008`

	// Position values setup
	__Position__000001_Position_0009.X = 520.000000
	__Position__000001_Position_0009.Y = 150.000000
	__Position__000001_Position_0009.Name = `Position-0009`

	// Position values setup
	__Position__000002_Position_0010.X = 360.000000
	__Position__000002_Position_0010.Y = 400.000000
	__Position__000002_Position_0010.Name = `Position-0010`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 87
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 47
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_PointExclusiveSet.Name = `PointExclusiveSet`
	__Reference__000002_PointExclusiveSet.NbInstances = 59
	__Reference__000002_PointExclusiveSet.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0007.X = 505.000000
	__Vertice__000000_Vertice_0007.Y = 309.000000
	__Vertice__000000_Vertice_0007.Name = `Vertice-0007`

	// Vertice values setup
	__Vertice__000001_Vertice_0008.X = 525.000000
	__Vertice__000001_Vertice_0008.Y = 99.000000
	__Vertice__000001_Vertice_0008.Name = `Vertice-0008`

	// Setup of pointers
	__Classdiagram__000000_Diagram4.Classshapes = append(__Classdiagram__000000_Diagram4.Classshapes, __Classshape__000000_Classshape0008)
	__Classdiagram__000000_Diagram4.Classshapes = append(__Classdiagram__000000_Diagram4.Classshapes, __Classshape__000001_Classshape0009)
	__Classdiagram__000000_Diagram4.Classshapes = append(__Classdiagram__000000_Diagram4.Classshapes, __Classshape__000002_Classshape0010)
	__Classdiagram__000000_Diagram4.Notes = append(__Classdiagram__000000_Diagram4.Notes, __NoteShape__000000_ShortNodeOnModels)
	__Classshape__000000_Classshape0008.Position = __Position__000000_Position_0008
	__Classshape__000000_Classshape0008.Reference = __Reference__000000_Line
	__Classshape__000000_Classshape0008.Fields = append(__Classshape__000000_Classshape0008.Fields, __Field__000000_JourneyTime)
	__Classshape__000000_Classshape0008.Fields = append(__Classshape__000000_Classshape0008.Fields, __Field__000001_Name)
	__Classshape__000000_Classshape0008.Fields = append(__Classshape__000000_Classshape0008.Fields, __Field__000003_Type)
	__Classshape__000000_Classshape0008.Links = append(__Classshape__000000_Classshape0008.Links, __Link__000000_End)
	__Classshape__000000_Classshape0008.Links = append(__Classshape__000000_Classshape0008.Links, __Link__000001_Start)
	__Classshape__000001_Classshape0009.Position = __Position__000001_Position_0009
	__Classshape__000001_Classshape0009.Reference = __Reference__000001_Point
	__Classshape__000001_Classshape0009.Fields = append(__Classshape__000001_Classshape0009.Fields, __Field__000002_Name)
	__Classshape__000002_Classshape0010.Position = __Position__000002_Position_0010
	__Classshape__000002_Classshape0010.Reference = __Reference__000002_PointExclusiveSet
	__Link__000000_End.Middlevertice = __Vertice__000000_Vertice_0007
	__Link__000001_Start.Middlevertice = __Vertice__000001_Vertice_0008
}


