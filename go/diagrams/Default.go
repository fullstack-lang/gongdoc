package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
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

	"ref_models.GongEnumShape.Heigth": (ref_models.GongEnumShape{}).Heigth,

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

	"ref_models.GongStructShape.Heigth": (ref_models.GongStructShape{}).Heigth,

	"ref_models.GongStructShape.Identifier": (ref_models.GongStructShape{}).Identifier,

	"ref_models.GongStructShape.IsSelected": (ref_models.GongStructShape{}).IsSelected,

	"ref_models.GongStructShape.Links": (ref_models.GongStructShape{}).Links,

	"ref_models.GongStructShape.Name": (ref_models.GongStructShape{}).Name,

	"ref_models.GongStructShape.NbInstances": (ref_models.GongStructShape{}).NbInstances,

	"ref_models.GongStructShape.Position": (ref_models.GongStructShape{}).Position,

	"ref_models.GongStructShape.ShowNbInstances": (ref_models.GongStructShape{}).ShowNbInstances,

	"ref_models.GongStructShape.Width": (ref_models.GongStructShape{}).Width,

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

	"ref_models.NoteShape": &(ref_models.NoteShape{}),

	"ref_models.NoteShape.Body": (ref_models.NoteShape{}).Body,

	"ref_models.NoteShape.BodyHTML": (ref_models.NoteShape{}).BodyHTML,

	"ref_models.NoteShape.Heigth": (ref_models.NoteShape{}).Heigth,

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

	"ref_models.String": ref_models.String,

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

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Classdiagram := (&models.GongStructShape{Name: `Default-Classdiagram`}).Stage(stage)
	__GongStructShape__000001_Default_DiagramPackage := (&models.GongStructShape{Name: `Default-DiagramPackage`}).Stage(stage)
	__GongStructShape__000002_Default_NoteShape := (&models.GongStructShape{Name: `Default-NoteShape`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Classdiagrams := (&models.Link{Name: `Classdiagrams`}).Stage(stage)
	__Link__000001_NoteShapes := (&models.Link{Name: `NoteShapes`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Classdiagram := (&models.Position{Name: `Pos-Default-Classdiagram`}).Stage(stage)
	__Position__000001_Pos_Default_DiagramPackage := (&models.Position{Name: `Pos-Default-DiagramPackage`}).Stage(stage)
	__Position__000002_Pos_Default_NoteShape := (&models.Position{Name: `Pos-Default-NoteShape`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Classdiagram and Default-NoteShape`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_Default_Classdiagram.Name = `Default-Classdiagram`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__GongStructShape__000000_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.ShowNbInstances = true
	__GongStructShape__000000_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_Classdiagram.Heigth = 63.000000
	__GongStructShape__000000_Default_Classdiagram.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_DiagramPackage.Name = `Default-DiagramPackage`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage]
	__GongStructShape__000001_Default_DiagramPackage.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000001_Default_DiagramPackage.ShowNbInstances = true
	__GongStructShape__000001_Default_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_DiagramPackage.Heigth = 63.000000
	__GongStructShape__000001_Default_DiagramPackage.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_NoteShape.Name = `Default-NoteShape`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteShape]
	__GongStructShape__000002_Default_NoteShape.Identifier = `ref_models.NoteShape`
	__GongStructShape__000002_Default_NoteShape.ShowNbInstances = true
	__GongStructShape__000002_Default_NoteShape.NbInstances = 0
	__GongStructShape__000002_Default_NoteShape.Width = 240.000000
	__GongStructShape__000002_Default_NoteShape.Heigth = 63.000000
	__GongStructShape__000002_Default_NoteShape.IsSelected = false

	// Link values setup
	__Link__000000_Classdiagrams.Name = `Classdiagrams`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DiagramPackage.Classdiagrams]
	__Link__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram]
	__Link__000000_Classdiagrams.Fieldtypename = `ref_models.Classdiagram`
	__Link__000000_Classdiagrams.FieldOffsetX = -113.000000
	__Link__000000_Classdiagrams.FieldOffsetY = -19.000000
	__Link__000000_Classdiagrams.TargetMultiplicity = models.MANY
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetX = -29.000000
	__Link__000000_Classdiagrams.TargetMultiplicityOffsetY = 27.000000
	__Link__000000_Classdiagrams.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetX = 14.000000
	__Link__000000_Classdiagrams.SourceMultiplicityOffsetY = 23.000000
	__Link__000000_Classdiagrams.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Classdiagrams.StartRatio = 0.500000
	__Link__000000_Classdiagrams.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Classdiagrams.EndRatio = 0.500000
	__Link__000000_Classdiagrams.CornerOffsetRatio = 1.341985

	// Link values setup
	__Link__000001_NoteShapes.Name = `NoteShapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Classdiagram.NoteShapes]
	__Link__000001_NoteShapes.Identifier = `ref_models.Classdiagram.NoteShapes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteShape]
	__Link__000001_NoteShapes.Fieldtypename = `ref_models.NoteShape`
	__Link__000001_NoteShapes.FieldOffsetX = -94.000000
	__Link__000001_NoteShapes.FieldOffsetY = -14.000000
	__Link__000001_NoteShapes.TargetMultiplicity = models.MANY
	__Link__000001_NoteShapes.TargetMultiplicityOffsetX = -27.000000
	__Link__000001_NoteShapes.TargetMultiplicityOffsetY = 21.000000
	__Link__000001_NoteShapes.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_NoteShapes.SourceMultiplicityOffsetX = 22.000000
	__Link__000001_NoteShapes.SourceMultiplicityOffsetY = 26.000000
	__Link__000001_NoteShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_NoteShapes.StartRatio = 0.500000
	__Link__000001_NoteShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_NoteShapes.EndRatio = 0.500000
	__Link__000001_NoteShapes.CornerOffsetRatio = 1.262818

	// Position values setup
	__Position__000000_Pos_Default_Classdiagram.X = 456.000000
	__Position__000000_Pos_Default_Classdiagram.Y = 92.000000
	__Position__000000_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	// Position values setup
	__Position__000001_Pos_Default_DiagramPackage.X = 68.000000
	__Position__000001_Pos_Default_DiagramPackage.Y = 193.000000
	__Position__000001_Pos_Default_DiagramPackage.Name = `Pos-Default-DiagramPackage`

	// Position values setup
	__Position__000002_Pos_Default_NoteShape.X = 878.000000
	__Position__000002_Pos_Default_NoteShape.Y = 212.000000
	__Position__000002_Pos_Default_NoteShape.Name = `Pos-Default-NoteShape`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.X = 867.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.Y = 264.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape.Name = `Verticle in class diagram Default in middle between Default-Classdiagram and Default-NoteShape`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.X = 599.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Y = 118.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram.Name = `Verticle in class diagram Default in middle between Default-DiagramPackage and Default-Classdiagram`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_NoteShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_DiagramPackage)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Classdiagram)
	__GongStructShape__000000_Default_Classdiagram.Position = __Position__000000_Pos_Default_Classdiagram
	__GongStructShape__000000_Default_Classdiagram.Links = append(__GongStructShape__000000_Default_Classdiagram.Links, __Link__000001_NoteShapes)
	__GongStructShape__000001_Default_DiagramPackage.Position = __Position__000001_Pos_Default_DiagramPackage
	__GongStructShape__000001_Default_DiagramPackage.Links = append(__GongStructShape__000001_Default_DiagramPackage.Links, __Link__000000_Classdiagrams)
	__GongStructShape__000002_Default_NoteShape.Position = __Position__000002_Pos_Default_NoteShape
	__Link__000000_Classdiagrams.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_DiagramPackage_and_Default_Classdiagram
	__Link__000001_NoteShapes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Classdiagram_and_Default_NoteShape
}


