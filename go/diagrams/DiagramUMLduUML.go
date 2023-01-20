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
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["DiagramUMLduUML"] = DiagramUMLduUMLInjection
// }

// DiagramUMLduUMLInjection will stage objects of database "DiagramUMLduUML"
func DiagramUMLduUMLInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_DiagramUMLduUML := (&models.Classdiagram{Name: `DiagramUMLduUML`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0000 := (&models.Classshape{Name: `Classshape0000`}).Stage()
	__Classshape__000001_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()
	__Classshape__000002_Classshape0002 := (&models.Classshape{Name: `Classshape0002`}).Stage()
	__Classshape__000003_Classshape0003 := (&models.Classshape{Name: `Classshape0003`}).Stage()
	__Classshape__000004_Classshape0004 := (&models.Classshape{Name: `Classshape0004`}).Stage()
	__Classshape__000005_Classshape0005 := (&models.Classshape{Name: `Classshape0005`}).Stage()
	__Classshape__000006_Classshape0006 := (&models.Classshape{Name: `Classshape0006`}).Stage()
	__Classshape__000007_Classshape0007 := (&models.Classshape{Name: `Classshape0007`}).Stage()
	__Classshape__000008_Classshape0008 := (&models.Classshape{Name: `Classshape0008`}).Stage()

	// Declarations of staged instances of DiagramPackage
	__DiagramPackage__000000_gongdoc_diagrams := (&models.DiagramPackage{Name: `gongdoc_diagrams`}).Stage()

	// Declarations of staged instances of Field
	__Field__000000_FieldTypeAsString := (&models.Field{Name: `FieldTypeAsString`}).Stage()
	__Field__000001_Fieldname := (&models.Field{Name: `Fieldname`}).Stage()
	__Field__000002_Fieldname := (&models.Field{Name: `Fieldname`}).Stage()
	__Field__000003_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage()
	__Field__000004_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage()
	__Field__000005_Heigth := (&models.Field{Name: `Heigth`}).Stage()
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000008_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000009_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000010_Path := (&models.Field{Name: `Path`}).Stage()

	__Field__000012_Structname := (&models.Field{Name: `Structname`}).Stage()
	__Field__000013_Structname := (&models.Field{Name: `Structname`}).Stage()
	__Field__000014_TargetMultiplicity := (&models.Field{Name: `TargetMultiplicity`}).Stage()
	__Field__000015_Width := (&models.Field{Name: `Width`}).Stage()
	__Field__000016_X := (&models.Field{Name: `X`}).Stage()
	__Field__000017_X := (&models.Field{Name: `X`}).Stage()
	__Field__000018_X := (&models.Field{Name: `X`}).Stage()
	__Field__000019_Y := (&models.Field{Name: `Y`}).Stage()
	__Field__000020_Y := (&models.Field{Name: `Y`}).Stage()
	__Field__000021_Y := (&models.Field{Name: `Y`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Classdiagrams := (&models.Link{Name: `Classdiagrams`}).Stage()
	__Link__000001_Classshapes := (&models.Link{Name: `Classshapes`}).Stage()
	__Link__000002_Fields := (&models.Link{Name: `Fields`}).Stage()
	__Link__000003_Links := (&models.Link{Name: `Links`}).Stage()
	__Link__000004_Middlevertice := (&models.Link{Name: `Middlevertice`}).Stage()
	__Link__000005_Position := (&models.Link{Name: `Position`}).Stage()
	__Link__000006_States := (&models.Link{Name: `States`}).Stage()
	__Link__000007_Umlscs := (&models.Link{Name: `Umlscs`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()
	__Position__000002_Position_0002 := (&models.Position{Name: `Position-0002`}).Stage()
	__Position__000003_Position_0003 := (&models.Position{Name: `Position-0003`}).Stage()
	__Position__000004_Position_0004 := (&models.Position{Name: `Position-0004`}).Stage()
	__Position__000005_Position_0005 := (&models.Position{Name: `Position-0005`}).Stage()
	__Position__000006_Position_0006 := (&models.Position{Name: `Position-0006`}).Stage()
	__Position__000007_Position_0007 := (&models.Position{Name: `Position-0007`}).Stage()
	__Position__000008_Position_0008 := (&models.Position{Name: `Position-0008`}).Stage()

	// Declarations of staged instances of Reference

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage()
	__Vertice__000001_Vertice_0001 := (&models.Vertice{Name: `Vertice-0001`}).Stage()
	__Vertice__000002_Vertice_0002 := (&models.Vertice{Name: `Vertice-0002`}).Stage()
	__Vertice__000003_Vertice_0003 := (&models.Vertice{Name: `Vertice-0003`}).Stage()
	__Vertice__000004_Vertice_0004 := (&models.Vertice{Name: `Vertice-0004`}).Stage()
	__Vertice__000005_Vertice_0005 := (&models.Vertice{Name: `Vertice-0005`}).Stage()
	__Vertice__000006_Vertice_0006 := (&models.Vertice{Name: `Vertice-0006`}).Stage()
	__Vertice__000007_Vertice_0007 := (&models.Vertice{Name: `Vertice-0007`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_DiagramUMLduUML.Name = `DiagramUMLduUML`
	__Classdiagram__000000_DiagramUMLduUML.IsInDrawMode = false

	// Classshape values setup
	__Classshape__000000_Classshape0000.Name = `Classshape0000`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__Classshape__000000_Classshape0000.Identifier = `ref_models.Classdiagram`
	__Classshape__000000_Classshape0000.ShowNbInstances = true
	__Classshape__000000_Classshape0000.NbInstances = 0
	__Classshape__000000_Classshape0000.Width = 240.000000
	__Classshape__000000_Classshape0000.Heigth = 63.000000
	__Classshape__000000_Classshape0000.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0001.Name = `Classshape0001`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape]
	__Classshape__000001_Classshape0001.Identifier = `ref_models.Classshape`
	__Classshape__000001_Classshape0001.ShowNbInstances = true
	__Classshape__000001_Classshape0001.NbInstances = 0
	__Classshape__000001_Classshape0001.Width = 240.000000
	__Classshape__000001_Classshape0001.Heigth = 108.000000
	__Classshape__000001_Classshape0001.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0002.Name = `Classshape0002`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field]
	__Classshape__000002_Classshape0002.Identifier = `ref_models.Field`
	__Classshape__000002_Classshape0002.ShowNbInstances = true
	__Classshape__000002_Classshape0002.NbInstances = 0
	__Classshape__000002_Classshape0002.Width = 240.000000
	__Classshape__000002_Classshape0002.Heigth = 108.000000
	__Classshape__000002_Classshape0002.IsSelected = false

	// Classshape values setup
	__Classshape__000003_Classshape0003.Name = `Classshape0003`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Classshape__000003_Classshape0003.Identifier = `ref_models.Link`
	__Classshape__000003_Classshape0003.ShowNbInstances = true
	__Classshape__000003_Classshape0003.NbInstances = 0
	__Classshape__000003_Classshape0003.Width = 240.000000
	__Classshape__000003_Classshape0003.Heigth = 108.000000
	__Classshape__000003_Classshape0003.IsSelected = false

	// Classshape values setup
	__Classshape__000004_Classshape0004.Name = `Classshape0004`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage]
	__Classshape__000004_Classshape0004.Identifier = `ref_models.DiagramPackage`
	__Classshape__000004_Classshape0004.ShowNbInstances = true
	__Classshape__000004_Classshape0004.NbInstances = 0
	__Classshape__000004_Classshape0004.Width = 240.000000
	__Classshape__000004_Classshape0004.Heigth = 78.000000
	__Classshape__000004_Classshape0004.IsSelected = false

	// Classshape values setup
	__Classshape__000005_Classshape0005.Name = `Classshape0005`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position]
	__Classshape__000005_Classshape0005.Identifier = `ref_models.Position`
	__Classshape__000005_Classshape0005.ShowNbInstances = true
	__Classshape__000005_Classshape0005.NbInstances = 0
	__Classshape__000005_Classshape0005.Width = 240.000000
	__Classshape__000005_Classshape0005.Heigth = 78.000000
	__Classshape__000005_Classshape0005.IsSelected = false

	// Classshape values setup
	__Classshape__000006_Classshape0006.Name = `Classshape0006`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState]
	__Classshape__000006_Classshape0006.Identifier = `ref_models.UmlState`
	__Classshape__000006_Classshape0006.ShowNbInstances = true
	__Classshape__000006_Classshape0006.NbInstances = 0
	__Classshape__000006_Classshape0006.Width = 240.000000
	__Classshape__000006_Classshape0006.Heigth = 93.000000
	__Classshape__000006_Classshape0006.IsSelected = false

	// Classshape values setup
	__Classshape__000007_Classshape0007.Name = `Classshape0007`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc]
	__Classshape__000007_Classshape0007.Identifier = `ref_models.Umlsc`
	__Classshape__000007_Classshape0007.ShowNbInstances = true
	__Classshape__000007_Classshape0007.NbInstances = 0
	__Classshape__000007_Classshape0007.Width = 240.000000
	__Classshape__000007_Classshape0007.Heigth = 63.000000
	__Classshape__000007_Classshape0007.IsSelected = false

	// Classshape values setup
	__Classshape__000008_Classshape0008.Name = `Classshape0008`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice]
	__Classshape__000008_Classshape0008.Identifier = `ref_models.Vertice`
	__Classshape__000008_Classshape0008.ShowNbInstances = true
	__Classshape__000008_Classshape0008.NbInstances = 0
	__Classshape__000008_Classshape0008.Width = 240.000000
	__Classshape__000008_Classshape0008.Heigth = 78.000000
	__Classshape__000008_Classshape0008.IsSelected = false

	// DiagramPackage values setup
	__DiagramPackage__000000_gongdoc_diagrams.Name = `gongdoc_diagrams`
	__DiagramPackage__000000_gongdoc_diagrams.Path = `github.com/fullstack-lang/gongdoc/go/models`
	__DiagramPackage__000000_gongdoc_diagrams.GongModelPath = `github.com/fullstack-lang/gongdoc/go/models`
	__DiagramPackage__000000_gongdoc_diagrams.IsEditable = true
	__DiagramPackage__000000_gongdoc_diagrams.IsReloaded = false
	__DiagramPackage__000000_gongdoc_diagrams.AbsolutePathToDiagramPackage = `/Users/thomaspeugeot/go/src/github.com/fullstack-lang/gongdoc/go/diagrams`

	// Field values setup
	__Field__000000_FieldTypeAsString.Name = `FieldTypeAsString`
	__Field__000000_FieldTypeAsString.Fieldname = `FieldTypeAsString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.FieldTypeAsString]
	__Field__000000_FieldTypeAsString.Identifier = `ref_models.Field.FieldTypeAsString`
	__Field__000000_FieldTypeAsString.FieldTypeAsString = ``
	__Field__000000_FieldTypeAsString.Structname = `Field`
	__Field__000000_FieldTypeAsString.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Fieldname.Name = `Fieldname`
	__Field__000001_Fieldname.Fieldname = `Fieldname`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Fieldname]
	__Field__000001_Fieldname.Identifier = `ref_models.Field.Fieldname`
	__Field__000001_Fieldname.FieldTypeAsString = ``
	__Field__000001_Fieldname.Structname = `Field`
	__Field__000001_Fieldname.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Fieldname.Name = `Fieldname`
	__Field__000002_Fieldname.Fieldname = `Fieldname`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Fieldname]
	__Field__000002_Fieldname.Identifier = `ref_models.Link.Fieldname`
	__Field__000002_Fieldname.FieldTypeAsString = ``
	__Field__000002_Fieldname.Structname = `Link`
	__Field__000002_Fieldname.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Fieldtypename.Name = `Fieldtypename`
	__Field__000003_Fieldtypename.Fieldname = `Fieldtypename`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Fieldtypename]
	__Field__000003_Fieldtypename.Identifier = `ref_models.Link.Fieldtypename`
	__Field__000003_Fieldtypename.FieldTypeAsString = ``
	__Field__000003_Fieldtypename.Structname = `Link`
	__Field__000003_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Fieldtypename.Name = `Fieldtypename`
	__Field__000004_Fieldtypename.Fieldname = `Fieldtypename`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Fieldtypename]
	__Field__000004_Fieldtypename.Identifier = `ref_models.Field.Fieldtypename`
	__Field__000004_Fieldtypename.FieldTypeAsString = ``
	__Field__000004_Fieldtypename.Structname = `Field`
	__Field__000004_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Heigth.Name = `Heigth`
	__Field__000005_Heigth.Fieldname = `Heigth`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Heigth]
	__Field__000005_Heigth.Identifier = `ref_models.Classshape.Heigth`
	__Field__000005_Heigth.FieldTypeAsString = ``
	__Field__000005_Heigth.Structname = `Classshape`
	__Field__000005_Heigth.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_Name.Name = `Name`
	__Field__000006_Name.Fieldname = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Name]
	__Field__000006_Name.Identifier = `ref_models.UmlState.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `UmlState`
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Name.Name = `Name`
	__Field__000007_Name.Fieldname = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.Name]
	__Field__000007_Name.Identifier = `ref_models.Classdiagram.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `Classdiagram`
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Name.Name = `Name`
	__Field__000008_Name.Fieldname = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.Name]
	__Field__000008_Name.Identifier = `ref_models.Umlsc.Name`
	__Field__000008_Name.FieldTypeAsString = ``
	__Field__000008_Name.Structname = `Umlsc`
	__Field__000008_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000009_Name.Name = `Name`
	__Field__000009_Name.Fieldname = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Name]
	__Field__000009_Name.Identifier = `ref_models.DiagramPackage.Name`
	__Field__000009_Name.FieldTypeAsString = ``
	__Field__000009_Name.Structname = `DiagramPackage`
	__Field__000009_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Path.Name = `Path`
	__Field__000010_Path.Fieldname = `Path`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Path]
	__Field__000010_Path.Identifier = `ref_models.DiagramPackage.Path`
	__Field__000010_Path.FieldTypeAsString = ``
	__Field__000010_Path.Structname = `DiagramPackage`
	__Field__000010_Path.Fieldtypename = `string`

	// Field values setup

	// comment added to overcome the problem with the comment map association

	// Field values setup
	__Field__000012_Structname.Name = `Structname`
	__Field__000012_Structname.Fieldname = `Structname`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Structname]
	__Field__000012_Structname.Identifier = `ref_models.Field.Structname`
	__Field__000012_Structname.FieldTypeAsString = ``
	__Field__000012_Structname.Structname = `Field`
	__Field__000012_Structname.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Structname.Name = `Structname`
	__Field__000013_Structname.Fieldname = `Structname`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Structname]
	__Field__000013_Structname.Identifier = `ref_models.Link.Structname`
	__Field__000013_Structname.FieldTypeAsString = ``
	__Field__000013_Structname.Structname = `Link`
	__Field__000013_Structname.Fieldtypename = `string`

	// Field values setup
	__Field__000014_TargetMultiplicity.Name = `TargetMultiplicity`
	__Field__000014_TargetMultiplicity.Fieldname = `TargetMultiplicity`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.TargetMultiplicity]
	__Field__000014_TargetMultiplicity.Identifier = `ref_models.Link.TargetMultiplicity`
	__Field__000014_TargetMultiplicity.FieldTypeAsString = ``
	__Field__000014_TargetMultiplicity.Structname = `Link`
	__Field__000014_TargetMultiplicity.Fieldtypename = `MultiplicityType`

	// Field values setup
	__Field__000015_Width.Name = `Width`
	__Field__000015_Width.Fieldname = `Width`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Width]
	__Field__000015_Width.Identifier = `ref_models.Classshape.Width`
	__Field__000015_Width.FieldTypeAsString = ``
	__Field__000015_Width.Structname = `Classshape`
	__Field__000015_Width.Fieldtypename = `float64`

	// Field values setup
	__Field__000016_X.Name = `X`
	__Field__000016_X.Fieldname = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.X]
	__Field__000016_X.Identifier = `ref_models.UmlState.X`
	__Field__000016_X.FieldTypeAsString = ``
	__Field__000016_X.Structname = `UmlState`
	__Field__000016_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000017_X.Name = `X`
	__Field__000017_X.Fieldname = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.X]
	__Field__000017_X.Identifier = `ref_models.Position.X`
	__Field__000017_X.FieldTypeAsString = ``
	__Field__000017_X.Structname = `Position`
	__Field__000017_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000018_X.Name = `X`
	__Field__000018_X.Fieldname = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.X]
	__Field__000018_X.Identifier = `ref_models.Vertice.X`
	__Field__000018_X.FieldTypeAsString = ``
	__Field__000018_X.Structname = `Vertice`
	__Field__000018_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000019_Y.Name = `Y`
	__Field__000019_Y.Fieldname = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.Y]
	__Field__000019_Y.Identifier = `ref_models.Position.Y`
	__Field__000019_Y.FieldTypeAsString = ``
	__Field__000019_Y.Structname = `Position`
	__Field__000019_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000020_Y.Name = `Y`
	__Field__000020_Y.Fieldname = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.Y]
	__Field__000020_Y.Identifier = `ref_models.Vertice.Y`
	__Field__000020_Y.FieldTypeAsString = ``
	__Field__000020_Y.Structname = `Vertice`
	__Field__000020_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000021_Y.Name = `Y`
	__Field__000021_Y.Fieldname = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Y]
	__Field__000021_Y.Identifier = `ref_models.UmlState.Y`
	__Field__000021_Y.FieldTypeAsString = ``
	__Field__000021_Y.Structname = `UmlState`
	__Field__000021_Y.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_Classdiagrams.Name = `Classdiagrams`
	__Link__000000_Classdiagrams.Fieldname = `Classdiagrams`
	__Link__000000_Classdiagrams.Structname = `DiagramPackage`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Classdiagrams]
	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`
	__Link__000000_Classdiagrams.Fieldtypename = `Classdiagram`
	__Link__000000_Classdiagrams.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Classshapes.Name = `Classshapes`
	__Link__000001_Classshapes.Fieldname = `Classshapes`
	__Link__000001_Classshapes.Structname = `Classdiagram`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.Classshapes]
	__Link__000001_Classshapes.Identifier = `ref_models.Classdiagram.Classshapes`
	__Link__000001_Classshapes.Fieldtypename = `Classshape`
	__Link__000001_Classshapes.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000002_Fields.Name = `Fields`
	__Link__000002_Fields.Fieldname = `Fields`
	__Link__000002_Fields.Structname = `Classshape`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Fields]
	__Link__000002_Fields.Identifier = `ref_models.Classshape.Fields`
	__Link__000002_Fields.Fieldtypename = `Field`
	__Link__000002_Fields.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000003_Links.Name = `Links`
	__Link__000003_Links.Fieldname = `Links`
	__Link__000003_Links.Structname = `Classshape`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Links]
	__Link__000003_Links.Identifier = `ref_models.Classshape.Links`
	__Link__000003_Links.Fieldtypename = `Link`
	__Link__000003_Links.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000004_Middlevertice.Name = `Middlevertice`
	__Link__000004_Middlevertice.Fieldname = `Middlevertice`
	__Link__000004_Middlevertice.Structname = `Link`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Middlevertice]
	__Link__000004_Middlevertice.Identifier = `ref_models.Link.Middlevertice`
	__Link__000004_Middlevertice.Fieldtypename = `Vertice`
	__Link__000004_Middlevertice.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000005_Position.Name = `Position`
	__Link__000005_Position.Fieldname = `Position`
	__Link__000005_Position.Structname = `Classshape`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Position]
	__Link__000005_Position.Identifier = `ref_models.Classshape.Position`
	__Link__000005_Position.Fieldtypename = `Position`
	__Link__000005_Position.TargetMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000006_States.Name = `States`
	__Link__000006_States.Fieldname = `States`
	__Link__000006_States.Structname = `Umlsc`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.States]
	__Link__000006_States.Identifier = `ref_models.Umlsc.States`
	__Link__000006_States.Fieldtypename = `UmlState`
	__Link__000006_States.TargetMultiplicity = models.MANY

	// Link values setup
	__Link__000007_Umlscs.Name = `Umlscs`
	__Link__000007_Umlscs.Fieldname = `Umlscs`
	__Link__000007_Umlscs.Structname = `DiagramPackage`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Umlscs]
	__Link__000007_Umlscs.Identifier = `ref_models.DiagramPackage.Umlscs`
	__Link__000007_Umlscs.Fieldtypename = `Umlsc`
	__Link__000007_Umlscs.TargetMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Position_0000.X = 80.000000
	__Position__000000_Position_0000.Y = 290.000000
	__Position__000000_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000001_Position_0001.X = 70.000000
	__Position__000001_Position_0001.Y = 490.000000
	__Position__000001_Position_0001.Name = `Position-0001`

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
	__Position__000004_Position_0004.Y = 90.000000
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

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Reference values setup

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 200.000000
	__Vertice__000000_Vertice_0000.Y = 420.000000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0001.X = 470.000000
	__Vertice__000001_Vertice_0001.Y = 610.000000
	__Vertice__000001_Vertice_0001.Name = `Vertice-0001`

	// Vertice values setup
	__Vertice__000002_Vertice_0002.X = 200.000000
	__Vertice__000002_Vertice_0002.Y = 620.000000
	__Vertice__000002_Vertice_0002.Name = `Vertice-0002`

	// Vertice values setup
	__Vertice__000003_Vertice_0003.X = 725.000000
	__Vertice__000003_Vertice_0003.Y = 609.000000
	__Vertice__000003_Vertice_0003.Name = `Vertice-0003`

	// Vertice values setup
	__Vertice__000004_Vertice_0004.X = 200.000000
	__Vertice__000004_Vertice_0004.Y = 790.000000
	__Vertice__000004_Vertice_0004.Name = `Vertice-0004`

	// Vertice values setup
	__Vertice__000005_Vertice_0005.X = 200.000000
	__Vertice__000005_Vertice_0005.Y = 170.000000
	__Vertice__000005_Vertice_0005.Name = `Vertice-0005`

	// Vertice values setup
	__Vertice__000006_Vertice_0006.X = 700.000000
	__Vertice__000006_Vertice_0006.Y = 180.000000
	__Vertice__000006_Vertice_0006.Name = `Vertice-0006`

	// Vertice values setup
	__Vertice__000007_Vertice_0007.X = 700.000000
	__Vertice__000007_Vertice_0007.Y = 420.000000
	__Vertice__000007_Vertice_0007.Name = `Vertice-0007`

	// Setup of pointers
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000000_Classshape0000)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000001_Classshape0001)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000002_Classshape0002)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000003_Classshape0003)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000004_Classshape0004)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000005_Classshape0005)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000006_Classshape0006)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000007_Classshape0007)
	__Classdiagram__000000_DiagramUMLduUML.Classshapes = append(__Classdiagram__000000_DiagramUMLduUML.Classshapes, __Classshape__000008_Classshape0008)
	__Classshape__000000_Classshape0000.Position = __Position__000000_Position_0000

	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000007_Name)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000001_Classshapes)
	__Classshape__000001_Classshape0001.Position = __Position__000001_Position_0001

	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000005_Heigth)

	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000015_Width)
	__Classshape__000001_Classshape0001.Links = append(__Classshape__000001_Classshape0001.Links, __Link__000002_Fields)
	__Classshape__000001_Classshape0001.Links = append(__Classshape__000001_Classshape0001.Links, __Link__000003_Links)
	__Classshape__000001_Classshape0001.Links = append(__Classshape__000001_Classshape0001.Links, __Link__000005_Position)
	__Classshape__000002_Classshape0002.Position = __Position__000002_Position_0002

	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000000_FieldTypeAsString)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000001_Fieldname)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000004_Fieldtypename)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000012_Structname)
	__Classshape__000003_Classshape0003.Position = __Position__000003_Position_0003

	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000002_Fieldname)
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000003_Fieldtypename)
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000014_TargetMultiplicity)
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000013_Structname)
	__Classshape__000003_Classshape0003.Links = append(__Classshape__000003_Classshape0003.Links, __Link__000004_Middlevertice)
	__Classshape__000004_Classshape0004.Position = __Position__000004_Position_0004

	__Classshape__000004_Classshape0004.Fields = append(__Classshape__000004_Classshape0004.Fields, __Field__000009_Name)
	__Classshape__000004_Classshape0004.Fields = append(__Classshape__000004_Classshape0004.Fields, __Field__000010_Path)
	__Classshape__000004_Classshape0004.Links = append(__Classshape__000004_Classshape0004.Links, __Link__000000_Classdiagrams)
	__Classshape__000004_Classshape0004.Links = append(__Classshape__000004_Classshape0004.Links, __Link__000007_Umlscs)
	__Classshape__000005_Classshape0005.Position = __Position__000005_Position_0005

	__Classshape__000005_Classshape0005.Fields = append(__Classshape__000005_Classshape0005.Fields, __Field__000017_X)
	__Classshape__000005_Classshape0005.Fields = append(__Classshape__000005_Classshape0005.Fields, __Field__000019_Y)
	__Classshape__000006_Classshape0006.Position = __Position__000006_Position_0006

	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000006_Name)
	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000016_X)
	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000021_Y)
	__Classshape__000007_Classshape0007.Position = __Position__000007_Position_0007

	__Classshape__000007_Classshape0007.Fields = append(__Classshape__000007_Classshape0007.Fields, __Field__000008_Name)
	__Classshape__000007_Classshape0007.Links = append(__Classshape__000007_Classshape0007.Links, __Link__000006_States)
	__Classshape__000008_Classshape0008.Position = __Position__000008_Position_0008

	__Classshape__000008_Classshape0008.Fields = append(__Classshape__000008_Classshape0008.Fields, __Field__000018_X)
	__Classshape__000008_Classshape0008.Fields = append(__Classshape__000008_Classshape0008.Fields, __Field__000020_Y)
	__DiagramPackage__000000_gongdoc_diagrams.Classdiagrams = append(__DiagramPackage__000000_gongdoc_diagrams.Classdiagrams, __Classdiagram__000000_DiagramUMLduUML)
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000005_Vertice_0005
	__Link__000001_Classshapes.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000002_Fields.Middlevertice = __Vertice__000001_Vertice_0001
	__Link__000003_Links.Middlevertice = __Vertice__000002_Vertice_0002
	__Link__000004_Middlevertice.Middlevertice = __Vertice__000004_Vertice_0004
	__Link__000005_Position.Middlevertice = __Vertice__000003_Vertice_0003
	__Link__000006_States.Middlevertice = __Vertice__000007_Vertice_0007
	__Link__000007_Umlscs.Middlevertice = __Vertice__000006_Vertice_0006
}
