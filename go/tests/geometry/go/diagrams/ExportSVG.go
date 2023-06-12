package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_ExportSVG models.StageStruct
var ___dummy__Time_ExportSVG time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_ExportSVG ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_ExportSVG map[string]any = map[string]any{
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
// 	InjectionGateway["ExportSVG"] = ExportSVGInjection
// }

// ExportSVGInjection will stage objects of database "ExportSVG"
func ExportSVGInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_ExportSVG := (&models.Classdiagram{Name: `ExportSVG`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_ExportSVG_Line := (&models.GongStructShape{Name: `ExportSVG-Line`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_ExportSVG_Line := (&models.Position{Name: `Pos-ExportSVG-Line`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_ExportSVG.Name = `ExportSVG`
	__Classdiagram__000000_ExportSVG.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_ExportSVG_Line.Name = `ExportSVG-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_ExportSVG_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_ExportSVG_Line.ShowNbInstances = false
	__GongStructShape__000000_ExportSVG_Line.NbInstances = 0
	__GongStructShape__000000_ExportSVG_Line.Width = 264.000000
	__GongStructShape__000000_ExportSVG_Line.Heigth = 127.000000
	__GongStructShape__000000_ExportSVG_Line.IsSelected = false

	// Position values setup
	__Position__000000_Pos_ExportSVG_Line.X = 288.000000
	__Position__000000_Pos_ExportSVG_Line.Y = 200.000000
	__Position__000000_Pos_ExportSVG_Line.Name = `Pos-ExportSVG-Line`

	// Setup of pointers
	__Classdiagram__000000_ExportSVG.GongStructShapes = append(__Classdiagram__000000_ExportSVG.GongStructShapes, __GongStructShape__000000_ExportSVG_Line)
	__GongStructShape__000000_ExportSVG_Line.Position = __Position__000000_Pos_ExportSVG_Line
}


