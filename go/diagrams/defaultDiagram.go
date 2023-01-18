package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_defaultDiagram models.StageStruct
var ___dummy__Time_defaultDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_defaultDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_defaultDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["defaultDiagram"] = defaultDiagramInjection
// }

// defaultDiagramInjection will stage objects of database "defaultDiagram"
func defaultDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_defaultDiagram := (&models.Classdiagram{Name: `defaultDiagram`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0009 := (&models.Classshape{Name: `Classshape0009`}).Stage()
	__Classshape__000001_Classshape0010 := (&models.Classshape{Name: `Classshape0010`}).Stage()
	__Classshape__000002_Classshape0011 := (&models.Classshape{Name: `Classshape0011`}).Stage()
	__Classshape__000003_Classshape0012 := (&models.Classshape{Name: `Classshape0012`}).Stage()
	__Classshape__000004_Classshape0013 := (&models.Classshape{Name: `Classshape0013`}).Stage()
	__Classshape__000005_Classshape0014 := (&models.Classshape{Name: `Classshape0014`}).Stage()
	__Classshape__000006_Classshape0015 := (&models.Classshape{Name: `Classshape0015`}).Stage()
	__Classshape__000007_Classshape0016 := (&models.Classshape{Name: `Classshape0016`}).Stage()
	__Classshape__000008_Classshape0017 := (&models.Classshape{Name: `Classshape0017`}).Stage()
	__Classshape__000009_Classshape0018 := (&models.Classshape{Name: `Classshape0018`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_ := (&models.Field{Name: ``}).Stage()
	__Field__000001_ := (&models.Field{Name: ``}).Stage()
	__Field__000002_ := (&models.Field{Name: ``}).Stage()
	__Field__000003_Activestate := (&models.Field{Name: `Activestate`}).Stage()
	__Field__000004_Fieldname := (&models.Field{Name: `Fieldname`}).Stage()
	__Field__000005_Fieldname := (&models.Field{Name: `Fieldname`}).Stage()
	__Field__000006_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage()
	__Field__000007_Fieldtypename := (&models.Field{Name: `Fieldtypename`}).Stage()
	__Field__000008_Heigth := (&models.Field{Name: `Heigth`}).Stage()
	__Field__000009_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000010_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000015_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000016_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000017_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000018_ReferenceName := (&models.Field{Name: `ReferenceName`}).Stage()
	__Field__000019_Structname := (&models.Field{Name: `Structname`}).Stage()
	__Field__000020_Structname := (&models.Field{Name: `Structname`}).Stage()
	__Field__000021_TargetMultiplicity := (&models.Field{Name: `TargetMultiplicity`}).Stage()
	__Field__000022_Width := (&models.Field{Name: `Width`}).Stage()
	__Field__000023_X := (&models.Field{Name: `X`}).Stage()
	__Field__000024_X := (&models.Field{Name: `X`}).Stage()
	__Field__000025_X := (&models.Field{Name: `X`}).Stage()
	__Field__000026_Y := (&models.Field{Name: `Y`}).Stage()
	__Field__000027_Y := (&models.Field{Name: `Y`}).Stage()
	__Field__000028_Y := (&models.Field{Name: `Y`}).Stage()

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
	__Position__000000_Position_0009 := (&models.Position{Name: `Position-0009`}).Stage()
	__Position__000001_Position_0010 := (&models.Position{Name: `Position-0010`}).Stage()
	__Position__000002_Position_0011 := (&models.Position{Name: `Position-0011`}).Stage()
	__Position__000003_Position_0012 := (&models.Position{Name: `Position-0012`}).Stage()
	__Position__000004_Position_0013 := (&models.Position{Name: `Position-0013`}).Stage()
	__Position__000005_Position_0014 := (&models.Position{Name: `Position-0014`}).Stage()
	__Position__000006_Position_0015 := (&models.Position{Name: `Position-0015`}).Stage()
	__Position__000007_Position_0016 := (&models.Position{Name: `Position-0016`}).Stage()
	__Position__000008_Position_0017 := (&models.Position{Name: `Position-0017`}).Stage()
	__Position__000009_Position_0018 := (&models.Position{Name: `Position-0018`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Classdiagram := (&models.Reference{Name: `Classdiagram`}).Stage()
	__Reference__000001_Classshape := (&models.Reference{Name: `Classshape`}).Stage()
	__Reference__000002_DiagramPackage := (&models.Reference{Name: `DiagramPackage`}).Stage()
	__Reference__000003_Field := (&models.Reference{Name: `Field`}).Stage()
	__Reference__000004_Link := (&models.Reference{Name: `Link`}).Stage()
	__Reference__000005_MultiplicityType := (&models.Reference{Name: `MultiplicityType`}).Stage()
	__Reference__000006_Position := (&models.Reference{Name: `Position`}).Stage()
	__Reference__000007_UmlState := (&models.Reference{Name: `UmlState`}).Stage()
	__Reference__000008_Umlsc := (&models.Reference{Name: `Umlsc`}).Stage()
	__Reference__000009_Vertice := (&models.Reference{Name: `Vertice`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0008 := (&models.Vertice{Name: `Vertice-0008`}).Stage()
	__Vertice__000001_Vertice_0009 := (&models.Vertice{Name: `Vertice-0009`}).Stage()
	__Vertice__000002_Vertice_0010 := (&models.Vertice{Name: `Vertice-0010`}).Stage()
	__Vertice__000003_Vertice_0011 := (&models.Vertice{Name: `Vertice-0011`}).Stage()
	__Vertice__000004_Vertice_0012 := (&models.Vertice{Name: `Vertice-0012`}).Stage()
	__Vertice__000005_Vertice_0013 := (&models.Vertice{Name: `Vertice-0013`}).Stage()
	__Vertice__000006_Vertice_0014 := (&models.Vertice{Name: `Vertice-0014`}).Stage()
	__Vertice__000007_Vertice_0015 := (&models.Vertice{Name: `Vertice-0015`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_defaultDiagram.Name = `defaultDiagram`
	__Classdiagram__000000_defaultDiagram.IsInDrawMode = true

	// Classshape values setup
	__Classshape__000000_Classshape0009.Name = `Classshape0009`
	__Classshape__000000_Classshape0009.ReferenceName = `Classdiagram`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__Classshape__000000_Classshape0009.Identifier = `ref_models.Classdiagram`
	__Classshape__000000_Classshape0009.ShowNbInstances = true
	__Classshape__000000_Classshape0009.NbInstances = 0
	__Classshape__000000_Classshape0009.Width = 240.000000
	__Classshape__000000_Classshape0009.Heigth = 63.000000
	__Classshape__000000_Classshape0009.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0010.Name = `Classshape0010`
	__Classshape__000001_Classshape0010.ReferenceName = `Classshape`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape]
	__Classshape__000001_Classshape0010.Identifier = `ref_models.Classshape`
	__Classshape__000001_Classshape0010.ShowNbInstances = true
	__Classshape__000001_Classshape0010.NbInstances = 0
	__Classshape__000001_Classshape0010.Width = 240.000000
	__Classshape__000001_Classshape0010.Heigth = 138.000000
	__Classshape__000001_Classshape0010.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0011.Name = `Classshape0011`
	__Classshape__000002_Classshape0011.ReferenceName = `Field`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field]
	__Classshape__000002_Classshape0011.Identifier = `ref_models.Field`
	__Classshape__000002_Classshape0011.ShowNbInstances = true
	__Classshape__000002_Classshape0011.NbInstances = 0
	__Classshape__000002_Classshape0011.Width = 240.000000
	__Classshape__000002_Classshape0011.Heigth = 123.000000
	__Classshape__000002_Classshape0011.IsSelected = false

	// Classshape values setup
	__Classshape__000003_Classshape0012.Name = `Classshape0012`
	__Classshape__000003_Classshape0012.ReferenceName = `Link`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Classshape__000003_Classshape0012.Identifier = `ref_models.Link`
	__Classshape__000003_Classshape0012.ShowNbInstances = true
	__Classshape__000003_Classshape0012.NbInstances = 0
	__Classshape__000003_Classshape0012.Width = 240.000000
	__Classshape__000003_Classshape0012.Heigth = 138.000000
	__Classshape__000003_Classshape0012.IsSelected = false

	// Classshape values setup
	__Classshape__000004_Classshape0013.Name = `Classshape0013`
	__Classshape__000004_Classshape0013.ReferenceName = `MultiplicityType`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MultiplicityType]
	__Classshape__000004_Classshape0013.Identifier = `ref_models.MultiplicityType`
	__Classshape__000004_Classshape0013.ShowNbInstances = false
	__Classshape__000004_Classshape0013.NbInstances = 0
	__Classshape__000004_Classshape0013.Width = 240.000000
	__Classshape__000004_Classshape0013.Heigth = 93.000000
	__Classshape__000004_Classshape0013.IsSelected = false

	// Classshape values setup
	__Classshape__000005_Classshape0014.Name = `Classshape0014`
	__Classshape__000005_Classshape0014.ReferenceName = `DiagramPackage`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage]
	__Classshape__000005_Classshape0014.Identifier = `ref_models.DiagramPackage`
	__Classshape__000005_Classshape0014.ShowNbInstances = true
	__Classshape__000005_Classshape0014.NbInstances = 0
	__Classshape__000005_Classshape0014.Width = 240.000000
	__Classshape__000005_Classshape0014.Heigth = 63.000000
	__Classshape__000005_Classshape0014.IsSelected = false

	// Classshape values setup
	__Classshape__000006_Classshape0015.Name = `Classshape0015`
	__Classshape__000006_Classshape0015.ReferenceName = `Position`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position]
	__Classshape__000006_Classshape0015.Identifier = `ref_models.Position`
	__Classshape__000006_Classshape0015.ShowNbInstances = true
	__Classshape__000006_Classshape0015.NbInstances = 0
	__Classshape__000006_Classshape0015.Width = 240.000000
	__Classshape__000006_Classshape0015.Heigth = 93.000000
	__Classshape__000006_Classshape0015.IsSelected = false

	// Classshape values setup
	__Classshape__000007_Classshape0016.Name = `Classshape0016`
	__Classshape__000007_Classshape0016.ReferenceName = `UmlState`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState]
	__Classshape__000007_Classshape0016.Identifier = `ref_models.UmlState`
	__Classshape__000007_Classshape0016.ShowNbInstances = true
	__Classshape__000007_Classshape0016.NbInstances = 0
	__Classshape__000007_Classshape0016.Width = 240.000000
	__Classshape__000007_Classshape0016.Heigth = 93.000000
	__Classshape__000007_Classshape0016.IsSelected = false

	// Classshape values setup
	__Classshape__000008_Classshape0017.Name = `Classshape0017`
	__Classshape__000008_Classshape0017.ReferenceName = `Umlsc`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc]
	__Classshape__000008_Classshape0017.Identifier = `ref_models.Umlsc`
	__Classshape__000008_Classshape0017.ShowNbInstances = true
	__Classshape__000008_Classshape0017.NbInstances = 0
	__Classshape__000008_Classshape0017.Width = 240.000000
	__Classshape__000008_Classshape0017.Heigth = 78.000000
	__Classshape__000008_Classshape0017.IsSelected = false

	// Classshape values setup
	__Classshape__000009_Classshape0018.Name = `Classshape0018`
	__Classshape__000009_Classshape0018.ReferenceName = `Vertice`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice]
	__Classshape__000009_Classshape0018.Identifier = `ref_models.Vertice`
	__Classshape__000009_Classshape0018.ShowNbInstances = true
	__Classshape__000009_Classshape0018.NbInstances = 0
	__Classshape__000009_Classshape0018.Width = 240.000000
	__Classshape__000009_Classshape0018.Heigth = 93.000000
	__Classshape__000009_Classshape0018.IsSelected = false

	// Field values setup
	__Field__000000_.Name = ``
	__Field__000000_.Fieldname = `MANY`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MANY]
	__Field__000000_.Identifier = `ref_models.MANY`
	__Field__000000_.FieldTypeAsString = ``
	__Field__000000_.Structname = `models`
	__Field__000000_.Fieldtypename = ``

	// Field values setup
	__Field__000001_.Name = ``
	__Field__000001_.Fieldname = `ZERO_ONE`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ZERO_ONE]
	__Field__000001_.Identifier = `ref_models.ZERO_ONE`
	__Field__000001_.FieldTypeAsString = ``
	__Field__000001_.Structname = `models`
	__Field__000001_.Fieldtypename = ``

	// Field values setup
	__Field__000002_.Name = ``
	__Field__000002_.Fieldname = `ONE`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ONE]
	__Field__000002_.Identifier = `ref_models.ONE`
	__Field__000002_.FieldTypeAsString = ``
	__Field__000002_.Structname = `models`
	__Field__000002_.Fieldtypename = ``

	// Field values setup
	__Field__000003_Activestate.Name = `Activestate`
	__Field__000003_Activestate.Fieldname = `Activestate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.Activestate]
	__Field__000003_Activestate.Identifier = `ref_models.Umlsc.Activestate`
	__Field__000003_Activestate.FieldTypeAsString = ``
	__Field__000003_Activestate.Structname = `Umlsc`
	__Field__000003_Activestate.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Fieldname.Name = `Fieldname`
	__Field__000004_Fieldname.Fieldname = `Fieldname`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Fieldname]
	__Field__000004_Fieldname.Identifier = `ref_models.Field.Fieldname`
	__Field__000004_Fieldname.FieldTypeAsString = ``
	__Field__000004_Fieldname.Structname = `Field`
	__Field__000004_Fieldname.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Fieldname.Name = `Fieldname`
	__Field__000005_Fieldname.Fieldname = `Fieldname`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Fieldname]
	__Field__000005_Fieldname.Identifier = `ref_models.Link.Fieldname`
	__Field__000005_Fieldname.FieldTypeAsString = ``
	__Field__000005_Fieldname.Structname = `Link`
	__Field__000005_Fieldname.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Fieldtypename.Name = `Fieldtypename`
	__Field__000006_Fieldtypename.Fieldname = `Fieldtypename`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Fieldtypename]
	__Field__000006_Fieldtypename.Identifier = `ref_models.Field.Fieldtypename`
	__Field__000006_Fieldtypename.FieldTypeAsString = ``
	__Field__000006_Fieldtypename.Structname = `Field`
	__Field__000006_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Fieldtypename.Name = `Fieldtypename`
	__Field__000007_Fieldtypename.Fieldname = `Fieldtypename`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Fieldtypename]
	__Field__000007_Fieldtypename.Identifier = `ref_models.Link.Fieldtypename`
	__Field__000007_Fieldtypename.FieldTypeAsString = ``
	__Field__000007_Fieldtypename.Structname = `Link`
	__Field__000007_Fieldtypename.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Heigth.Name = `Heigth`
	__Field__000008_Heigth.Fieldname = `Heigth`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Heigth]
	__Field__000008_Heigth.Identifier = `ref_models.Classshape.Heigth`
	__Field__000008_Heigth.FieldTypeAsString = ``
	__Field__000008_Heigth.Structname = `Classshape`
	__Field__000008_Heigth.Fieldtypename = `float64`

	// Field values setup
	__Field__000009_Name.Name = `Name`
	__Field__000009_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Name]
	__Field__000009_Name.Identifier = `ref_models.Link.Name`
	__Field__000009_Name.FieldTypeAsString = ``
	__Field__000009_Name.Structname = `Link`
	__Field__000009_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Name.Name = `Name`
	__Field__000010_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Name]
	__Field__000010_Name.Identifier = `ref_models.UmlState.Name`
	__Field__000010_Name.FieldTypeAsString = ``
	__Field__000010_Name.Structname = `UmlState`
	__Field__000010_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000011_Name.Name = `Name`
	__Field__000011_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Name]
	__Field__000011_Name.Identifier = `ref_models.Classshape.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `Classshape`
	__Field__000011_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000012_Name.Name = `Name`
	__Field__000012_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.Name]
	__Field__000012_Name.Identifier = `ref_models.Classdiagram.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Classdiagram`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`
	__Field__000013_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Umlsc.Name]
	__Field__000013_Name.Identifier = `ref_models.Umlsc.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Umlsc`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Name.Name = `Name`
	__Field__000014_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.Name]
	__Field__000014_Name.Identifier = `ref_models.Position.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `Position`
	__Field__000014_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000015_Name.Name = `Name`
	__Field__000015_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Name]
	__Field__000015_Name.Identifier = `ref_models.Field.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `Field`
	__Field__000015_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000016_Name.Name = `Name`
	__Field__000016_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.Name]
	__Field__000016_Name.Identifier = `ref_models.Vertice.Name`
	__Field__000016_Name.FieldTypeAsString = ``
	__Field__000016_Name.Structname = `Vertice`
	__Field__000016_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000017_Name.Name = `Name`
	__Field__000017_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Name]
	__Field__000017_Name.Identifier = `ref_models.DiagramPackage.Name`
	__Field__000017_Name.FieldTypeAsString = ``
	__Field__000017_Name.Structname = `DiagramPackage`
	__Field__000017_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000018_ReferenceName.Name = `ReferenceName`
	__Field__000018_ReferenceName.Fieldname = `ReferenceName`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.ReferenceName]
	__Field__000018_ReferenceName.Identifier = `ref_models.Classshape.ReferenceName`
	__Field__000018_ReferenceName.FieldTypeAsString = ``
	__Field__000018_ReferenceName.Structname = `Classshape`
	__Field__000018_ReferenceName.Fieldtypename = `string`

	// Field values setup
	__Field__000019_Structname.Name = `Structname`
	__Field__000019_Structname.Fieldname = `Structname`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Structname]
	__Field__000019_Structname.Identifier = `ref_models.Link.Structname`
	__Field__000019_Structname.FieldTypeAsString = ``
	__Field__000019_Structname.Structname = `Link`
	__Field__000019_Structname.Fieldtypename = `string`

	// Field values setup
	__Field__000020_Structname.Name = `Structname`
	__Field__000020_Structname.Fieldname = `Structname`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Field.Structname]
	__Field__000020_Structname.Identifier = `ref_models.Field.Structname`
	__Field__000020_Structname.FieldTypeAsString = ``
	__Field__000020_Structname.Structname = `Field`
	__Field__000020_Structname.Fieldtypename = `string`

	// Field values setup
	__Field__000021_TargetMultiplicity.Name = `TargetMultiplicity`
	__Field__000021_TargetMultiplicity.Fieldname = `TargetMultiplicity`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.TargetMultiplicity]
	__Field__000021_TargetMultiplicity.Identifier = `ref_models.Link.TargetMultiplicity`
	__Field__000021_TargetMultiplicity.FieldTypeAsString = ``
	__Field__000021_TargetMultiplicity.Structname = `Link`
	__Field__000021_TargetMultiplicity.Fieldtypename = `MultiplicityType`

	// Field values setup
	__Field__000022_Width.Name = `Width`
	__Field__000022_Width.Fieldname = `Width`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classshape.Width]
	__Field__000022_Width.Identifier = `ref_models.Classshape.Width`
	__Field__000022_Width.FieldTypeAsString = ``
	__Field__000022_Width.Structname = `Classshape`
	__Field__000022_Width.Fieldtypename = `float64`

	// Field values setup
	__Field__000023_X.Name = `X`
	__Field__000023_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.X]
	__Field__000023_X.Identifier = `ref_models.Vertice.X`
	__Field__000023_X.FieldTypeAsString = ``
	__Field__000023_X.Structname = `Vertice`
	__Field__000023_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000024_X.Name = `X`
	__Field__000024_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.X]
	__Field__000024_X.Identifier = `ref_models.UmlState.X`
	__Field__000024_X.FieldTypeAsString = ``
	__Field__000024_X.Structname = `UmlState`
	__Field__000024_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000025_X.Name = `X`
	__Field__000025_X.Fieldname = `X`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.X]
	__Field__000025_X.Identifier = `ref_models.Position.X`
	__Field__000025_X.FieldTypeAsString = ``
	__Field__000025_X.Structname = `Position`
	__Field__000025_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000026_Y.Name = `Y`
	__Field__000026_Y.Fieldname = `Y`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Position.Y]
	__Field__000026_Y.Identifier = `ref_models.Position.Y`
	__Field__000026_Y.FieldTypeAsString = ``
	__Field__000026_Y.Structname = `Position`
	__Field__000026_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000027_Y.Name = `Y`
	__Field__000027_Y.Fieldname = `Y`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UmlState.Y]
	__Field__000027_Y.Identifier = `ref_models.UmlState.Y`
	__Field__000027_Y.FieldTypeAsString = ``
	__Field__000027_Y.Structname = `UmlState`
	__Field__000027_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000028_Y.Name = `Y`
	__Field__000028_Y.Fieldname = `Y`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Vertice.Y]
	__Field__000028_Y.Identifier = `ref_models.Vertice.Y`
	__Field__000028_Y.FieldTypeAsString = ``
	__Field__000028_Y.Structname = `Vertice`
	__Field__000028_Y.Fieldtypename = `float64`

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
	__Link__000004_Middlevertice.TargetMultiplicity = models.ZERO_ONE

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
	__Position__000000_Position_0009.X = 40.000000
	__Position__000000_Position_0009.Y = 40.000000
	__Position__000000_Position_0009.Name = `Position-0009`

	// Position values setup
	__Position__000001_Position_0010.X = 340.000000
	__Position__000001_Position_0010.Y = 40.000000
	__Position__000001_Position_0010.Name = `Position-0010`

	// Position values setup
	__Position__000002_Position_0011.X = 640.000000
	__Position__000002_Position_0011.Y = 40.000000
	__Position__000002_Position_0011.Name = `Position-0011`

	// Position values setup
	__Position__000003_Position_0012.X = 1240.000000
	__Position__000003_Position_0012.Y = 40.000000
	__Position__000003_Position_0012.Name = `Position-0012`

	// Position values setup
	__Position__000004_Position_0013.X = 340.000000
	__Position__000004_Position_0013.Y = 500.000000
	__Position__000004_Position_0013.Name = `Position-0013`

	// Position values setup
	__Position__000005_Position_0014.X = 1540.000000
	__Position__000005_Position_0014.Y = 40.000000
	__Position__000005_Position_0014.Name = `Position-0014`

	// Position values setup
	__Position__000006_Position_0015.X = 1840.000000
	__Position__000006_Position_0015.Y = 40.000000
	__Position__000006_Position_0015.Name = `Position-0015`

	// Position values setup
	__Position__000007_Position_0016.X = 2140.000000
	__Position__000007_Position_0016.Y = 40.000000
	__Position__000007_Position_0016.Name = `Position-0016`

	// Position values setup
	__Position__000008_Position_0017.X = 2440.000000
	__Position__000008_Position_0017.Y = 40.000000
	__Position__000008_Position_0017.Name = `Position-0017`

	// Position values setup
	__Position__000009_Position_0018.X = 2740.000000
	__Position__000009_Position_0018.Y = 40.000000
	__Position__000009_Position_0018.Name = `Position-0018`

	// Reference values setup
	__Reference__000000_Classdiagram.Name = `Classdiagram`
	__Reference__000000_Classdiagram.NbInstances = 0
	__Reference__000000_Classdiagram.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Classshape.Name = `Classshape`
	__Reference__000001_Classshape.NbInstances = 0
	__Reference__000001_Classshape.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_DiagramPackage.Name = `DiagramPackage`
	__Reference__000002_DiagramPackage.NbInstances = 0
	__Reference__000002_DiagramPackage.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000003_Field.Name = `Field`
	__Reference__000003_Field.NbInstances = 0
	__Reference__000003_Field.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000004_Link.Name = `Link`
	__Reference__000004_Link.NbInstances = 0
	__Reference__000004_Link.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000005_MultiplicityType.Name = `MultiplicityType`
	__Reference__000005_MultiplicityType.NbInstances = 0
	__Reference__000005_MultiplicityType.Type = models.REFERENCE_GONG_ENUM

	// Reference values setup
	__Reference__000006_Position.Name = `Position`
	__Reference__000006_Position.NbInstances = 0
	__Reference__000006_Position.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000007_UmlState.Name = `UmlState`
	__Reference__000007_UmlState.NbInstances = 0
	__Reference__000007_UmlState.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000008_Umlsc.Name = `Umlsc`
	__Reference__000008_Umlsc.NbInstances = 0
	__Reference__000008_Umlsc.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000009_Vertice.Name = `Vertice`
	__Reference__000009_Vertice.NbInstances = 0
	__Reference__000009_Vertice.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0008.X = 290.000000
	__Vertice__000000_Vertice_0008.Y = 200.000000
	__Vertice__000000_Vertice_0008.Name = `Vertice-0008`

	// Vertice values setup
	__Vertice__000001_Vertice_0009.X = 590.000000
	__Vertice__000001_Vertice_0009.Y = 200.000000
	__Vertice__000001_Vertice_0009.Name = `Vertice-0009`

	// Vertice values setup
	__Vertice__000002_Vertice_0010.X = 590.000000
	__Vertice__000002_Vertice_0010.Y = 250.000000
	__Vertice__000002_Vertice_0010.Name = `Vertice-0010`

	// Vertice values setup
	__Vertice__000003_Vertice_0011.X = 590.000000
	__Vertice__000003_Vertice_0011.Y = 300.000000
	__Vertice__000003_Vertice_0011.Name = `Vertice-0011`

	// Vertice values setup
	__Vertice__000004_Vertice_0012.X = 1490.000000
	__Vertice__000004_Vertice_0012.Y = 200.000000
	__Vertice__000004_Vertice_0012.Name = `Vertice-0012`

	// Vertice values setup
	__Vertice__000005_Vertice_0013.X = 1790.000000
	__Vertice__000005_Vertice_0013.Y = 200.000000
	__Vertice__000005_Vertice_0013.Name = `Vertice-0013`

	// Vertice values setup
	__Vertice__000006_Vertice_0014.X = 1790.000000
	__Vertice__000006_Vertice_0014.Y = 250.000000
	__Vertice__000006_Vertice_0014.Name = `Vertice-0014`

	// Vertice values setup
	__Vertice__000007_Vertice_0015.X = 2690.000000
	__Vertice__000007_Vertice_0015.Y = 200.000000
	__Vertice__000007_Vertice_0015.Name = `Vertice-0015`

	// Setup of pointers
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000000_Classshape0009)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000001_Classshape0010)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000002_Classshape0011)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000003_Classshape0012)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000004_Classshape0013)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000005_Classshape0014)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000006_Classshape0015)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000007_Classshape0016)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000008_Classshape0017)
	__Classdiagram__000000_defaultDiagram.Classshapes = append(__Classdiagram__000000_defaultDiagram.Classshapes, __Classshape__000009_Classshape0018)
	__Classshape__000000_Classshape0009.Position = __Position__000000_Position_0009
	__Classshape__000000_Classshape0009.Reference = __Reference__000000_Classdiagram
	__Classshape__000000_Classshape0009.Fields = append(__Classshape__000000_Classshape0009.Fields, __Field__000012_Name)
	__Classshape__000000_Classshape0009.Links = append(__Classshape__000000_Classshape0009.Links, __Link__000001_Classshapes)
	__Classshape__000001_Classshape0010.Position = __Position__000001_Position_0010
	__Classshape__000001_Classshape0010.Reference = __Reference__000001_Classshape
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000008_Heigth)
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000011_Name)
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000018_ReferenceName)
	__Classshape__000001_Classshape0010.Fields = append(__Classshape__000001_Classshape0010.Fields, __Field__000022_Width)
	__Classshape__000001_Classshape0010.Links = append(__Classshape__000001_Classshape0010.Links, __Link__000002_Fields)
	__Classshape__000001_Classshape0010.Links = append(__Classshape__000001_Classshape0010.Links, __Link__000003_Links)
	__Classshape__000001_Classshape0010.Links = append(__Classshape__000001_Classshape0010.Links, __Link__000005_Position)
	__Classshape__000002_Classshape0011.Position = __Position__000002_Position_0011
	__Classshape__000002_Classshape0011.Reference = __Reference__000003_Field
	__Classshape__000002_Classshape0011.Fields = append(__Classshape__000002_Classshape0011.Fields, __Field__000004_Fieldname)
	__Classshape__000002_Classshape0011.Fields = append(__Classshape__000002_Classshape0011.Fields, __Field__000006_Fieldtypename)
	__Classshape__000002_Classshape0011.Fields = append(__Classshape__000002_Classshape0011.Fields, __Field__000015_Name)
	__Classshape__000002_Classshape0011.Fields = append(__Classshape__000002_Classshape0011.Fields, __Field__000020_Structname)
	__Classshape__000003_Classshape0012.Position = __Position__000003_Position_0012
	__Classshape__000003_Classshape0012.Reference = __Reference__000004_Link
	__Classshape__000003_Classshape0012.Fields = append(__Classshape__000003_Classshape0012.Fields, __Field__000005_Fieldname)
	__Classshape__000003_Classshape0012.Fields = append(__Classshape__000003_Classshape0012.Fields, __Field__000007_Fieldtypename)
	__Classshape__000003_Classshape0012.Fields = append(__Classshape__000003_Classshape0012.Fields, __Field__000021_TargetMultiplicity)
	__Classshape__000003_Classshape0012.Fields = append(__Classshape__000003_Classshape0012.Fields, __Field__000009_Name)
	__Classshape__000003_Classshape0012.Fields = append(__Classshape__000003_Classshape0012.Fields, __Field__000019_Structname)
	__Classshape__000003_Classshape0012.Links = append(__Classshape__000003_Classshape0012.Links, __Link__000004_Middlevertice)
	__Classshape__000004_Classshape0013.Position = __Position__000004_Position_0013
	__Classshape__000004_Classshape0013.Reference = __Reference__000005_MultiplicityType
	__Classshape__000004_Classshape0013.Fields = append(__Classshape__000004_Classshape0013.Fields, __Field__000000_)
	__Classshape__000004_Classshape0013.Fields = append(__Classshape__000004_Classshape0013.Fields, __Field__000002_)
	__Classshape__000004_Classshape0013.Fields = append(__Classshape__000004_Classshape0013.Fields, __Field__000001_)
	__Classshape__000005_Classshape0014.Position = __Position__000005_Position_0014
	__Classshape__000005_Classshape0014.Reference = __Reference__000002_DiagramPackage
	__Classshape__000005_Classshape0014.Fields = append(__Classshape__000005_Classshape0014.Fields, __Field__000017_Name)
	__Classshape__000005_Classshape0014.Links = append(__Classshape__000005_Classshape0014.Links, __Link__000000_Classdiagrams)
	__Classshape__000005_Classshape0014.Links = append(__Classshape__000005_Classshape0014.Links, __Link__000007_Umlscs)
	__Classshape__000006_Classshape0015.Position = __Position__000006_Position_0015
	__Classshape__000006_Classshape0015.Reference = __Reference__000006_Position
	__Classshape__000006_Classshape0015.Fields = append(__Classshape__000006_Classshape0015.Fields, __Field__000014_Name)
	__Classshape__000006_Classshape0015.Fields = append(__Classshape__000006_Classshape0015.Fields, __Field__000025_X)
	__Classshape__000006_Classshape0015.Fields = append(__Classshape__000006_Classshape0015.Fields, __Field__000026_Y)
	__Classshape__000007_Classshape0016.Position = __Position__000007_Position_0016
	__Classshape__000007_Classshape0016.Reference = __Reference__000007_UmlState
	__Classshape__000007_Classshape0016.Fields = append(__Classshape__000007_Classshape0016.Fields, __Field__000010_Name)
	__Classshape__000007_Classshape0016.Fields = append(__Classshape__000007_Classshape0016.Fields, __Field__000024_X)
	__Classshape__000007_Classshape0016.Fields = append(__Classshape__000007_Classshape0016.Fields, __Field__000027_Y)
	__Classshape__000008_Classshape0017.Position = __Position__000008_Position_0017
	__Classshape__000008_Classshape0017.Reference = __Reference__000008_Umlsc
	__Classshape__000008_Classshape0017.Fields = append(__Classshape__000008_Classshape0017.Fields, __Field__000003_Activestate)
	__Classshape__000008_Classshape0017.Fields = append(__Classshape__000008_Classshape0017.Fields, __Field__000013_Name)
	__Classshape__000008_Classshape0017.Links = append(__Classshape__000008_Classshape0017.Links, __Link__000006_States)
	__Classshape__000009_Classshape0018.Position = __Position__000009_Position_0018
	__Classshape__000009_Classshape0018.Reference = __Reference__000009_Vertice
	__Classshape__000009_Classshape0018.Fields = append(__Classshape__000009_Classshape0018.Fields, __Field__000016_Name)
	__Classshape__000009_Classshape0018.Fields = append(__Classshape__000009_Classshape0018.Fields, __Field__000023_X)
	__Classshape__000009_Classshape0018.Fields = append(__Classshape__000009_Classshape0018.Fields, __Field__000028_Y)
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000005_Vertice_0013
	__Link__000001_Classshapes.Middlevertice = __Vertice__000000_Vertice_0008
	__Link__000002_Fields.Middlevertice = __Vertice__000001_Vertice_0009
	__Link__000003_Links.Middlevertice = __Vertice__000002_Vertice_0010
	__Link__000004_Middlevertice.Middlevertice = __Vertice__000004_Vertice_0012
	__Link__000005_Position.Middlevertice = __Vertice__000003_Vertice_0011
	__Link__000006_States.Middlevertice = __Vertice__000007_Vertice_0015
	__Link__000007_Umlscs.Middlevertice = __Vertice__000006_Vertice_0014
}


