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

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_LineTypeInt := (&models.GongEnumShape{Name: `NewDiagram-LineTypeInt`}).Stage()

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_CONTINUOUS_ZERO := (&models.GongEnumValueEntry{Name: `CONTINUOUS_ZERO`}).Stage()

	// Declarations of staged instances of GongStructShape

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram-LineTypeInt`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_LineTypeInt.Name = `NewDiagram-LineTypeInt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__GongEnumShape__000000_NewDiagram_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__GongEnumShape__000000_NewDiagram_LineTypeInt.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_LineTypeInt.Heigth = 78.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_CONTINUOUS_ZERO.Name = `CONTINUOUS_ZERO`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.CONTINUOUS_ZERO]
	__GongEnumValueEntry__000000_CONTINUOUS_ZERO.Identifier = `ref_models.LineTypeInt.CONTINUOUS_ZERO`

	// Position values setup
	__Position__000000_Pos_NewDiagram_LineTypeInt.X = 70.000000
	__Position__000000_Pos_NewDiagram_LineTypeInt.Y = 104.000000
	__Position__000000_Pos_NewDiagram_LineTypeInt.Name = `Pos-NewDiagram-LineTypeInt`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_LineTypeInt)
	__GongEnumShape__000000_NewDiagram_LineTypeInt.Position = __Position__000000_Pos_NewDiagram_LineTypeInt
	__GongEnumShape__000000_NewDiagram_LineTypeInt.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_LineTypeInt.GongEnumValueEntrys, __GongEnumValueEntry__000000_CONTINUOUS_ZERO)
}


