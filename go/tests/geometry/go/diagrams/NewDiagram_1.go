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

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Start": (ref_models.Line{}).Start,

	"ref_models.Line.Type": (ref_models.Line{}).Type,

	"ref_models.LineTypeInt": ref_models.LineTypeInt(0),

	"ref_models.LineTypeString": ref_models.LineTypeString(""),

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.ShortNodeOnModels": ref_models.ShortNodeOnModels,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram_1"] = NewDiagram_1Injection
// }

// NewDiagram_1Injection will stage objects of database "NewDiagram_1"
func NewDiagram_1Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_1 := (&models.Classdiagram{Name: `NewDiagram_1`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000002_Type := (&models.Field{Name: `Type`}).Stage()

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt := (&models.GongEnumShape{Name: `NewDiagram_1-LineTypeInt`}).Stage()
	__GongEnumShape__000001_NewDiagram_1_LineTypeString := (&models.GongEnumShape{Name: `NewDiagram_1-LineTypeString`}).Stage()

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_CONTINUOUS := (&models.GongEnumValueEntry{Name: `CONTINUOUS`}).Stage()
	__GongEnumValueEntry__000001_CONTINUOUS_ZERO := (&models.GongEnumValueEntry{Name: `CONTINUOUS_ZERO`}).Stage()
	__GongEnumValueEntry__000002_DOTTED := (&models.GongEnumValueEntry{Name: `DOTTED`}).Stage()
	__GongEnumValueEntry__000003_DOTTED_ONE := (&models.GongEnumValueEntry{Name: `DOTTED_ONE`}).Stage()

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_1_Line := (&models.GongStructShape{Name: `NewDiagram_1-Line`}).Stage()
	__GongStructShape__000001_NewDiagram_1_Point := (&models.GongStructShape{Name: `NewDiagram_1-Point`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Line := (&models.NoteShapeLink{Name: `Line`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Line := (&models.Position{Name: `Pos-NewDiagram_1-Line`}).Stage()
	__Position__000001_Pos_NewDiagram_1_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram_1-LineTypeInt`}).Stage()
	__Position__000002_Pos_NewDiagram_1_LineTypeString := (&models.Position{Name: `Pos-NewDiagram_1-LineTypeString`}).Stage()
	__Position__000003_Pos_NewDiagram_1_Point := (&models.Position{Name: `Pos-NewDiagram_1-Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_1.Name = `NewDiagram_1`
	__Classdiagram__000000_NewDiagram_1.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000000_Name.Identifier = `ref_models.Line.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Line`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.Name]
	__Field__000001_Name.Identifier = `ref_models.Point.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Point`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Type.Name = `Type`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Type]
	__Field__000002_Type.Identifier = `ref_models.Line.Type`
	__Field__000002_Type.FieldTypeAsString = ``
	__Field__000002_Type.Structname = `Line`
	__Field__000002_Type.Fieldtypename = `LineTypeString`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Name = `NewDiagram_1-LineTypeInt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_1_LineTypeInt.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Name = `NewDiagram_1-LineTypeString`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString]
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Identifier = `ref_models.LineTypeString`
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_1_LineTypeString.Heigth = 93.000000

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
	__GongStructShape__000000_NewDiagram_1_Line.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_1_Line.NbInstances = 0
	__GongStructShape__000000_NewDiagram_1_Line.Width = 240.000000
	__GongStructShape__000000_NewDiagram_1_Line.Heigth = 93.000000
	__GongStructShape__000000_NewDiagram_1_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_1_Point.Name = `NewDiagram_1-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_NewDiagram_1_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_NewDiagram_1_Point.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_1_Point.NbInstances = 0
	__GongStructShape__000001_NewDiagram_1_Point.Width = 240.000000
	__GongStructShape__000001_NewDiagram_1_Point.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_1_Point.IsSelected = false

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
	__NoteShape__000000_ShortNodeOnModels.X = 90.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 560.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Line.Name = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__NoteShapeLink__000000_Line.Identifier = `ref_models.Line`
	__NoteShapeLink__000000_Line.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Line.X = 40.000000
	__Position__000000_Pos_NewDiagram_1_Line.Y = 300.000000
	__Position__000000_Pos_NewDiagram_1_Line.Name = `Pos-NewDiagram_1-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.X = 70.000000
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.Y = 104.000000
	__Position__000001_Pos_NewDiagram_1_LineTypeInt.Name = `Pos-NewDiagram_1-LineTypeInt`

	// Position values setup
	__Position__000002_Pos_NewDiagram_1_LineTypeString.X = 430.000000
	__Position__000002_Pos_NewDiagram_1_LineTypeString.Y = 410.000000
	__Position__000002_Pos_NewDiagram_1_LineTypeString.Name = `Pos-NewDiagram_1-LineTypeString`

	// Position values setup
	__Position__000003_Pos_NewDiagram_1_Point.X = 370.000000
	__Position__000003_Pos_NewDiagram_1_Point.Y = 130.000000
	__Position__000003_Pos_NewDiagram_1_Point.Name = `Pos-NewDiagram_1-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.X = 565.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Y = 261.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000000_NewDiagram_1_Line)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000001_NewDiagram_1_Point)
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
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000000_Name)
	__GongStructShape__000000_NewDiagram_1_Line.Fields = append(__GongStructShape__000000_NewDiagram_1_Line.Fields, __Field__000002_Type)
	__GongStructShape__000000_NewDiagram_1_Line.Links = append(__GongStructShape__000000_NewDiagram_1_Line.Links, __Link__000000_Start)
	__GongStructShape__000001_NewDiagram_1_Point.Position = __Position__000003_Pos_NewDiagram_1_Point
	__GongStructShape__000001_NewDiagram_1_Point.Fields = append(__GongStructShape__000001_NewDiagram_1_Point.Fields, __Field__000001_Name)
	__Link__000000_Start.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000000_Line)
}


