package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Foo models.StageStruct
var ___dummy__Time_Foo time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Foo ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Foo map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.AFTER": ref_models.AFTER,

	"ref_models.BEFORE": ref_models.BEFORE,

	"ref_models.CONTINUOUS": ref_models.CONTINUOUS,

	"ref_models.CONTINUOUS_ZERO": ref_models.CONTINUOUS_ZERO,

	"ref_models.DOTTED": ref_models.DOTTED,

	"ref_models.DOTTED_ONE": ref_models.DOTTED_ONE,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.CreationDate": (ref_models.Line{}).CreationDate,

	"ref_models.Line.End": (ref_models.Line{}).End,

	"ref_models.Line.JourneyTime": (ref_models.Line{}).JourneyTime,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Start": (ref_models.Line{}).Start,

	"ref_models.Line.Type": (ref_models.Line{}).Type,

	"ref_models.Line.VeryLongLongLongLongLongLongField": (ref_models.Line{}).VeryLongLongLongLongLongLongField,

	"ref_models.LineTypeInt": ref_models.LineTypeInt(0),

	"ref_models.LineTypeString": ref_models.LineTypeString(""),

	"ref_models.LongNodeOnModels": ref_models.LongNodeOnModels,

	"ref_models.MarkdownNodeOnModels": ref_models.MarkdownNodeOnModels,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.CreatedAt": (ref_models.Point{}).CreatedAt,

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Point.Z": (ref_models.Point{}).Z,

	"ref_models.PointExclusiveSet": &(ref_models.PointExclusiveSet{}),

	"ref_models.PointExclusiveSet.Name": (ref_models.PointExclusiveSet{}).Name,

	"ref_models.PointExclusiveSet.Points": (ref_models.PointExclusiveSet{}).Points,

	"ref_models.PointNonExclusiveSet": &(ref_models.PointNonExclusiveSet{}),

	"ref_models.PointNonExclusiveSet.Name": (ref_models.PointNonExclusiveSet{}).Name,

	"ref_models.PointNonExclusiveSet.Points": (ref_models.PointNonExclusiveSet{}).Points,

	"ref_models.PointUse": &(ref_models.PointUse{}),

	"ref_models.PointUse.Name": (ref_models.PointUse{}).Name,

	"ref_models.PointUse.Points": (ref_models.PointUse{}).Points,

	"ref_models.ShortNodeOnModels": ref_models.ShortNodeOnModels,

	"ref_models.SimulationStage": ref_models.SimulationStage(""),
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Foo"] = FooInjection
// }

// FooInjection will stage objects of database "Foo"
func FooInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Foo := (&models.Classdiagram{Name: `Foo`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_2_Line := (&models.GongStructShape{Name: `NewDiagram_2-Line`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_2_Point := (&models.GongStructShape{Name: `NewDiagram_2-Point`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_2_Line := (&models.Position{Name: `Pos-NewDiagram_2-Line`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_2_Point := (&models.Position{Name: `Pos-NewDiagram_2-Point`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_2_in_middle_between_NewDiagram_2_Line_and_NewDiagram_2_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_2 in middle between NewDiagram_2-Line and NewDiagram_2-Point`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Foo.Name = `Foo`
	__Classdiagram__000000_Foo.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_2_Line.Name = `NewDiagram_2-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_NewDiagram_2_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_NewDiagram_2_Line.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_2_Line.NbInstances = 54
	__GongStructShape__000000_NewDiagram_2_Line.Width = 240.000000
	__GongStructShape__000000_NewDiagram_2_Line.Heigth = 114.999992
	__GongStructShape__000000_NewDiagram_2_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_2_Point.Name = `NewDiagram_2-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_NewDiagram_2_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_NewDiagram_2_Point.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_2_Point.NbInstances = 33
	__GongStructShape__000001_NewDiagram_2_Point.Width = 240.000000
	__GongStructShape__000001_NewDiagram_2_Point.Heigth = 114.999992
	__GongStructShape__000001_NewDiagram_2_Point.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000000_End.Fieldtypename = `ref_models.Point`
	__Link__000000_End.FieldOffsetX = -51.000000
	__Link__000000_End.FieldOffsetY = -19.000000
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.TargetMultiplicityOffsetX = -51.000000
	__Link__000000_End.TargetMultiplicityOffsetY = 27.000000
	__Link__000000_End.SourceMultiplicity = models.MANY
	__Link__000000_End.SourceMultiplicityOffsetX = 12.000000
	__Link__000000_End.SourceMultiplicityOffsetY = 23.000000
	__Link__000000_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_End.StartRatio = 0.547860
	__Link__000000_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_End.EndRatio = 0.460903
	__Link__000000_End.CornerOffsetRatio = 1.598584

	// Position values setup
	__Position__000000_Pos_NewDiagram_2_Line.X = 138.000000
	__Position__000000_Pos_NewDiagram_2_Line.Y = 128.000008
	__Position__000000_Pos_NewDiagram_2_Line.Name = `Pos-NewDiagram_2-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_2_Point.X = 693.000000
	__Position__000001_Pos_NewDiagram_2_Point.Y = 212.000008
	__Position__000001_Pos_NewDiagram_2_Point.Name = `Pos-NewDiagram_2-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_2_in_middle_between_NewDiagram_2_Line_and_NewDiagram_2_Point.X = 515.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_2_in_middle_between_NewDiagram_2_Line_and_NewDiagram_2_Point.Y = 161.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_2_in_middle_between_NewDiagram_2_Line_and_NewDiagram_2_Point.Name = `Verticle in class diagram NewDiagram_2 in middle between NewDiagram_2-Line and NewDiagram_2-Point`

	// Setup of pointers
	__Classdiagram__000000_Foo.GongStructShapes = append(__Classdiagram__000000_Foo.GongStructShapes, __GongStructShape__000000_NewDiagram_2_Line)
	__Classdiagram__000000_Foo.GongStructShapes = append(__Classdiagram__000000_Foo.GongStructShapes, __GongStructShape__000001_NewDiagram_2_Point)
	__GongStructShape__000000_NewDiagram_2_Line.Position = __Position__000000_Pos_NewDiagram_2_Line
	__GongStructShape__000000_NewDiagram_2_Line.Links = append(__GongStructShape__000000_NewDiagram_2_Line.Links, __Link__000000_End)
	__GongStructShape__000001_NewDiagram_2_Point.Position = __Position__000001_Pos_NewDiagram_2_Point
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_2_in_middle_between_NewDiagram_2_Line_and_NewDiagram_2_Point
}

