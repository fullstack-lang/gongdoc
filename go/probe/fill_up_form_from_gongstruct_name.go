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
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Classdiagram":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Classdiagram Form",
			OnSave: __gong__New__ClassdiagramFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		classdiagram := new(models.Classdiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(classdiagram, formGroup, probe)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DiagramPackage Form",
			OnSave: __gong__New__DiagramPackageFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		diagrampackage := new(models.DiagramPackage)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagrampackage, formGroup, probe)
	case "Field":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Field Form",
			OnSave: __gong__New__FieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		field := new(models.Field)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(field, formGroup, probe)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GongEnumShape Form",
			OnSave: __gong__New__GongEnumShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenumshape := new(models.GongEnumShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumshape, formGroup, probe)
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GongEnumValueEntry Form",
			OnSave: __gong__New__GongEnumValueEntryFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumvalueentry, formGroup, probe)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GongStructShape Form",
			OnSave: __gong__New__GongStructShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongstructshape := new(models.GongStructShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongstructshape, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Link Form",
			OnSave: __gong__New__LinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "NoteShape Form",
			OnSave: __gong__New__NoteShapeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "NoteShapeLink Form",
			OnSave: __gong__New__NoteShapeLinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		noteshapelink := new(models.NoteShapeLink)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshapelink, formGroup, probe)
	case "Position":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Position Form",
			OnSave: __gong__New__PositionFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		position := new(models.Position)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(position, formGroup, probe)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "UmlState Form",
			OnSave: __gong__New__UmlStateFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		umlstate := new(models.UmlState)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(umlstate, formGroup, probe)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Umlsc Form",
			OnSave: __gong__New__UmlscFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		umlsc := new(models.Umlsc)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(umlsc, formGroup, probe)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Vertice Form",
			OnSave: __gong__New__VerticeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vertice, formGroup, probe)
	}
	formStage.Commit()
}
