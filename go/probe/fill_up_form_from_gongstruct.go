// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Classdiagram Form",
			OnSave: NewClassdiagramFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.DiagramPackage:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update DiagramPackage Form",
			OnSave: NewDiagramPackageFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Field:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Field Form",
			OnSave: NewFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongEnumShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnumShape Form",
			OnSave: NewGongEnumShapeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongEnumValueEntry:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnumValueEntry Form",
			OnSave: NewGongEnumValueEntryFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongStructShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongStructShape Form",
			OnSave: NewGongStructShapeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Link Form",
			OnSave: NewLinkFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.NoteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update NoteShape Form",
			OnSave: NewNoteShapeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.NoteShapeLink:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update NoteShapeLink Form",
			OnSave: NewNoteShapeLinkFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Position:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Position Form",
			OnSave: NewPositionFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.UmlState:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update UmlState Form",
			OnSave: NewUmlStateFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Umlsc:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Umlsc Form",
			OnSave: NewUmlscFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Vertice:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Vertice Form",
			OnSave: NewVerticeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
