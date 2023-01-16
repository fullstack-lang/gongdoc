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
	__Classshape__000000_Classshape0004 := (&models.Classshape{Name: `Classshape0004`}).Stage()
	__Classshape__000001_Classshape0005 := (&models.Classshape{Name: `Classshape0005`}).Stage()
	__Classshape__000002_Classshape0006 := (&models.Classshape{Name: `Classshape0006`}).Stage()

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
	__Position__000000_Position_0004 := (&models.Position{Name: `Position-0004`}).Stage()
	__Position__000001_Position_0005 := (&models.Position{Name: `Position-0005`}).Stage()
	__Position__000002_Position_0006 := (&models.Position{Name: `Position-0006`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()
	__Reference__000002_PointExclusiveSet := (&models.Reference{Name: `PointExclusiveSet`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0002 := (&models.Vertice{Name: `Vertice-0002`}).Stage()
	__Vertice__000001_Vertice_0003 := (&models.Vertice{Name: `Vertice-0003`}).Stage()
	__Vertice__000002_Vertice_0004 := (&models.Vertice{Name: `Vertice-0004`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0004.Name = `Classshape0004`
	__Classshape__000000_Classshape0004.ReferenceName = `Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line]
	__Classshape__000000_Classshape0004.Identifier = `Ref_models.Line`
	__Classshape__000000_Classshape0004.ShowNbInstances = true
	__Classshape__000000_Classshape0004.NbInstances = 81
	__Classshape__000000_Classshape0004.Width = 240.000000
	__Classshape__000000_Classshape0004.Heigth = 78.000000
	__Classshape__000000_Classshape0004.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0005.Name = `Classshape0005`
	__Classshape__000001_Classshape0005.ReferenceName = `Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000001_Classshape0005.Identifier = `Ref_models.Point`
	__Classshape__000001_Classshape0005.ShowNbInstances = true
	__Classshape__000001_Classshape0005.NbInstances = 81
	__Classshape__000001_Classshape0005.Width = 240.000000
	__Classshape__000001_Classshape0005.Heigth = 63.000000
	__Classshape__000001_Classshape0005.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0006.Name = `Classshape0006`
	__Classshape__000002_Classshape0006.ReferenceName = `PointExclusiveSet`

	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet]
	__Classshape__000002_Classshape0006.Identifier = `Ref_models.PointExclusiveSet`
	__Classshape__000002_Classshape0006.ShowNbInstances = true
	__Classshape__000002_Classshape0006.NbInstances = 87
	__Classshape__000002_Classshape0006.Width = 240.000000
	__Classshape__000002_Classshape0006.Heigth = 63.000000
	__Classshape__000002_Classshape0006.IsSelected = false

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
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Points.Name = `Points`
	__Link__000001_Points.Fieldname = `Points`
	__Link__000001_Points.Structname = `PointExclusiveSet`
	__Link__000001_Points.Fieldtypename = `Point`
	__Link__000001_Points.TargetMultiplicity = models.MANY
	__Link__000001_Points.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_Start.Name = `Start`
	__Link__000002_Start.Fieldname = `Start`
	__Link__000002_Start.Structname = `Line`
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
	__Position__000000_Position_0004.X = 50.000000
	__Position__000000_Position_0004.Y = 210.000000
	__Position__000000_Position_0004.Name = `Position-0004`

	// Position values setup
	__Position__000001_Position_0005.X = 680.000000
	__Position__000001_Position_0005.Y = 220.000000
	__Position__000001_Position_0005.Name = `Position-0005`

	// Position values setup
	__Position__000002_Position_0006.X = 30.000000
	__Position__000002_Position_0006.Y = 350.000000
	__Position__000002_Position_0006.Name = `Position-0006`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 81
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 81
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_PointExclusiveSet.Name = `PointExclusiveSet`
	__Reference__000002_PointExclusiveSet.NbInstances = 87
	__Reference__000002_PointExclusiveSet.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0002.X = 449.000000
	__Vertice__000000_Vertice_0002.Y = 329.500000
	__Vertice__000000_Vertice_0002.Name = `Vertice-0002`

	// Vertice values setup
	__Vertice__000001_Vertice_0003.X = 485.000000
	__Vertice__000001_Vertice_0003.Y = 201.500000
	__Vertice__000001_Vertice_0003.Name = `Vertice-0003`

	// Vertice values setup
	__Vertice__000002_Vertice_0004.X = 699.000000
	__Vertice__000002_Vertice_0004.Y = 389.500000
	__Vertice__000002_Vertice_0004.Name = `Vertice-0004`

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Classshape0004)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000001_Classshape0005)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000002_Classshape0006)
	__Classdiagram__000000_Diagram2.Notes = append(__Classdiagram__000000_Diagram2.Notes, __NoteShape__000000_ShortNodeOnModels)
	__Classshape__000000_Classshape0004.Position = __Position__000000_Position_0004
	__Classshape__000000_Classshape0004.Reference = __Reference__000000_Line
	__Classshape__000000_Classshape0004.Fields = append(__Classshape__000000_Classshape0004.Fields, __Field__000000_CreationDate)
	__Classshape__000000_Classshape0004.Links = append(__Classshape__000000_Classshape0004.Links, __Link__000000_End)
	__Classshape__000000_Classshape0004.Links = append(__Classshape__000000_Classshape0004.Links, __Link__000002_Start)
	__Classshape__000001_Classshape0005.Position = __Position__000001_Position_0005
	__Classshape__000001_Classshape0005.Reference = __Reference__000001_Point
	__Classshape__000002_Classshape0006.Position = __Position__000002_Position_0006
	__Classshape__000002_Classshape0006.Reference = __Reference__000002_PointExclusiveSet
	__Classshape__000002_Classshape0006.Links = append(__Classshape__000002_Classshape0006.Links, __Link__000001_Points)
	__Link__000000_End.Middlevertice = __Vertice__000000_Vertice_0002
	__Link__000001_Points.Middlevertice = __Vertice__000002_Vertice_0004
	__Link__000002_Start.Middlevertice = __Vertice__000001_Vertice_0003
}
