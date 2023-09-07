// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

type ButtonImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	Icon               gongtree_buttons.ButtonType
	tableStage         *form.StageStruct
	formStage          *form.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	tableStage *form.StageStruct,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.tableStage = tableStage
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r
	buttonImplGongstruct.backRepoOfInterest = backRepoOfInterest

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Classdiagram":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewClassdiagramFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDiagramPackageFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Field":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		field := new(models.Field)
		FillUpForm(field, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumShapeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueEntryFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructShapeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Link":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeLinkFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Position":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPositionFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		position := new(models.Position)
		FillUpForm(position, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewUmlStateFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewUmlscFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewVerticeFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	stageOfInterest *models.StageStruct,
	formStage *form.StageStruct,
	formGroup *form.FormGroup,
	r *gin.Engine,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("NoteShapes", instanceWithInferedType, &instanceWithInferedType.NoteShapes, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, formStage, formGroup)

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, stageOfInterest, formStage, formGroup, r)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, stageOfInterest, formStage, formGroup)
		AssociationSliceToForm("Umlscs", instanceWithInferedType, &instanceWithInferedType.Umlscs, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("IsReloaded", instanceWithInferedType.IsReloaded, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, formStage, formGroup)

	case *models.Field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, formStage, formGroup)

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("GongEnumValueEntrys", instanceWithInferedType, &instanceWithInferedType.GongEnumValueEntrys, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, formStage, formGroup)

	case *models.GongEnumValueEntry:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("NbInstances", instanceWithInferedType.NbInstances, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Fields", instanceWithInferedType, &instanceWithInferedType.Fields, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, formStage, formGroup)

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Middlevertice", instanceWithInferedType.Middlevertice, stageOfInterest, formStage, formGroup)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, formStage, formGroup)

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("NoteShapeLinks", instanceWithInferedType, &instanceWithInferedType.NoteShapeLinks, stageOfInterest, formStage, formGroup, r)

	case *models.NoteShapeLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, formStage, formGroup)

	case *models.Position:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.UmlState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)

	case *models.Umlsc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("Activestate", instanceWithInferedType.Activestate, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, formStage, formGroup)

	case *models.Vertice:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

