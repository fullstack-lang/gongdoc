package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram_1 models.StageStruct
var ___dummy__Time_NewDiagram_1 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram_1 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram_1 map[string]any = map[string]any{
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

	"ref_models.MarkdownNodeOnModels": ref_models.MarkdownNodeOnModels,

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
// 	InjectionGateway["NewDiagram_1"] = NewDiagram_1Injection
// }

// NewDiagram_1Injection will stage objects of database "NewDiagram_1"
func NewDiagram_1Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_1 := (&models.Classdiagram{Name: `NewDiagram_1`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage(stage)
	__Field__000001_JourneyTime := (&models.Field{Name: `JourneyTime`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Type := (&models.Field{Name: `Type`}).Stage(stage)
	__Field__000005_VeryLongLongLongLongLongLongField := (&models.Field{Name: `VeryLongLongLongLongLongLongField`}).Stage(stage)
	__Field__000006_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000007_Y := (&models.Field{Name: `Y`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt := (&models.GongEnumShape{Name: `NewDiagram_1-LineTypeInt`}).Stage(stage)
	__GongEnumShape__000001_NewDiagram_1_LineTypeString := (&models.GongEnumShape{Name: `NewDiagram_1-LineTypeString`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_CONTINUOUS := (&models.GongEnumValueEntry{Name: `CONTINUOUS`}).Stage(stage)
	__GongEnumValueEntry__000001_CONTINUOUS_ZERO := (&models.GongEnumValueEntry{Name: `CONTINUOUS_ZERO`}).Stage(stage)
	__GongEnumValueEntry__000002_DOTTED := (&models.GongEnumValueEntry{Name: `DOTTED`}).Stage(stage)
	__GongEnumValueEntry__000003_DOTTED_ONE := (&models.GongEnumValueEntry{Name: `DOTTED_ONE`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_1_Line := (&models.GongStructShape{Name: `NewDiagram_1-Line`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_1_Point := (&models.GongStructShape{Name: `NewDiagram_1-Point`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000001_Start := (&models.Link{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Line := (&models.NoteShapeLink{Name: `Line`}).Stage(stage)
	__NoteShapeLink__000001_Line_Start := (&models.NoteShapeLink{Name: `Line.Start`}).Stage(stage)
	__NoteShapeLink__000002_LineTypeString := (&models.NoteShapeLink{Name: `LineTypeString`}).Stage(stage)
	__NoteShapeLink__000003_Point := (&models.NoteShapeLink{Name: `Point`}).Stage(stage)

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Line := (&models.Position{Name: `Pos-NewDiagram_1-Line`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_1_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram_1-LineTypeInt`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_1_LineTypeString := (&models.Position{Name: `Pos-NewDiagram_1-LineTypeString`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_1_Point := (&models.Position{Name: `Pos-NewDiagram_1-Point`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_1.Name = `NewDiagram_1`
	__Classdiagram__000000_NewDiagram_1.IsInDrawMode = true

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.CreationDate]
	__Field__000000_CreationDate.Identifier = `ref_models.Line.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Field values setup
	__Field__000001_JourneyTime.Name = `JourneyTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.JourneyTime]
	__Field__000001_JourneyTime.Identifier = `ref_models.Line.JourneyTime`
	__Field__000001_JourneyTime.FieldTypeAsString = ``
	__Field__000001_JourneyTime.Structname = `Line`
	__Field__000001_JourneyTime.Fieldtypename = `Duration`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000002_Name.Identifier = `ref_models.Line.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Line`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.Name]
	__Field__000003_Name.Identifier = `ref_models.Point.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Point`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Type.Name = `Type`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Type]
	__Field__000004_Type.Identifier = `ref_models.Line.Type`
	__Field__000004_Type.FieldTypeAsString = ``
	__Field__000004_Type.Structname = `Line`
	__Field__000004_Type.Fieldtypename = `LineTypeString`

	// Field values setup
	__Field__000005_VeryLongLongLongLongLongLongField.Name = `VeryLongLongLongLongLongLongField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.VeryLongLongLongLongLongLongField]
	__Field__000005_VeryLongLongLongLongLongLongField.Identifier = `ref_models.Line.VeryLongLongLongLongLongLongField`
	__Field__000005_VeryLongLongLongLongLongLongField.FieldTypeAsString = ``
	__Field__000005_VeryLongLongLongLongLongLongField.Structname = `Line`
	__Field__000005_VeryLongLongLongLongLongLongField.Fieldtypename = `string`

	// Field values setup
	__Field__000006_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.X]
	__Field__000006_X.Identifier = `ref_models.Point.X`
	__Field__000006_X.FieldTypeAsString = ``
	__Field__000006_X.Structname = `Point`
	__Field__000006_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000007_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.Y]
	__Field__000007_Y.Identifier = `ref_models.Point.Y`
	__Field__000007_Y.FieldTypeAsString = ``
	__Field__000007_Y.Structname = `Point`
	__Field__000007_Y.Fieldtypename = `float64`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Name = `NewDiagram_1-LineTypeInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Height = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Name = `NewDiagram_1-LineTypeString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString]
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Identifier = `ref_models.LineTypeString`
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Width = 335.000000
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Height = 138.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_CONTINUOUS.Name = `CONTINUOUS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString.CONTINUOUS]
	__GongEnumValueEntry__000000_CONTINUOUS.Identifier = `ref_models.LineTypeString.CONTINUOUS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_CONTINUOUS_ZERO.Name = `CONTINUOUS_ZERO`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.CONTINUOUS_ZERO]
	__GongEnumValueEntry__000001_CONTINUOUS_ZERO.Identifier = `ref_models.LineTypeInt.CONTINUOUS_ZERO`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_DOTTED.Name = `DOTTED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString.DOTTED]
	__GongEnumValueEntry__000002_DOTTED.Identifier = `ref_models.LineTypeString.DOTTED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_DOTTED_ONE.Name = `DOTTED_ONE`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.DOTTED_ONE]
	__GongEnumValueEntry__000003_DOTTED_ONE.Identifier = `ref_models.LineTypeInt.DOTTED_ONE`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_1_Line.Name = `NewDiagram_1-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_NewDiagram_1_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_NewDiagram_1_Line.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_1_Line.NbInstances = 0
	__GongStructShape__000000_NewDiagram_1_Line.Width = 335.000000
	__GongStructShape__000000_NewDiagram_1_Line.Height = 138.000000
	__GongStructShape__000000_NewDiagram_1_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_1_Point.Name = `NewDiagram_1-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_NewDiagram_1_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_NewDiagram_1_Point.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_1_Point.NbInstances = 0
	__GongStructShape__000001_NewDiagram_1_Point.Width = 320.000000
	__GongStructShape__000001_NewDiagram_1_Point.Height = 123.000000
	__GongStructShape__000001_NewDiagram_1_Point.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000000_End.Fieldtypename = `ref_models.Point`
	__Link__000000_End.FieldOffsetX = 20.000000
	__Link__000000_End.FieldOffsetY = -26.000000
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.TargetMultiplicityOffsetX = -35.000000
	__Link__000000_End.TargetMultiplicityOffsetY = -28.000000
	__Link__000000_End.SourceMultiplicity = models.MANY
	__Link__000000_End.SourceMultiplicityOffsetX = 16.000000
	__Link__000000_End.SourceMultiplicityOffsetY = 28.000000
	__Link__000000_End.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_End.StartRatio = 0.686567
	__Link__000000_End.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_End.EndRatio = 0.700906
	__Link__000000_End.CornerOffsetRatio = 1.910256

	// Link values setup
	__Link__000001_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Start]
	__Link__000001_Start.Identifier = `ref_models.Line.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000001_Start.Fieldtypename = `ref_models.Point`
	__Link__000001_Start.FieldOffsetX = -48.000000
	__Link__000001_Start.FieldOffsetY = -25.000000
	__Link__000001_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Start.TargetMultiplicityOffsetX = 17.000000
	__Link__000001_Start.TargetMultiplicityOffsetY = -25.000000
	__Link__000001_Start.SourceMultiplicity = models.MANY
	__Link__000001_Start.SourceMultiplicityOffsetX = 13.000000
	__Link__000001_Start.SourceMultiplicityOffsetY = 22.000000
	__Link__000001_Start.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Start.StartRatio = 0.200000
	__Link__000001_Start.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Start.EndRatio = 0.217523
	__Link__000001_Start.CornerOffsetRatio = 1.256410

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ShortNodeOnModels]
	__NoteShape__000000_ShortNodeOnModels.Identifier = `ref_models.ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [models.Line], [models.Point] and [models.Line.Start] [models.LineTypeString]
are referenced in the go code
`
	__NoteShape__000000_ShortNodeOnModels.BodyHTML = `<p>this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case <a href="/models#Line">models.Line</a>, <a href="/models#Point">models.Point</a> and <a href="/models#Line.Start">models.Line.Start</a> <a href="/models#LineTypeString">models.LineTypeString</a>
are referenced in the go code
`
	__NoteShape__000000_ShortNodeOnModels.X = 496.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 531.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 626.000000
	__NoteShape__000000_ShortNodeOnModels.Height = 111.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Line.Name = `Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__NoteShapeLink__000000_Line.Identifier = `ref_models.Line`
	__NoteShapeLink__000000_Line.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000001_Line_Start.Name = `Line.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Start]
	__NoteShapeLink__000001_Line_Start.Identifier = `ref_models.Line.Start`
	__NoteShapeLink__000001_Line_Start.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	// NoteShapeLink values setup
	__NoteShapeLink__000002_LineTypeString.Name = `LineTypeString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString]
	__NoteShapeLink__000002_LineTypeString.Identifier = `ref_models.LineTypeString`
	__NoteShapeLink__000002_LineTypeString.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000003_Point.Name = `Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__NoteShapeLink__000003_Point.Identifier = `ref_models.Point`
	__NoteShapeLink__000003_Point.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Line.X = 48.000000
	__Position__000000_Pos_NewDiagram_1_Line.Y = 88.000000
	__Position__000000_Pos_NewDiagram_1_Line.Name = `Pos-NewDiagram_1-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.X = 440.000000
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.Y = 103.000000
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.Name = `Pos-NewDiagram_1-LineTypeInt`

	// Position values setup
	__Position__000002_Pos_NewDiagram_1_LineTypeString.X = 732.000000
	__Position__000002_Pos_NewDiagram_1_LineTypeString.Y = 106.000000
	__Position__000002_Pos_NewDiagram_1_LineTypeString.Name = `Pos-NewDiagram_1-LineTypeString`

	// Position values setup
	__Position__000003_Pos_NewDiagram_1_Point.X = 49.000000
	__Position__000003_Pos_NewDiagram_1_Point.Y = 529.000000
	__Position__000003_Pos_NewDiagram_1_Point.Name = `Pos-NewDiagram_1-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.X = 474.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Y = 397.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.X = 680.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Y = 424.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000001_NewDiagram_1_Point)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000000_NewDiagram_1_Line)
	__Classdiagram__000000_NewDiagram_1.GongEnumShapes = append(__Classdiagram__000000_NewDiagram_1.GongEnumShapes, __GongEnumShape__000000_NewDiagram_1_LineTypeInt)
	__Classdiagram__000000_NewDiagram_1.GongEnumShapes = append(__Classdiagram__000000_NewDiagram_1.GongEnumShapes, __GongEnumShape__000001_NewDiagram_1_LineTypeString)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000000_ShortNodeOnModels)
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Position = __Position__000001_Pos_NewDiagram_1_LineTypeInt
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_1_LineTypeInt.GongEnumValueEntrys, __GongEnumValueEntry__000001_CONTINUOUS_ZERO)
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_1_LineTypeInt.GongEnumValueEntrys, __GongEnumValueEntry__000003_DOTTED_ONE)
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Position = __Position__000002_Pos_NewDiagram_1_LineTypeString
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_1_LineTypeString.GongEnumValueEntrys, __GongEnumValueEntry__000000_CONTINUOUS)
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_1_LineTypeString.GongEnumValueEntrys, __GongEnumValueEntry__000002_DOTTED)
	__GongStructShape__000000_NewDiagram_1_Line.Position = __Position__000000_Pos_NewDiagram_1_Line
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000002_Name)
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000000_CreationDate)
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000001_JourneyTime)
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000004_Type)
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000005_VeryLongLongLongLongLongLongField)
	__GongStructShape__000000_NewDiagram_1_Line.Links = append(__GongStructShape__000000_NewDiagram_1_Line.Links, __Link__000000_End)
	__GongStructShape__000000_NewDiagram_1_Line.Links = append(__GongStructShape__000000_NewDiagram_1_Line.Links, __Link__000001_Start)
	__GongStructShape__000001_NewDiagram_1_Point.Position = __Position__000003_Pos_NewDiagram_1_Point
	__GongStructShape__000001_NewDiagram_1_Point.Fields = append(__GongStructShape__000001_NewDiagram_1_Point.Fields, __Field__000003_Name)
	__GongStructShape__000001_NewDiagram_1_Point.Fields = append(__GongStructShape__000001_NewDiagram_1_Point.Fields, __Field__000006_X)
	__GongStructShape__000001_NewDiagram_1_Point.Fields = append(__GongStructShape__000001_NewDiagram_1_Point.Fields, __Field__000007_Y)
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point
	__Link__000001_Start.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000000_Line)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000002_LineTypeString)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000001_Line_Start)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000003_Point)
}
