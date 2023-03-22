package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/dummy/go/models"
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

	"ref_models.Dummy": &(ref_models.Dummy{}),

	"ref_models.Dummy.Name": (ref_models.Dummy{}).Name,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Dummy := (&models.GongStructShape{Name: `NewDiagram-Dummy`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Dummy := (&models.Position{Name: `Pos-NewDiagram-Dummy`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Dummy.Name]
	__Field__000000_Name.Identifier = `ref_models.Dummy.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Dummy`
	__Field__000000_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Dummy.Name = `NewDiagram-Dummy`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Dummy]
	__GongStructShape__000000_NewDiagram_Dummy.Identifier = `ref_models.Dummy`
	__GongStructShape__000000_NewDiagram_Dummy.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Dummy.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Dummy.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Dummy.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_Dummy.IsSelected = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_Dummy.X = 91.000000
	__Position__000000_Pos_NewDiagram_Dummy.Y = 43.000000
	__Position__000000_Pos_NewDiagram_Dummy.Name = `Pos-NewDiagram-Dummy`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Dummy)
	__GongStructShape__000000_NewDiagram_Dummy.Position = __Position__000000_Pos_NewDiagram_Dummy
	__GongStructShape__000000_NewDiagram_Dummy.Fields = append(__GongStructShape__000000_NewDiagram_Dummy.Fields, __Field__000000_Name)
}


