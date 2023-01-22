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

	"ref_models.PointUse.Points": (ref_models.PointUse{}).Points,

	"ref_models.Line.CreationDate": (ref_models.Line{}).CreationDate,

	"ref_models.Line.VeryLongLongLongLongLongLongField": (ref_models.Line{}).VeryLongLongLongLongLongLongField,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Line.End": (ref_models.Line{}).End,

	"ref_models.PointNonExclusiveSet": &(ref_models.PointNonExclusiveSet{}),

	"ref_models.PointNonExclusiveSet.Points": (ref_models.PointNonExclusiveSet{}).Points,

	"ref_models.Line.Start": (ref_models.Line{}).Start,

	"ref_models.Line.JourneyTime": (ref_models.Line{}).JourneyTime,

	"ref_models.PointUse.Name": (ref_models.PointUse{}).Name,

	"ref_models.PointNonExclusiveSet.Name": (ref_models.PointNonExclusiveSet{}).Name,

	"ref_models.Point.CreatedAt": (ref_models.Point{}).CreatedAt,

	"ref_models.PointExclusiveSet": &(ref_models.PointExclusiveSet{}),

	"ref_models.PointExclusiveSet.Name": (ref_models.PointExclusiveSet{}).Name,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Type": (ref_models.Line{}).Type,

	"ref_models.Point.Z": (ref_models.Point{}).Z,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.PointExclusiveSet.Points": (ref_models.PointExclusiveSet{}).Points,

	"ref_models.PointUse": &(ref_models.PointUse{}),
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
	__Classshape__000001_NewDiagram_Point := (&models.Classshape{Name: `NewDiagram-Point`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreatedAt := (&models.Field{Name: `CreatedAt`}).Stage()
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000002_Type := (&models.Field{Name: `Type`}).Stage()
	__Field__000003_VeryLongLongLongLongLongLongField := (&models.Field{Name: `VeryLongLongLongLongLongLongField`}).Stage()
	__Field__000004_X := (&models.Field{Name: `X`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()
	__Link__000001_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage()
	__Position__000001_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`}).Stage()
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`}).Stage()

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
	__Classshape__000000_NewDiagram_Line.Heigth = 93.000000
	__Classshape__000000_NewDiagram_Line.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_Point.Name = `NewDiagram-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000001_NewDiagram_Point.Identifier = `ref_models.Point`
	__Classshape__000001_NewDiagram_Point.ShowNbInstances = true
	__Classshape__000001_NewDiagram_Point.NbInstances = 0
	__Classshape__000001_NewDiagram_Point.Width = 240.000000
	__Classshape__000001_NewDiagram_Point.Heigth = 108.000000
	__Classshape__000001_NewDiagram_Point.IsSelected = false

	// Field values setup
	__Field__000000_CreatedAt.Name = `CreatedAt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.CreatedAt]
	__Field__000000_CreatedAt.Identifier = `ref_models.Point.CreatedAt`
	__Field__000000_CreatedAt.FieldTypeAsString = ``
	__Field__000000_CreatedAt.Structname = `Point`
	__Field__000000_CreatedAt.Fieldtypename = `Time`

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
	__Field__000002_Type.Fieldtypename = `LineType`

	// Field values setup
	__Field__000003_VeryLongLongLongLongLongLongField.Name = `VeryLongLongLongLongLongLongField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.VeryLongLongLongLongLongLongField]
	__Field__000003_VeryLongLongLongLongLongLongField.Identifier = `ref_models.Line.VeryLongLongLongLongLongLongField`
	__Field__000003_VeryLongLongLongLongLongLongField.FieldTypeAsString = ``
	__Field__000003_VeryLongLongLongLongLongLongField.Structname = `Line`
	__Field__000003_VeryLongLongLongLongLongLongField.Fieldtypename = `string`

	// Field values setup
	__Field__000004_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.X]
	__Field__000004_X.Identifier = `ref_models.Point.X`
	__Field__000004_X.FieldTypeAsString = ``
	__Field__000004_X.Structname = `Point`
	__Field__000004_X.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Structname = `Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Start.Name = `Start`
	__Link__000001_Start.Structname = `Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Start]
	__Link__000001_Start.Identifier = `ref_models.Line.Start`
	__Link__000001_Start.Fieldtypename = `Point`
	__Link__000001_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Start.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_Line.X = 310.000000
	__Position__000000_Pos_NewDiagram_Line.Y = 230.000000
	__Position__000000_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Point.X = 70.000000
	__Position__000001_Pos_NewDiagram_Point.Y = 104.000000
	__Position__000001_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 550.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 198.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 550.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 198.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000001_NewDiagram_Point)
	__Classdiagram__000000_NewDiagram.Classshapes = append(__Classdiagram__000000_NewDiagram.Classshapes, __Classshape__000000_NewDiagram_Line)
	__Classshape__000000_NewDiagram_Line.Position = __Position__000000_Pos_NewDiagram_Line
	__Classshape__000000_NewDiagram_Line.Fields = append(__Classshape__000000_NewDiagram_Line.Fields, __Field__000002_Type)
	__Classshape__000000_NewDiagram_Line.Fields = append(__Classshape__000000_NewDiagram_Line.Fields, __Field__000003_VeryLongLongLongLongLongLongField)
	__Classshape__000000_NewDiagram_Line.Links = append(__Classshape__000000_NewDiagram_Line.Links, __Link__000000_End)
	__Classshape__000000_NewDiagram_Line.Links = append(__Classshape__000000_NewDiagram_Line.Links, __Link__000001_Start)
	__Classshape__000001_NewDiagram_Point.Position = __Position__000001_Pos_NewDiagram_Point
	__Classshape__000001_NewDiagram_Point.Fields = append(__Classshape__000001_NewDiagram_Point.Fields, __Field__000001_Name)
	__Classshape__000001_NewDiagram_Point.Fields = append(__Classshape__000001_NewDiagram_Point.Fields, __Field__000004_X)
	__Classshape__000001_NewDiagram_Point.Fields = append(__Classshape__000001_NewDiagram_Point.Fields, __Field__000000_CreatedAt)
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
	__Link__000001_Start.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
}
