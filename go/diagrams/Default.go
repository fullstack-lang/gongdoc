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

	// Declarations of staged instances of Button

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

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Classdiagram := (&models.Position{Name: `Pos-Default-Classdiagram`}).Stage(stage)
	__Position__000001_Pos_Default_DiagramPackage := (&models.Position{Name: `Pos-Default-DiagramPackage`}).Stage(stage)
	__Position__000002_Pos_Default_NoteShape := (&models.Position{Name: `Pos-Default-NoteShape`}).Stage(stage)

	// Declarations of staged instances of Tree

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
	__Position__000001_Pos_Default_DiagramPackage.X = 41.000000
	__Position__000001_Pos_Default_DiagramPackage.Y = 88.000000
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
