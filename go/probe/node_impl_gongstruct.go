// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Classdiagram" {
		fillUpTable[models.Classdiagram](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DiagramPackage" {
		fillUpTable[models.DiagramPackage](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Field" {
		fillUpTable[models.Field](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumShape" {
		fillUpTable[models.GongEnumShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValueEntry" {
		fillUpTable[models.GongEnumValueEntry](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStructShape" {
		fillUpTable[models.GongStructShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShape" {
		fillUpTable[models.NoteShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShapeLink" {
		fillUpTable[models.NoteShapeLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Position" {
		fillUpTable[models.Position](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UmlState" {
		fillUpTable[models.UmlState](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Umlsc" {
		fillUpTable[models.Umlsc](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Vertice" {
		fillUpTable[models.Vertice](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Classdiagram:
		fillUpTable[models.Classdiagram](playground)
	case *models.DiagramPackage:
		fillUpTable[models.DiagramPackage](playground)
	case *models.Field:
		fillUpTable[models.Field](playground)
	case *models.GongEnumShape:
		fillUpTable[models.GongEnumShape](playground)
	case *models.GongEnumValueEntry:
		fillUpTable[models.GongEnumValueEntry](playground)
	case *models.GongStructShape:
		fillUpTable[models.GongStructShape](playground)
	case *models.Link:
		fillUpTable[models.Link](playground)
	case *models.NoteShape:
		fillUpTable[models.NoteShape](playground)
	case *models.NoteShapeLink:
		fillUpTable[models.NoteShapeLink](playground)
	case *models.Position:
		fillUpTable[models.Position](playground)
	case *models.UmlState:
		fillUpTable[models.UmlState](playground)
	case *models.Umlsc:
		fillUpTable[models.Umlsc](playground)
	case *models.Vertice:
		fillUpTable[models.Vertice](playground)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	playground *Playground,
) {

	playground.tableStage.Reset()
	playground.tableStage.Commit()

	table := new(gongtable.Table).Stage(playground.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	playground.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](playground.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			playground.stageOfInterest,
			playground.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(playground.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, playground)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				structInstance,
			),
		}).Stage(playground.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(playground.tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(playground.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(playground.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.playground = playground
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Classdiagram:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewClassdiagramFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.DiagramPackage:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewDiagramPackageFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Field:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFieldFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongEnumShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumShapeFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongEnumValueEntry:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueEntryFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongStructShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructShapeFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.NoteShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.NoteShapeLink:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeLinkFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Position:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPositionFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.UmlState:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewUmlStateFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Umlsc:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewUmlscFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Vertice:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewVerticeFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	}
	formStage.Commit()

}
