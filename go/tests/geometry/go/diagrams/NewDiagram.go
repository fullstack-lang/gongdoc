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
	__Classshape__000000_NewDiagram_LineTypeInt := (&models.Classshape{Name: `NewDiagram-LineTypeInt`}).Stage()
	__Classshape__000001_NewDiagram_LineTypeString := (&models.Classshape{Name: `NewDiagram-LineTypeString`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CONTINUOUS_ZERO := (&models.Field{Name: `CONTINUOUS_ZERO`}).Stage()
	__Field__000001_DOTTED := (&models.Field{Name: `DOTTED`}).Stage()
	__Field__000002_DOTTED_ONE := (&models.Field{Name: `DOTTED_ONE`}).Stage()

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram-LineTypeInt`}).Stage()
	__Position__000001_Pos_NewDiagram_LineTypeString := (&models.Position{Name: `Pos-NewDiagram-LineTypeString`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_LineTypeInt.Name = `NewDiagram-LineTypeInt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__Classshape__000000_NewDiagram_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__Classshape__000000_NewDiagram_LineTypeInt.ShowNbInstances = false
	__Classshape__000000_NewDiagram_LineTypeInt.NbInstances = 0
	__Classshape__000000_NewDiagram_LineTypeInt.Width = 240.000000
	__Classshape__000000_NewDiagram_LineTypeInt.Heigth = 93.000000
	__Classshape__000000_NewDiagram_LineTypeInt.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_LineTypeString.Name = `NewDiagram-LineTypeString`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString]
	__Classshape__000001_NewDiagram_LineTypeString.Identifier = `ref_models.LineTypeString`
	__Classshape__000001_NewDiagram_LineTypeString.ShowNbInstances = false
	__Classshape__000001_NewDiagram_LineTypeString.NbInstances = 0
	__Classshape__000001_NewDiagram_LineTypeString.Width = 240.000000
	__Classshape__000001_NewDiagram_LineTypeString.Heigth = 78.000000
	__Classshape__000001_NewDiagram_LineTypeString.IsSelected = false

	// Field values setup
	__Field__000000_CONTINUOUS_ZERO.Name = `CONTINUOUS_ZERO`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.CONTINUOUS_ZERO]
	__Field__000000_CONTINUOUS_ZERO.Identifier = `ref_models.LineTypeInt.CONTINUOUS_ZERO`
	__Field__000000_CONTINUOUS_ZERO.FieldTypeAsString = ``
	__Field__000000_CONTINUOUS_ZERO.Structname = ``
	__Field__000000_CONTINUOUS_ZERO.Fieldtypename = ``

	// Field values setup
	__Field__000001_DOTTED.Name = `DOTTED`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeString.DOTTED]
	__Field__000001_DOTTED.Identifier = `ref_models.LineTypeString.DOTTED`
	__Field__000001_DOTTED.FieldTypeAsString = ``
	__Field__000001_DOTTED.Structname = ``
	__Field__000001_DOTTED.Fieldtypename = ``

	// Field values setup
	__Field__000002_DOTTED_ONE.Name = `DOTTED_ONE`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.DOTTED_ONE]
	__Field__000002_DOTTED_ONE.Identifier = `ref_models.LineTypeInt.DOTTED_ONE`
	__Field__000002_DOTTED_ONE.FieldTypeAsString = ``
	__Field__000002_DOTTED_ONE.Structname = ``
	__Field__000002_DOTTED_ONE.Fieldtypename = ``

	// Position values setup
	__Position__000000_Pos_NewDiagram_LineTypeInt.X = 110.000000
	__Position__000000_Pos_NewDiagram_LineTypeInt.Y = 210.000000
	__Position__000000_Pos_NewDiagram_LineTypeInt.Name = `Pos-NewDiagram-LineTypeInt`

	// Position values setup
	__Position__000001_Pos_NewDiagram_LineTypeString.X = 140.000000
	__Position__000001_Pos_NewDiagram_LineTypeString.Y = 340.000000
	__Position__000001_Pos_NewDiagram_LineTypeString.Name = `Pos-NewDiagram-LineTypeString`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000000_NewDiagram_LineTypeInt)
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000001_NewDiagram_LineTypeString)
	__Classshape__000000_NewDiagram_LineTypeInt.Position = __Position__000000_Pos_NewDiagram_LineTypeInt
	__Classshape__000000_NewDiagram_LineTypeInt.Fields = append(__Classshape__000000_NewDiagram_LineTypeInt.Fields, __Field__000000_CONTINUOUS_ZERO)
	__Classshape__000000_NewDiagram_LineTypeInt.Fields = append(__Classshape__000000_NewDiagram_LineTypeInt.Fields, __Field__000002_DOTTED_ONE)
	__Classshape__000001_NewDiagram_LineTypeString.Position = __Position__000001_Pos_NewDiagram_LineTypeString
	__Classshape__000001_NewDiagram_LineTypeString.Fields = append(__Classshape__000001_NewDiagram_LineTypeString.Fields, __Field__000001_DOTTED)
}


