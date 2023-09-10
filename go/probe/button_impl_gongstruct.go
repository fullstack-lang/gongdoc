// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Classdiagram":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewClassdiagramFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, formGroup, buttonImpl.playground)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDiagramPackageFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, formGroup, buttonImpl.playground)
	case "Field":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		field := new(models.Field)
		FillUpForm(field, formGroup, buttonImpl.playground)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumShapeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, formGroup, buttonImpl.playground)
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueEntryFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, formGroup, buttonImpl.playground)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructShapeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, formGroup, buttonImpl.playground)
	case "Link":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLinkFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, formGroup, buttonImpl.playground)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, formGroup, buttonImpl.playground)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewNoteShapeLinkFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, formGroup, buttonImpl.playground)
	case "Position":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPositionFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		position := new(models.Position)
		FillUpForm(position, formGroup, buttonImpl.playground)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewUmlStateFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, formGroup, buttonImpl.playground)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewUmlscFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, formGroup, buttonImpl.playground)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewVerticeFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, formGroup, buttonImpl.playground)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, formGroup, playground)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, formGroup, playground)
		AssociationSliceToForm("NoteShapes", instanceWithInferedType, &instanceWithInferedType.NoteShapes, formGroup, playground)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup)

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, formGroup, playground)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, formGroup, playground)
		AssociationSliceToForm("Umlscs", instanceWithInferedType, &instanceWithInferedType.Umlscs, formGroup, playground)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsReloaded", instanceWithInferedType.IsReloaded, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongEnumValueEntrys", instanceWithInferedType, &instanceWithInferedType.GongEnumValueEntrys, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongEnumValueEntry:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("NbInstances", instanceWithInferedType.NbInstances, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Fields", instanceWithInferedType, &instanceWithInferedType.Fields, formGroup, playground)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Middlevertice", instanceWithInferedType.Middlevertice, formGroup, playground)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, playground.formStage, formGroup)

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("NoteShapeLinks", instanceWithInferedType, &instanceWithInferedType.NoteShapeLinks, formGroup, playground)

	case *models.NoteShapeLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Position:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.UmlState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Umlsc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, playground)
		BasicFieldtoForm("Activestate", instanceWithInferedType.Activestate, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Vertice:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

