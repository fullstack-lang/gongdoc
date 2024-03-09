package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default_2 models.StageStruct
var ___dummy__Time_Default_2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default_2 ref_models.StageStruct

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default_2"] = Default_2Injection
// }

// Default_2Injection will stage objects of database "Default_2"
func Default_2Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default_2 := (&models.Classdiagram{Name: `Default_2`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_2_Classdiagram := (&models.GongStructShape{Name: `Default_2-Classdiagram`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_2_Classdiagram := (&models.Position{Name: `Pos-Default_2-Classdiagram`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default_2.Name = `Default_2`
	__Classdiagram__000000_Default_2.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_Default_2_Classdiagram.Name = `Default_2-Classdiagram`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__GongStructShape__000000_Default_2_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_2_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_2_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_2_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_2_Classdiagram.Height = 63.000000
	__GongStructShape__000000_Default_2_Classdiagram.IsSelected = false

	// Position values setup
	__Position__000000_Pos_Default_2_Classdiagram.X = 78.000000
	__Position__000000_Pos_Default_2_Classdiagram.Y = 27.000000
	__Position__000000_Pos_Default_2_Classdiagram.Name = `Pos-Default_2-Classdiagram`

	// Setup of pointers
	__Classdiagram__000000_Default_2.GongStructShapes = append(__Classdiagram__000000_Default_2.GongStructShapes, __GongStructShape__000000_Default_2_Classdiagram)
	__GongStructShape__000000_Default_2_Classdiagram.Position = __Position__000000_Pos_Default_2_Classdiagram
}
