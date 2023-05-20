package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram_2 models.StageStruct
var ___dummy__Time_NewDiagram_2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram_2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram_2 map[string]any = map[string]any{
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
// 	InjectionGateway["NewDiagram_2"] = NewDiagram_2Injection
// }

// NewDiagram_2Injection will stage objects of database "NewDiagram_2"
func NewDiagram_2Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_2 := (&models.Classdiagram{Name: `NewDiagram_2`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt := (&models.GongEnumShape{Name: `NewDiagram_2-LineTypeInt`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_2_Line := (&models.GongStructShape{Name: `NewDiagram_2-Line`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_2_Point := (&models.GongStructShape{Name: `NewDiagram_2-Point`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_2_Line := (&models.Position{Name: `Pos-NewDiagram_2-Line`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_2_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram_2-LineTypeInt`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_2_Point := (&models.Position{Name: `Pos-NewDiagram_2-Point`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_2.Name = `NewDiagram_2`
	__Classdiagram__000000_NewDiagram_2.IsInDrawMode = true

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt.Name = `NewDiagram_2-LineTypeInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt.Heigth = 63.000000

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_2_Line.Name = `NewDiagram_2-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_NewDiagram_2_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_NewDiagram_2_Line.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_2_Line.NbInstances = 0
	__GongStructShape__000000_NewDiagram_2_Line.Width = 240.000000
	__GongStructShape__000000_NewDiagram_2_Line.Heigth = 63.000000
	__GongStructShape__000000_NewDiagram_2_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_2_Point.Name = `NewDiagram_2-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_NewDiagram_2_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_NewDiagram_2_Point.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_2_Point.NbInstances = 0
	__GongStructShape__000001_NewDiagram_2_Point.Width = 240.000000
	__GongStructShape__000001_NewDiagram_2_Point.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_2_Point.IsSelected = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_2_Line.X = 71.000000
	__Position__000000_Pos_NewDiagram_2_Line.Y = 25.000000
	__Position__000000_Pos_NewDiagram_2_Line.Name = `Pos-NewDiagram_2-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_2_LineTypeInt.X = 93.000000
	__Position__000001_Pos_NewDiagram_2_LineTypeInt.Y = 71.000000
	__Position__000001_Pos_NewDiagram_2_LineTypeInt.Name = `Pos-NewDiagram_2-LineTypeInt`

	// Position values setup
	__Position__000002_Pos_NewDiagram_2_Point.X = 67.000000
	__Position__000002_Pos_NewDiagram_2_Point.Y = 106.000000
	__Position__000002_Pos_NewDiagram_2_Point.Name = `Pos-NewDiagram_2-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000000_NewDiagram_2_Line)
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000001_NewDiagram_2_Point)
	__Classdiagram__000000_NewDiagram_2.GongEnumShapes = append(__Classdiagram__000000_NewDiagram_2.GongEnumShapes, __GongEnumShape__000000_NewDiagram_2_LineTypeInt)
	__GongEnumShape__000000_NewDiagram_2_LineTypeInt.Position = __Position__000001_Pos_NewDiagram_2_LineTypeInt
	__GongStructShape__000000_NewDiagram_2_Line.Position = __Position__000000_Pos_NewDiagram_2_Line
	__GongStructShape__000001_NewDiagram_2_Point.Position = __Position__000002_Pos_NewDiagram_2_Point
}


