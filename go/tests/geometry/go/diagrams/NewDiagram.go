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
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Type := (&models.Field{Name: `Type`}).Stage(stage)
	__Field__000002_VeryLongLongLongLongLongLongField := (&models.Field{Name: `VeryLongLongLongLongLongLongField`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_LineTypeString := (&models.GongEnumShape{Name: `NewDiagram-LineTypeString`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Line := (&models.GongStructShape{Name: `NewDiagram-Line`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Point := (&models.GongStructShape{Name: `NewDiagram-Point`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Start := (&models.Link{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Line := (&models.NoteShapeLink{Name: `Line`}).Stage(stage)
	__NoteShapeLink__000001_Line_Start := (&models.NoteShapeLink{Name: `Line.Start`}).Stage(stage)
	__NoteShapeLink__000002_LineTypeString := (&models.NoteShapeLink{Name: `LineTypeString`}).Stage(stage)
	__NoteShapeLink__000003_Point := (&models.NoteShapeLink{Name: `Point`}).Stage(stage)

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_LineTypeString := (&models.Position{Name: `Pos-NewDiagram-LineTypeString`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000000_Name.Identifier = `ref_models.Line.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Line`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Type.Name = `Type`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Type]
	__Field__000001_Type.Identifier = `ref_models.Line.Type`
	__Field__000001_Type.FieldTypeAsString = ``
	__Field__000001_Type.Structname = `Line`
	__Field__000001_Type.Fieldtypename = `LineTypeString`

	// Field values setup
	__Field__000002_VeryLongLongLongLongLongLongField.Name = `VeryLongLongLongLongLongLongField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.VeryLongLongLongLongLongLongField]
	__Field__000002_VeryLongLongLongLongLongLongField.Identifier = `ref_models.Line.VeryLongLongLongLongLongLongField`
	__Field__000002_VeryLongLongLongLongLongLongField.FieldTypeAsString = ``
	__Field__000002_VeryLongLongLongLongLongLongField.Structname = `Line`
	__Field__000002_VeryLongLongLongLongLongLongField.Fieldtypename = `string`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_LineTypeString.Name = `NewDiagram-LineTypeString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString]
	__GongEnumShape__000000_NewDiagram_LineTypeString.Identifier = `ref_models.LineTypeString`
	__GongEnumShape__000000_NewDiagram_LineTypeString.Width = 376.016602
	__GongEnumShape__000000_NewDiagram_LineTypeString.Heigth = 106.000000

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Line.Name = `NewDiagram-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_NewDiagram_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_NewDiagram_Line.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Line.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Line.Width = 316.633301
	__GongStructShape__000000_NewDiagram_Line.Heigth = 108.000000
	__GongStructShape__000000_NewDiagram_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Point.Name = `NewDiagram-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_NewDiagram_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_NewDiagram_Point.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Point.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Point.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Point.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_Point.IsSelected = false

	// Link values setup
	__Link__000000_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Start]
	__Link__000000_Start.Identifier = `ref_models.Line.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000000_Start.Fieldtypename = `ref_models.Point`
	__Link__000000_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Start.SourceMultiplicity = models.MANY

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
	__NoteShape__000000_ShortNodeOnModels.X = 180.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 300.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
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
	__Position__000000_Pos_NewDiagram_Line.X = 115.000000
	__Position__000000_Pos_NewDiagram_Line.Y = 52.000000
	__Position__000000_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_LineTypeString.X = 148.616699
	__Position__000001_Pos_NewDiagram_LineTypeString.Y = 454.000000
	__Position__000001_Pos_NewDiagram_LineTypeString.Name = `Pos-NewDiagram-LineTypeString`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Point.X = 540.000000
	__Position__000002_Pos_NewDiagram_Point.Y = 190.000000
	__Position__000002_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 473.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 100.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Line)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Point)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_LineTypeString)
	__Classdiagram__000000_NewDiagram.NoteShapes = append(__Classdiagram__000000_NewDiagram.NoteShapes, __NoteShape__000000_ShortNodeOnModels)
	__GongEnumShape__000000_NewDiagram_LineTypeString.Position = __Position__000001_Pos_NewDiagram_LineTypeString
	__GongStructShape__000000_NewDiagram_Line.Position = __Position__000000_Pos_NewDiagram_Line
	__GongStructShape__000000_NewDiagram_Line.Fields = append(__GongStructShape__000000_NewDiagram_Line.Fields, __Field__000000_Name)
	__GongStructShape__000000_NewDiagram_Line.Fields = append(__GongStructShape__000000_NewDiagram_Line.Fields, __Field__000001_Type)
	__GongStructShape__000000_NewDiagram_Line.Fields = append(__GongStructShape__000000_NewDiagram_Line.Fields, __Field__000002_VeryLongLongLongLongLongLongField)
	__GongStructShape__000000_NewDiagram_Line.Links = append(__GongStructShape__000000_NewDiagram_Line.Links, __Link__000000_Start)
	__GongStructShape__000001_NewDiagram_Point.Position = __Position__000002_Pos_NewDiagram_Point
	__Link__000000_Start.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000000_Line)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000001_Line_Start)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000003_Point)
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000002_LineTypeString)
}


