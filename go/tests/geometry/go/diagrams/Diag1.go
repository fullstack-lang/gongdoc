package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Diag1 models.StageStruct
var ___dummy__Time_Diag1 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Diag1 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Diag1 map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Diag1"] = Diag1Injection
// }

// Diag1Injection will stage objects of database "Diag1"
func Diag1Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Diag1 := (&models.Classdiagram{Name: `Diag1`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_NewDiagram_Line := (&models.Classshape{Name: `NewDiagram-Line`}).Stage()
	__Classshape__000001_NewDiagram_Point := (&models.Classshape{Name: `NewDiagram-Point`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()

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
	__Vertice__000000_Verticle_in_class_diagram_Diag1_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram Diag1 in middle between NewDiagram-Line and NewDiagram-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diag1.Name = `Diag1`
	__Classdiagram__000000_Diag1.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_Line.Name = `NewDiagram-Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Classshape__000000_NewDiagram_Line.Identifier = `ref_models.Line`
	__Classshape__000000_NewDiagram_Line.ShowNbInstances = true
	__Classshape__000000_NewDiagram_Line.NbInstances = 0
	__Classshape__000000_NewDiagram_Line.Width = 240.000000
	__Classshape__000000_NewDiagram_Line.Heigth = 78.000000
	__Classshape__000000_NewDiagram_Line.IsSelected = false

	// Classshape values setup
	__Classshape__000001_NewDiagram_Point.Name = `NewDiagram-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Classshape__000001_NewDiagram_Point.Identifier = `ref_models.Point`
	__Classshape__000001_NewDiagram_Point.ShowNbInstances = true
	__Classshape__000001_NewDiagram_Point.NbInstances = 0
	__Classshape__000001_NewDiagram_Point.Width = 240.000000
	__Classshape__000001_NewDiagram_Point.Heigth = 63.000000
	__Classshape__000001_NewDiagram_Point.IsSelected = false

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`
	__Field__000000_CreationDate.Fieldname = `CreationDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CreationDate]
	__Field__000000_CreationDate.Identifier = `ref_models.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Fieldname = `End`
	__Link__000000_End.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.End]
	__Link__000000_End.Identifier = `ref_models.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_Line.X = 60.000000
	__Position__000000_Pos_NewDiagram_Line.Y = 60.000000
	__Position__000000_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Point.X = 230.000000
	__Position__000001_Pos_NewDiagram_Point.Y = 230.000000
	__Position__000001_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Diag1_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 505.000000
	__Vertice__000000_Verticle_in_class_diagram_Diag1_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 184.000000
	__Vertice__000000_Verticle_in_class_diagram_Diag1_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram Diag1 in middle between NewDiagram-Line and NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_Diag1.Classshapes = append(__Classdiagram__000000_Diag1.Classshapes, __Classshape__000000_NewDiagram_Line)
	__Classdiagram__000000_Diag1.Classshapes = append(__Classdiagram__000000_Diag1.Classshapes, __Classshape__000001_NewDiagram_Point)
	__Classshape__000000_NewDiagram_Line.Position = __Position__000000_Pos_NewDiagram_Line
	__Classshape__000000_NewDiagram_Line.Fields = append(__Classshape__000000_NewDiagram_Line.Fields, __Field__000000_CreationDate)
	__Classshape__000000_NewDiagram_Line.Links = append(__Classshape__000000_NewDiagram_Line.Links, __Link__000000_End)
	__Classshape__000001_NewDiagram_Point.Position = __Position__000001_Pos_NewDiagram_Point
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Diag1_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
}


