package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Toto models.StageStruct
var ___dummy__Time_Toto time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Toto ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Toto map[string]any = map[string]any{
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
// 	InjectionGateway["Toto"] = TotoInjection
// }

// TotoInjection will stage objects of database "Toto"
func TotoInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Toto := (&models.Classdiagram{Name: `Toto`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CreationDate := (&models.Field{Name: `CreationDate`}).Stage(stage)
	__Field__000001_JourneyTime := (&models.Field{Name: `JourneyTime`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Line := (&models.GongStructShape{Name: `Default-Line`}).Stage(stage)
	__GongStructShape__000001_Default_Point := (&models.GongStructShape{Name: `Default-Point`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000001_Start := (&models.Link{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Line := (&models.Position{Name: `Pos-Default-Line`}).Stage(stage)
	__Position__000001_Pos_Default_Point := (&models.Position{Name: `Pos-Default-Point`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Line and Default-Point`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Line and Default-Point`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Toto.Name = `Toto`
	__Classdiagram__000000_Toto.IsInDrawMode = false

	// Field values setup
	__Field__000000_CreationDate.Name = `CreationDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.CreationDate]
	__Field__000000_CreationDate.Identifier = `ref_models.Line.CreationDate`
	__Field__000000_CreationDate.FieldTypeAsString = ``
	__Field__000000_CreationDate.Structname = `Line`
	__Field__000000_CreationDate.Fieldtypename = `Time`

	// Field values setup
	__Field__000001_JourneyTime.Name = `JourneyTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.JourneyTime]
	__Field__000001_JourneyTime.Identifier = `ref_models.Line.JourneyTime`
	__Field__000001_JourneyTime.FieldTypeAsString = ``
	__Field__000001_JourneyTime.Structname = `Line`
	__Field__000001_JourneyTime.Fieldtypename = `Duration`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000002_Name.Identifier = `ref_models.Line.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Line`
	__Field__000002_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Line.Name = `Default-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_Default_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_Default_Line.ShowNbInstances = false
	__GongStructShape__000000_Default_Line.NbInstances = 0
	__GongStructShape__000000_Default_Line.Width = 240.000000
	__GongStructShape__000000_Default_Line.Heigth = 108.000000
	__GongStructShape__000000_Default_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_Point.Name = `Default-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000001_Default_Point.Identifier = `ref_models.Point`
	__GongStructShape__000001_Default_Point.ShowNbInstances = false
	__GongStructShape__000001_Default_Point.NbInstances = 0
	__GongStructShape__000001_Default_Point.Width = 240.000000
	__GongStructShape__000001_Default_Point.Heigth = 63.000000
	__GongStructShape__000001_Default_Point.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000000_End.Fieldtypename = `ref_models.Point`
	__Link__000000_End.FieldOffsetX = -50.000000
	__Link__000000_End.FieldOffsetY = -16.000000
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_End.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_End.SourceMultiplicity = models.MANY
	__Link__000000_End.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_End.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_End.StartRatio = 0.500000
	__Link__000000_End.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_End.EndRatio = 0.476111
	__Link__000000_End.CornerOffsetRatio = 3.412698

	// Link values setup
	__Link__000001_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Start]
	__Link__000001_Start.Identifier = `ref_models.Line.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__Link__000001_Start.Fieldtypename = `ref_models.Point`
	__Link__000001_Start.FieldOffsetX = -50.000000
	__Link__000001_Start.FieldOffsetY = -16.000000
	__Link__000001_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Start.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Start.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Start.SourceMultiplicity = models.MANY
	__Link__000001_Start.SourceMultiplicityOffsetX = -44.000000
	__Link__000001_Start.SourceMultiplicityOffsetY = 27.000000
	__Link__000001_Start.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Start.StartRatio = 0.492778
	__Link__000001_Start.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Start.EndRatio = 0.507937
	__Link__000001_Start.CornerOffsetRatio = 1.046296

	// Position values setup
	__Position__000000_Pos_Default_Line.X = 63.000000
	__Position__000000_Pos_Default_Line.Y = 94.000000
	__Position__000000_Pos_Default_Line.Name = `Pos-Default-Line`

	// Position values setup
	__Position__000001_Pos_Default_Point.X = 427.000000
	__Position__000001_Pos_Default_Point.Y = 270.000000
	__Position__000001_Pos_Default_Point.Name = `Pos-Default-Point`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.X = 566.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.Y = 169.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.Name = `Verticle in class diagram Default in middle between Default-Line and Default-Point`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.X = 462.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.Y = 261.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point.Name = `Verticle in class diagram Default in middle between Default-Line and Default-Point`

	// Setup of pointers
	__Classdiagram__000000_Toto.GongStructShapes = append(__Classdiagram__000000_Toto.GongStructShapes, __GongStructShape__000000_Default_Line)
	__Classdiagram__000000_Toto.GongStructShapes = append(__Classdiagram__000000_Toto.GongStructShapes, __GongStructShape__000001_Default_Point)
	__GongStructShape__000000_Default_Line.Position = __Position__000000_Pos_Default_Line
	__GongStructShape__000000_Default_Line.Fields = append(__GongStructShape__000000_Default_Line.Fields, __Field__000002_Name)
	__GongStructShape__000000_Default_Line.Fields = append(__GongStructShape__000000_Default_Line.Fields, __Field__000000_CreationDate)
	__GongStructShape__000000_Default_Line.Fields = append(__GongStructShape__000000_Default_Line.Fields, __Field__000001_JourneyTime)
	__GongStructShape__000000_Default_Line.Links = append(__GongStructShape__000000_Default_Line.Links, __Link__000000_End)
	__GongStructShape__000000_Default_Line.Links = append(__GongStructShape__000000_Default_Line.Links, __Link__000001_Start)
	__GongStructShape__000001_Default_Point.Position = __Position__000001_Pos_Default_Point
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point
	__Link__000001_Start.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Line_and_Default_Point
}


