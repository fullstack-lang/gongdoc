package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdoc/go/tests/geometry/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
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
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CONTINUOUS_ZERO := (&models.Field{Name: `CONTINUOUS_ZERO`}).Stage()
	__Field__000001_DOTTED_ONE := (&models.Field{Name: `DOTTED_ONE`}).Stage()
	__Field__000002_JourneyTime := (&models.Field{Name: `JourneyTime`}).Stage()
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000005_Y := (&models.Field{Name: `Y`}).Stage()

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Line := (&models.GongStructShape{Name: `NewDiagram-Line`}).Stage()
	__GongStructShape__000001_NewDiagram_LineTypeInt := (&models.GongStructShape{Name: `NewDiagram-LineTypeInt`}).Stage()
	__GongStructShape__000002_NewDiagram_Point := (&models.GongStructShape{Name: `NewDiagram-Point`}).Stage()
	__GongStructShape__000003_NewDiagram_PointExclusiveSet := (&models.GongStructShape{Name: `NewDiagram-PointExclusiveSet`}).Stage()
	__GongStructShape__000004_NewDiagram_PointUse := (&models.GongStructShape{Name: `NewDiagram-PointUse`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_ShortNodeOnModels := (&models.NoteShape{Name: `ShortNodeOnModels`}).Stage()

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Line := (&models.NoteShapeLink{Name: `Line`}).Stage()

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage()
	__Position__000001_Pos_NewDiagram_LineTypeInt := (&models.Position{Name: `Pos-NewDiagram-LineTypeInt`}).Stage()
	__Position__000002_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage()
	__Position__000003_Pos_NewDiagram_PointExclusiveSet := (&models.Position{Name: `Pos-NewDiagram-PointExclusiveSet`}).Stage()
	__Position__000004_Pos_NewDiagram_PointUse := (&models.Position{Name: `Pos-NewDiagram-PointUse`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_CONTINUOUS_ZERO.Name = `CONTINUOUS_ZERO`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.CONTINUOUS_ZERO]
	__Field__000000_CONTINUOUS_ZERO.Identifier = `ref_models.LineTypeInt.CONTINUOUS_ZERO`
	__Field__000000_CONTINUOUS_ZERO.FieldTypeAsString = ``
	__Field__000000_CONTINUOUS_ZERO.Structname = ``
	__Field__000000_CONTINUOUS_ZERO.Fieldtypename = ``

	// Field values setup
	__Field__000001_DOTTED_ONE.Name = `DOTTED_ONE`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt.DOTTED_ONE]
	__Field__000001_DOTTED_ONE.Identifier = `ref_models.LineTypeInt.DOTTED_ONE`
	__Field__000001_DOTTED_ONE.FieldTypeAsString = ``
	__Field__000001_DOTTED_ONE.Structname = ``
	__Field__000001_DOTTED_ONE.Fieldtypename = ``

	// Field values setup
	__Field__000002_JourneyTime.Name = `JourneyTime`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.JourneyTime]
	__Field__000002_JourneyTime.Identifier = `ref_models.Line.JourneyTime`
	__Field__000002_JourneyTime.FieldTypeAsString = ``
	__Field__000002_JourneyTime.Structname = `Line`
	__Field__000002_JourneyTime.Fieldtypename = `Duration`

	// Field values setup
	__Field__000003_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.Name]
	__Field__000003_Name.Identifier = `ref_models.Point.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Point`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Name]
	__Field__000004_Name.Identifier = `ref_models.Line.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `Line`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Y.Name = `Y`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point.Y]
	__Field__000005_Y.Identifier = `ref_models.Point.Y`
	__Field__000005_Y.FieldTypeAsString = ``
	__Field__000005_Y.Structname = `Point`
	__Field__000005_Y.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Line.Name = `NewDiagram-Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000000_NewDiagram_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_NewDiagram_Line.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Line.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Line.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Line.Heigth = 93.000000
	__GongStructShape__000000_NewDiagram_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_LineTypeInt.Name = `NewDiagram-LineTypeInt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LineTypeInt]
	__GongStructShape__000001_NewDiagram_LineTypeInt.Identifier = `ref_models.LineTypeInt`
	__GongStructShape__000001_NewDiagram_LineTypeInt.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_LineTypeInt.NbInstances = 0
	__GongStructShape__000001_NewDiagram_LineTypeInt.Width = 240.000000
	__GongStructShape__000001_NewDiagram_LineTypeInt.Heigth = 93.000000
	__GongStructShape__000001_NewDiagram_LineTypeInt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Point.Name = `NewDiagram-Point`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000002_NewDiagram_Point.Identifier = `ref_models.Point`
	__GongStructShape__000002_NewDiagram_Point.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Point.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Point.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Point.Heigth = 93.000000
	__GongStructShape__000002_NewDiagram_Point.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.Name = `NewDiagram-PointExclusiveSet`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.PointExclusiveSet]
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.Identifier = `ref_models.PointExclusiveSet`
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.NbInstances = 0
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.Width = 240.000000
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.Heigth = 63.000000
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_PointUse.Name = `NewDiagram-PointUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.PointUse]
	__GongStructShape__000004_NewDiagram_PointUse.Identifier = `ref_models.PointUse`
	__GongStructShape__000004_NewDiagram_PointUse.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_PointUse.NbInstances = 0
	__GongStructShape__000004_NewDiagram_PointUse.Width = 240.000000
	__GongStructShape__000004_NewDiagram_PointUse.Heigth = 63.000000
	__GongStructShape__000004_NewDiagram_PointUse.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`
	__Link__000000_End.Structname = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.End]
	__Link__000000_End.Identifier = `ref_models.Line.End`
	__Link__000000_End.Fieldtypename = `Point`
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.SourceMultiplicity = models.MANY

	// NoteShape values setup
	__NoteShape__000000_ShortNodeOnModels.Name = `ShortNodeOnModels`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ShortNodeOnModels]
	__NoteShape__000000_ShortNodeOnModels.Identifier = `ref_models.ShortNodeOnModels`
	__NoteShape__000000_ShortNodeOnModels.Body = `this is an example of a short note
It uses the DocLink convention for referencing Identifiers
In this case [Line], [Point] and [Line.Start]
`
	__NoteShape__000000_ShortNodeOnModels.X = 80.000000
	__NoteShape__000000_ShortNodeOnModels.Y = 480.000000
	__NoteShape__000000_ShortNodeOnModels.Width = 240.000000
	__NoteShape__000000_ShortNodeOnModels.Heigth = 63.000000
	__NoteShape__000000_ShortNodeOnModels.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Line.Name = `Line`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ShortNodeOnModels.Line]
	__NoteShapeLink__000000_Line.Identifier = `ref_models.ShortNodeOnModels.Line`

	// Position values setup
	__Position__000000_Pos_NewDiagram_Line.X = 340.000000
	__Position__000000_Pos_NewDiagram_Line.Y = 270.000000
	__Position__000000_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000001_Pos_NewDiagram_LineTypeInt.X = 660.000000
	__Position__000001_Pos_NewDiagram_LineTypeInt.Y = 110.000000
	__Position__000001_Pos_NewDiagram_LineTypeInt.Name = `Pos-NewDiagram-LineTypeInt`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Point.X = 60.000000
	__Position__000002_Pos_NewDiagram_Point.Y = 130.000000
	__Position__000002_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Position values setup
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.X = 730.000000
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.Y = 400.000000
	__Position__000003_Pos_NewDiagram_PointExclusiveSet.Name = `Pos-NewDiagram-PointExclusiveSet`

	// Position values setup
	__Position__000004_Pos_NewDiagram_PointUse.X = 30.000000
	__Position__000004_Pos_NewDiagram_PointUse.Y = 330.000000
	__Position__000004_Pos_NewDiagram_PointUse.Name = `Pos-NewDiagram-PointUse`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.X = 540.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Y = 126.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Line and NewDiagram-Point`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Line)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_LineTypeInt)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_PointExclusiveSet)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_PointUse)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Point)
	__Classdiagram__000000_NewDiagram.NoteShapes = append(__Classdiagram__000000_NewDiagram.NoteShapes, __NoteShape__000000_ShortNodeOnModels)
	__GongStructShape__000000_NewDiagram_Line.Position = __Position__000000_Pos_NewDiagram_Line
	__GongStructShape__000000_NewDiagram_Line.Fields = append(__GongStructShape__000000_NewDiagram_Line.Fields, __Field__000004_Name)
	__GongStructShape__000000_NewDiagram_Line.Fields = append(__GongStructShape__000000_NewDiagram_Line.Fields, __Field__000002_JourneyTime)
	__GongStructShape__000000_NewDiagram_Line.Links = append(__GongStructShape__000000_NewDiagram_Line.Links, __Link__000000_End)
	__GongStructShape__000001_NewDiagram_LineTypeInt.Position = __Position__000001_Pos_NewDiagram_LineTypeInt
	__GongStructShape__000001_NewDiagram_LineTypeInt.Fields = append(__GongStructShape__000001_NewDiagram_LineTypeInt.Fields, __Field__000000_CONTINUOUS_ZERO)
	__GongStructShape__000001_NewDiagram_LineTypeInt.Fields = append(__GongStructShape__000001_NewDiagram_LineTypeInt.Fields, __Field__000001_DOTTED_ONE)
	__GongStructShape__000002_NewDiagram_Point.Position = __Position__000002_Pos_NewDiagram_Point
	__GongStructShape__000002_NewDiagram_Point.Fields = append(__GongStructShape__000002_NewDiagram_Point.Fields, __Field__000003_Name)
	__GongStructShape__000002_NewDiagram_Point.Fields = append(__GongStructShape__000002_NewDiagram_Point.Fields, __Field__000005_Y)
	__GongStructShape__000003_NewDiagram_PointExclusiveSet.Position = __Position__000003_Pos_NewDiagram_PointExclusiveSet
	__GongStructShape__000004_NewDiagram_PointUse.Position = __Position__000004_Pos_NewDiagram_PointUse
	__Link__000000_End.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Line_and_NewDiagram_Point
	__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks = append(__NoteShape__000000_ShortNodeOnModels.NoteShapeLinks, __NoteShapeLink__000000_Line)
}


