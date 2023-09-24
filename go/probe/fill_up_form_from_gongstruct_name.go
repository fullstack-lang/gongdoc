// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
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
				playground,
			),
		}).Stage(formStage)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, formGroup, playground)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DiagramPackage Form",
			OnSave: __gong__New__DiagramPackageFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, formGroup, playground)
	case "Field":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Field Form",
			OnSave: __gong__New__FieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		field := new(models.Field)
		FillUpForm(field, formGroup, playground)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumShape Form",
			OnSave: __gong__New__GongEnumShapeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, formGroup, playground)
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumValueEntry Form",
			OnSave: __gong__New__GongEnumValueEntryFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, formGroup, playground)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongStructShape Form",
			OnSave: __gong__New__GongStructShapeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, formGroup, playground)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Link Form",
			OnSave: __gong__New__LinkFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, formGroup, playground)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " NoteShape Form",
			OnSave: __gong__New__NoteShapeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, formGroup, playground)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " NoteShapeLink Form",
			OnSave: __gong__New__NoteShapeLinkFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, formGroup, playground)
	case "Position":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Position Form",
			OnSave: __gong__New__PositionFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		position := new(models.Position)
		FillUpForm(position, formGroup, playground)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " UmlState Form",
			OnSave: __gong__New__UmlStateFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, formGroup, playground)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Umlsc Form",
			OnSave: __gong__New__UmlscFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, formGroup, playground)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Vertice Form",
			OnSave: __gong__New__VerticeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, formGroup, playground)
	}
	formStage.Commit()
}
