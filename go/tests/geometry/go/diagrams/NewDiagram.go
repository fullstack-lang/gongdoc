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

	"ref_models.Line.Type": (ref_models.Line{}).Type,

	"ref_models.Point.CreatedAt": (ref_models.Point{}).CreatedAt,

	"ref_models.PointUse.Points": (ref_models.PointUse{}).Points,

	"ref_models.Line.CreationDate": (ref_models.Line{}).CreationDate,

	"ref_models.PointExclusiveSet.Name": (ref_models.PointExclusiveSet{}).Name,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Line.VeryLongLongLongLongLongLongField": (ref_models.Line{}).VeryLongLongLongLongLongLongField,

	"ref_models.Point.Z": (ref_models.Point{}).Z,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.PointNonExclusiveSet": &(ref_models.PointNonExclusiveSet{}),

	"ref_models.PointNonExclusiveSet.Points": (ref_models.PointNonExclusiveSet{}).Points,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.End": (ref_models.Line{}).End,

	"ref_models.PointUse": &(ref_models.PointUse{}),

	"ref_models.PointUse.Name": (ref_models.PointUse{}).Name,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Start": (ref_models.Line{}).Start,

	"ref_models.Line.JourneyTime": (ref_models.Line{}).JourneyTime,

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.PointExclusiveSet": &(ref_models.PointExclusiveSet{}),

	"ref_models.PointExclusiveSet.Points": (ref_models.PointExclusiveSet{}).Points,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.PointNonExclusiveSet.Name": (ref_models.PointNonExclusiveSet{}).Name,
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
	__Classshape__000000_NewDiagram_Point := (&models.Classshape{Name: `NewDiagram-Point`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_X := (&models.Field{Name: `X`}).Stage()

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_Point.Name = `NewDiagram-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000000_NewDiagram_Point.Identifier = `ref_models.Point`
	__Classshape__000000_NewDiagram_Point.ShowNbInstances = false
	__Classshape__000000_NewDiagram_Point.NbInstances = 0
	__Classshape__000000_NewDiagram_Point.Width = 240.000000
	__Classshape__000000_NewDiagram_Point.Heigth = 78.000000
	__Classshape__000000_NewDiagram_Point.IsSelected = false

	// Field values setup
	__Field__000000_X.Name = `X`
	__Field__000000_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.X]
	__Field__000000_X.Identifier = `ref_models.Point.X`
	__Field__000000_X.FieldTypeAsString = ``
	__Field__000000_X.Structname = `Point`
	__Field__000000_X.Fieldtypename = `float64`

	// Position values setup
	__Position__000000_Pos_NewDiagram_Point.X = 76.000000
	__Position__000000_Pos_NewDiagram_Point.Y = 53.000000
	__Position__000000_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000000_NewDiagram_Point)
	__Classshape__000000_NewDiagram_Point.Position = __Position__000000_Pos_NewDiagram_Point
	__Classshape__000000_NewDiagram_Point.Fields = append(__Classshape__000000_NewDiagram_Point.Fields, __Field__000000_X)
}


