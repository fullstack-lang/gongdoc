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

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Point := (&models.Reference{Name: `Point`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0001.Name = `Classshape0001`
	__Classshape__000000_Classshape0001.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000000_Classshape0001.Identifier = `Ref_models.Point`
	__Classshape__000000_Classshape0001.ShowNbInstances = true
	__Classshape__000000_Classshape0001.NbInstances = 0
	__Classshape__000000_Classshape0001.Width = 240.000000
	__Classshape__000000_Classshape0001.Heigth = 63.000000
	__Classshape__000000_Classshape0001.IsSelected = false

	// Position values setup
	__Position__000000_Position_0001.X = 140.000000
	__Position__000000_Position_0001.Y = 50.000000
	__Position__000000_Position_0001.Name = `Position-0001`

	// Reference values setup
	__Reference__000000_Point.Name = `Point`
	__Reference__000000_Point.NbInstances = 0
	__Reference__000000_Point.Type = models.REFERENCE_GONG_STRUCT

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Classshape0001)
	__Classshape__000000_Classshape0001.Position = __Position__000000_Position_0001
	__Classshape__000000_Classshape0001.Reference = __Reference__000000_Point
}


