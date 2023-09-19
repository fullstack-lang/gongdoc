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
			OnSave: NewClassdiagramFormCallback(
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
			OnSave: NewDiagramPackageFormCallback(
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
			OnSave: NewFieldFormCallback(
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
			OnSave: NewGongEnumShapeFormCallback(
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
			OnSave: NewGongEnumValueEntryFormCallback(
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
			OnSave: NewGongStructShapeFormCallback(
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
			OnSave: NewLinkFormCallback(
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
			OnSave: NewNoteShapeFormCallback(
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
			OnSave: NewNoteShapeLinkFormCallback(
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
			OnSave: NewPositionFormCallback(
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
			OnSave: NewUmlStateFormCallback(
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
			OnSave: NewUmlscFormCallback(
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
			OnSave: NewVerticeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, formGroup, playground)
	}
	formStage.Commit()
}
