package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_DiagramEnums models.StageStruct
var ___dummy__Time_DiagramEnums time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_DiagramEnums ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_DiagramEnums map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["DiagramEnums"] = DiagramEnumsInjection
// }

// DiagramEnumsInjection will stage objects of database "DiagramEnums"
func DiagramEnumsInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_DiagramEnums := (&models.Classdiagram{Name: `DiagramEnums`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0011 := (&models.Classshape{Name: `Classshape0011`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_ := (&models.Field{Name: ``}).Stage()

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0011 := (&models.Position{Name: `Position-0011`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_LineType := (&models.Reference{Name: `LineType`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_DiagramEnums.Name = `DiagramEnums`
	__Classdiagram__000000_DiagramEnums.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0011.Name = `Classshape0011`
	__Classshape__000000_Classshape0011.ReferenceName = `LineType`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.LineType]
	__Classshape__000000_Classshape0011.Identifier = `Ref_models.LineType`
	__Classshape__000000_Classshape0011.ShowNbInstances = false
	__Classshape__000000_Classshape0011.NbInstances = 0
	__Classshape__000000_Classshape0011.Width = 240.000000
	__Classshape__000000_Classshape0011.Heigth = 78.000000
	__Classshape__000000_Classshape0011.IsSelected = false

	// Field values setup
	__Field__000000_.Name = ``
	__Field__000000_.Fieldname = `CONTINUOUS`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.CONTINUOUS]
	__Field__000000_.Identifier = `Ref_models.CONTINUOUS`
	__Field__000000_.FieldTypeAsString = ``
	__Field__000000_.Structname = `models`
	__Field__000000_.Fieldtypename = ``

	// Position values setup
	__Position__000000_Position_0011.X = 150.000000
	__Position__000000_Position_0011.Y = 40.000000
	__Position__000000_Position_0011.Name = `Position-0011`

	// Reference values setup
	__Reference__000000_LineType.Name = `LineType`
	__Reference__000000_LineType.NbInstances = 0
	__Reference__000000_LineType.Type = models.REFERENCE_GONG_ENUM

	// Setup of pointers
	__Classdiagram__000000_DiagramEnums.Classshapes = append(__Classdiagram__000000_DiagramEnums.Classshapes, __Classshape__000000_Classshape0011)
	__Classshape__000000_Classshape0011.Position = __Position__000000_Position_0011
	__Classshape__000000_Classshape0011.Reference = __Reference__000000_LineType
	__Classshape__000000_Classshape0011.Fields = append(__Classshape__000000_Classshape0011.Fields, __Field__000000_)
}


