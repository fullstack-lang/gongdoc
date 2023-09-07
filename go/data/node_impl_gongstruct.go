// generated code - do not edit
package data

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	r                  *gin.Engine
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct,
	r *gin.Engine,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.tableStage = tableStage
	nodeImplGongstruct.formStage = formStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	nodeImplGongstruct.backRepoOfInterest = backRepoOfInterest
	nodeImplGongstruct.r = r
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

	tableStage := nodeImplGongstruct.tableStage

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Classdiagram" {
		fillUpTable[models.Classdiagram](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DiagramPackage" {
		fillUpTable[models.DiagramPackage](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Field" {
		fillUpTable[models.Field](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumShape" {
		fillUpTable[models.GongEnumShape](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValueEntry" {
		fillUpTable[models.GongEnumValueEntry](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStructShape" {
		fillUpTable[models.GongStructShape](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShape" {
		fillUpTable[models.NoteShape](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShapeLink" {
		fillUpTable[models.NoteShapeLink](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Position" {
		fillUpTable[models.Position](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UmlState" {
		fillUpTable[models.UmlState](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Umlsc" {
		fillUpTable[models.Umlsc](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Vertice" {
		fillUpTable[models.Vertice](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	stageOfInterest *models.StageStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) {

	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			stageOfInterest,
			backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance,
			tableStage,
			formStage,
			stageOfInterest,
			r,
			backRepoOfInterest)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				structInstance,
			),
		}).Stage(tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.tableStage = tableStage
	rowUpdate.formStage = formStage
	rowUpdate.stageOfInterest = stageOfInterest
	rowUpdate.r = r
	rowUpdate.backRepoOfInterest = backRepoOfInterest
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance           *T
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Classdiagram:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewClassdiagramFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.DiagramPackage:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewDiagramPackageFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Field:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongEnumShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumShapeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongEnumValueEntry:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueEntryFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongStructShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructShapeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.NoteShape:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.NoteShapeLink:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeLinkFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Position:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPositionFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.UmlState:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewUmlStateFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Umlsc:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewUmlscFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Vertice:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewVerticeFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	}
	formStage.Commit()

}
