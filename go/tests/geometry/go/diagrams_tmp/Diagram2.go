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
	__Classshape__000000_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()
	__Classshape__000001_Diagram2_Line := (&models.Classshape{Name: `Diagram2-Line`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_X := (&models.Field{Name: `X`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_Diagram2_Line := (&models.Position{Name: `Pos-Diagram2-Line`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Line := (&models.Reference{Name: `Line`}).Stage()
	__Reference__000001_Point := (&models.Reference{Name: `Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_Diagram2_Line_and_Classshape0001 := (&models.Vertice{Name: `Verticle in class diagram Diagram2 in middle between Diagram2-Line and Classshape0001`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0001.Name = `Classshape0001`
	__Classshape__000000_Classshape0001.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident []
	__Classshape__000000_Classshape0001.Identifier = ``
	__Classshape__000000_Classshape0001.ShowNbInstances = true
	__Classshape__000000_Classshape0001.NbInstances = 0
	__Classshape__000000_Classshape0001.Width = 240.000000
	__Classshape__000000_Classshape0001.Heigth = 78.000000
	__Classshape__000000_Classshape0001.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Diagram2_Line.Name = `Diagram2-Line`
	__Classshape__000001_Diagram2_Line.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Classshape__000001_Diagram2_Line.Identifier = `ref_models.Line`
	__Classshape__000001_Diagram2_Line.ShowNbInstances = false
	__Classshape__000001_Diagram2_Line.NbInstances = 0
	__Classshape__000001_Diagram2_Line.Width = 240.000000
	__Classshape__000001_Diagram2_Line.Heigth = 63.000000
	__Classshape__000001_Diagram2_Line.IsSelected = false

	// Field values setup
	__Field__000000_X.Name = `X`
	__Field__000000_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.X]
	__Field__000000_X.Identifier = `ref_models.X`
	__Field__000000_X.FieldTypeAsString = ``
	__Field__000000_X.Structname = `Point`
	__Field__000000_X.Fieldtypename = `float64`

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
	__Position__000000_Pos_Diagram2_Line.X = 410.000000
	__Position__000000_Pos_Diagram2_Line.Y = 200.000000
	__Position__000000_Pos_Diagram2_Line.Name = `Pos-Diagram2-Line`

	// Position values setup
	__Position__000001_Position_0001.X = 140.000000
	__Position__000001_Position_0001.Y = 50.000000
	__Position__000001_Position_0001.Name = `Position-0001`

	// Reference values setup
	__Reference__000000_Line.Name = `Line`
	__Reference__000000_Line.NbInstances = 0
	__Reference__000000_Line.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Point.Name = `Point`
	__Reference__000001_Point.NbInstances = 0
	__Reference__000001_Point.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_Diagram2_Line_and_Classshape0001.X = 635.000000
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_Diagram2_Line_and_Classshape0001.Y = 156.500000
	__Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_Diagram2_Line_and_Classshape0001.Name = `Verticle in class diagram Diagram2 in middle between Diagram2-Line and Classshape0001`

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000001_Diagram2_Line)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Classshape0001)
	__Classshape__000000_Classshape0001.Position = __Position__000001_Position_0001
	__Classshape__000000_Classshape0001.Reference = __Reference__000001_Point
	__Classshape__000000_Classshape0001.Fields = append(__Classshape__000000_Classshape0001.Fields, __Field__000000_X)
	__Classshape__000001_Diagram2_Line.Position = __Position__000000_Pos_Diagram2_Line
	__Classshape__000001_Diagram2_Line.Reference = __Reference__000000_Line
	__Classshape__000001_Diagram2_Line.Links = append(__Classshape__000001_Diagram2_Line.Links, __Link__000000_Start)
	__Link__000000_Start.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Diagram2_in_middle_between_Diagram2_Line_and_Classshape0001
}


