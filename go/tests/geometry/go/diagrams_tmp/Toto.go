package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Toto models.StageStruct
var ___dummy__Time_Toto time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Toto ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Toto map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Toto"] = TotoInjection
// }

// TotoInjection will stage objects of database "Toto"
func TotoInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Toto := (&models.Classdiagram{Name: `Toto`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_NewDiagram_Point := (&models.Classshape{Name: `NewDiagram-Point`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Point := (&models.Reference{Name: `Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Toto.Name = `Toto`
	__Classdiagram__000000_Toto.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_NewDiagram_Point.Name = `NewDiagram-Point`
	__Classshape__000000_NewDiagram_Point.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000000_NewDiagram_Point.Identifier = `Ref_models.Point`
	__Classshape__000000_NewDiagram_Point.ShowNbInstances = true
	__Classshape__000000_NewDiagram_Point.NbInstances = 47
	__Classshape__000000_NewDiagram_Point.Width = 240.000000
	__Classshape__000000_NewDiagram_Point.Heigth = 63.000000
	__Classshape__000000_NewDiagram_Point.IsSelected = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_Point.X = 40.000000
	__Position__000000_Pos_NewDiagram_Point.Y = 61.000000
	__Position__000000_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Reference values setup
	__Reference__000000_Point.Name = `Point`
	__Reference__000000_Point.NbInstances = 47
	__Reference__000000_Point.Type = models.REFERENCE_GONG_STRUCT

	// Setup of pointers
	__Classdiagram__000000_Toto.Classshapes = append(__Classdiagram__000000_Toto.Classshapes, __Classshape__000000_NewDiagram_Point)
	__Classshape__000000_NewDiagram_Point.Position = __Position__000000_Pos_NewDiagram_Point
	__Classshape__000000_NewDiagram_Point.Reference = __Reference__000000_Point
}


