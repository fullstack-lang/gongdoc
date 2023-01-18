package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Diagram2 models.StageStruct
var ___dummy__Time_Diagram2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Diagram2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Diagram2 map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Diagram2"] = Diagram2Injection
// }

// Diagram2Injection will stage objects of database "Diagram2"
func Diagram2Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Diagram2 := (&models.Classdiagram{Name: `Diagram2`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Diagram2_Point := (&models.Classshape{Name: `Diagram2-Point`}).Stage()
	__Classshape__000001_NewDiagram_Line := (&models.Classshape{Name: `NewDiagram-Line`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000002_X := (&models.Field{Name: `X`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_Diagram2_Point := (&models.Position{Name: `Pos-Diagram2-Point`}).Stage()
	__Position__000001_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_NewDiagram_Line_and_Diagram2_Point := (&models.Vertice{Name: `Verticle in class diagram Diagram2 in middle between NewDiagram-Line and Diagram2-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Diagram2_Point.Name = `Diagram2-Point`
	__Classshape__000000_Diagram2_Point.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000000_Diagram2_Point.Identifier = `ref_models.Point`
	__Classshape__000000_Diagram2_Point.ShowNbInstances = false
	__Classshape__000000_Diagram2_Point.NbInstances = 0
	__Classshape__000000_Diagram2_Point.Width = 240.000000
	__Classshape__000000_Diagram2_Point.Heigth = 93.000000
	__Classshape__000000_Diagram2_Point.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_Line.Name = `NewDiagram-Line`
	__Classshape__000001_NewDiagram_Line.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Classshape__000001_NewDiagram_Line.Identifier = `ref_models.Line`
	__Classshape__000001_NewDiagram_Line.ShowNbInstances = false
	__Classshape__000001_NewDiagram_Line.NbInstances = 0
	__Classshape__000001_NewDiagram_Line.Width = 240.000000
	__Classshape__000001_NewDiagram_Line.Heigth = 78.000000
	__Classshape__000001_NewDiagram_Line.IsSelected = false

	// Field values setup
	__Field__000000_Name.Name = `Name`
	__Field__000000_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Name]
	__Field__000000_Name.Identifier = `ref_models.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Line`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`
	__Field__000001_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Name]
	__Field__000001_Name.Identifier = `ref_models.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Point`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_X.Name = `X`
	__Field__000002_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.X]
	__Field__000002_X.Identifier = `ref_models.X`
	__Field__000002_X.FieldTypeAsString = ``
	__Field__000002_X.Structname = `Point`
	__Field__000002_X.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_Start.Name = `Start`
	__Link__000000_Start.Fieldname = `Start`
	__Link__000000_Start.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Start]
	__Link__000000_Start.Identifier = `ref_models.Start`
	__Link__000000_Start.Fieldtypename = `Point`
	__Link__000000_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Start.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_Diagram2_Point.X = 480.000000
	__Position__000000_Pos_Diagram2_Point.Y = 250.000000
	__Position__000000_Pos_Diagram2_Point.Name = `Pos-Diagram2-Point`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Line.X = 70.000000
	__Position__000001_Pos_NewDiagram_Line.Y = 60.000000
	__Position__000001_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 0
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 0
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_NewDiagram_Line_and_Diagram2_Point.X = 300.000000
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_NewDiagram_Line_and_Diagram2_Point.Y = 241.000000
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_NewDiagram_Line_and_Diagram2_Point.Name = `Verticle in class diagram Diagram2 in middle between NewDiagram-Line and Diagram2-Point`

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000001_NewDiagram_Line)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Diagram2_Point)
	__Classshape__000000_Diagram2_Point.Position = __Position__000000_Pos_Diagram2_Point
	__Classshape__000000_Diagram2_Point.Reference = __Reference__000001_Point
	__Classshape__000000_Diagram2_Point.Fields = append(__Classshape__000000_Diagram2_Point.Fields, __Field__000001_Name)
	__Classshape__000000_Diagram2_Point.Fields = append(__Classshape__000000_Diagram2_Point.Fields, __Field__000002_X)
	__Classshape__000001_NewDiagram_Line.Position = __Position__000001_Pos_NewDiagram_Line
	__Classshape__000001_NewDiagram_Line.Reference = __Reference__000000_Line
	__Classshape__000001_NewDiagram_Line.Fields = append(__Classshape__000001_NewDiagram_Line.Fields, __Field__000000_Name)
	__Classshape__000001_NewDiagram_Line.Links = append(__Classshape__000001_NewDiagram_Line.Links, __Link__000000_Start)
	__Link__000000_Start.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_NewDiagram_Line_and_Diagram2_Point
}


