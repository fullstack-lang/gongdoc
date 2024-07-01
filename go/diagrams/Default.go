package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Classdiagram": &(ref_models.Classdiagram{}),

	"ref_models.Classdiagram.GongEnumShapes": (ref_models.Classdiagram{}).GongEnumShapes,

	"ref_models.Classdiagram.GongStructShapes": (ref_models.Classdiagram{}).GongStructShapes,

	"ref_models.Classdiagram.IsInDrawMode": (ref_models.Classdiagram{}).IsInDrawMode,

	"ref_models.Classdiagram.Name": (ref_models.Classdiagram{}).Name,

	"ref_models.Classdiagram.NoteShapes": (ref_models.Classdiagram{}).NoteShapes,

	"ref_models.DiagramPackage": &(ref_models.DiagramPackage{}),

	"ref_models.DiagramPackage.AbsolutePathToDiagramPackage": (ref_models.DiagramPackage{}).AbsolutePathToDiagramPackage,

	"ref_models.DiagramPackage.Classdiagrams": (ref_models.DiagramPackage{}).Classdiagrams,

	"ref_models.DiagramPackage.GongModelPath": (ref_models.DiagramPackage{}).GongModelPath,

	"ref_models.DiagramPackage.IsEditable": (ref_models.DiagramPackage{}).IsEditable,

	"ref_models.DiagramPackage.IsReloaded": (ref_models.DiagramPackage{}).IsReloaded,

	"ref_models.DiagramPackage.Name": (ref_models.DiagramPackage{}).Name,

	"ref_models.DiagramPackage.Path": (ref_models.DiagramPackage{}).Path,

	"ref_models.DiagramPackage.SelectedClassdiagram": (ref_models.DiagramPackage{}).SelectedClassdiagram,

	"ref_models.DiagramPackage.Umlscs": (ref_models.DiagramPackage{}).Umlscs,

	"ref_models.Field": &(ref_models.Field{}),

	"ref_models.Field.FieldTypeAsString": (ref_models.Field{}).FieldTypeAsString,

	"ref_models.Field.Fieldtypename": (ref_models.Field{}).Fieldtypename,

	"ref_models.Field.Identifier": (ref_models.Field{}).Identifier,

	"ref_models.Field.Name": (ref_models.Field{}).Name,

	"ref_models.Field.Structname": (ref_models.Field{}).Structname,

	"ref_models.GongEnumShape": &(ref_models.GongEnumShape{}),

	"ref_models.GongEnumShape.GongEnumValueEntrys": (ref_models.GongEnumShape{}).GongEnumValueEntrys,

	"ref_models.GongEnumShape.Height": (ref_models.GongEnumShape{}).Height,

	"ref_models.GongEnumShape.Identifier": (ref_models.GongEnumShape{}).Identifier,

	"ref_models.GongEnumShape.Name": (ref_models.GongEnumShape{}).Name,

	"ref_models.GongEnumShape.Position": (ref_models.GongEnumShape{}).Position,

	"ref_models.GongEnumShape.Width": (ref_models.GongEnumShape{}).Width,

	"ref_models.GongEnumShapeType": ref_models.GongEnumShapeType(0),

	"ref_models.GongEnumValueEntry": &(ref_models.GongEnumValueEntry{}),

	"ref_models.GongEnumValueEntry.Identifier": (ref_models.GongEnumValueEntry{}).Identifier,

	"ref_models.GongEnumValueEntry.Name": (ref_models.GongEnumValueEntry{}).Name,

	"ref_models.GongStructShape": &(ref_models.GongStructShape{}),

	"ref_models.GongStructShape.Fields": (ref_models.GongStructShape{}).Fields,

	"ref_models.GongStructShape.Height": (ref_models.GongStructShape{}).Height,

	"ref_models.GongStructShape.Identifier": (ref_models.GongStructShape{}).Identifier,

	"ref_models.GongStructShape.IsSelected": (ref_models.GongStructShape{}).IsSelected,

	"ref_models.GongStructShape.Links": (ref_models.GongStructShape{}).Links,

	"ref_models.GongStructShape.Name": (ref_models.GongStructShape{}).Name,

	"ref_models.GongStructShape.NbInstances": (ref_models.GongStructShape{}).NbInstances,

	"ref_models.GongStructShape.Position": (ref_models.GongStructShape{}).Position,

	"ref_models.GongStructShape.ShowNbInstances": (ref_models.GongStructShape{}).ShowNbInstances,

	"ref_models.GongStructShape.Width": (ref_models.GongStructShape{}).Width,

	"ref_models.GongdocStackName": ref_models.GongdocStackName,

	"ref_models.Int": ref_models.Int,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.CornerOffsetRatio": (ref_models.Link{}).CornerOffsetRatio,

	"ref_models.Link.EndOrientation": (ref_models.Link{}).EndOrientation,

	"ref_models.Link.EndRatio": (ref_models.Link{}).EndRatio,

	"ref_models.Link.FieldOffsetX": (ref_models.Link{}).FieldOffsetX,

	"ref_models.Link.FieldOffsetY": (ref_models.Link{}).FieldOffsetY,

	"ref_models.Link.Fieldtypename": (ref_models.Link{}).Fieldtypename,

	"ref_models.Link.Identifier": (ref_models.Link{}).Identifier,

	"ref_models.Link.Middlevertice": (ref_models.Link{}).Middlevertice,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.SourceMultiplicity": (ref_models.Link{}).SourceMultiplicity,

	"ref_models.Link.SourceMultiplicityOffsetX": (ref_models.Link{}).SourceMultiplicityOffsetX,

	"ref_models.Link.SourceMultiplicityOffsetY": (ref_models.Link{}).SourceMultiplicityOffsetY,

	"ref_models.Link.StartOrientation": (ref_models.Link{}).StartOrientation,

	"ref_models.Link.StartRatio": (ref_models.Link{}).StartRatio,

	"ref_models.Link.TargetMultiplicity": (ref_models.Link{}).TargetMultiplicity,

	"ref_models.Link.TargetMultiplicityOffsetX": (ref_models.Link{}).TargetMultiplicityOffsetX,

	"ref_models.Link.TargetMultiplicityOffsetY": (ref_models.Link{}).TargetMultiplicityOffsetY,

	"ref_models.MANY": ref_models.MANY,

	"ref_models.MultiplicityType": ref_models.MultiplicityType(""),

	"ref_models.NOTE_SHAPE_LINK_TO_GONG_FIELD": ref_models.NOTE_SHAPE_LINK_TO_GONG_FIELD,

	"ref_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE": ref_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,

	"ref_models.NoteOnGongdoc": ref_models.NoteOnGongdoc,

	"ref_models.NoteShape": &(ref_models.NoteShape{}),

	"ref_models.NoteShape.Body": (ref_models.NoteShape{}).Body,

	"ref_models.NoteShape.BodyHTML": (ref_models.NoteShape{}).BodyHTML,

	"ref_models.NoteShape.Height": (ref_models.NoteShape{}).Height,

	"ref_models.NoteShape.Identifier": (ref_models.NoteShape{}).Identifier,

	"ref_models.NoteShape.Matched": (ref_models.NoteShape{}).Matched,

	"ref_models.NoteShape.Name": (ref_models.NoteShape{}).Name,

	"ref_models.NoteShape.NoteShapeLinks": (ref_models.NoteShape{}).NoteShapeLinks,

	"ref_models.NoteShape.Width": (ref_models.NoteShape{}).Width,

	"ref_models.NoteShape.X": (ref_models.NoteShape{}).X,

	"ref_models.NoteShape.Y": (ref_models.NoteShape{}).Y,

	"ref_models.NoteShapeLink": &(ref_models.NoteShapeLink{}),

	"ref_models.NoteShapeLink.Identifier": (ref_models.NoteShapeLink{}).Identifier,

	"ref_models.NoteShapeLink.Name": (ref_models.NoteShapeLink{}).Name,

	"ref_models.NoteShapeLink.Type": (ref_models.NoteShapeLink{}).Type,

	"ref_models.NoteShapeLinkType": ref_models.NoteShapeLinkType(""),

	"ref_models.ONE": ref_models.ONE,

	"ref_models.ORIENTATION_HORIZONTAL": ref_models.ORIENTATION_HORIZONTAL,

	"ref_models.ORIENTATION_VERTICAL": ref_models.ORIENTATION_VERTICAL,

	"ref_models.OrientationType": ref_models.OrientationType(""),

	"ref_models.Position": &(ref_models.Position{}),

	"ref_models.Position.Name": (ref_models.Position{}).Name,

	"ref_models.Position.X": (ref_models.Position{}).X,

	"ref_models.Position.Y": (ref_models.Position{}).Y,

	"ref_models.StacksNames": ref_models.StacksNames(""),

	"ref_models.String": ref_models.String,

	"ref_models.SvgStackName": ref_models.SvgStackName,

	"ref_models.UmlState": &(ref_models.UmlState{}),

	"ref_models.UmlState.Name": (ref_models.UmlState{}).Name,

	"ref_models.UmlState.X": (ref_models.UmlState{}).X,

	"ref_models.UmlState.Y": (ref_models.UmlState{}).Y,

	"ref_models.Umlsc": &(ref_models.Umlsc{}),

	"ref_models.Umlsc.Activestate": (ref_models.Umlsc{}).Activestate,

	"ref_models.Umlsc.IsInDrawMode": (ref_models.Umlsc{}).IsInDrawMode,

	"ref_models.Umlsc.Name": (ref_models.Umlsc{}).Name,

	"ref_models.Umlsc.States": (ref_models.Umlsc{}).States,

	"ref_models.Vertice": &(ref_models.Vertice{}),

	"ref_models.Vertice.Name": (ref_models.Vertice{}).Name,

	"ref_models.Vertice.X": (ref_models.Vertice{}).X,

	"ref_models.Vertice.Y": (ref_models.Vertice{}).Y,

	"ref_models.ZERO_ONE": ref_models.ZERO_ONE,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)

	__GongEnumShape__000000_Default_GongEnumShapeType := (&models.GongEnumShape{Name: `Default-GongEnumShapeType`}).Stage(stage)
	__GongEnumShape__000001_Default_MultiplicityType := (&models.GongEnumShape{Name: `Default-MultiplicityType`}).Stage(stage)

	__GongEnumValueEntry__000000_Int := (&models.GongEnumValueEntry{Name: `Int`}).Stage(stage)
	__GongEnumValueEntry__000001_ONE := (&models.GongEnumValueEntry{Name: `ONE`}).Stage(stage)
	__GongEnumValueEntry__000002_String := (&models.GongEnumValueEntry{Name: `String`}).Stage(stage)
	__GongEnumValueEntry__000003_ZERO_ONE := (&models.GongEnumValueEntry{Name: `ZERO_ONE`}).Stage(stage)

	__GongStructShape__000000_Default_Classdiagram := (&models.GongStructShape{Name: `Default-Classdiagram`}).Stage(stage)
	__GongStructShape__000001_Default_DiagramPackage := (&models.GongStructShape{Name: `Default-DiagramPackage`}).Stage(stage)
	__GongStructShape__000002_Default_Field := (&models.GongStructShape{Name: `Default-Field`}).Stage(stage)
	__GongStructShape__000003_Default_GongEnumShape := (&models.GongStructShape{Name: `Default-GongEnumShape`}).Stage(stage)
	__GongStructShape__000004_Default_GongStructShape := (&models.GongStructShape{Name: `Default-GongStructShape`}).Stage(stage)
	__GongStructShape__000005_Default_NoteShape := (&models.GongStructShape{Name: `Default-NoteShape`}).Stage(stage)

	__Link__000000_Classdiagrams := (&models.Link{Name: `Classdiagrams`}).Stage(stage)
	__Link__000001_Fields := (&models.Link{Name: `Fields`}).Stage(stage)
	__Link__000002_NoteShapes := (&models.Link{Name: `NoteShapes`}).Stage(stage)
	__Link__000003_SelectedClassdiagram := (&models.Link{Name: `SelectedClassdiagram`}).Stage(stage)

	__NoteShape__000000_NoteOnGongdoc := (&models.NoteShape{Name: `NoteOnGongdoc`}).Stage(stage)

	__NoteShapeLink__000000_Classdiagram_NoteShapes := (&models.NoteShapeLink{Name: `Classdiagram.NoteShapes`}).Stage(stage)
	__NoteShapeLink__000001_NoteShape := (&models.NoteShapeLink{Name: `NoteShape`}).Stage(stage)

	__Position__000000_Pos_Default_Classdiagram := (&models.Position{Name: `Pos-Default-Classdiagram`}).Stage(stage)
	__Position__000001_Pos_Default_DiagramPackage := (&models.Position{Name: `Pos-Default-DiagramPackage`}).Stage(stage)
	__Position__000002_Pos_Default_Field := (&models.Position{Name: `Pos-Default-Field`}).Stage(stage)
	__Position__000003_Pos_Default_GongEnumShape := (&models.Position{Name: `Pos-Default-GongEnumShape`}).Stage(stage)
	__Position__000004_Pos_Default_GongEnumShapeType := (&models.Position{Name: `Pos-Default-GongEnumShapeType`}).Stage(stage)
	__Position__000005_Pos_Default_GongStructShape := (&models.Position{Name: `Pos-Default-GongStructShape`}).Stage(stage)
	__Position__000006_Pos_Default_MultiplicityType := (&models.Position{Name: `Pos-Default-MultiplicityType`}).Stage(stage)
	__Position__000007_Pos_Default_NoteShape := (&models.Position{Name: `Pos-Default-NoteShape`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Classdiagram and Default-NoteShape`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStructShape_and_Default_Field := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongStructShape and Default-Field`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.GongStructShape.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.GongStructShape.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `GongStructShape`
	__Field__000000_Name.Fieldtypename = `string`

	__GongEnumShape__000000_Default_GongEnumShapeType.Name = `Default-GongEnumShapeType`

	//gong:ident [ref_models.GongEnumShapeType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_GongEnumShapeType.Identifier = `ref_models.GongEnumShapeType`
	__GongEnumShape__000000_Default_GongEnumShapeType.Width = 240.000000
	__GongEnumShape__000000_Default_GongEnumShapeType.Height = 93.000000

	__GongEnumShape__000001_Default_MultiplicityType.Name = `Default-MultiplicityType`

	//gong:ident [ref_models.MultiplicityType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000001_Default_MultiplicityType.Identifier = `ref_models.MultiplicityType`
	__GongEnumShape__000001_Default_MultiplicityType.Width = 240.000000
	__GongEnumShape__000001_Default_MultiplicityType.Height = 93.000000

	__GongEnumValueEntry__000000_Int.Name = `Int`

	//gong:ident [ref_models.GongEnumShapeType.Int] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000000_Int.Identifier = `ref_models.GongEnumShapeType.Int`

	__GongEnumValueEntry__000001_ONE.Name = `ONE`

	//gong:ident [ref_models.MultiplicityType.ONE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000001_ONE.Identifier = `ref_models.MultiplicityType.ONE`

	__GongEnumValueEntry__000002_String.Name = `String`

	//gong:ident [ref_models.GongEnumShapeType.String] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000002_String.Identifier = `ref_models.GongEnumShapeType.String`

	__GongEnumValueEntry__000003_ZERO_ONE.Name = `ZERO_ONE`

	//gong:ident [ref_models.MultiplicityType.ZERO_ONE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000003_ZERO_ONE.Identifier = `ref_models.MultiplicityType.ZERO_ONE`

	__GongStructShape__000000_Default_Classdiagram.Name = `Default-Classdiagram`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_Classdiagram.Height = 63.000000
	__GongStructShape__000000_Default_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_DiagramPackage.Name = `Default-DiagramPackage`

	//gong:ident [ref_models.DiagramPackage] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_DiagramPackage.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000001_Default_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000001_Default_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_DiagramPackage.Height = 63.000000
	__GongStructShape__000001_Default_DiagramPackage.IsSelected = false

	__GongStructShape__000002_Default_Field.Name = `Default-Field`

	//gong:ident [ref_models.Field] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Field.Identifier = `ref_models.Field`
	__GongStructShape__000002_Default_Field.ShowNbInstances = false
	__GongStructShape__000002_Default_Field.NbInstances = 0
	__GongStructShape__000002_Default_Field.Width = 240.000000
	__GongStructShape__000002_Default_Field.Height = 63.000000
	__GongStructShape__000002_Default_Field.IsSelected = false

	__GongStructShape__000003_Default_GongEnumShape.Name = `Default-GongEnumShape`

	//gong:ident [ref_models.GongEnumShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_GongEnumShape.Identifier = `ref_models.GongEnumShape`
	__GongStructShape__000003_Default_GongEnumShape.ShowNbInstances = false
	__GongStructShape__000003_Default_GongEnumShape.NbInstances = 0
	__GongStructShape__000003_Default_GongEnumShape.Width = 240.000000
	__GongStructShape__000003_Default_GongEnumShape.Height = 63.000000
	__GongStructShape__000003_Default_GongEnumShape.IsSelected = false

	__GongStructShape__000004_Default_GongStructShape.Name = `Default-GongStructShape`

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_GongStructShape.Identifier = `ref_models.GongStructShape`
	__GongStructShape__000004_Default_GongStructShape.ShowNbInstances = false
	__GongStructShape__000004_Default_GongStructShape.NbInstances = 0
	__GongStructShape__000004_Default_GongStructShape.Width = 240.000000
	__GongStructShape__000004_Default_GongStructShape.Height = 78.000000
	__GongStructShape__000004_Default_GongStructShape.IsSelected = false

	__GongStructShape__000005_Default_NoteShape.Name = `Default-NoteShape`

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_NoteShape.Identifier = `ref_models.NoteShape`
	__GongStructShape__000005_Default_NoteShape.ShowNbInstances = false
	__GongStructShape__000005_Default_NoteShape.NbInstances = 0
	__GongStructShape__000005_Default_NoteShape.Width = 240.000000
	__GongStructShape__000005_Default_NoteShape.Height = 63.000000
	__GongStructShape__000005_Default_NoteShape.IsSelected = false

	__Link__000000_Classdiagrams.Name = `Classdiagrams`

	//gong:ident [ref_models.DiagramPackage.Classdiagrams] comment added to overcome the problem with the comment map association
	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__Link__000000_Classdiagrams.Fieldtypename = `ref_models.Classdiagram`
	__Link__000000_Classdiagrams.FieldOffsetX = -50.000000
	__Link__000000_Classdiagrams.FieldOffsetY = -16.000000
	__Link__000000_Classdiagrams.TargetMultiplicity = models.MANY
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Classdiagrams.SourceMultiplicity = models.MANY
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Classdiagrams.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Classdiagrams.StartRatio = 0.322222
	__Link__000000_Classdiagrams.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Classdiagrams.EndRatio = 0.500000
	__Link__000000_Classdiagrams.CornerOffsetRatio = 1.079365

	__Link__000001_Fields.Name = `Fields`

	//gong:ident [ref_models.GongStructShape.Fields] comment added to overcome the problem with the comment map association
	__Link__000001_Fields.Identifier = `ref_models.GongStructShape.Fields`

	//gong:ident [ref_models.Field] comment added to overcome the problem with the comment map association
	__Link__000001_Fields.Fieldtypename = `ref_models.Field`
	__Link__000001_Fields.FieldOffsetX = -50.000000
	__Link__000001_Fields.FieldOffsetY = -16.000000
	__Link__000001_Fields.TargetMultiplicity = models.MANY
	__Link__000001_Fields.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Fields.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Fields.SourceMultiplicity = models.MANY
	__Link__000001_Fields.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_Fields.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_Fields.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Fields.StartRatio = 0.500000
	__Link__000001_Fields.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Fields.EndRatio = 0.500000
	__Link__000001_Fields.CornerOffsetRatio = 1.380000

	__Link__000002_NoteShapes.Name = `NoteShapes`

	//gong:ident [ref_models.Classdiagram.NoteShapes] comment added to overcome the problem with the comment map association
	__Link__000002_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
	__Link__000002_NoteShapes.Fieldtypename = `ref_models.NoteShape`
	__Link__000002_NoteShapes.FieldOffsetX = -94.000000
	__Link__000002_NoteShapes.FieldOffsetY = -14.000000
	__Link__000002_NoteShapes.TargetMultiplicity = models.MANY
	__Link__000002_NoteShapes.TargetMultiplicityOffsetX = -27.000000
	__Link__000002_NoteShapes.TargetMultiplicityOffsetY = 21.000000
	__Link__000002_NoteShapes.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_NoteShapes.SourceMultiplicityOffsetX = 22.000000
	__Link__000002_NoteShapes.SourceMultiplicityOffsetY = 26.000000
	__Link__000002_NoteShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_NoteShapes.StartRatio = 0.500000
	__Link__000002_NoteShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_NoteShapes.EndRatio = 0.500000
	__Link__000002_NoteShapes.CornerOffsetRatio = 1.262818

	__Link__000003_SelectedClassdiagram.Name = `SelectedClassdiagram`

	//gong:ident [ref_models.DiagramPackage.SelectedClassdiagram] comment added to overcome the problem with the comment map association
	__Link__000003_SelectedClassdiagram.Identifier = `ref_models.DiagramPackage.SelectedClassdiagram`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__Link__000003_SelectedClassdiagram.Fieldtypename = `ref_models.Classdiagram`
	__Link__000003_SelectedClassdiagram.FieldOffsetX = -50.000000
	__Link__000003_SelectedClassdiagram.FieldOffsetY = -16.000000
	__Link__000003_SelectedClassdiagram.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_SelectedClassdiagram.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_SelectedClassdiagram.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_SelectedClassdiagram.SourceMultiplicity = models.MANY
	__Link__000003_SelectedClassdiagram.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_SelectedClassdiagram.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_SelectedClassdiagram.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SelectedClassdiagram.StartRatio = 0.500000
	__Link__000003_SelectedClassdiagram.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_SelectedClassdiagram.EndRatio = 0.562500
	__Link__000003_SelectedClassdiagram.CornerOffsetRatio = 2.174603

	__NoteShape__000000_NoteOnGongdoc.Name = `NoteOnGongdoc`

	//gong:ident [ref_models.NoteOnGongdoc] comment added to overcome the problem with the comment map association
	__NoteShape__000000_NoteOnGongdoc.Identifier = `ref_models.NoteOnGongdoc`
	__NoteShape__000000_NoteOnGongdoc.Body = `Note Example

This note can refers to [models.NoteShape]
or to [models.Classdiagram.NoteShapes]
or to [models.OrientationType]
`
	__NoteShape__000000_NoteOnGongdoc.BodyHTML = `<p>Note Example
<p>This note can refers to <a href="/models#NoteShape">models.NoteShape</a>
or to <a href="/models#Classdiagram.NoteShapes">models.Classdiagram.NoteShapes</a>
or to <a href="/models#OrientationType">models.OrientationType</a>
`
	__NoteShape__000000_NoteOnGongdoc.X = 759.999939
	__NoteShape__000000_NoteOnGongdoc.Y = 48.000000
	__NoteShape__000000_NoteOnGongdoc.Width = 347.000000
	__NoteShape__000000_NoteOnGongdoc.Height = 123.000000
	__NoteShape__000000_NoteOnGongdoc.Matched = false

	__NoteShapeLink__000000_Classdiagram_NoteShapes.Name = `Classdiagram.NoteShapes`

	//gong:ident [ref_models.Classdiagram.NoteShapes] comment added to overcome the problem with the comment map association
	__NoteShapeLink__000000_Classdiagram_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`
	__NoteShapeLink__000000_Classdiagram_NoteShapes.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	__NoteShapeLink__000001_NoteShape.Name = `NoteShape`

	//gong:ident [ref_models.NoteShape] comment added to overcome the problem with the comment map association
	__NoteShapeLink__000001_NoteShape.Identifier = `ref_models.NoteShape`
	__NoteShapeLink__000001_NoteShape.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	__Position__000000_Pos_Default_Classdiagram.X = 558.000031
	__Position__000000_Pos_Default_Classdiagram.Y = 319.000000
	__Position__000000_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000001_Pos_Default_DiagramPackage.X = 121.000000
	__Position__000001_Pos_Default_DiagramPackage.Y = 181.000000
	__Position__000001_Pos_Default_DiagramPackage.Name = `Pos-Default-DiagramPackage`

	__Position__000002_Pos_Default_Field.X = 595.000000
	__Position__000002_Pos_Default_Field.Y = 558.000000
	__Position__000002_Pos_Default_Field.Name = `Pos-Default-Field`

	__Position__000003_Pos_Default_GongEnumShape.X = 35.000000
	__Position__000003_Pos_Default_GongEnumShape.Y = 26.000000
	__Position__000003_Pos_Default_GongEnumShape.Name = `Pos-Default-GongEnumShape`

	__Position__000004_Pos_Default_GongEnumShapeType.X = 54.000000
	__Position__000004_Pos_Default_GongEnumShapeType.Y = 390.000000
	__Position__000004_Pos_Default_GongEnumShapeType.Name = `Pos-Default-GongEnumShapeType`

	__Position__000005_Pos_Default_GongStructShape.X = 175.000000
	__Position__000005_Pos_Default_GongStructShape.Y = 544.000000
	__Position__000005_Pos_Default_GongStructShape.Name = `Pos-Default-GongStructShape`

	__Position__000006_Pos_Default_MultiplicityType.X = 409.000000
	__Position__000006_Pos_Default_MultiplicityType.Y = 46.000000
	__Position__000006_Pos_Default_MultiplicityType.Name = `Pos-Default-MultiplicityType`

	__Position__000007_Pos_Default_NoteShape.X = 1031.000061
	__Position__000007_Pos_Default_NoteShape.Y = 309.000000
	__Position__000007_Pos_Default_NoteShape.Name = `Pos-Default-NoteShape`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.X = 867.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.Y = 264.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.Name = `Verticle in class diagram Default in middle between Default-Classdiagram and Default-NoteShape`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.X = 700.000016
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Y = 281.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Name = `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.X = 700.000016
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Y = 281.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Name = `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStructShape_and_Default_Field.X = 745.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStructShape_and_Default_Field.Y = 590.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStructShape_and_Default_Field.Name = `Verticle in class diagram Default in middle between Default-GongStructShape and Default-Field`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_NoteShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_DiagramPackage)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Classdiagram)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_GongStructShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_GongEnumShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Field)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_GongEnumShapeType)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000001_Default_MultiplicityType)
	__Classdiagram__000000_Default.NoteShapes = append(__Classdiagram__000000_Default.NoteShapes, __NoteShape__000000_NoteOnGongdoc)
	__GongEnumShape__000000_Default_GongEnumShapeType.Position = __Position__000004_Pos_Default_GongEnumShapeType
	__GongEnumShape__000000_Default_GongEnumShapeType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_GongEnumShapeType.GongEnumValueEntrys, __GongEnumValueEntry__000000_Int)
	__GongEnumShape__000000_Default_GongEnumShapeType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_GongEnumShapeType.GongEnumValueEntrys, __GongEnumValueEntry__000002_String)
	__GongEnumShape__000001_Default_MultiplicityType.Position = __Position__000006_Pos_Default_MultiplicityType
	__GongEnumShape__000001_Default_MultiplicityType.GongEnumValueEntrys = append(__GongEnumShape__000001_Default_MultiplicityType.GongEnumValueEntrys, __GongEnumValueEntry__000003_ZERO_ONE)
	__GongEnumShape__000001_Default_MultiplicityType.GongEnumValueEntrys = append(__GongEnumShape__000001_Default_MultiplicityType.GongEnumValueEntrys, __GongEnumValueEntry__000001_ONE)
	__GongStructShape__000000_Default_Classdiagram.Position = __Position__000000_Pos_Default_Classdiagram
	__GongStructShape__000000_Default_Classdiagram.Links = append(__GongStructShape__000000_Default_Classdiagram.Links, __Link__000002_NoteShapes)
	__GongStructShape__000001_Default_DiagramPackage.Position = __Position__000001_Pos_Default_DiagramPackage
	__GongStructShape__000001_Default_DiagramPackage.Links = append(__GongStructShape__000001_Default_DiagramPackage.Links, __Link__000003_SelectedClassdiagram)
	__GongStructShape__000001_Default_DiagramPackage.Links = append(__GongStructShape__000001_Default_DiagramPackage.Links, __Link__000000_Classdiagrams)
	__GongStructShape__000002_Default_Field.Position = __Position__000002_Pos_Default_Field
	__GongStructShape__000003_Default_GongEnumShape.Position = __Position__000003_Pos_Default_GongEnumShape
	__GongStructShape__000004_Default_GongStructShape.Position = __Position__000005_Pos_Default_GongStructShape
	__GongStructShape__000004_Default_GongStructShape.Fields = append(__GongStructShape__000004_Default_GongStructShape.Fields, __Field__000000_Name)
	__GongStructShape__000004_Default_GongStructShape.Links = append(__GongStructShape__000004_Default_GongStructShape.Links, __Link__000001_Fields)
	__GongStructShape__000005_Default_NoteShape.Position = __Position__000007_Pos_Default_NoteShape
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram
	__Link__000001_Fields.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStructShape_and_Default_Field
	__Link__000002_NoteShapes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape
	__Link__000003_SelectedClassdiagram.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000001_NoteShape)
	__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks = append(__NoteShape__000000_NoteOnGongdoc.NoteShapeLinks, __NoteShapeLink__000000_Classdiagram_NoteShapes)
}
