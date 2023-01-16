package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Diagram3 models.StageStruct
var ___dummy__Time_Diagram3 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Diagram3 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Diagram3 map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Diagram3"] = Diagram3Injection
// }

// Diagram3Injection will stage objects of database "Diagram3"
func Diagram3Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Diagram3 := (&models.Classdiagram{Name: `Diagram3`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0003 := (&models.Classshape{Name: `Classshape0003`}).Stage()
	__Classshape__000001_Classshape0004 := (&models.Classshape{Name: `Classshape0004`}).Stage()
	__Classshape__000002_Classshape0005 := (&models.Classshape{Name: `Classshape0005`}).Stage()
	__Classshape__000003_Classshape0006 := (&models.Classshape{Name: `Classshape0006`}).Stage()
	__Classshape__000004_Classshape0007 := (&models.Classshape{Name: `Classshape0007`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage()
	__Field__000001_JourneyTime := (&models.Field{Name: `JourneyTime`}).Stage()
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000007_Type := (&models.Field{Name: `Type`}).Stage()
	__Field__000008_VeryLongLongLongLongLongLongField := (&models.Field{Name: `VeryLongLongLongLongLongLongField`}).Stage()
	__Field__000009_X := (&models.Field{Name: `X`}).Stage()
	__Field__000010_Y := (&models.Field{Name: `Y`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Points := (&models.Link{Name: `Points`}).Stage()
	__Link__000001_Points := (&models.Link{Name: `Points`}).Stage()
	__Link__000002_Points := (&models.Link{Name: `Points`}).Stage()
	__Link__000003_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_LongNodeOnModels := (&models.NoteShape{Name: `LongNodeOnModels`}).Stage()
	__NoteShape__000001_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Position_0003 := (&models.Position{Name: `Position-0003`}).Stage()
	__Position__000001_Position_0004 := (&models.Position{Name: `Position-0004`}).Stage()
	__Position__000002_Position_0005 := (&models.Position{Name: `Position-0005`}).Stage()
	__Position__000003_Position_0006 := (&models.Position{Name: `Position-0006`}).Stage()
	__Position__000004_Position_0007 := (&models.Position{Name: `Position-0007`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()
	__Reference__000002_PointExclusiveSet := (&models.Reference{Name: `PointExclusiveSet`}).Stage()
	__Reference__000003_PointNonExclusiveSet := (&models.Reference{Name: `PointNonExclusiveSet`}).Stage()
	__Reference__000004_PointUse := (&models.Reference{Name: `PointUse`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0003 := (&models.Vertice{Name: `Vertice-0003`}).Stage()
	__Vertice__000001_Vertice_0004 := (&models.Vertice{Name: `Vertice-0004`}).Stage()
	__Vertice__000002_Vertice_0005 := (&models.Vertice{Name: `Vertice-0005`}).Stage()
	__Vertice__000003_Vertice_0006 := (&models.Vertice{Name: `Vertice-0006`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram3.Name = `Diagram3`
	__Classdiagram__000000_Diagram3.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0003.Name = `Classshape0003`
	__Classshape__000000_Classshape0003.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line]
	__Classshape__000000_Classshape0003.Identifier = `Ref_models.Line`
	__Classshape__000000_Classshape0003.ShowNbInstances = true
	__Classshape__000000_Classshape0003.NbInstances = 0
	__Classshape__000000_Classshape0003.Width = 400.000000
	__Classshape__000000_Classshape0003.Heigth = 123.000000
	__Classshape__000000_Classshape0003.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0004.Name = `Classshape0004`
	__Classshape__000001_Classshape0004.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000001_Classshape0004.Identifier = `Ref_models.Point`
	__Classshape__000001_Classshape0004.ShowNbInstances = true
	__Classshape__000001_Classshape0004.NbInstances = 0
	__Classshape__000001_Classshape0004.Width = 240.000000
	__Classshape__000001_Classshape0004.Heigth = 93.000000
	__Classshape__000001_Classshape0004.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0005.Name = `Classshape0005`
	__Classshape__000002_Classshape0005.ReferenceName = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet]
	__Classshape__000002_Classshape0005.Identifier = `Ref_models.PointExclusiveSet`
	__Classshape__000002_Classshape0005.ShowNbInstances = true
	__Classshape__000002_Classshape0005.NbInstances = 0
	__Classshape__000002_Classshape0005.Width = 240.000000
	__Classshape__000002_Classshape0005.Heigth = 63.000000
	__Classshape__000002_Classshape0005.IsSelected = false

	// Classshape values setup
	__Classshape__000003_Classshape0006.Name = `Classshape0006`
	__Classshape__000003_Classshape0006.ReferenceName = `PointNonExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointNonExclusiveSet]
	__Classshape__000003_Classshape0006.Identifier = `Ref_models.PointNonExclusiveSet`
	__Classshape__000003_Classshape0006.ShowNbInstances = true
	__Classshape__000003_Classshape0006.NbInstances = 0
	__Classshape__000003_Classshape0006.Width = 240.000000
	__Classshape__000003_Classshape0006.Heigth = 63.000000
	__Classshape__000003_Classshape0006.IsSelected = false

	// Classshape values setup
	__Classshape__000004_Classshape0007.Name = `Classshape0007`
	__Classshape__000004_Classshape0007.ReferenceName = `PointUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointUse]
	__Classshape__000004_Classshape0007.Identifier = `Ref_models.PointUse`
	__Classshape__000004_Classshape0007.ShowNbInstances = true
	__Classshape__000004_Classshape0007.NbInstances = 0
	__Classshape__000004_Classshape0007.Width = 240.000000
	__Classshape__000004_Classshape0007.Heigth = 63.000000
	__Classshape__000004_Classshape0007.IsSelected = false

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`
	__Field__000000_CreationDate.Fieldname = `CreationDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.CreationDate]
	__Field__000000_CreationDate.Identifier = `Ref_models.Line.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Field values setup
	__Field__000001_JourneyTime.Name = `JourneyTime`
	__Field__000001_JourneyTime.Fieldname = `JourneyTime`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.JourneyTime]
	__Field__000001_JourneyTime.Identifier = `Ref_models.Line.JourneyTime`
	__Field__000001_JourneyTime.FieldTypeAsString = ``
	__Field__000001_JourneyTime.Structname = `Line`
	__Field__000001_JourneyTime.Fieldtypename = `Duration`

	// Field values setup
	__Field__000002_Name.Name = `Name`
	__Field__000002_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Name]
	__Field__000002_Name.Identifier = `Ref_models.Line.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Line`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`
	__Field__000003_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointNonExclusiveSet.Name]
	__Field__000003_Name.Identifier = `Ref_models.PointNonExclusiveSet.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `PointNonExclusiveSet`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`
	__Field__000004_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointUse.Name]
	__Field__000004_Name.Identifier = `Ref_models.PointUse.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `PointUse`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Name.Name = `Name`
	__Field__000005_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point.Name]
	__Field__000005_Name.Identifier = `Ref_models.Point.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Point`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Name.Name = `Name`
	__Field__000006_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet.Name]
	__Field__000006_Name.Identifier = `Ref_models.PointExclusiveSet.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `PointExclusiveSet`
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Type.Name = `Type`
	__Field__000007_Type.Fieldname = `Type`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Type]
	__Field__000007_Type.Identifier = `Ref_models.Line.Type`
	__Field__000007_Type.FieldTypeAsString = ``
	__Field__000007_Type.Structname = `Line`
	__Field__000007_Type.Fieldtypename = `LineType`

	// Field values setup
	__Field__000008_VeryLongLongLongLongLongLongField.Name = `VeryLongLongLongLongLongLongField`
	__Field__000008_VeryLongLongLongLongLongLongField.Fieldname = `VeryLongLongLongLongLongLongField`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.VeryLongLongLongLongLongLongField]
	__Field__000008_VeryLongLongLongLongLongLongField.Identifier = `Ref_models.Line.VeryLongLongLongLongLongLongField`
	__Field__000008_VeryLongLongLongLongLongLongField.FieldTypeAsString = ``
	__Field__000008_VeryLongLongLongLongLongLongField.Structname = `Line`
	__Field__000008_VeryLongLongLongLongLongLongField.Fieldtypename = `string`

	// Field values setup
	__Field__000009_X.Name = `X`
	__Field__000009_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point.X]
	__Field__000009_X.Identifier = `Ref_models.Point.X`
	__Field__000009_X.FieldTypeAsString = ``
	__Field__000009_X.Structname = `Point`
	__Field__000009_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_Y.Name = `Y`
	__Field__000010_Y.Fieldname = `Y`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point.Y]
	__Field__000010_Y.Identifier = `Ref_models.Point.Y`
	__Field__000010_Y.FieldTypeAsString = ``
	__Field__000010_Y.Structname = `Point`
	__Field__000010_Y.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_Points.Name = `Points`
	__Link__000000_Points.Fieldname = `Points`
	__Link__000000_Points.Structname = `PointUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointUse.Points]
	__Link__000000_Points.Identifier = `Ref_models.PointUse.Points`
	__Link__000000_Points.Fieldtypename = `Point`
	__Link__000000_Points.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Points.SourceMultiplicity = models.MANY

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
	__Link__000002_Points.Name = `Points`
	__Link__000002_Points.Fieldname = `Points`
	__Link__000002_Points.Structname = `PointNonExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointNonExclusiveSet.Points]
	__Link__000002_Points.Identifier = `Ref_models.PointNonExclusiveSet.Points`
	__Link__000002_Points.Fieldtypename = `PointUse`
	__Link__000002_Points.TargetMultiplicity = models.MANY
	__Link__000002_Points.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000003_Start.Name = `Start`
	__Link__000003_Start.Fieldname = `Start`
	__Link__000003_Start.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Start]
	__Link__000003_Start.Identifier = `Ref_models.Line.Start`
	__Link__000003_Start.Fieldtypename = `Point`
	__Link__000003_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_Start.SourceMultiplicity = models.MANY

	// NoteShape values setup
	__NoteShape__000000_LongNodeOnModels.Name = `LongNodeOnModels`
	__NoteShape__000000_LongNodeOnModels.Body = `This is an example of a note that
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
`
	__NoteShape__000000_LongNodeOnModels.X = 70.000000
	__NoteShape__000000_LongNodeOnModels.Y = 550.000000
	__NoteShape__000000_LongNodeOnModels.Width = 240.000000
	__NoteShape__000000_LongNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_LongNodeOnModels.Matched = true

	// NoteShape values setup
	__NoteShape__000001_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__NoteShape__000001_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000001_ShortNodeOnModels.X = 90.000000
	__NoteShape__000001_ShortNodeOnModels.Y = 220.000000
	__NoteShape__000001_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000001_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000001_ShortNodeOnModels.Matched = true

	// Position values setup
	__Position__000000_Position_0003.X = 110.000000
	__Position__000000_Position_0003.Y = 40.000000
	__Position__000000_Position_0003.Name = `Position-0003`

	// Position values setup
	__Position__000001_Position_0004.X = 610.000000
	__Position__000001_Position_0004.Y = 150.000000
	__Position__000001_Position_0004.Name = `Position-0004`

	// Position values setup
	__Position__000002_Position_0005.X = 140.000000
	__Position__000002_Position_0005.Y = 290.000000
	__Position__000002_Position_0005.Name = `Position-0005`

	// Position values setup
	__Position__000003_Position_0006.X = 140.000000
	__Position__000003_Position_0006.Y = 420.000000
	__Position__000003_Position_0006.Name = `Position-0006`

	// Position values setup
	__Position__000004_Position_0007.X = 580.000000
	__Position__000004_Position_0007.Y = 420.000000
	__Position__000004_Position_0007.Name = `Position-0007`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 0
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 0
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_PointExclusiveSet.Name = `PointExclusiveSet`
	__Reference__000002_PointExclusiveSet.NbInstances = 0
	__Reference__000002_PointExclusiveSet.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000003_PointNonExclusiveSet.Name = `PointNonExclusiveSet`
	__Reference__000003_PointNonExclusiveSet.NbInstances = 0
	__Reference__000003_PointNonExclusiveSet.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000004_PointUse.Name = `PointUse`
	__Reference__000004_PointUse.NbInstances = 0
	__Reference__000004_PointUse.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0003.X = 695.000000
	__Vertice__000000_Vertice_0003.Y = 99.000000
	__Vertice__000000_Vertice_0003.Name = `Vertice-0003`

	// Vertice values setup
	__Vertice__000001_Vertice_0004.X = 620.000000
	__Vertice__000001_Vertice_0004.Y = 321.500000
	__Vertice__000001_Vertice_0004.Name = `Vertice-0004`

	// Vertice values setup
	__Vertice__000002_Vertice_0005.X = 515.000000
	__Vertice__000002_Vertice_0005.Y = 459.000000
	__Vertice__000002_Vertice_0005.Name = `Vertice-0005`

	// Vertice values setup
	__Vertice__000003_Vertice_0006.X = 770.000000
	__Vertice__000003_Vertice_0006.Y = 364.000000
	__Vertice__000003_Vertice_0006.Name = `Vertice-0006`

	// Setup of pointers
	__Classdiagram__000000_Diagram3.Classshapes = append(__Classdiagram__000000_Diagram3.Classshapes, __Classshape__000000_Classshape0003)
	__Classdiagram__000000_Diagram3.Classshapes = append(__Classdiagram__000000_Diagram3.Classshapes, __Classshape__000001_Classshape0004)
	__Classdiagram__000000_Diagram3.Classshapes = append(__Classdiagram__000000_Diagram3.Classshapes, __Classshape__000002_Classshape0005)
	__Classdiagram__000000_Diagram3.Classshapes = append(__Classdiagram__000000_Diagram3.Classshapes, __Classshape__000003_Classshape0006)
	__Classdiagram__000000_Diagram3.Classshapes = append(__Classdiagram__000000_Diagram3.Classshapes, __Classshape__000004_Classshape0007)
	__Classdiagram__000000_Diagram3.Notes = append(__Classdiagram__000000_Diagram3.Notes, __NoteShape__000000_LongNodeOnModels)
	__Classdiagram__000000_Diagram3.Notes = append(__Classdiagram__000000_Diagram3.Notes, __NoteShape__000001_ShortNodeOnModels)
	__Classshape__000000_Classshape0003.Position = __Position__000000_Position_0003
	__Classshape__000000_Classshape0003.Reference = __Reference__000000_Line
	__Classshape__000000_Classshape0003.Fields = append(__Classshape__000000_Classshape0003.Fields, __Field__000000_CreationDate)
	__Classshape__000000_Classshape0003.Fields = append(__Classshape__000000_Classshape0003.Fields, __Field__000001_JourneyTime)
	__Classshape__000000_Classshape0003.Fields = append(__Classshape__000000_Classshape0003.Fields, __Field__000002_Name)
	__Classshape__000000_Classshape0003.Fields = append(__Classshape__000000_Classshape0003.Fields, __Field__000007_Type)
	__Classshape__000000_Classshape0003.Fields = append(__Classshape__000000_Classshape0003.Fields, __Field__000008_VeryLongLongLongLongLongLongField)
	__Classshape__000000_Classshape0003.Links = append(__Classshape__000000_Classshape0003.Links, __Link__000003_Start)
	__Classshape__000001_Classshape0004.Position = __Position__000001_Position_0004
	__Classshape__000001_Classshape0004.Reference = __Reference__000001_Point
	__Classshape__000001_Classshape0004.Fields = append(__Classshape__000001_Classshape0004.Fields, __Field__000005_Name)
	__Classshape__000001_Classshape0004.Fields = append(__Classshape__000001_Classshape0004.Fields, __Field__000009_X)
	__Classshape__000001_Classshape0004.Fields = append(__Classshape__000001_Classshape0004.Fields, __Field__000010_Y)
	__Classshape__000002_Classshape0005.Position = __Position__000002_Position_0005
	__Classshape__000002_Classshape0005.Reference = __Reference__000002_PointExclusiveSet
	__Classshape__000002_Classshape0005.Fields = append(__Classshape__000002_Classshape0005.Fields, __Field__000006_Name)
	__Classshape__000002_Classshape0005.Links = append(__Classshape__000002_Classshape0005.Links, __Link__000001_Points)
	__Classshape__000003_Classshape0006.Position = __Position__000003_Position_0006
	__Classshape__000003_Classshape0006.Reference = __Reference__000003_PointNonExclusiveSet
	__Classshape__000003_Classshape0006.Fields = append(__Classshape__000003_Classshape0006.Fields, __Field__000003_Name)
	__Classshape__000003_Classshape0006.Links = append(__Classshape__000003_Classshape0006.Links, __Link__000002_Points)
	__Classshape__000004_Classshape0007.Position = __Position__000004_Position_0007
	__Classshape__000004_Classshape0007.Reference = __Reference__000004_PointUse
	__Classshape__000004_Classshape0007.Fields = append(__Classshape__000004_Classshape0007.Fields, __Field__000004_Name)
	__Classshape__000004_Classshape0007.Links = append(__Classshape__000004_Classshape0007.Links, __Link__000000_Points)
	__Link__000000_Points.Middlevertice = __Vertice__000003_Vertice_0006
	__Link__000001_Points.Middlevertice = __Vertice__000001_Vertice_0004
	__Link__000002_Points.Middlevertice = __Vertice__000002_Vertice_0005
	__Link__000003_Start.Middlevertice = __Vertice__000000_Vertice_0003
}


