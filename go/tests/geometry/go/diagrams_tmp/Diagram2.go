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
	__Classdiagram__000001_Diagram3 := (&models.Classdiagram{Name: `Diagram3`}).Stage()
	__Classdiagram__000002_Diagram4 := (&models.Classdiagram{Name: `Diagram4`}).Stage()
	__Classdiagram__000003_DiagramEnums := (&models.Classdiagram{Name: `DiagramEnums`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0000 := (&models.Classshape{Name: `Classshape0000`}).Stage()
	__Classshape__000001_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()
	__Classshape__000002_Classshape0002 := (&models.Classshape{Name: `Classshape0002`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()
	__Link__000001_Points := (&models.Link{Name: `Points`}).Stage()
	__Link__000002_Start := (&models.Link{Name: `Start`}).Stage()

	// Declarations of staged instances of Node
	__Node__000000_AFTER := (&models.Node{Name: `AFTER`}).Stage()
	__Node__000001_BEFORE := (&models.Node{Name: `BEFORE`}).Stage()
	__Node__000002_CONTINUOUS := (&models.Node{Name: `CONTINUOUS`}).Stage()
	__Node__000003_CreatedAt := (&models.Node{Name: `CreatedAt`}).Stage()
	__Node__000004_CreationDate := (&models.Node{Name: `CreationDate`}).Stage()
	__Node__000005_DOTTED := (&models.Node{Name: `DOTTED`}).Stage()
	__Node__000006_Diagram2 := (&models.Node{Name: `Diagram2`}).Stage()
	__Node__000007_Diagram3 := (&models.Node{Name: `Diagram3`}).Stage()
	__Node__000008_Diagram4 := (&models.Node{Name: `Diagram4`}).Stage()
	__Node__000009_DiagramEnums := (&models.Node{Name: `DiagramEnums`}).Stage()
	__Node__000010_End := (&models.Node{Name: `End`}).Stage()
	__Node__000011_JourneyTime := (&models.Node{Name: `JourneyTime`}).Stage()
	__Node__000012_Line := (&models.Node{Name: `Line`}).Stage()
	__Node__000013_Line := (&models.Node{Name: `Line`}).Stage()
	__Node__000014_LineType := (&models.Node{Name: `LineType`}).Stage()
	__Node__000015_LongNodeOnModels := (&models.Node{Name: `LongNodeOnModels`}).Stage()
	__Node__000016_Name := (&models.Node{Name: `Name`}).Stage()
	__Node__000017_Name := (&models.Node{Name: `Name`}).Stage()
	__Node__000018_Name := (&models.Node{Name: `Name`}).Stage()
	__Node__000019_Name := (&models.Node{Name: `Name`}).Stage()
	__Node__000020_Name := (&models.Node{Name: `Name`}).Stage()
	__Node__000021_Point := (&models.Node{Name: `Point`}).Stage()
	__Node__000022_Point := (&models.Node{Name: `Point`}).Stage()
	__Node__000023_PointExclusiveSet := (&models.Node{Name: `PointExclusiveSet`}).Stage()
	__Node__000024_PointNonExclusiveSet := (&models.Node{Name: `PointNonExclusiveSet`}).Stage()
	__Node__000025_PointUse := (&models.Node{Name: `PointUse`}).Stage()
	__Node__000026_Points := (&models.Node{Name: `Points`}).Stage()
	__Node__000027_Points := (&models.Node{Name: `Points`}).Stage()
	__Node__000028_Points := (&models.Node{Name: `Points`}).Stage()
	__Node__000029_ShortNodeOnModels := (&models.Node{Name: `ShortNodeOnModels`}).Stage()
	__Node__000030_SimulationStage := (&models.Node{Name: `SimulationStage`}).Stage()
	__Node__000031_Start := (&models.Node{Name: `Start`}).Stage()
	__Node__000032_Type := (&models.Node{Name: `Type`}).Stage()
	__Node__000033_VeryLongLongLongLongLongLongField := (&models.Node{Name: `VeryLongLongLongLongLongLongField`}).Stage()
	__Node__000034_X := (&models.Node{Name: `X`}).Stage()
	__Node__000035_Y := (&models.Node{Name: `Y`}).Stage()
	__Node__000036_Z := (&models.Node{Name: `Z`}).Stage()
	__Node__000037_class_diagrams := (&models.Node{Name: `class diagrams`}).Stage()
	__Node__000038_gongenums := (&models.Node{Name: `gongenums`}).Stage()
	__Node__000039_gongstructs := (&models.Node{Name: `gongstructs`}).Stage()
	__Node__000040_notes := (&models.Node{Name: `notes`}).Stage()

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()
	__Position__000002_Position_0002 := (&models.Position{Name: `Position-0002`}).Stage()

	// Declarations of staged instances of Reference

	// Declarations of staged instances of Tree
	__Tree__000000_gong := (&models.Tree{Name: `gong`}).Stage()
	__Tree__000001_gongdoc := (&models.Tree{Name: `gongdoc`}).Stage()

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage()
	__Vertice__000001_Vertice_0001 := (&models.Vertice{Name: `Vertice-0001`}).Stage()
	__Vertice__000002_Vertice_0002 := (&models.Vertice{Name: `Vertice-0002`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Diagram2.Name = `Diagram2`
	__Classdiagram__000000_Diagram2.IsInDrawMode = true

	// Classdiagram values setup
	__Classdiagram__000001_Diagram3.Name = `Diagram3`
	__Classdiagram__000001_Diagram3.IsInDrawMode = false

	// Classdiagram values setup
	__Classdiagram__000002_Diagram4.Name = `Diagram4`
	__Classdiagram__000002_Diagram4.IsInDrawMode = false

	// Classdiagram values setup
	__Classdiagram__000003_DiagramEnums.Name = `DiagramEnums`
	__Classdiagram__000003_DiagramEnums.IsInDrawMode = false

	// Classshape values setup
	__Classshape__000000_Classshape0000.Name = `Classshape0000`
	__Classshape__000000_Classshape0000.ReferenceName = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line]
	__Classshape__000000_Classshape0000.Identifier = `Ref_models.Line`
	__Classshape__000000_Classshape0000.ShowNbInstances = true
	__Classshape__000000_Classshape0000.NbInstances = 59
	__Classshape__000000_Classshape0000.Width = 240.000000
	__Classshape__000000_Classshape0000.Heigth = 78.000000
	__Classshape__000000_Classshape0000.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0001.Name = `Classshape0001`
	__Classshape__000001_Classshape0001.ReferenceName = `Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Point]
	__Classshape__000001_Classshape0001.Identifier = `Ref_models.Point`
	__Classshape__000001_Classshape0001.ShowNbInstances = true
	__Classshape__000001_Classshape0001.NbInstances = 81
	__Classshape__000001_Classshape0001.Width = 240.000000
	__Classshape__000001_Classshape0001.Heigth = 63.000000
	__Classshape__000001_Classshape0001.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0002.Name = `Classshape0002`
	__Classshape__000002_Classshape0002.ReferenceName = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet]
	__Classshape__000002_Classshape0002.Identifier = `Ref_models.PointExclusiveSet`
	__Classshape__000002_Classshape0002.ShowNbInstances = true
	__Classshape__000002_Classshape0002.NbInstances = 87
	__Classshape__000002_Classshape0002.Width = 240.000000
	__Classshape__000002_Classshape0002.Heigth = 63.000000
	__Classshape__000002_Classshape0002.IsSelected = false

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`
	__Field__000000_CreationDate.Fieldname = `CreationDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.CreationDate]
	__Field__000000_CreationDate.Identifier = `Ref_models.Line.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Fieldname = `End`
	__Link__000000_End.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.End]
	__Link__000000_End.Identifier = `Ref_models.Line.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Points.Name = `Points`
	__Link__000001_Points.Fieldname = `Points`
	__Link__000001_Points.Structname = `PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.PointExclusiveSet.Points]
	__Link__000001_Points.Identifier = `Ref_models.PointExclusiveSet.Points`
	__Link__000001_Points.Fieldtypename = `Point`
	__Link__000001_Points.TargetMultiplicity = models.MANY
	__Link__000001_Points.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_Start.Name = `Start`
	__Link__000002_Start.Fieldname = `Start`
	__Link__000002_Start.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [Ref_models.Line.Start]
	__Link__000002_Start.Identifier = `Ref_models.Line.Start`
	__Link__000002_Start.Fieldtypename = `Point`
	__Link__000002_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_Start.SourceMultiplicity = models.MANY

	// Node values setup
	__Node__000000_AFTER.Name = `AFTER`
	__Node__000000_AFTER.Type = models.GONG_ENUM_VALUE
	__Node__000000_AFTER.IsExpanded = false
	__Node__000000_AFTER.HasCheckboxButton = true
	__Node__000000_AFTER.IsChecked = false
	__Node__000000_AFTER.IsCheckboxDisabled = true
	__Node__000000_AFTER.HasAddChildButton = false
	__Node__000000_AFTER.HasEditButton = false
	__Node__000000_AFTER.IsInEditMode = false
	__Node__000000_AFTER.HasDrawButton = false
	__Node__000000_AFTER.HasDrawOffButton = false
	__Node__000000_AFTER.IsInDrawMode = false
	__Node__000000_AFTER.IsSaved = false
	__Node__000000_AFTER.HasDeleteButton = false

	// Node values setup
	__Node__000001_BEFORE.Name = `BEFORE`
	__Node__000001_BEFORE.Type = models.GONG_ENUM_VALUE
	__Node__000001_BEFORE.IsExpanded = false
	__Node__000001_BEFORE.HasCheckboxButton = true
	__Node__000001_BEFORE.IsChecked = false
	__Node__000001_BEFORE.IsCheckboxDisabled = true
	__Node__000001_BEFORE.HasAddChildButton = false
	__Node__000001_BEFORE.HasEditButton = false
	__Node__000001_BEFORE.IsInEditMode = false
	__Node__000001_BEFORE.HasDrawButton = false
	__Node__000001_BEFORE.HasDrawOffButton = false
	__Node__000001_BEFORE.IsInDrawMode = false
	__Node__000001_BEFORE.IsSaved = false
	__Node__000001_BEFORE.HasDeleteButton = false

	// Node values setup
	__Node__000002_CONTINUOUS.Name = `CONTINUOUS`
	__Node__000002_CONTINUOUS.Type = models.GONG_ENUM_VALUE
	__Node__000002_CONTINUOUS.IsExpanded = false
	__Node__000002_CONTINUOUS.HasCheckboxButton = true
	__Node__000002_CONTINUOUS.IsChecked = false
	__Node__000002_CONTINUOUS.IsCheckboxDisabled = true
	__Node__000002_CONTINUOUS.HasAddChildButton = false
	__Node__000002_CONTINUOUS.HasEditButton = false
	__Node__000002_CONTINUOUS.IsInEditMode = false
	__Node__000002_CONTINUOUS.HasDrawButton = false
	__Node__000002_CONTINUOUS.HasDrawOffButton = false
	__Node__000002_CONTINUOUS.IsInDrawMode = false
	__Node__000002_CONTINUOUS.IsSaved = false
	__Node__000002_CONTINUOUS.HasDeleteButton = false

	// Node values setup
	__Node__000003_CreatedAt.Name = `CreatedAt`
	__Node__000003_CreatedAt.Type = models.GONG_STRUCT_FIELD
	__Node__000003_CreatedAt.IsExpanded = false
	__Node__000003_CreatedAt.HasCheckboxButton = true
	__Node__000003_CreatedAt.IsChecked = false
	__Node__000003_CreatedAt.IsCheckboxDisabled = false
	__Node__000003_CreatedAt.HasAddChildButton = false
	__Node__000003_CreatedAt.HasEditButton = false
	__Node__000003_CreatedAt.IsInEditMode = false
	__Node__000003_CreatedAt.HasDrawButton = false
	__Node__000003_CreatedAt.HasDrawOffButton = false
	__Node__000003_CreatedAt.IsInDrawMode = false
	__Node__000003_CreatedAt.IsSaved = false
	__Node__000003_CreatedAt.HasDeleteButton = false

	// Node values setup
	__Node__000004_CreationDate.Name = `CreationDate`
	__Node__000004_CreationDate.Type = models.GONG_STRUCT_FIELD
	__Node__000004_CreationDate.IsExpanded = false
	__Node__000004_CreationDate.HasCheckboxButton = true
	__Node__000004_CreationDate.IsChecked = true
	__Node__000004_CreationDate.IsCheckboxDisabled = false
	__Node__000004_CreationDate.HasAddChildButton = false
	__Node__000004_CreationDate.HasEditButton = false
	__Node__000004_CreationDate.IsInEditMode = false
	__Node__000004_CreationDate.HasDrawButton = false
	__Node__000004_CreationDate.HasDrawOffButton = false
	__Node__000004_CreationDate.IsInDrawMode = false
	__Node__000004_CreationDate.IsSaved = false
	__Node__000004_CreationDate.HasDeleteButton = false

	// Node values setup
	__Node__000005_DOTTED.Name = `DOTTED`
	__Node__000005_DOTTED.Type = models.GONG_ENUM_VALUE
	__Node__000005_DOTTED.IsExpanded = false
	__Node__000005_DOTTED.HasCheckboxButton = true
	__Node__000005_DOTTED.IsChecked = false
	__Node__000005_DOTTED.IsCheckboxDisabled = true
	__Node__000005_DOTTED.HasAddChildButton = false
	__Node__000005_DOTTED.HasEditButton = false
	__Node__000005_DOTTED.IsInEditMode = false
	__Node__000005_DOTTED.HasDrawButton = false
	__Node__000005_DOTTED.HasDrawOffButton = false
	__Node__000005_DOTTED.IsInDrawMode = false
	__Node__000005_DOTTED.IsSaved = false
	__Node__000005_DOTTED.HasDeleteButton = false

	// Node values setup
	__Node__000006_Diagram2.Name = `Diagram2`
	__Node__000006_Diagram2.Type = models.CLASS_DIAGRAM
	__Node__000006_Diagram2.IsExpanded = false
	__Node__000006_Diagram2.HasCheckboxButton = true
	__Node__000006_Diagram2.IsChecked = true
	__Node__000006_Diagram2.IsCheckboxDisabled = false
	__Node__000006_Diagram2.HasAddChildButton = false
	__Node__000006_Diagram2.HasEditButton = false
	__Node__000006_Diagram2.IsInEditMode = false
	__Node__000006_Diagram2.HasDrawButton = false
	__Node__000006_Diagram2.HasDrawOffButton = false
	__Node__000006_Diagram2.IsInDrawMode = true
	__Node__000006_Diagram2.IsSaved = true
	__Node__000006_Diagram2.HasDeleteButton = false

	// Node values setup
	__Node__000007_Diagram3.Name = `Diagram3`
	__Node__000007_Diagram3.Type = models.CLASS_DIAGRAM
	__Node__000007_Diagram3.IsExpanded = false
	__Node__000007_Diagram3.HasCheckboxButton = true
	__Node__000007_Diagram3.IsChecked = false
	__Node__000007_Diagram3.IsCheckboxDisabled = false
	__Node__000007_Diagram3.HasAddChildButton = false
	__Node__000007_Diagram3.HasEditButton = false
	__Node__000007_Diagram3.IsInEditMode = false
	__Node__000007_Diagram3.HasDrawButton = false
	__Node__000007_Diagram3.HasDrawOffButton = false
	__Node__000007_Diagram3.IsInDrawMode = false
	__Node__000007_Diagram3.IsSaved = false
	__Node__000007_Diagram3.HasDeleteButton = false

	// Node values setup
	__Node__000008_Diagram4.Name = `Diagram4`
	__Node__000008_Diagram4.Type = models.CLASS_DIAGRAM
	__Node__000008_Diagram4.IsExpanded = false
	__Node__000008_Diagram4.HasCheckboxButton = true
	__Node__000008_Diagram4.IsChecked = false
	__Node__000008_Diagram4.IsCheckboxDisabled = false
	__Node__000008_Diagram4.HasAddChildButton = false
	__Node__000008_Diagram4.HasEditButton = false
	__Node__000008_Diagram4.IsInEditMode = false
	__Node__000008_Diagram4.HasDrawButton = false
	__Node__000008_Diagram4.HasDrawOffButton = false
	__Node__000008_Diagram4.IsInDrawMode = false
	__Node__000008_Diagram4.IsSaved = false
	__Node__000008_Diagram4.HasDeleteButton = false

	// Node values setup
	__Node__000009_DiagramEnums.Name = `DiagramEnums`
	__Node__000009_DiagramEnums.Type = models.CLASS_DIAGRAM
	__Node__000009_DiagramEnums.IsExpanded = false
	__Node__000009_DiagramEnums.HasCheckboxButton = true
	__Node__000009_DiagramEnums.IsChecked = false
	__Node__000009_DiagramEnums.IsCheckboxDisabled = false
	__Node__000009_DiagramEnums.HasAddChildButton = false
	__Node__000009_DiagramEnums.HasEditButton = false
	__Node__000009_DiagramEnums.IsInEditMode = false
	__Node__000009_DiagramEnums.HasDrawButton = false
	__Node__000009_DiagramEnums.HasDrawOffButton = false
	__Node__000009_DiagramEnums.IsInDrawMode = false
	__Node__000009_DiagramEnums.IsSaved = false
	__Node__000009_DiagramEnums.HasDeleteButton = false

	// Node values setup
	__Node__000010_End.Name = `End`
	__Node__000010_End.Type = models.GONG_STRUCT_FIELD
	__Node__000010_End.IsExpanded = false
	__Node__000010_End.HasCheckboxButton = true
	__Node__000010_End.IsChecked = true
	__Node__000010_End.IsCheckboxDisabled = false
	__Node__000010_End.HasAddChildButton = false
	__Node__000010_End.HasEditButton = false
	__Node__000010_End.IsInEditMode = false
	__Node__000010_End.HasDrawButton = false
	__Node__000010_End.HasDrawOffButton = false
	__Node__000010_End.IsInDrawMode = false
	__Node__000010_End.IsSaved = false
	__Node__000010_End.HasDeleteButton = false

	// Node values setup
	__Node__000011_JourneyTime.Name = `JourneyTime`
	__Node__000011_JourneyTime.Type = models.GONG_STRUCT_FIELD
	__Node__000011_JourneyTime.IsExpanded = false
	__Node__000011_JourneyTime.HasCheckboxButton = true
	__Node__000011_JourneyTime.IsChecked = false
	__Node__000011_JourneyTime.IsCheckboxDisabled = false
	__Node__000011_JourneyTime.HasAddChildButton = false
	__Node__000011_JourneyTime.HasEditButton = false
	__Node__000011_JourneyTime.IsInEditMode = false
	__Node__000011_JourneyTime.HasDrawButton = false
	__Node__000011_JourneyTime.HasDrawOffButton = false
	__Node__000011_JourneyTime.IsInDrawMode = false
	__Node__000011_JourneyTime.IsSaved = false
	__Node__000011_JourneyTime.HasDeleteButton = false

	// Node values setup
	__Node__000012_Line.Name = `Line`
	__Node__000012_Line.Type = models.GONG_STRUCT
	__Node__000012_Line.IsExpanded = false
	__Node__000012_Line.HasCheckboxButton = true
	__Node__000012_Line.IsChecked = true
	__Node__000012_Line.IsCheckboxDisabled = false
	__Node__000012_Line.HasAddChildButton = false
	__Node__000012_Line.HasEditButton = false
	__Node__000012_Line.IsInEditMode = false
	__Node__000012_Line.HasDrawButton = false
	__Node__000012_Line.HasDrawOffButton = false
	__Node__000012_Line.IsInDrawMode = false
	__Node__000012_Line.IsSaved = false
	__Node__000012_Line.HasDeleteButton = false

	// Node values setup
	__Node__000013_Line.Name = `Line`
	__Node__000013_Line.Type = models.GONG_NOTE_LINK
	__Node__000013_Line.IsExpanded = false
	__Node__000013_Line.HasCheckboxButton = true
	__Node__000013_Line.IsChecked = false
	__Node__000013_Line.IsCheckboxDisabled = false
	__Node__000013_Line.HasAddChildButton = false
	__Node__000013_Line.HasEditButton = false
	__Node__000013_Line.IsInEditMode = false
	__Node__000013_Line.HasDrawButton = false
	__Node__000013_Line.HasDrawOffButton = false
	__Node__000013_Line.IsInDrawMode = false
	__Node__000013_Line.IsSaved = false
	__Node__000013_Line.HasDeleteButton = false

	// Node values setup
	__Node__000014_LineType.Name = `LineType`
	__Node__000014_LineType.Type = models.GONG_ENUM
	__Node__000014_LineType.IsExpanded = false
	__Node__000014_LineType.HasCheckboxButton = true
	__Node__000014_LineType.IsChecked = false
	__Node__000014_LineType.IsCheckboxDisabled = false
	__Node__000014_LineType.HasAddChildButton = false
	__Node__000014_LineType.HasEditButton = false
	__Node__000014_LineType.IsInEditMode = false
	__Node__000014_LineType.HasDrawButton = false
	__Node__000014_LineType.HasDrawOffButton = false
	__Node__000014_LineType.IsInDrawMode = false
	__Node__000014_LineType.IsSaved = false
	__Node__000014_LineType.HasDeleteButton = false

	// Node values setup
	__Node__000015_LongNodeOnModels.Name = `LongNodeOnModels`
	__Node__000015_LongNodeOnModels.Type = models.GONG_NOTE
	__Node__000015_LongNodeOnModels.IsExpanded = true
	__Node__000015_LongNodeOnModels.HasCheckboxButton = true
	__Node__000015_LongNodeOnModels.IsChecked = false
	__Node__000015_LongNodeOnModels.IsCheckboxDisabled = false
	__Node__000015_LongNodeOnModels.HasAddChildButton = false
	__Node__000015_LongNodeOnModels.HasEditButton = false
	__Node__000015_LongNodeOnModels.IsInEditMode = false
	__Node__000015_LongNodeOnModels.HasDrawButton = false
	__Node__000015_LongNodeOnModels.HasDrawOffButton = false
	__Node__000015_LongNodeOnModels.IsInDrawMode = false
	__Node__000015_LongNodeOnModels.IsSaved = false
	__Node__000015_LongNodeOnModels.HasDeleteButton = false

	// Node values setup
	__Node__000016_Name.Name = `Name`
	__Node__000016_Name.Type = models.GONG_STRUCT_FIELD
	__Node__000016_Name.IsExpanded = false
	__Node__000016_Name.HasCheckboxButton = true
	__Node__000016_Name.IsChecked = false
	__Node__000016_Name.IsCheckboxDisabled = false
	__Node__000016_Name.HasAddChildButton = false
	__Node__000016_Name.HasEditButton = false
	__Node__000016_Name.IsInEditMode = false
	__Node__000016_Name.HasDrawButton = false
	__Node__000016_Name.HasDrawOffButton = false
	__Node__000016_Name.IsInDrawMode = false
	__Node__000016_Name.IsSaved = false
	__Node__000016_Name.HasDeleteButton = false

	// Node values setup
	__Node__000017_Name.Name = `Name`
	__Node__000017_Name.Type = models.GONG_STRUCT_FIELD
	__Node__000017_Name.IsExpanded = false
	__Node__000017_Name.HasCheckboxButton = true
	__Node__000017_Name.IsChecked = false
	__Node__000017_Name.IsCheckboxDisabled = false
	__Node__000017_Name.HasAddChildButton = false
	__Node__000017_Name.HasEditButton = false
	__Node__000017_Name.IsInEditMode = false
	__Node__000017_Name.HasDrawButton = false
	__Node__000017_Name.HasDrawOffButton = false
	__Node__000017_Name.IsInDrawMode = false
	__Node__000017_Name.IsSaved = false
	__Node__000017_Name.HasDeleteButton = false

	// Node values setup
	__Node__000018_Name.Name = `Name`
	__Node__000018_Name.Type = models.GONG_STRUCT_FIELD
	__Node__000018_Name.IsExpanded = false
	__Node__000018_Name.HasCheckboxButton = true
	__Node__000018_Name.IsChecked = false
	__Node__000018_Name.IsCheckboxDisabled = false
	__Node__000018_Name.HasAddChildButton = false
	__Node__000018_Name.HasEditButton = false
	__Node__000018_Name.IsInEditMode = false
	__Node__000018_Name.HasDrawButton = false
	__Node__000018_Name.HasDrawOffButton = false
	__Node__000018_Name.IsInDrawMode = false
	__Node__000018_Name.IsSaved = false
	__Node__000018_Name.HasDeleteButton = false

	// Node values setup
	__Node__000019_Name.Name = `Name`
	__Node__000019_Name.Type = models.GONG_STRUCT_FIELD
	__Node__000019_Name.IsExpanded = false
	__Node__000019_Name.HasCheckboxButton = true
	__Node__000019_Name.IsChecked = false
	__Node__000019_Name.IsCheckboxDisabled = true
	__Node__000019_Name.HasAddChildButton = false
	__Node__000019_Name.HasEditButton = false
	__Node__000019_Name.IsInEditMode = false
	__Node__000019_Name.HasDrawButton = false
	__Node__000019_Name.HasDrawOffButton = false
	__Node__000019_Name.IsInDrawMode = false
	__Node__000019_Name.IsSaved = false
	__Node__000019_Name.HasDeleteButton = false

	// Node values setup
	__Node__000020_Name.Name = `Name`
	__Node__000020_Name.Type = models.GONG_STRUCT_FIELD
	__Node__000020_Name.IsExpanded = false
	__Node__000020_Name.HasCheckboxButton = true
	__Node__000020_Name.IsChecked = false
	__Node__000020_Name.IsCheckboxDisabled = true
	__Node__000020_Name.HasAddChildButton = false
	__Node__000020_Name.HasEditButton = false
	__Node__000020_Name.IsInEditMode = false
	__Node__000020_Name.HasDrawButton = false
	__Node__000020_Name.HasDrawOffButton = false
	__Node__000020_Name.IsInDrawMode = false
	__Node__000020_Name.IsSaved = false
	__Node__000020_Name.HasDeleteButton = false

	// Node values setup
	__Node__000021_Point.Name = `Point`
	__Node__000021_Point.Type = models.GONG_NOTE_LINK
	__Node__000021_Point.IsExpanded = false
	__Node__000021_Point.HasCheckboxButton = true
	__Node__000021_Point.IsChecked = false
	__Node__000021_Point.IsCheckboxDisabled = false
	__Node__000021_Point.HasAddChildButton = false
	__Node__000021_Point.HasEditButton = false
	__Node__000021_Point.IsInEditMode = false
	__Node__000021_Point.HasDrawButton = false
	__Node__000021_Point.HasDrawOffButton = false
	__Node__000021_Point.IsInDrawMode = false
	__Node__000021_Point.IsSaved = false
	__Node__000021_Point.HasDeleteButton = false

	// Node values setup
	__Node__000022_Point.Name = `Point`
	__Node__000022_Point.Type = models.GONG_STRUCT
	__Node__000022_Point.IsExpanded = false
	__Node__000022_Point.HasCheckboxButton = true
	__Node__000022_Point.IsChecked = true
	__Node__000022_Point.IsCheckboxDisabled = false
	__Node__000022_Point.HasAddChildButton = false
	__Node__000022_Point.HasEditButton = false
	__Node__000022_Point.IsInEditMode = false
	__Node__000022_Point.HasDrawButton = false
	__Node__000022_Point.HasDrawOffButton = false
	__Node__000022_Point.IsInDrawMode = false
	__Node__000022_Point.IsSaved = false
	__Node__000022_Point.HasDeleteButton = false

	// Node values setup
	__Node__000023_PointExclusiveSet.Name = `PointExclusiveSet`
	__Node__000023_PointExclusiveSet.Type = models.GONG_STRUCT
	__Node__000023_PointExclusiveSet.IsExpanded = false
	__Node__000023_PointExclusiveSet.HasCheckboxButton = true
	__Node__000023_PointExclusiveSet.IsChecked = true
	__Node__000023_PointExclusiveSet.IsCheckboxDisabled = false
	__Node__000023_PointExclusiveSet.HasAddChildButton = false
	__Node__000023_PointExclusiveSet.HasEditButton = false
	__Node__000023_PointExclusiveSet.IsInEditMode = false
	__Node__000023_PointExclusiveSet.HasDrawButton = false
	__Node__000023_PointExclusiveSet.HasDrawOffButton = false
	__Node__000023_PointExclusiveSet.IsInDrawMode = false
	__Node__000023_PointExclusiveSet.IsSaved = false
	__Node__000023_PointExclusiveSet.HasDeleteButton = false

	// Node values setup
	__Node__000024_PointNonExclusiveSet.Name = `PointNonExclusiveSet`
	__Node__000024_PointNonExclusiveSet.Type = models.GONG_STRUCT
	__Node__000024_PointNonExclusiveSet.IsExpanded = false
	__Node__000024_PointNonExclusiveSet.HasCheckboxButton = true
	__Node__000024_PointNonExclusiveSet.IsChecked = false
	__Node__000024_PointNonExclusiveSet.IsCheckboxDisabled = false
	__Node__000024_PointNonExclusiveSet.HasAddChildButton = false
	__Node__000024_PointNonExclusiveSet.HasEditButton = false
	__Node__000024_PointNonExclusiveSet.IsInEditMode = false
	__Node__000024_PointNonExclusiveSet.HasDrawButton = false
	__Node__000024_PointNonExclusiveSet.HasDrawOffButton = false
	__Node__000024_PointNonExclusiveSet.IsInDrawMode = false
	__Node__000024_PointNonExclusiveSet.IsSaved = false
	__Node__000024_PointNonExclusiveSet.HasDeleteButton = false

	// Node values setup
	__Node__000025_PointUse.Name = `PointUse`
	__Node__000025_PointUse.Type = models.GONG_STRUCT
	__Node__000025_PointUse.IsExpanded = false
	__Node__000025_PointUse.HasCheckboxButton = true
	__Node__000025_PointUse.IsChecked = false
	__Node__000025_PointUse.IsCheckboxDisabled = false
	__Node__000025_PointUse.HasAddChildButton = false
	__Node__000025_PointUse.HasEditButton = false
	__Node__000025_PointUse.IsInEditMode = false
	__Node__000025_PointUse.HasDrawButton = false
	__Node__000025_PointUse.HasDrawOffButton = false
	__Node__000025_PointUse.IsInDrawMode = false
	__Node__000025_PointUse.IsSaved = false
	__Node__000025_PointUse.HasDeleteButton = false

	// Node values setup
	__Node__000026_Points.Name = `Points`
	__Node__000026_Points.Type = models.GONG_STRUCT_FIELD
	__Node__000026_Points.IsExpanded = false
	__Node__000026_Points.HasCheckboxButton = true
	__Node__000026_Points.IsChecked = false
	__Node__000026_Points.IsCheckboxDisabled = true
	__Node__000026_Points.HasAddChildButton = false
	__Node__000026_Points.HasEditButton = false
	__Node__000026_Points.IsInEditMode = false
	__Node__000026_Points.HasDrawButton = false
	__Node__000026_Points.HasDrawOffButton = false
	__Node__000026_Points.IsInDrawMode = false
	__Node__000026_Points.IsSaved = false
	__Node__000026_Points.HasDeleteButton = false

	// Node values setup
	__Node__000027_Points.Name = `Points`
	__Node__000027_Points.Type = models.GONG_STRUCT_FIELD
	__Node__000027_Points.IsExpanded = false
	__Node__000027_Points.HasCheckboxButton = true
	__Node__000027_Points.IsChecked = false
	__Node__000027_Points.IsCheckboxDisabled = true
	__Node__000027_Points.HasAddChildButton = false
	__Node__000027_Points.HasEditButton = false
	__Node__000027_Points.IsInEditMode = false
	__Node__000027_Points.HasDrawButton = false
	__Node__000027_Points.HasDrawOffButton = false
	__Node__000027_Points.IsInDrawMode = false
	__Node__000027_Points.IsSaved = false
	__Node__000027_Points.HasDeleteButton = false

	// Node values setup
	__Node__000028_Points.Name = `Points`
	__Node__000028_Points.Type = models.GONG_STRUCT_FIELD
	__Node__000028_Points.IsExpanded = false
	__Node__000028_Points.HasCheckboxButton = true
	__Node__000028_Points.IsChecked = true
	__Node__000028_Points.IsCheckboxDisabled = false
	__Node__000028_Points.HasAddChildButton = false
	__Node__000028_Points.HasEditButton = false
	__Node__000028_Points.IsInEditMode = false
	__Node__000028_Points.HasDrawButton = false
	__Node__000028_Points.HasDrawOffButton = false
	__Node__000028_Points.IsInDrawMode = false
	__Node__000028_Points.IsSaved = false
	__Node__000028_Points.HasDeleteButton = false

	// Node values setup
	__Node__000029_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__Node__000029_ShortNodeOnModels.Type = models.GONG_NOTE
	__Node__000029_ShortNodeOnModels.IsExpanded = true
	__Node__000029_ShortNodeOnModels.HasCheckboxButton = true
	__Node__000029_ShortNodeOnModels.IsChecked = true
	__Node__000029_ShortNodeOnModels.IsCheckboxDisabled = false
	__Node__000029_ShortNodeOnModels.HasAddChildButton = false
	__Node__000029_ShortNodeOnModels.HasEditButton = false
	__Node__000029_ShortNodeOnModels.IsInEditMode = false
	__Node__000029_ShortNodeOnModels.HasDrawButton = false
	__Node__000029_ShortNodeOnModels.HasDrawOffButton = false
	__Node__000029_ShortNodeOnModels.IsInDrawMode = false
	__Node__000029_ShortNodeOnModels.IsSaved = false
	__Node__000029_ShortNodeOnModels.HasDeleteButton = false

	// Node values setup
	__Node__000030_SimulationStage.Name = `SimulationStage`
	__Node__000030_SimulationStage.Type = models.GONG_ENUM
	__Node__000030_SimulationStage.IsExpanded = false
	__Node__000030_SimulationStage.HasCheckboxButton = true
	__Node__000030_SimulationStage.IsChecked = false
	__Node__000030_SimulationStage.IsCheckboxDisabled = false
	__Node__000030_SimulationStage.HasAddChildButton = false
	__Node__000030_SimulationStage.HasEditButton = false
	__Node__000030_SimulationStage.IsInEditMode = false
	__Node__000030_SimulationStage.HasDrawButton = false
	__Node__000030_SimulationStage.HasDrawOffButton = false
	__Node__000030_SimulationStage.IsInDrawMode = false
	__Node__000030_SimulationStage.IsSaved = false
	__Node__000030_SimulationStage.HasDeleteButton = false

	// Node values setup
	__Node__000031_Start.Name = `Start`
	__Node__000031_Start.Type = models.GONG_STRUCT_FIELD
	__Node__000031_Start.IsExpanded = false
	__Node__000031_Start.HasCheckboxButton = true
	__Node__000031_Start.IsChecked = true
	__Node__000031_Start.IsCheckboxDisabled = false
	__Node__000031_Start.HasAddChildButton = false
	__Node__000031_Start.HasEditButton = false
	__Node__000031_Start.IsInEditMode = false
	__Node__000031_Start.HasDrawButton = false
	__Node__000031_Start.HasDrawOffButton = false
	__Node__000031_Start.IsInDrawMode = false
	__Node__000031_Start.IsSaved = false
	__Node__000031_Start.HasDeleteButton = false

	// Node values setup
	__Node__000032_Type.Name = `Type`
	__Node__000032_Type.Type = models.GONG_STRUCT_FIELD
	__Node__000032_Type.IsExpanded = false
	__Node__000032_Type.HasCheckboxButton = true
	__Node__000032_Type.IsChecked = false
	__Node__000032_Type.IsCheckboxDisabled = false
	__Node__000032_Type.HasAddChildButton = false
	__Node__000032_Type.HasEditButton = false
	__Node__000032_Type.IsInEditMode = false
	__Node__000032_Type.HasDrawButton = false
	__Node__000032_Type.HasDrawOffButton = false
	__Node__000032_Type.IsInDrawMode = false
	__Node__000032_Type.IsSaved = false
	__Node__000032_Type.HasDeleteButton = false

	// Node values setup
	__Node__000033_VeryLongLongLongLongLongLongField.Name = `VeryLongLongLongLongLongLongField`
	__Node__000033_VeryLongLongLongLongLongLongField.Type = models.GONG_STRUCT_FIELD
	__Node__000033_VeryLongLongLongLongLongLongField.IsExpanded = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasCheckboxButton = true
	__Node__000033_VeryLongLongLongLongLongLongField.IsChecked = false
	__Node__000033_VeryLongLongLongLongLongLongField.IsCheckboxDisabled = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasAddChildButton = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasEditButton = false
	__Node__000033_VeryLongLongLongLongLongLongField.IsInEditMode = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasDrawButton = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasDrawOffButton = false
	__Node__000033_VeryLongLongLongLongLongLongField.IsInDrawMode = false
	__Node__000033_VeryLongLongLongLongLongLongField.IsSaved = false
	__Node__000033_VeryLongLongLongLongLongLongField.HasDeleteButton = false

	// Node values setup
	__Node__000034_X.Name = `X`
	__Node__000034_X.Type = models.GONG_STRUCT_FIELD
	__Node__000034_X.IsExpanded = false
	__Node__000034_X.HasCheckboxButton = true
	__Node__000034_X.IsChecked = false
	__Node__000034_X.IsCheckboxDisabled = false
	__Node__000034_X.HasAddChildButton = false
	__Node__000034_X.HasEditButton = false
	__Node__000034_X.IsInEditMode = false
	__Node__000034_X.HasDrawButton = false
	__Node__000034_X.HasDrawOffButton = false
	__Node__000034_X.IsInDrawMode = false
	__Node__000034_X.IsSaved = false
	__Node__000034_X.HasDeleteButton = false

	// Node values setup
	__Node__000035_Y.Name = `Y`
	__Node__000035_Y.Type = models.GONG_STRUCT_FIELD
	__Node__000035_Y.IsExpanded = false
	__Node__000035_Y.HasCheckboxButton = true
	__Node__000035_Y.IsChecked = false
	__Node__000035_Y.IsCheckboxDisabled = false
	__Node__000035_Y.HasAddChildButton = false
	__Node__000035_Y.HasEditButton = false
	__Node__000035_Y.IsInEditMode = false
	__Node__000035_Y.HasDrawButton = false
	__Node__000035_Y.HasDrawOffButton = false
	__Node__000035_Y.IsInDrawMode = false
	__Node__000035_Y.IsSaved = false
	__Node__000035_Y.HasDeleteButton = false

	// Node values setup
	__Node__000036_Z.Name = `Z`
	__Node__000036_Z.Type = models.GONG_STRUCT_FIELD
	__Node__000036_Z.IsExpanded = false
	__Node__000036_Z.HasCheckboxButton = true
	__Node__000036_Z.IsChecked = false
	__Node__000036_Z.IsCheckboxDisabled = false
	__Node__000036_Z.HasAddChildButton = false
	__Node__000036_Z.HasEditButton = false
	__Node__000036_Z.IsInEditMode = false
	__Node__000036_Z.HasDrawButton = false
	__Node__000036_Z.HasDrawOffButton = false
	__Node__000036_Z.IsInDrawMode = false
	__Node__000036_Z.IsSaved = false
	__Node__000036_Z.HasDeleteButton = false

	// Node values setup
	__Node__000037_class_diagrams.Name = `class diagrams`
	__Node__000037_class_diagrams.Type = models.ROOT_OF_CLASS_DIAGRAMS
	__Node__000037_class_diagrams.IsExpanded = true
	__Node__000037_class_diagrams.HasCheckboxButton = false
	__Node__000037_class_diagrams.IsChecked = false
	__Node__000037_class_diagrams.IsCheckboxDisabled = false
	__Node__000037_class_diagrams.HasAddChildButton = true
	__Node__000037_class_diagrams.HasEditButton = false
	__Node__000037_class_diagrams.IsInEditMode = false
	__Node__000037_class_diagrams.HasDrawButton = false
	__Node__000037_class_diagrams.HasDrawOffButton = false
	__Node__000037_class_diagrams.IsInDrawMode = false
	__Node__000037_class_diagrams.IsSaved = false
	__Node__000037_class_diagrams.HasDeleteButton = false

	// Node values setup
	__Node__000038_gongenums.Name = `gongenums`
	__Node__000038_gongenums.IsExpanded = true
	__Node__000038_gongenums.HasCheckboxButton = false
	__Node__000038_gongenums.IsChecked = false
	__Node__000038_gongenums.IsCheckboxDisabled = true
	__Node__000038_gongenums.HasAddChildButton = false
	__Node__000038_gongenums.HasEditButton = false
	__Node__000038_gongenums.IsInEditMode = false
	__Node__000038_gongenums.HasDrawButton = false
	__Node__000038_gongenums.HasDrawOffButton = false
	__Node__000038_gongenums.IsInDrawMode = false
	__Node__000038_gongenums.IsSaved = false
	__Node__000038_gongenums.HasDeleteButton = false

	// Node values setup
	__Node__000039_gongstructs.Name = `gongstructs`
	__Node__000039_gongstructs.IsExpanded = true
	__Node__000039_gongstructs.HasCheckboxButton = false
	__Node__000039_gongstructs.IsChecked = false
	__Node__000039_gongstructs.IsCheckboxDisabled = true
	__Node__000039_gongstructs.HasAddChildButton = false
	__Node__000039_gongstructs.HasEditButton = false
	__Node__000039_gongstructs.IsInEditMode = false
	__Node__000039_gongstructs.HasDrawButton = false
	__Node__000039_gongstructs.HasDrawOffButton = false
	__Node__000039_gongstructs.IsInDrawMode = false
	__Node__000039_gongstructs.IsSaved = false
	__Node__000039_gongstructs.HasDeleteButton = false

	// Node values setup
	__Node__000040_notes.Name = `notes`
	__Node__000040_notes.IsExpanded = true
	__Node__000040_notes.HasCheckboxButton = false
	__Node__000040_notes.IsChecked = false
	__Node__000040_notes.IsCheckboxDisabled = true
	__Node__000040_notes.HasAddChildButton = false
	__Node__000040_notes.HasEditButton = false
	__Node__000040_notes.IsInEditMode = false
	__Node__000040_notes.HasDrawButton = false
	__Node__000040_notes.HasDrawOffButton = false
	__Node__000040_notes.IsInDrawMode = false
	__Node__000040_notes.IsSaved = false
	__Node__000040_notes.HasDeleteButton = false

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 90.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 60.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = true

	// Position values setup
	__Position__000000_Position_0000.X = 50.000000
	__Position__000000_Position_0000.Y = 210.000000
	__Position__000000_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000001_Position_0001.X = 680.000000
	__Position__000001_Position_0001.Y = 220.000000
	__Position__000001_Position_0001.Name = `Position-0001`

	// Position values setup
	__Position__000002_Position_0002.X = 30.000000
	__Position__000002_Position_0002.Y = 350.000000
	__Position__000002_Position_0002.Name = `Position-0002`

	// Tree values setup
	__Tree__000000_gong.Name = `gong`
	__Tree__000000_gong.Type = models.TREE_OF_IDENTIFIERS

	// Tree values setup
	__Tree__000001_gongdoc.Name = `gongdoc`
	__Tree__000001_gongdoc.Type = models.TREE_OF_DIAGRAMS

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 449.000000
	__Vertice__000000_Vertice_0000.Y = 329.500000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0001.X = 485.000000
	__Vertice__000001_Vertice_0001.Y = 201.500000
	__Vertice__000001_Vertice_0001.Name = `Vertice-0001`

	// Vertice values setup
	__Vertice__000002_Vertice_0002.X = 699.000000
	__Vertice__000002_Vertice_0002.Y = 389.500000
	__Vertice__000002_Vertice_0002.Name = `Vertice-0002`

	// Setup of pointers
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000000_Classshape0000)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000001_Classshape0001)
	__Classdiagram__000000_Diagram2.Classshapes = append(__Classdiagram__000000_Diagram2.Classshapes, __Classshape__000002_Classshape0002)
	__Classdiagram__000000_Diagram2.Notes = append(__Classdiagram__000000_Diagram2.Notes, __NoteShape__000000_ShortNodeOnModels)
	__Classdiagram__000001_Diagram3.Classshapes = append(__Classdiagram__000001_Diagram3.Classshapes, )
	__Classdiagram__000001_Diagram3.Classshapes = append(__Classdiagram__000001_Diagram3.Classshapes, )
	__Classdiagram__000001_Diagram3.Classshapes = append(__Classdiagram__000001_Diagram3.Classshapes, )
	__Classdiagram__000001_Diagram3.Classshapes = append(__Classdiagram__000001_Diagram3.Classshapes, )
	__Classdiagram__000001_Diagram3.Classshapes = append(__Classdiagram__000001_Diagram3.Classshapes, )
	__Classdiagram__000001_Diagram3.Notes = append(__Classdiagram__000001_Diagram3.Notes, )
	__Classdiagram__000001_Diagram3.Notes = append(__Classdiagram__000001_Diagram3.Notes, )
	__Classdiagram__000002_Diagram4.Classshapes = append(__Classdiagram__000002_Diagram4.Classshapes, )
	__Classdiagram__000002_Diagram4.Classshapes = append(__Classdiagram__000002_Diagram4.Classshapes, )
	__Classdiagram__000002_Diagram4.Classshapes = append(__Classdiagram__000002_Diagram4.Classshapes, )
	__Classdiagram__000002_Diagram4.Notes = append(__Classdiagram__000002_Diagram4.Notes, )
	__Classdiagram__000003_DiagramEnums.Classshapes = append(__Classdiagram__000003_DiagramEnums.Classshapes, )
	__Classshape__000000_Classshape0000.Position = __Position__000000_Position_0000
	__Classshape__000000_Classshape0000.Reference = 
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000000_CreationDate)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000000_End)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000002_Start)
	__Classshape__000001_Classshape0001.Position = __Position__000001_Position_0001
	__Classshape__000001_Classshape0001.Reference = 
	__Classshape__000002_Classshape0002.Position = __Position__000002_Position_0002
	__Classshape__000002_Classshape0002.Reference = 
	__Classshape__000002_Classshape0002.Links = append(__Classshape__000002_Classshape0002.Links, __Link__000001_Points)
	__Link__000000_End.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000001_Points.Middlevertice = __Vertice__000002_Vertice_0002
	__Link__000002_Start.Middlevertice = __Vertice__000001_Vertice_0001
	__Node__000006_Diagram2.Classdiagram = __Classdiagram__000000_Diagram2
	__Node__000007_Diagram3.Classdiagram = __Classdiagram__000001_Diagram3
	__Node__000008_Diagram4.Classdiagram = __Classdiagram__000002_Diagram4
	__Node__000009_DiagramEnums.Classdiagram = __Classdiagram__000003_DiagramEnums
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000017_Name)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000031_Start)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000010_End)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000004_CreationDate)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000011_JourneyTime)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000032_Type)
	__Node__000012_Line.Children = append(__Node__000012_Line.Children, __Node__000033_VeryLongLongLongLongLongLongField)
	__Node__000014_LineType.Children = append(__Node__000014_LineType.Children, __Node__000002_CONTINUOUS)
	__Node__000014_LineType.Children = append(__Node__000014_LineType.Children, __Node__000005_DOTTED)
	__Node__000022_Point.Children = append(__Node__000022_Point.Children, __Node__000016_Name)
	__Node__000022_Point.Children = append(__Node__000022_Point.Children, __Node__000036_Z)
	__Node__000022_Point.Children = append(__Node__000022_Point.Children, __Node__000034_X)
	__Node__000022_Point.Children = append(__Node__000022_Point.Children, __Node__000035_Y)
	__Node__000022_Point.Children = append(__Node__000022_Point.Children, __Node__000003_CreatedAt)
	__Node__000023_PointExclusiveSet.Children = append(__Node__000023_PointExclusiveSet.Children, __Node__000018_Name)
	__Node__000023_PointExclusiveSet.Children = append(__Node__000023_PointExclusiveSet.Children, __Node__000028_Points)
	__Node__000024_PointNonExclusiveSet.Children = append(__Node__000024_PointNonExclusiveSet.Children, __Node__000020_Name)
	__Node__000024_PointNonExclusiveSet.Children = append(__Node__000024_PointNonExclusiveSet.Children, __Node__000026_Points)
	__Node__000025_PointUse.Children = append(__Node__000025_PointUse.Children, __Node__000019_Name)
	__Node__000025_PointUse.Children = append(__Node__000025_PointUse.Children, __Node__000027_Points)
	__Node__000029_ShortNodeOnModels.Children = append(__Node__000029_ShortNodeOnModels.Children, __Node__000013_Line)
	__Node__000029_ShortNodeOnModels.Children = append(__Node__000029_ShortNodeOnModels.Children, __Node__000021_Point)
	__Node__000030_SimulationStage.Children = append(__Node__000030_SimulationStage.Children, __Node__000001_BEFORE)
	__Node__000030_SimulationStage.Children = append(__Node__000030_SimulationStage.Children, __Node__000000_AFTER)
	__Node__000037_class_diagrams.Children = append(__Node__000037_class_diagrams.Children, __Node__000008_Diagram4)
	__Node__000037_class_diagrams.Children = append(__Node__000037_class_diagrams.Children, __Node__000009_DiagramEnums)
	__Node__000037_class_diagrams.Children = append(__Node__000037_class_diagrams.Children, __Node__000006_Diagram2)
	__Node__000037_class_diagrams.Children = append(__Node__000037_class_diagrams.Children, __Node__000007_Diagram3)
	__Node__000038_gongenums.Children = append(__Node__000038_gongenums.Children, __Node__000014_LineType)
	__Node__000038_gongenums.Children = append(__Node__000038_gongenums.Children, __Node__000030_SimulationStage)
	__Node__000039_gongstructs.Children = append(__Node__000039_gongstructs.Children, __Node__000022_Point)
	__Node__000039_gongstructs.Children = append(__Node__000039_gongstructs.Children, __Node__000023_PointExclusiveSet)
	__Node__000039_gongstructs.Children = append(__Node__000039_gongstructs.Children, __Node__000024_PointNonExclusiveSet)
	__Node__000039_gongstructs.Children = append(__Node__000039_gongstructs.Children, __Node__000012_Line)
	__Node__000039_gongstructs.Children = append(__Node__000039_gongstructs.Children, __Node__000025_PointUse)
	__Node__000040_notes.Children = append(__Node__000040_notes.Children, __Node__000015_LongNodeOnModels)
	__Node__000040_notes.Children = append(__Node__000040_notes.Children, __Node__000029_ShortNodeOnModels)
	__Tree__000000_gong.RootNodes = append(__Tree__000000_gong.RootNodes, __Node__000039_gongstructs)
	__Tree__000000_gong.RootNodes = append(__Tree__000000_gong.RootNodes, __Node__000038_gongenums)
	__Tree__000000_gong.RootNodes = append(__Tree__000000_gong.RootNodes, __Node__000040_notes)
	__Tree__000001_gongdoc.RootNodes = append(__Tree__000001_gongdoc.RootNodes, __Node__000037_class_diagrams)
}


