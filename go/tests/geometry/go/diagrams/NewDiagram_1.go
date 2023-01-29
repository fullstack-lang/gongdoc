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
func NewDiagram_1Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_1 := (&models.Classdiagram{Name: `NewDiagram_1`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_NewDiagram_1_Line := (&models.Classshape{Name: `NewDiagram_1-Line`}).Stage()
	__Classshape__000001_NewDiagram_1_Point := (&models.Classshape{Name: `NewDiagram_1-Point`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Line := (&models.NoteShapeLink{Name: `Line`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Line := (&models.Position{Name: `Pos-NewDiagram_1-Line`}).Stage()
	__Position__000001_Pos_NewDiagram_1_Point := (&models.Position{Name: `Pos-NewDiagram_1-Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_1.Name = `NewDiagram_1`
	__Classdiagram__000000_NewDiagram_1.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_1_Line.Name = `NewDiagram_1-Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Classshape__000000_NewDiagram_1_Line.Identifier = `ref_models.Line`
	__Classshape__000000_NewDiagram_1_Line.ShowNbInstances = false
	__Classshape__000000_NewDiagram_1_Line.NbInstances = 0
	__Classshape__000000_NewDiagram_1_Line.Width = 240.000000
	__Classshape__000000_NewDiagram_1_Line.Heigth = 78.000000
	__Classshape__000000_NewDiagram_1_Line.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_1_Point.Name = `NewDiagram_1-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000001_NewDiagram_1_Point.Identifier = `ref_models.Point`
	__Classshape__000001_NewDiagram_1_Point.ShowNbInstances = false
	__Classshape__000001_NewDiagram_1_Point.NbInstances = 0
	__Classshape__000001_NewDiagram_1_Point.Width = 240.000000
	__Classshape__000001_NewDiagram_1_Point.Heigth = 63.000000
	__Classshape__000001_NewDiagram_1_Point.IsSelected = false

	// Field values setup
	__Field__000000_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000000_Name.Identifier = `ref_models.Line.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Line`
	__Field__000000_Name.Fieldtypename = `string`

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
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ShortNodeOnModels]
	__NoteShape__000000_ShortNodeOnModels.Identifier = `ref_models.ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 200.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 280.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Line.Name = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ShortNodeOnModels.Line]
	__NoteShapeLink__000000_Line.Identifier = `ref_models.ShortNodeOnModels.Line`

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Line.X = 480.000000
	__Position__000000_Pos_NewDiagram_1_Line.Y = 160.000000
	__Position__000000_Pos_NewDiagram_1_Line.Name = `Pos-NewDiagram_1-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_Point.X = 20.000000
	__Position__000001_Pos_NewDiagram_1_Point.Y = 110.000000
	__Position__000001_Pos_NewDiagram_1_Point.Name = `Pos-NewDiagram_1-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.X = 433.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Y = 110.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Line and NewDiagram_1-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_1.Classshapes = append(__Classdiagram__000000_NewDiagram_1.Classshapes, __Classshape__000000_NewDiagram_1_Line)
	__Classdiagram__000000_NewDiagram_1.Classshapes = append(__Classdiagram__000000_NewDiagram_1.Classshapes, __Classshape__000001_NewDiagram_1_Point)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000000_ShortNodeOnModels)
	__Classshape__000000_NewDiagram_1_Line.Position = __Position__000000_Pos_NewDiagram_1_Line
	__Classshape__000000_NewDiagram_1_Line.Fields = append(__Classshape__000000_NewDiagram_1_Line.Fields, __Field__000000_Name)
	__Classshape__000000_NewDiagram_1_Line.Links = append(__Classshape__000000_NewDiagram_1_Line.Links, __Link__000000_End)
	__Classshape__000001_NewDiagram_1_Point.Position = __Position__000001_Pos_NewDiagram_1_Point
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Line_and_NewDiagram_1_Point
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000000_Line)
}


