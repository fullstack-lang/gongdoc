// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Classdiagram Form",
			OnSave: __gong__New__ClassdiagramFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DiagramPackage:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "DiagramPackage Form",
			OnSave: __gong__New__DiagramPackageFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Field:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Field Form",
			OnSave: __gong__New__FieldFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "GongEnumShape Form",
			OnSave: __gong__New__GongEnumShapeFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumValueEntry:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "GongEnumValueEntry Form",
			OnSave: __gong__New__GongEnumValueEntryFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongStructShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "GongStructShape Form",
			OnSave: __gong__New__GongStructShapeFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Link Form",
			OnSave: __gong__New__LinkFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "NoteShape Form",
			OnSave: __gong__New__NoteShapeFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShapeLink:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "NoteShapeLink Form",
			OnSave: __gong__New__NoteShapeLinkFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Position:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Position Form",
			OnSave: __gong__New__PositionFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UmlState:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "UmlState Form",
			OnSave: __gong__New__UmlStateFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Umlsc:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Umlsc Form",
			OnSave: __gong__New__UmlscFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Vertice:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Vertice Form",
			OnSave: __gong__New__VerticeFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
