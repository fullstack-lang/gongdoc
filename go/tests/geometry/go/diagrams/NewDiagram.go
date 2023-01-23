package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.AFTER": ref_models.AFTER,

	"ref_models.BEFORE": ref_models.BEFORE,

	"ref_models.CONTINUOUS": ref_models.CONTINUOUS,

	"ref_models.CONTINUOUS_ZERO": ref_models.CONTINUOUS_ZERO,

	"ref_models.DOTTED": ref_models.DOTTED,

	"ref_models.DOTTED_ONE": ref_models.DOTTED_ONE,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.CreationDate": (ref_models.Line{}).CreationDate,

	"ref_models.Line.End": (ref_models.Line{}).End,

	"ref_models.Line.JourneyTime": (ref_models.Line{}).JourneyTime,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Start": (ref_models.Line{}).Start,

	"ref_models.Line.Type": (ref_models.Line{}).Type,

	"ref_models.Line.VeryLongLongLongLongLongLongField": (ref_models.Line{}).VeryLongLongLongLongLongLongField,

	"ref_models.LineTypeInt": ref_models.LineTypeInt(0),

	"ref_models.LineTypeString": ref_models.LineTypeString(""),

	"ref_models.LongNodeOnModels": ref_models.LongNodeOnModels,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.CreatedAt": (ref_models.Point{}).CreatedAt,

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Point.Z": (ref_models.Point{}).Z,

	"ref_models.PointExclusiveSet": &(ref_models.PointExclusiveSet{}),

	"ref_models.PointExclusiveSet.Name": (ref_models.PointExclusiveSet{}).Name,

	"ref_models.PointExclusiveSet.Points": (ref_models.PointExclusiveSet{}).Points,

	"ref_models.PointNonExclusiveSet": &(ref_models.PointNonExclusiveSet{}),

	"ref_models.PointNonExclusiveSet.Name": (ref_models.PointNonExclusiveSet{}).Name,

	"ref_models.PointNonExclusiveSet.Points": (ref_models.PointNonExclusiveSet{}).Points,

	"ref_models.PointUse": &(ref_models.PointUse{}),

	"ref_models.PointUse.Name": (ref_models.PointUse{}).Name,

	"ref_models.PointUse.Points": (ref_models.PointUse{}).Points,

	"ref_models.ShortNodeOnModels": ref_models.ShortNodeOnModels,

	"ref_models.SimulationStage": ref_models.SimulationStage(""),
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_NewDiagram_Line := (&models.Classshape{Name: `NewDiagram-Line`}).Stage()
	__Classshape__000001_NewDiagram_LineType := (&models.Classshape{Name: `NewDiagram-LineType`}).Stage()
	__Classshape__000002_NewDiagram_Point := (&models.Classshape{Name: `NewDiagram-Point`}).Stage()
	__Classshape__000003_NewDiagram_PointExclusiveSet := (&models.Classshape{Name: `NewDiagram-PointExclusiveSet`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CONTINUOUS := (&models.Field{Name: `CONTINUOUS`}).Stage()
	__Field__000001_DOTTED := (&models.Field{Name: `DOTTED`}).Stage()
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage()
	__Position__000001_Pos_NewDiagram_LineType := (&models.Position{Name: `Pos-NewDiagram-LineType`}).Stage()
	__Position__000002_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage()
	__Position__000003_Pos_NewDiagram_PointExclusiveSet := (&models.Position{Name: `Pos-NewDiagram-PointExclusiveSet`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_Line.Name = `NewDiagram-Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Classshape__000000_NewDiagram_Line.Identifier = `ref_models.Line`
	__Classshape__000000_NewDiagram_Line.ShowNbInstances = true
	__Classshape__000000_NewDiagram_Line.NbInstances = 0
	__Classshape__000000_NewDiagram_Line.Width = 240.000000
	__Classshape__000000_NewDiagram_Line.Heigth = 78.000000
	__Classshape__000000_NewDiagram_Line.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_LineType.Name = `NewDiagram-LineType`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineType]
	__Classshape__000001_NewDiagram_LineType.Identifier = `ref_models.LineType`
	__Classshape__000001_NewDiagram_LineType.ShowNbInstances = false
	__Classshape__000001_NewDiagram_LineType.NbInstances = 0
	__Classshape__000001_NewDiagram_LineType.Width = 240.000000
	__Classshape__000001_NewDiagram_LineType.Heigth = 93.000000
	__Classshape__000001_NewDiagram_LineType.IsSelected = false

	// Classshape values setup
	__Classshape__000002_NewDiagram_Point.Name = `NewDiagram-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000002_NewDiagram_Point.Identifier = `ref_models.Point`
	__Classshape__000002_NewDiagram_Point.ShowNbInstances = true
	__Classshape__000002_NewDiagram_Point.NbInstances = 0
	__Classshape__000002_NewDiagram_Point.Width = 240.000000
	__Classshape__000002_NewDiagram_Point.Heigth = 63.000000
	__Classshape__000002_NewDiagram_Point.IsSelected = false

	// Classshape values setup
	__Classshape__000003_NewDiagram_PointExclusiveSet.Name = `NewDiagram-PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.PointExclusiveSet]
	__Classshape__000003_NewDiagram_PointExclusiveSet.Identifier = `ref_models.PointExclusiveSet`
	__Classshape__000003_NewDiagram_PointExclusiveSet.ShowNbInstances = true
	__Classshape__000003_NewDiagram_PointExclusiveSet.NbInstances = 0
	__Classshape__000003_NewDiagram_PointExclusiveSet.Width = 240.000000
	__Classshape__000003_NewDiagram_PointExclusiveSet.Heigth = 63.000000
	__Classshape__000003_NewDiagram_PointExclusiveSet.IsSelected = false

	// Field values setup
	__Field__000000_CONTINUOUS.Name = `CONTINUOUS`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineType.CONTINUOUS]
	__Field__000000_CONTINUOUS.Identifier = `ref_models.LineType.CONTINUOUS`
	__Field__000000_CONTINUOUS.FieldTypeAsString = ``
	__Field__000000_CONTINUOUS.Structname = ``
	__Field__000000_CONTINUOUS.Fieldtypename = ``

	// Field values setup
	__Field__000001_DOTTED.Name = `DOTTED`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineType.DOTTED]
	__Field__000001_DOTTED.Identifier = `ref_models.LineType.DOTTED`
	__Field__000001_DOTTED.FieldTypeAsString = ``
	__Field__000001_DOTTED.Structname = ``
	__Field__000001_DOTTED.Fieldtypename = ``

	// Field values setup
	__Field__000002_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000002_Name.Identifier = `ref_models.Line.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Line`
	__Field__000002_Name.Fieldtypename = `string`

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 30.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 270.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_Line.X = 70.000000
	__Position__000000_Pos_NewDiagram_Line.Y = 104.000000
	__Position__000000_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_LineType.X = 440.000000
	__Position__000001_Pos_NewDiagram_LineType.Y = 90.000000
	__Position__000001_Pos_NewDiagram_LineType.Name = `Pos-NewDiagram-LineType`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Point.X = 480.000000
	__Position__000002_Pos_NewDiagram_Point.Y = 290.000000
	__Position__000002_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Position values setup
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.X = 110.000000
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.Y = 190.000000
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.Name = `Pos-NewDiagram-PointExclusiveSet`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 635.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 228.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000001_NewDiagram_LineType)
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000000_NewDiagram_Line)
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000002_NewDiagram_Point)
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000003_NewDiagram_PointExclusiveSet)
	__Classdiagram__000000_NewDiagram.Notes = append(__Classdiagram__000000_NewDiagram.Notes, __NoteShape__000000_ShortNodeOnModels)
	__Classshape__000000_NewDiagram_Line.Position = __Position__000000_Pos_NewDiagram_Line
	__Classshape__000000_NewDiagram_Line.Fields = append(__Classshape__000000_NewDiagram_Line.Fields, __Field__000002_Name)
	__Classshape__000000_NewDiagram_Line.Links = append(__Classshape__000000_NewDiagram_Line.Links, __Link__000000_End)
	__Classshape__000001_NewDiagram_LineType.Position = __Position__000001_Pos_NewDiagram_LineType
	__Classshape__000001_NewDiagram_LineType.Fields = append(__Classshape__000001_NewDiagram_LineType.Fields, __Field__000000_CONTINUOUS)
	__Classshape__000001_NewDiagram_LineType.Fields = append(__Classshape__000001_NewDiagram_LineType.Fields, __Field__000001_DOTTED)
	__Classshape__000002_NewDiagram_Point.Position = __Position__000002_Pos_NewDiagram_Point
	__Classshape__000003_NewDiagram_PointExclusiveSet.Position = __Position__000003_Pos_NewDiagram_PointExclusiveSet
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
}


