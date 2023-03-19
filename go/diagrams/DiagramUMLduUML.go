package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_DiagramUMLduUML models.StageStruct
var ___dummy__Time_DiagramUMLduUML time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_DiagramUMLduUML ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_DiagramUMLduUML map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Classdiagram": &(ref_models.Classdiagram{}),

	"ref_models.Classdiagram.Classshapes": (ref_models.Classdiagram{}).Classshapes,

	"ref_models.Classdiagram.GongStructShapes": (ref_models.Classdiagram{}).GongStructShapes,

	"ref_models.Classdiagram.Name": (ref_models.Classdiagram{}).Name,

	"ref_models.DiagramPackage": &(ref_models.DiagramPackage{}),

	"ref_models.DiagramPackage.Classdiagrams": (ref_models.DiagramPackage{}).Classdiagrams,

	"ref_models.DiagramPackage.Name": (ref_models.DiagramPackage{}).Name,

	"ref_models.DiagramPackage.Path": (ref_models.DiagramPackage{}).Path,

	"ref_models.DiagramPackage.Umlscs": (ref_models.DiagramPackage{}).Umlscs,

	"ref_models.Field": &(ref_models.Field{}),

	"ref_models.Field.FieldTypeAsString": (ref_models.Field{}).FieldTypeAsString,

	"ref_models.Field.Fieldtypename": (ref_models.Field{}).Fieldtypename,

	"ref_models.GongStructShape": &(ref_models.GongStructShape{}),

	"ref_models.GongStructShape.Fields": (ref_models.GongStructShape{}).Fields,

	"ref_models.GongStructShape.Links": (ref_models.GongStructShape{}).Links,

	"ref_models.GongStructShape.Name": (ref_models.GongStructShape{}).Name,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.Fieldtypename": (ref_models.Link{}).Fieldtypename,

	"ref_models.Link.Middlevertice": (ref_models.Link{}).Middlevertice,

	"ref_models.Link.TargetMultiplicity": (ref_models.Link{}).TargetMultiplicity,

	"ref_models.Position": &(ref_models.Position{}),

	"ref_models.Position.X": (ref_models.Position{}).X,

	"ref_models.Position.Y": (ref_models.Position{}).Y,

	"ref_models.UmlState": &(ref_models.UmlState{}),

	"ref_models.UmlState.Name": (ref_models.UmlState{}).Name,

	"ref_models.UmlState.X": (ref_models.UmlState{}).X,

	"ref_models.UmlState.Y": (ref_models.UmlState{}).Y,

	"ref_models.Umlsc": &(ref_models.Umlsc{}),

	"ref_models.Umlsc.Name": (ref_models.Umlsc{}).Name,

	"ref_models.Umlsc.States": (ref_models.Umlsc{}).States,

	"ref_models.Vertice": &(ref_models.Vertice{}),

	"ref_models.Vertice.X": (ref_models.Vertice{}).X,

	"ref_models.Vertice.Y": (ref_models.Vertice{}).Y,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["DiagramUMLduUML"] = DiagramUMLduUMLInjection
// }

// DiagramUMLduUMLInjection will stage objects of database "DiagramUMLduUML"
func DiagramUMLduUMLInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_DiagramUMLduUML := (&models.Classdiagram{Name: `DiagramUMLduUML`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_FieldTypeAsString := (&models.Field{Name: `FieldTypeAsString`}).Stage(stage)
	__Field__000001_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage(stage)
	__Field__000002_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000008_Path := (&models.Field{Name: `Path`}).Stage(stage)
	__Field__000009_TargetMultiplicity := (&models.Field{Name: `TargetMultiplicity`}).Stage(stage)
	__Field__000010_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000011_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000012_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000013_Y := (&models.Field{Name: `Y`}).Stage(stage)
	__Field__000014_Y := (&models.Field{Name: `Y`}).Stage(stage)
	__Field__000015_Y := (&models.Field{Name: `Y`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Classshape0000 := (&models.GongStructShape{Name: `Classshape0000`}).Stage(stage)
	__GongStructShape__000001_Classshape0002 := (&models.GongStructShape{Name: `Classshape0002`}).Stage(stage)
	__GongStructShape__000002_Classshape0003 := (&models.GongStructShape{Name: `Classshape0003`}).Stage(stage)
	__GongStructShape__000003_Classshape0004 := (&models.GongStructShape{Name: `Classshape0004`}).Stage(stage)
	__GongStructShape__000004_Classshape0005 := (&models.GongStructShape{Name: `Classshape0005`}).Stage(stage)
	__GongStructShape__000005_Classshape0006 := (&models.GongStructShape{Name: `Classshape0006`}).Stage(stage)
	__GongStructShape__000006_Classshape0007 := (&models.GongStructShape{Name: `Classshape0007`}).Stage(stage)
	__GongStructShape__000007_Classshape0008 := (&models.GongStructShape{Name: `Classshape0008`}).Stage(stage)
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape := (&models.GongStructShape{Name: `DiagramUMLduUML-GongStructShape`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Classdiagrams := (&models.Link{Name: `Classdiagrams`}).Stage(stage)
	__Link__000001_Classshapes := (&models.Link{Name: `Classshapes`}).Stage(stage)
	__Link__000002_Fields := (&models.Link{Name: `Fields`}).Stage(stage)
	__Link__000003_GongStructShapes := (&models.Link{Name: `GongStructShapes`}).Stage(stage)
	__Link__000004_Links := (&models.Link{Name: `Links`}).Stage(stage)
	__Link__000005_Middlevertice := (&models.Link{Name: `Middlevertice`}).Stage(stage)
	__Link__000006_States := (&models.Link{Name: `States`}).Stage(stage)
	__Link__000007_Umlscs := (&models.Link{Name: `Umlscs`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_DiagramUMLduUML_GongStructShape := (&models.Position{Name: `Pos-DiagramUMLduUML-GongStructShape`}).Stage(stage)
	__Position__000001_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage(stage)
	__Position__000002_Position_0002 := (&models.Position{Name: `Position-0002`}).Stage(stage)
	__Position__000003_Position_0003 := (&models.Position{Name: `Position-0003`}).Stage(stage)
	__Position__000004_Position_0004 := (&models.Position{Name: `Position-0004`}).Stage(stage)
	__Position__000005_Position_0005 := (&models.Position{Name: `Position-0005`}).Stage(stage)
	__Position__000006_Position_0006 := (&models.Position{Name: `Position-0006`}).Stage(stage)
	__Position__000007_Position_0007 := (&models.Position{Name: `Position-0007`}).Stage(stage)
	__Position__000008_Position_0008 := (&models.Position{Name: `Position-0008`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage(stage)
	__Vertice__000001_Vertice_0004 := (&models.Vertice{Name: `Vertice-0004`}).Stage(stage)
	__Vertice__000002_Vertice_0005 := (&models.Vertice{Name: `Vertice-0005`}).Stage(stage)
	__Vertice__000003_Vertice_0006 := (&models.Vertice{Name: `Vertice-0006`}).Stage(stage)
	__Vertice__000004_Vertice_0007 := (&models.Vertice{Name: `Vertice-0007`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_Classshape0000_and_DiagramUMLduUML_GongStructShape := (&models.Vertice{Name: `Verticle in class diagram DiagramUMLduUML in middle between Classshape0000 and DiagramUMLduUML-GongStructShape`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0002 := (&models.Vertice{Name: `Verticle in class diagram DiagramUMLduUML in middle between DiagramUMLduUML-GongStructShape and Classshape0002`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0003 := (&models.Vertice{Name: `Verticle in class diagram DiagramUMLduUML in middle between DiagramUMLduUML-GongStructShape and Classshape0003`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_DiagramUMLduUML.Name = `DiagramUMLduUML`
	__Classdiagram__000000_DiagramUMLduUML.IsInDrawMode = true

	// Field values setup
	__Field__000000_FieldTypeAsString.Name = `FieldTypeAsString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.FieldTypeAsString]
	__Field__000000_FieldTypeAsString.Identifier = `ref_models.Field.FieldTypeAsString`
	__Field__000000_FieldTypeAsString.FieldTypeAsString = ``
	__Field__000000_FieldTypeAsString.Structname = ``
	__Field__000000_FieldTypeAsString.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Fieldtypename.Name = `Fieldtypename`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Fieldtypename]
	__Field__000001_Fieldtypename.Identifier = `ref_models.Field.Fieldtypename`
	__Field__000001_Fieldtypename.FieldTypeAsString = ``
	__Field__000001_Fieldtypename.Structname = ``
	__Field__000001_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Fieldtypename.Name = `Fieldtypename`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Fieldtypename]
	__Field__000002_Fieldtypename.Identifier = `ref_models.Link.Fieldtypename`
	__Field__000002_Fieldtypename.FieldTypeAsString = ``
	__Field__000002_Fieldtypename.Structname = ``
	__Field__000002_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.Name]
	__Field__000003_Name.Identifier = `ref_models.Classdiagram.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = ``
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStructShape.Name]
	__Field__000004_Name.Identifier = `ref_models.GongStructShape.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `GongStructShape`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.Name]
	__Field__000005_Name.Identifier = `ref_models.Umlsc.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = ``
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Name]
	__Field__000006_Name.Identifier = `ref_models.UmlState.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = ``
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Name]
	__Field__000007_Name.Identifier = `ref_models.DiagramPackage.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = ``
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Path.Name = `Path`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Path]
	__Field__000008_Path.Identifier = `ref_models.DiagramPackage.Path`
	__Field__000008_Path.FieldTypeAsString = ``
	__Field__000008_Path.Structname = ``
	__Field__000008_Path.Fieldtypename = `string`

	// Field values setup
	__Field__000009_TargetMultiplicity.Name = `TargetMultiplicity`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.TargetMultiplicity]
	__Field__000009_TargetMultiplicity.Identifier = `ref_models.Link.TargetMultiplicity`
	__Field__000009_TargetMultiplicity.FieldTypeAsString = ``
	__Field__000009_TargetMultiplicity.Structname = ``
	__Field__000009_TargetMultiplicity.Fieldtypename = `MultiplicityType`

	// Field values setup
	__Field__000010_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.X]
	__Field__000010_X.Identifier = `ref_models.Vertice.X`
	__Field__000010_X.FieldTypeAsString = ``
	__Field__000010_X.Structname = ``
	__Field__000010_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000011_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.X]
	__Field__000011_X.Identifier = `ref_models.Position.X`
	__Field__000011_X.FieldTypeAsString = ``
	__Field__000011_X.Structname = ``
	__Field__000011_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000012_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.X]
	__Field__000012_X.Identifier = `ref_models.UmlState.X`
	__Field__000012_X.FieldTypeAsString = ``
	__Field__000012_X.Structname = ``
	__Field__000012_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.Y]
	__Field__000013_Y.Identifier = `ref_models.Position.Y`
	__Field__000013_Y.FieldTypeAsString = ``
	__Field__000013_Y.Structname = ``
	__Field__000013_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000014_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Y]
	__Field__000014_Y.Identifier = `ref_models.UmlState.Y`
	__Field__000014_Y.FieldTypeAsString = ``
	__Field__000014_Y.Structname = ``
	__Field__000014_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.Y]
	__Field__000015_Y.Identifier = `ref_models.Vertice.Y`
	__Field__000015_Y.FieldTypeAsString = ``
	__Field__000015_Y.Structname = ``
	__Field__000015_Y.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_Classshape0000.Name = `Classshape0000`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__GongStructShape__000000_Classshape0000.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Classshape0000.ShowNbInstances = true
	__GongStructShape__000000_Classshape0000.NbInstances = 0
	__GongStructShape__000000_Classshape0000.Width = 240.000000
	__GongStructShape__000000_Classshape0000.Heigth = 63.000000
	__GongStructShape__000000_Classshape0000.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Classshape0002.Name = `Classshape0002`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field]
	__GongStructShape__000001_Classshape0002.Identifier = `ref_models.Field`
	__GongStructShape__000001_Classshape0002.ShowNbInstances = true
	__GongStructShape__000001_Classshape0002.NbInstances = 0
	__GongStructShape__000001_Classshape0002.Width = 240.000000
	__GongStructShape__000001_Classshape0002.Heigth = 108.000000
	__GongStructShape__000001_Classshape0002.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Classshape0003.Name = `Classshape0003`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__GongStructShape__000002_Classshape0003.Identifier = `ref_models.Link`
	__GongStructShape__000002_Classshape0003.ShowNbInstances = true
	__GongStructShape__000002_Classshape0003.NbInstances = 0
	__GongStructShape__000002_Classshape0003.Width = 240.000000
	__GongStructShape__000002_Classshape0003.Heigth = 108.000000
	__GongStructShape__000002_Classshape0003.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Classshape0004.Name = `Classshape0004`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage]
	__GongStructShape__000003_Classshape0004.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000003_Classshape0004.ShowNbInstances = true
	__GongStructShape__000003_Classshape0004.NbInstances = 0
	__GongStructShape__000003_Classshape0004.Width = 240.000000
	__GongStructShape__000003_Classshape0004.Heigth = 78.000000
	__GongStructShape__000003_Classshape0004.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Classshape0005.Name = `Classshape0005`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position]
	__GongStructShape__000004_Classshape0005.Identifier = `ref_models.Position`
	__GongStructShape__000004_Classshape0005.ShowNbInstances = true
	__GongStructShape__000004_Classshape0005.NbInstances = 0
	__GongStructShape__000004_Classshape0005.Width = 240.000000
	__GongStructShape__000004_Classshape0005.Heigth = 78.000000
	__GongStructShape__000004_Classshape0005.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Classshape0006.Name = `Classshape0006`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState]
	__GongStructShape__000005_Classshape0006.Identifier = `ref_models.UmlState`
	__GongStructShape__000005_Classshape0006.ShowNbInstances = true
	__GongStructShape__000005_Classshape0006.NbInstances = 0
	__GongStructShape__000005_Classshape0006.Width = 240.000000
	__GongStructShape__000005_Classshape0006.Heigth = 93.000000
	__GongStructShape__000005_Classshape0006.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_Classshape0007.Name = `Classshape0007`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc]
	__GongStructShape__000006_Classshape0007.Identifier = `ref_models.Umlsc`
	__GongStructShape__000006_Classshape0007.ShowNbInstances = true
	__GongStructShape__000006_Classshape0007.NbInstances = 0
	__GongStructShape__000006_Classshape0007.Width = 240.000000
	__GongStructShape__000006_Classshape0007.Heigth = 63.000000
	__GongStructShape__000006_Classshape0007.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_Classshape0008.Name = `Classshape0008`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice]
	__GongStructShape__000007_Classshape0008.Identifier = `ref_models.Vertice`
	__GongStructShape__000007_Classshape0008.ShowNbInstances = true
	__GongStructShape__000007_Classshape0008.NbInstances = 0
	__GongStructShape__000007_Classshape0008.Width = 240.000000
	__GongStructShape__000007_Classshape0008.Heigth = 78.000000
	__GongStructShape__000007_Classshape0008.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Name = `DiagramUMLduUML-GongStructShape`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStructShape]
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Identifier = `ref_models.GongStructShape`
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.ShowNbInstances = false
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.NbInstances = 0
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Width = 240.000000
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Heigth = 78.000000
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.IsSelected = false

	// Link values setup
	__Link__000000_Classdiagrams.Name = `Classdiagrams`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Classdiagrams]
	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__Link__000000_Classdiagrams.Fieldtypename = `ref_models.Classdiagram`
	__Link__000000_Classdiagrams.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Classshapes.Name = `Classshapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.Classshapes]
	__Link__000001_Classshapes.Identifier = `ref_models.Classdiagram.Classshapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape]
	__Link__000001_Classshapes.Fieldtypename = `ref_models.Classshape`
	__Link__000001_Classshapes.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000002_Fields.Name = `Fields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStructShape.Fields]
	__Link__000002_Fields.Identifier = `ref_models.GongStructShape.Fields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field]
	__Link__000002_Fields.Fieldtypename = `ref_models.Field`
	__Link__000002_Fields.TargetMultiplicity = models.MANY
	__Link__000002_Fields.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000003_GongStructShapes.Name = `GongStructShapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.GongStructShapes]
	__Link__000003_GongStructShapes.Identifier = `ref_models.Classdiagram.GongStructShapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStructShape]
	__Link__000003_GongStructShapes.Fieldtypename = `ref_models.GongStructShape`
	__Link__000003_GongStructShapes.TargetMultiplicity = models.MANY
	__Link__000003_GongStructShapes.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000004_Links.Name = `Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStructShape.Links]
	__Link__000004_Links.Identifier = `ref_models.GongStructShape.Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Link__000004_Links.Fieldtypename = `ref_models.Link`
	__Link__000004_Links.TargetMultiplicity = models.MANY
	__Link__000004_Links.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000005_Middlevertice.Name = `Middlevertice`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Middlevertice]
	__Link__000005_Middlevertice.Identifier = `ref_models.Link.Middlevertice`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice]
	__Link__000005_Middlevertice.Fieldtypename = `ref_models.Vertice`
	__Link__000005_Middlevertice.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000006_States.Name = `States`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.States]
	__Link__000006_States.Identifier = `ref_models.Umlsc.States`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState]
	__Link__000006_States.Fieldtypename = `ref_models.UmlState`
	__Link__000006_States.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000007_Umlscs.Name = `Umlscs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Umlscs]
	__Link__000007_Umlscs.Identifier = `ref_models.DiagramPackage.Umlscs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc]
	__Link__000007_Umlscs.Fieldtypename = `ref_models.Umlsc`
	__Link__000007_Umlscs.TargetMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_DiagramUMLduUML_GongStructShape.X = 80.000000
	__Position__000000_Pos_DiagramUMLduUML_GongStructShape.Y = 430.000000
	__Position__000000_Pos_DiagramUMLduUML_GongStructShape.Name = `Pos-DiagramUMLduUML-GongStructShape`

	// Position values setup
	__Position__000001_Position_0000.X = 80.000000
	__Position__000001_Position_0000.Y = 290.000000
	__Position__000001_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000002_Position_0002.X = 350.000000
	__Position__000002_Position_0002.Y = 680.000000
	__Position__000002_Position_0002.Name = `Position-0002`

	// Position values setup
	__Position__000003_Position_0003.X = 80.000000
	__Position__000003_Position_0003.Y = 680.000000
	__Position__000003_Position_0003.Name = `Position-0003`

	// Position values setup
	__Position__000004_Position_0004.X = 350.000000
	__Position__000004_Position_0004.Y = 80.000000
	__Position__000004_Position_0004.Name = `Position-0004`

	// Position values setup
	__Position__000005_Position_0005.X = 600.000000
	__Position__000005_Position_0005.Y = 680.000000
	__Position__000005_Position_0005.Name = `Position-0005`

	// Position values setup
	__Position__000006_Position_0006.X = 580.000000
	__Position__000006_Position_0006.Y = 480.000000
	__Position__000006_Position_0006.Name = `Position-0006`

	// Position values setup
	__Position__000007_Position_0007.X = 580.000000
	__Position__000007_Position_0007.Y = 290.000000
	__Position__000007_Position_0007.Name = `Position-0007`

	// Position values setup
	__Position__000008_Position_0008.X = 80.000000
	__Position__000008_Position_0008.Y = 850.000000
	__Position__000008_Position_0008.Name = `Position-0008`

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 200.000000
	__Vertice__000000_Vertice_0000.Y = 420.000000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0004.X = 200.000000
	__Vertice__000001_Vertice_0004.Y = 790.000000
	__Vertice__000001_Vertice_0004.Name = `Vertice-0004`

	// Vertice values setup
	__Vertice__000002_Vertice_0005.X = 200.000000
	__Vertice__000002_Vertice_0005.Y = 170.000000
	__Vertice__000002_Vertice_0005.Name = `Vertice-0005`

	// Vertice values setup
	__Vertice__000003_Vertice_0006.X = 700.000000
	__Vertice__000003_Vertice_0006.Y = 180.000000
	__Vertice__000003_Vertice_0006.Name = `Vertice-0006`

	// Vertice values setup
	__Vertice__000004_Vertice_0007.X = 700.000000
	__Vertice__000004_Vertice_0007.Y = 420.000000
	__Vertice__000004_Vertice_0007.Name = `Vertice-0007`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_Classshape0000_and_DiagramUMLduUML_GongStructShape.X = 440.000000
	__Vertice__000005_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_Classshape0000_and_DiagramUMLduUML_GongStructShape.Y = 391.500000
	__Vertice__000005_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_Classshape0000_and_DiagramUMLduUML_GongStructShape.Name = `Verticle in class diagram DiagramUMLduUML in middle between Classshape0000 and DiagramUMLduUML-GongStructShape`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0002.X = 476.000000
	__Vertice__000006_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0002.Y = 608.000000
	__Vertice__000006_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0002.Name = `Verticle in class diagram DiagramUMLduUML in middle between DiagramUMLduUML-GongStructShape and Classshape0002`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0003.X = 201.000000
	__Vertice__000007_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0003.Y = 600.500000
	__Vertice__000007_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0003.Name = `Verticle in class diagram DiagramUMLduUML in middle between DiagramUMLduUML-GongStructShape and Classshape0003`

	// Setup of pointers
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000000_Classshape0000)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000001_Classshape0002)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000002_Classshape0003)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000003_Classshape0004)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000004_Classshape0005)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000005_Classshape0006)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000006_Classshape0007)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000007_Classshape0008)
	__Classdiagram__000000_DiagramUMLduUML.GongStructShapes = append(__Classdiagram__000000_DiagramUMLduUML.GongStructShapes, __GongStructShape__000008_DiagramUMLduUML_GongStructShape)
	__GongStructShape__000000_Classshape0000.Position = __Position__000001_Position_0000
	__GongStructShape__000000_Classshape0000.Fields = append(__GongStructShape__000000_Classshape0000.Fields, __Field__000003_Name)
	__GongStructShape__000000_Classshape0000.Links = append(__GongStructShape__000000_Classshape0000.Links, __Link__000001_Classshapes)
	__GongStructShape__000000_Classshape0000.Links = append(__GongStructShape__000000_Classshape0000.Links, __Link__000003_GongStructShapes)
	__GongStructShape__000001_Classshape0002.Position = __Position__000002_Position_0002
	__GongStructShape__000001_Classshape0002.Fields = append(__GongStructShape__000001_Classshape0002.Fields, __Field__000000_FieldTypeAsString)
	__GongStructShape__000001_Classshape0002.Fields = append(__GongStructShape__000001_Classshape0002.Fields, __Field__000001_Fieldtypename)
	__GongStructShape__000002_Classshape0003.Position = __Position__000003_Position_0003
	__GongStructShape__000002_Classshape0003.Fields = append(__GongStructShape__000002_Classshape0003.Fields, __Field__000002_Fieldtypename)
	__GongStructShape__000002_Classshape0003.Fields = append(__GongStructShape__000002_Classshape0003.Fields, __Field__000009_TargetMultiplicity)
	__GongStructShape__000002_Classshape0003.Links = append(__GongStructShape__000002_Classshape0003.Links, __Link__000005_Middlevertice)
	__GongStructShape__000003_Classshape0004.Position = __Position__000004_Position_0004
	__GongStructShape__000003_Classshape0004.Fields = append(__GongStructShape__000003_Classshape0004.Fields, __Field__000007_Name)
	__GongStructShape__000003_Classshape0004.Fields = append(__GongStructShape__000003_Classshape0004.Fields, __Field__000008_Path)
	__GongStructShape__000003_Classshape0004.Links = append(__GongStructShape__000003_Classshape0004.Links, __Link__000000_Classdiagrams)
	__GongStructShape__000003_Classshape0004.Links = append(__GongStructShape__000003_Classshape0004.Links, __Link__000007_Umlscs)
	__GongStructShape__000004_Classshape0005.Position = __Position__000005_Position_0005
	__GongStructShape__000004_Classshape0005.Fields = append(__GongStructShape__000004_Classshape0005.Fields, __Field__000011_X)
	__GongStructShape__000004_Classshape0005.Fields = append(__GongStructShape__000004_Classshape0005.Fields, __Field__000013_Y)
	__GongStructShape__000005_Classshape0006.Position = __Position__000006_Position_0006
	__GongStructShape__000005_Classshape0006.Fields = append(__GongStructShape__000005_Classshape0006.Fields, __Field__000006_Name)
	__GongStructShape__000005_Classshape0006.Fields = append(__GongStructShape__000005_Classshape0006.Fields, __Field__000012_X)
	__GongStructShape__000005_Classshape0006.Fields = append(__GongStructShape__000005_Classshape0006.Fields, __Field__000014_Y)
	__GongStructShape__000006_Classshape0007.Position = __Position__000007_Position_0007
	__GongStructShape__000006_Classshape0007.Fields = append(__GongStructShape__000006_Classshape0007.Fields, __Field__000005_Name)
	__GongStructShape__000006_Classshape0007.Links = append(__GongStructShape__000006_Classshape0007.Links, __Link__000006_States)
	__GongStructShape__000007_Classshape0008.Position = __Position__000008_Position_0008
	__GongStructShape__000007_Classshape0008.Fields = append(__GongStructShape__000007_Classshape0008.Fields, __Field__000010_X)
	__GongStructShape__000007_Classshape0008.Fields = append(__GongStructShape__000007_Classshape0008.Fields, __Field__000015_Y)
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Position = __Position__000000_Pos_DiagramUMLduUML_GongStructShape
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Fields = append(__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Fields, __Field__000004_Name)
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Links = append(__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Links, __Link__000004_Links)
	__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Links = append(__GongStructShape__000008_DiagramUMLduUML_GongStructShape.Links, __Link__000002_Fields)
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000002_Vertice_0005
	__Link__000001_Classshapes.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000002_Fields.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0002
	__Link__000003_GongStructShapes.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_Classshape0000_and_DiagramUMLduUML_GongStructShape
	__Link__000004_Links.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_DiagramUMLduUML_in_middle_between_DiagramUMLduUML_GongStructShape_and_Classshape0003
	__Link__000005_Middlevertice.Middlevertice = __Vertice__000001_Vertice_0004
	__Link__000006_States.Middlevertice = __Vertice__000004_Vertice_0007
	__Link__000007_Umlscs.Middlevertice = __Vertice__000003_Vertice_0006
}


