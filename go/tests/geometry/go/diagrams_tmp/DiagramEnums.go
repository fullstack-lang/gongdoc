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
	__Classshape__000000_Classshape0009 := (&models.Classshape{Name: `Classshape0009`}).Stage()
	__Classshape__000001_Classshape0010 := (&models.Classshape{Name: `Classshape0010`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_ := (&models.Field{Name: ``}).Stage()
	__Field__000001_ := (&models.Field{Name: ``}).Stage()
	__Field__000002_ := (&models.Field{Name: ``}).Stage()
	__Field__000003_ := (&models.Field{Name: ``}).Stage()

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0009 := (&models.Position{Name: `Position-0009`}).Stage()
	__Position__000001_Position_0010 := (&models.Position{Name: `Position-0010`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_LineType := (&models.Reference{Name: `LineType`}).Stage()
	__Reference__000001_SimulationStage := (&models.Reference{Name: `SimulationStage`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_DiagramEnums.Name = `DiagramEnums`
	__Classdiagram__000000_DiagramEnums.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0009.Name = `Classshape0009`
	__Classshape__000000_Classshape0009.ReferenceName = `LineType`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineType]
	__Classshape__000000_Classshape0009.Identifier = `ref_models.LineType`
	__Classshape__000000_Classshape0009.ShowNbInstances = false
	__Classshape__000000_Classshape0009.NbInstances = 0
	__Classshape__000000_Classshape0009.Width = 240.000000
	__Classshape__000000_Classshape0009.Heigth = 93.000000
	__Classshape__000000_Classshape0009.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0010.Name = `Classshape0010`
	__Classshape__000001_Classshape0010.ReferenceName = `SimulationStage`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SimulationStage]
	__Classshape__000001_Classshape0010.Identifier = `ref_models.SimulationStage`
	__Classshape__000001_Classshape0010.ShowNbInstances = false
	__Classshape__000001_Classshape0010.NbInstances = 0
	__Classshape__000001_Classshape0010.Width = 240.000000
	__Classshape__000001_Classshape0010.Heigth = 93.000000
	__Classshape__000001_Classshape0010.IsSelected = false

	// Field values setup
	__Field__000000_.Name = ``
	__Field__000000_.Fieldname = `DOTTED`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DOTTED]
	__Field__000000_.Identifier = `ref_models.DOTTED`
	__Field__000000_.FieldTypeAsString = ``
	__Field__000000_.Structname = `models`
	__Field__000000_.Fieldtypename = ``

	// Field values setup
	__Field__000001_.Name = ``
	__Field__000001_.Fieldname = `AFTER`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AFTER]
	__Field__000001_.Identifier = `ref_models.AFTER`
	__Field__000001_.FieldTypeAsString = ``
	__Field__000001_.Structname = `models`
	__Field__000001_.Fieldtypename = ``

	// Field values setup
	__Field__000002_.Name = ``
	__Field__000002_.Fieldname = `CONTINUOUS`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CONTINUOUS]
	__Field__000002_.Identifier = `ref_models.CONTINUOUS`
	__Field__000002_.FieldTypeAsString = ``
	__Field__000002_.Structname = `models`
	__Field__000002_.Fieldtypename = ``

	// Field values setup
	__Field__000003_.Name = ``
	__Field__000003_.Fieldname = `BEFORE`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.BEFORE]
	__Field__000003_.Identifier = `ref_models.BEFORE`
	__Field__000003_.FieldTypeAsString = ``
	__Field__000003_.Structname = `models`
	__Field__000003_.Fieldtypename = ``

	// Position values setup
	__Position__000000_Position_0009.X = 150.000000
	__Position__000000_Position_0009.Y = 40.000000
	__Position__000000_Position_0009.Name = `Position-0009`

	// Position values setup
	__Position__000001_Position_0010.X = 240.000000
	__Position__000001_Position_0010.Y = 130.000000
	__Position__000001_Position_0010.Name = `Position-0010`

	// Reference values setup
	__Reference__000000_LineType.Name = `LineType`
	__Reference__000000_LineType.NbInstances = 0
	__Reference__000000_LineType.Type = models.REFERENCE_GONG_ENUM

	// Reference values setup
	__Reference__000001_SimulationStage.Name = `SimulationStage`
	__Reference__000001_SimulationStage.NbInstances = 0
	__Reference__000001_SimulationStage.Type = models.REFERENCE_GONG_ENUM

	// Setup of pointers
	__Classdiagram__000000_DiagramEnums.Classshapes = append(__Classdiagram__000000_DiagramEnums.Classshapes, __Classshape__000000_Classshape0009)
	__Classdiagram__000000_DiagramEnums.Classshapes = append(__Classdiagram__000000_DiagramEnums.Classshapes, __Classshape__000001_Classshape0010)
	__Classshape__000000_Classshape0009.Position = __Position__000000_Position_0009
	__Classshape__000000_Classshape0009.Reference = __Reference__000000_LineType
	__Classshape__000000_Classshape0009.Fields = append(__Classshape__000000_Classshape0009.Fields, __Field__000002_)
	__Classshape__000000_Classshape0009.Fields = append(__Classshape__000000_Classshape0009.Fields, __Field__000000_)
	__Classshape__000001_Classshape0010.Position = __Position__000001_Position_0010
	__Classshape__000001_Classshape0010.Reference = __Reference__000001_SimulationStage
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000001_)
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000003_)
}


