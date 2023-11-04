// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Classdiagram":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Classdiagram Form",
			OnSave: __gong__New__ClassdiagramFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, formGroup, probe)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DiagramPackage Form",
			OnSave: __gong__New__DiagramPackageFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, formGroup, probe)
	case "Field":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Field Form",
			OnSave: __gong__New__FieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		field := new(models.Field)
		FillUpForm(field, formGroup, probe)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumShape Form",
			OnSave: __gong__New__GongEnumShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, formGroup, probe)
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumValueEntry Form",
			OnSave: __gong__New__GongEnumValueEntryFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, formGroup, probe)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongStructShape Form",
			OnSave: __gong__New__GongStructShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Link Form",
			OnSave: __gong__New__LinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " NoteShape Form",
			OnSave: __gong__New__NoteShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, formGroup, probe)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " NoteShapeLink Form",
			OnSave: __gong__New__NoteShapeLinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, formGroup, probe)
	case "Position":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Position Form",
			OnSave: __gong__New__PositionFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		position := new(models.Position)
		FillUpForm(position, formGroup, probe)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " UmlState Form",
			OnSave: __gong__New__UmlStateFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, formGroup, probe)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Umlsc Form",
			OnSave: __gong__New__UmlscFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, formGroup, probe)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Vertice Form",
			OnSave: __gong__New__VerticeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, formGroup, probe)
	}
	formStage.Commit()
}
