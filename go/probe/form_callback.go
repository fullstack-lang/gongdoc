// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ClassdiagramFormCallback(
	classdiagram *models.Classdiagram,
	probe *Probe,
) (classdiagramFormCallback *ClassdiagramFormCallback) {
	classdiagramFormCallback = new(ClassdiagramFormCallback)
	classdiagramFormCallback.probe = probe
	classdiagramFormCallback.classdiagram = classdiagram

	classdiagramFormCallback.CreationMode = (classdiagram == nil)

	return
}

type ClassdiagramFormCallback struct {
	classdiagram *models.Classdiagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (classdiagramFormCallback *ClassdiagramFormCallback) OnSave() {

	log.Println("ClassdiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.probe.formStage.Checkout()

	if classdiagramFormCallback.classdiagram == nil {
		classdiagramFormCallback.classdiagram = new(models.Classdiagram).Stage(classdiagramFormCallback.probe.stageOfInterest)
	}
	classdiagram_ := classdiagramFormCallback.classdiagram
	_ = classdiagram_

	// get the formGroup
	formGroup := classdiagramFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(classdiagram_.Name), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInDrawMode), formDiv)
		case "DiagramPackage:Classdiagrams":
			// we need to retrieve the field owner before the change
			var pastDiagramPackageOwner *models.DiagramPackage
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Classdiagrams"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				classdiagramFormCallback.probe.stageOfInterest,
				classdiagramFormCallback.probe.backRepoOfInterest,
				classdiagram_,
				&rf)

			if reverseFieldOwner != nil {
				pastDiagramPackageOwner = reverseFieldOwner.(*models.DiagramPackage)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDiagramPackageOwner != nil {
					idx := slices.Index(pastDiagramPackageOwner.Classdiagrams, classdiagram_)
					pastDiagramPackageOwner.Classdiagrams = slices.Delete(pastDiagramPackageOwner.Classdiagrams, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _diagrampackage.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDiagramPackageOwner := _diagrampackage // we have a match
						if pastDiagramPackageOwner != nil {
							if newDiagramPackageOwner != pastDiagramPackageOwner {
								idx := slices.Index(pastDiagramPackageOwner.Classdiagrams, classdiagram_)
								pastDiagramPackageOwner.Classdiagrams = slices.Delete(pastDiagramPackageOwner.Classdiagrams, idx, idx+1)
								newDiagramPackageOwner.Classdiagrams = append(newDiagramPackageOwner.Classdiagrams, classdiagram_)
							}
						} else {
							newDiagramPackageOwner.Classdiagrams = append(newDiagramPackageOwner.Classdiagrams, classdiagram_)
						}
					}
				}
			}
		}
	}

	classdiagramFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Classdiagram](
		classdiagramFormCallback.probe,
	)
	classdiagramFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if classdiagramFormCallback.CreationMode {
		classdiagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ClassdiagramFormCallback(
				nil,
				classdiagramFormCallback.probe,
			),
		}).Stage(classdiagramFormCallback.probe.formStage)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, newFormGroup, classdiagramFormCallback.probe)
		classdiagramFormCallback.probe.formStage.Commit()
	}

	fillUpTree(classdiagramFormCallback.probe)
}
func __gong__New__DiagramPackageFormCallback(
	diagrampackage *models.DiagramPackage,
	probe *Probe,
) (diagrampackageFormCallback *DiagramPackageFormCallback) {
	diagrampackageFormCallback = new(DiagramPackageFormCallback)
	diagrampackageFormCallback.probe = probe
	diagrampackageFormCallback.diagrampackage = diagrampackage

	diagrampackageFormCallback.CreationMode = (diagrampackage == nil)

	return
}

type DiagramPackageFormCallback struct {
	diagrampackage *models.DiagramPackage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (diagrampackageFormCallback *DiagramPackageFormCallback) OnSave() {

	log.Println("DiagramPackageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagrampackageFormCallback.probe.formStage.Checkout()

	if diagrampackageFormCallback.diagrampackage == nil {
		diagrampackageFormCallback.diagrampackage = new(models.DiagramPackage).Stage(diagrampackageFormCallback.probe.stageOfInterest)
	}
	diagrampackage_ := diagrampackageFormCallback.diagrampackage
	_ = diagrampackage_

	// get the formGroup
	formGroup := diagrampackageFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagrampackage_.Name), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(diagrampackage_.Path), formDiv)
		case "GongModelPath":
			FormDivBasicFieldToField(&(diagrampackage_.GongModelPath), formDiv)
		case "SelectedClassdiagram":
			FormDivSelectFieldToField(&(diagrampackage_.SelectedClassdiagram), diagrampackageFormCallback.probe.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(diagrampackage_.IsEditable), formDiv)
		case "IsReloaded":
			FormDivBasicFieldToField(&(diagrampackage_.IsReloaded), formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	diagrampackageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DiagramPackage](
		diagrampackageFormCallback.probe,
	)
	diagrampackageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if diagrampackageFormCallback.CreationMode {
		diagrampackageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DiagramPackageFormCallback(
				nil,
				diagrampackageFormCallback.probe,
			),
		}).Stage(diagrampackageFormCallback.probe.formStage)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, newFormGroup, diagrampackageFormCallback.probe)
		diagrampackageFormCallback.probe.formStage.Commit()
	}

	fillUpTree(diagrampackageFormCallback.probe)
}
func __gong__New__FieldFormCallback(
	field *models.Field,
	probe *Probe,
) (fieldFormCallback *FieldFormCallback) {
	fieldFormCallback = new(FieldFormCallback)
	fieldFormCallback.probe = probe
	fieldFormCallback.field = field

	fieldFormCallback.CreationMode = (field == nil)

	return
}

type FieldFormCallback struct {
	field *models.Field

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (fieldFormCallback *FieldFormCallback) OnSave() {

	log.Println("FieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fieldFormCallback.probe.formStage.Checkout()

	if fieldFormCallback.field == nil {
		fieldFormCallback.field = new(models.Field).Stage(fieldFormCallback.probe.stageOfInterest)
	}
	field_ := fieldFormCallback.field
	_ = field_

	// get the formGroup
	formGroup := fieldFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(field_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(field_.Identifier), formDiv)
		case "FieldTypeAsString":
			FormDivBasicFieldToField(&(field_.FieldTypeAsString), formDiv)
		case "Structname":
			FormDivBasicFieldToField(&(field_.Structname), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(field_.Fieldtypename), formDiv)
		case "GongStructShape:Fields":
			// we need to retrieve the field owner before the change
			var pastGongStructShapeOwner *models.GongStructShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Fields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				fieldFormCallback.probe.stageOfInterest,
				fieldFormCallback.probe.backRepoOfInterest,
				field_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructShapeOwner = reverseFieldOwner.(*models.GongStructShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.Fields, field_)
					pastGongStructShapeOwner.Fields = slices.Delete(pastGongStructShapeOwner.Fields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](fieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstructshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructShapeOwner := _gongstructshape // we have a match
						if pastGongStructShapeOwner != nil {
							if newGongStructShapeOwner != pastGongStructShapeOwner {
								idx := slices.Index(pastGongStructShapeOwner.Fields, field_)
								pastGongStructShapeOwner.Fields = slices.Delete(pastGongStructShapeOwner.Fields, idx, idx+1)
								newGongStructShapeOwner.Fields = append(newGongStructShapeOwner.Fields, field_)
							}
						} else {
							newGongStructShapeOwner.Fields = append(newGongStructShapeOwner.Fields, field_)
						}
					}
				}
			}
		}
	}

	fieldFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Field](
		fieldFormCallback.probe,
	)
	fieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fieldFormCallback.CreationMode {
		fieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FieldFormCallback(
				nil,
				fieldFormCallback.probe,
			),
		}).Stage(fieldFormCallback.probe.formStage)
		field := new(models.Field)
		FillUpForm(field, newFormGroup, fieldFormCallback.probe)
		fieldFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fieldFormCallback.probe)
}
func __gong__New__GongEnumShapeFormCallback(
	gongenumshape *models.GongEnumShape,
	probe *Probe,
) (gongenumshapeFormCallback *GongEnumShapeFormCallback) {
	gongenumshapeFormCallback = new(GongEnumShapeFormCallback)
	gongenumshapeFormCallback.probe = probe
	gongenumshapeFormCallback.gongenumshape = gongenumshape

	gongenumshapeFormCallback.CreationMode = (gongenumshape == nil)

	return
}

type GongEnumShapeFormCallback struct {
	gongenumshape *models.GongEnumShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (gongenumshapeFormCallback *GongEnumShapeFormCallback) OnSave() {

	log.Println("GongEnumShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumshapeFormCallback.probe.formStage.Checkout()

	if gongenumshapeFormCallback.gongenumshape == nil {
		gongenumshapeFormCallback.gongenumshape = new(models.GongEnumShape).Stage(gongenumshapeFormCallback.probe.stageOfInterest)
	}
	gongenumshape_ := gongenumshapeFormCallback.gongenumshape
	_ = gongenumshape_

	// get the formGroup
	formGroup := gongenumshapeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongenumshape_.Position), gongenumshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumshape_.Identifier), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(gongenumshape_.Heigth), formDiv)
		case "Classdiagram:GongEnumShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongEnumShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongenumshapeFormCallback.probe.stageOfInterest,
				gongenumshapeFormCallback.probe.backRepoOfInterest,
				gongenumshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongEnumShapes, gongenumshape_)
					pastClassdiagramOwner.GongEnumShapes = slices.Delete(pastClassdiagramOwner.GongEnumShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
						if pastClassdiagramOwner != nil {
							if newClassdiagramOwner != pastClassdiagramOwner {
								idx := slices.Index(pastClassdiagramOwner.GongEnumShapes, gongenumshape_)
								pastClassdiagramOwner.GongEnumShapes = slices.Delete(pastClassdiagramOwner.GongEnumShapes, idx, idx+1)
								newClassdiagramOwner.GongEnumShapes = append(newClassdiagramOwner.GongEnumShapes, gongenumshape_)
							}
						} else {
							newClassdiagramOwner.GongEnumShapes = append(newClassdiagramOwner.GongEnumShapes, gongenumshape_)
						}
					}
				}
			}
		}
	}

	gongenumshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongEnumShape](
		gongenumshapeFormCallback.probe,
	)
	gongenumshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumshapeFormCallback.CreationMode {
		gongenumshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongEnumShapeFormCallback(
				nil,
				gongenumshapeFormCallback.probe,
			),
		}).Stage(gongenumshapeFormCallback.probe.formStage)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, newFormGroup, gongenumshapeFormCallback.probe)
		gongenumshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongenumshapeFormCallback.probe)
}
func __gong__New__GongEnumValueEntryFormCallback(
	gongenumvalueentry *models.GongEnumValueEntry,
	probe *Probe,
) (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) {
	gongenumvalueentryFormCallback = new(GongEnumValueEntryFormCallback)
	gongenumvalueentryFormCallback.probe = probe
	gongenumvalueentryFormCallback.gongenumvalueentry = gongenumvalueentry

	gongenumvalueentryFormCallback.CreationMode = (gongenumvalueentry == nil)

	return
}

type GongEnumValueEntryFormCallback struct {
	gongenumvalueentry *models.GongEnumValueEntry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) OnSave() {

	log.Println("GongEnumValueEntryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueentryFormCallback.probe.formStage.Checkout()

	if gongenumvalueentryFormCallback.gongenumvalueentry == nil {
		gongenumvalueentryFormCallback.gongenumvalueentry = new(models.GongEnumValueEntry).Stage(gongenumvalueentryFormCallback.probe.stageOfInterest)
	}
	gongenumvalueentry_ := gongenumvalueentryFormCallback.gongenumvalueentry
	_ = gongenumvalueentry_

	// get the formGroup
	formGroup := gongenumvalueentryFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Identifier), formDiv)
		case "GongEnumShape:GongEnumValueEntrys":
			// we need to retrieve the field owner before the change
			var pastGongEnumShapeOwner *models.GongEnumShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnumShape"
			rf.Fieldname = "GongEnumValueEntrys"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongenumvalueentryFormCallback.probe.stageOfInterest,
				gongenumvalueentryFormCallback.probe.backRepoOfInterest,
				gongenumvalueentry_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongEnumShapeOwner = reverseFieldOwner.(*models.GongEnumShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongEnumShapeOwner != nil {
					idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
					pastGongEnumShapeOwner.GongEnumValueEntrys = slices.Delete(pastGongEnumShapeOwner.GongEnumValueEntrys, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongenumshape := range *models.GetGongstructInstancesSet[models.GongEnumShape](gongenumvalueentryFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongenumshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongEnumShapeOwner := _gongenumshape // we have a match
						if pastGongEnumShapeOwner != nil {
							if newGongEnumShapeOwner != pastGongEnumShapeOwner {
								idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
								pastGongEnumShapeOwner.GongEnumValueEntrys = slices.Delete(pastGongEnumShapeOwner.GongEnumValueEntrys, idx, idx+1)
								newGongEnumShapeOwner.GongEnumValueEntrys = append(newGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
							}
						} else {
							newGongEnumShapeOwner.GongEnumValueEntrys = append(newGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
						}
					}
				}
			}
		}
	}

	gongenumvalueentryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValueEntry](
		gongenumvalueentryFormCallback.probe,
	)
	gongenumvalueentryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueentryFormCallback.CreationMode {
		gongenumvalueentryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongEnumValueEntryFormCallback(
				nil,
				gongenumvalueentryFormCallback.probe,
			),
		}).Stage(gongenumvalueentryFormCallback.probe.formStage)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, newFormGroup, gongenumvalueentryFormCallback.probe)
		gongenumvalueentryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongenumvalueentryFormCallback.probe)
}
func __gong__New__GongStructShapeFormCallback(
	gongstructshape *models.GongStructShape,
	probe *Probe,
) (gongstructshapeFormCallback *GongStructShapeFormCallback) {
	gongstructshapeFormCallback = new(GongStructShapeFormCallback)
	gongstructshapeFormCallback.probe = probe
	gongstructshapeFormCallback.gongstructshape = gongstructshape

	gongstructshapeFormCallback.CreationMode = (gongstructshape == nil)

	return
}

type GongStructShapeFormCallback struct {
	gongstructshape *models.GongStructShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (gongstructshapeFormCallback *GongStructShapeFormCallback) OnSave() {

	log.Println("GongStructShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructshapeFormCallback.probe.formStage.Checkout()

	if gongstructshapeFormCallback.gongstructshape == nil {
		gongstructshapeFormCallback.gongstructshape = new(models.GongStructShape).Stage(gongstructshapeFormCallback.probe.stageOfInterest)
	}
	gongstructshape_ := gongstructshapeFormCallback.gongstructshape
	_ = gongstructshape_

	// get the formGroup
	formGroup := gongstructshapeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstructshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongstructshape_.Position), gongstructshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongstructshape_.Identifier), formDiv)
		case "ShowNbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.ShowNbInstances), formDiv)
		case "NbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.NbInstances), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongstructshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongstructshape_.Height), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(gongstructshape_.IsSelected), formDiv)
		case "Classdiagram:GongStructShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongStructShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				gongstructshapeFormCallback.probe.stageOfInterest,
				gongstructshapeFormCallback.probe.backRepoOfInterest,
				gongstructshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongStructShapes, gongstructshape_)
					pastClassdiagramOwner.GongStructShapes = slices.Delete(pastClassdiagramOwner.GongStructShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
						if pastClassdiagramOwner != nil {
							if newClassdiagramOwner != pastClassdiagramOwner {
								idx := slices.Index(pastClassdiagramOwner.GongStructShapes, gongstructshape_)
								pastClassdiagramOwner.GongStructShapes = slices.Delete(pastClassdiagramOwner.GongStructShapes, idx, idx+1)
								newClassdiagramOwner.GongStructShapes = append(newClassdiagramOwner.GongStructShapes, gongstructshape_)
							}
						} else {
							newClassdiagramOwner.GongStructShapes = append(newClassdiagramOwner.GongStructShapes, gongstructshape_)
						}
					}
				}
			}
		}
	}

	gongstructshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongStructShape](
		gongstructshapeFormCallback.probe,
	)
	gongstructshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongstructshapeFormCallback.CreationMode {
		gongstructshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongStructShapeFormCallback(
				nil,
				gongstructshapeFormCallback.probe,
			),
		}).Stage(gongstructshapeFormCallback.probe.formStage)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, newFormGroup, gongstructshapeFormCallback.probe)
		gongstructshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongstructshapeFormCallback.probe)
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(link_.Identifier), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(link_.Fieldtypename), formDiv)
		case "FieldOffsetX":
			FormDivBasicFieldToField(&(link_.FieldOffsetX), formDiv)
		case "FieldOffsetY":
			FormDivBasicFieldToField(&(link_.FieldOffsetY), formDiv)
		case "TargetMultiplicity":
			FormDivEnumStringFieldToField(&(link_.TargetMultiplicity), formDiv)
		case "TargetMultiplicityOffsetX":
			FormDivBasicFieldToField(&(link_.TargetMultiplicityOffsetX), formDiv)
		case "TargetMultiplicityOffsetY":
			FormDivBasicFieldToField(&(link_.TargetMultiplicityOffsetY), formDiv)
		case "SourceMultiplicity":
			FormDivEnumStringFieldToField(&(link_.SourceMultiplicity), formDiv)
		case "SourceMultiplicityOffsetX":
			FormDivBasicFieldToField(&(link_.SourceMultiplicityOffsetX), formDiv)
		case "SourceMultiplicityOffsetY":
			FormDivBasicFieldToField(&(link_.SourceMultiplicityOffsetY), formDiv)
		case "Middlevertice":
			FormDivSelectFieldToField(&(link_.Middlevertice), linkFormCallback.probe.stageOfInterest, formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "GongStructShape:Links":
			// we need to retrieve the field owner before the change
			var pastGongStructShapeOwner *models.GongStructShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkFormCallback.probe.stageOfInterest,
				linkFormCallback.probe.backRepoOfInterest,
				link_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructShapeOwner = reverseFieldOwner.(*models.GongStructShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.Links, link_)
					pastGongStructShapeOwner.Links = slices.Delete(pastGongStructShapeOwner.Links, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](linkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstructshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructShapeOwner := _gongstructshape // we have a match
						if pastGongStructShapeOwner != nil {
							if newGongStructShapeOwner != pastGongStructShapeOwner {
								idx := slices.Index(pastGongStructShapeOwner.Links, link_)
								pastGongStructShapeOwner.Links = slices.Delete(pastGongStructShapeOwner.Links, idx, idx+1)
								newGongStructShapeOwner.Links = append(newGongStructShapeOwner.Links, link_)
							}
						} else {
							newGongStructShapeOwner.Links = append(newGongStructShapeOwner.Links, link_)
						}
					}
				}
			}
		}
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.probe,
	)
	linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LinkFormCallback(
				nil,
				linkFormCallback.probe,
			),
		}).Stage(linkFormCallback.probe.formStage)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(linkFormCallback.probe)
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.probe = probe
	noteshapeFormCallback.noteshape = noteshape

	noteshapeFormCallback.CreationMode = (noteshape == nil)

	return
}

type NoteShapeFormCallback struct {
	noteshape *models.NoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {

	log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.probe.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.probe.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	// get the formGroup
	formGroup := noteshapeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(noteshape_.Identifier), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(noteshape_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(noteshape_.BodyHTML), formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(noteshape_.Heigth), formDiv)
		case "Matched":
			FormDivBasicFieldToField(&(noteshape_.Matched), formDiv)
		case "Classdiagram:NoteShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "NoteShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				noteshapeFormCallback.probe.stageOfInterest,
				noteshapeFormCallback.probe.backRepoOfInterest,
				noteshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.NoteShapes, noteshape_)
					pastClassdiagramOwner.NoteShapes = slices.Delete(pastClassdiagramOwner.NoteShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](noteshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
						if pastClassdiagramOwner != nil {
							if newClassdiagramOwner != pastClassdiagramOwner {
								idx := slices.Index(pastClassdiagramOwner.NoteShapes, noteshape_)
								pastClassdiagramOwner.NoteShapes = slices.Delete(pastClassdiagramOwner.NoteShapes, idx, idx+1)
								newClassdiagramOwner.NoteShapes = append(newClassdiagramOwner.NoteShapes, noteshape_)
							}
						} else {
							newClassdiagramOwner.NoteShapes = append(newClassdiagramOwner.NoteShapes, noteshape_)
						}
					}
				}
			}
		}
	}

	noteshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.NoteShape](
		noteshapeFormCallback.probe,
	)
	noteshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteshapeFormCallback.CreationMode {
		noteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NoteShapeFormCallback(
				nil,
				noteshapeFormCallback.probe,
			),
		}).Stage(noteshapeFormCallback.probe.formStage)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, newFormGroup, noteshapeFormCallback.probe)
		noteshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteshapeFormCallback.probe)
}
func __gong__New__NoteShapeLinkFormCallback(
	noteshapelink *models.NoteShapeLink,
	probe *Probe,
) (noteshapelinkFormCallback *NoteShapeLinkFormCallback) {
	noteshapelinkFormCallback = new(NoteShapeLinkFormCallback)
	noteshapelinkFormCallback.probe = probe
	noteshapelinkFormCallback.noteshapelink = noteshapelink

	noteshapelinkFormCallback.CreationMode = (noteshapelink == nil)

	return
}

type NoteShapeLinkFormCallback struct {
	noteshapelink *models.NoteShapeLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (noteshapelinkFormCallback *NoteShapeLinkFormCallback) OnSave() {

	log.Println("NoteShapeLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapelinkFormCallback.probe.formStage.Checkout()

	if noteshapelinkFormCallback.noteshapelink == nil {
		noteshapelinkFormCallback.noteshapelink = new(models.NoteShapeLink).Stage(noteshapelinkFormCallback.probe.stageOfInterest)
	}
	noteshapelink_ := noteshapelinkFormCallback.noteshapelink
	_ = noteshapelink_

	// get the formGroup
	formGroup := noteshapelinkFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshapelink_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(noteshapelink_.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(noteshapelink_.Type), formDiv)
		case "NoteShape:NoteShapeLinks":
			// we need to retrieve the field owner before the change
			var pastNoteShapeOwner *models.NoteShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "NoteShape"
			rf.Fieldname = "NoteShapeLinks"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				noteshapelinkFormCallback.probe.stageOfInterest,
				noteshapelinkFormCallback.probe.backRepoOfInterest,
				noteshapelink_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteShapeOwner = reverseFieldOwner.(*models.NoteShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteShapeOwner != nil {
					idx := slices.Index(pastNoteShapeOwner.NoteShapeLinks, noteshapelink_)
					pastNoteShapeOwner.NoteShapeLinks = slices.Delete(pastNoteShapeOwner.NoteShapeLinks, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _noteshape := range *models.GetGongstructInstancesSet[models.NoteShape](noteshapelinkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _noteshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteShapeOwner := _noteshape // we have a match
						if pastNoteShapeOwner != nil {
							if newNoteShapeOwner != pastNoteShapeOwner {
								idx := slices.Index(pastNoteShapeOwner.NoteShapeLinks, noteshapelink_)
								pastNoteShapeOwner.NoteShapeLinks = slices.Delete(pastNoteShapeOwner.NoteShapeLinks, idx, idx+1)
								newNoteShapeOwner.NoteShapeLinks = append(newNoteShapeOwner.NoteShapeLinks, noteshapelink_)
							}
						} else {
							newNoteShapeOwner.NoteShapeLinks = append(newNoteShapeOwner.NoteShapeLinks, noteshapelink_)
						}
					}
				}
			}
		}
	}

	noteshapelinkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.NoteShapeLink](
		noteshapelinkFormCallback.probe,
	)
	noteshapelinkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteshapelinkFormCallback.CreationMode {
		noteshapelinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NoteShapeLinkFormCallback(
				nil,
				noteshapelinkFormCallback.probe,
			),
		}).Stage(noteshapelinkFormCallback.probe.formStage)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, newFormGroup, noteshapelinkFormCallback.probe)
		noteshapelinkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteshapelinkFormCallback.probe)
}
func __gong__New__PositionFormCallback(
	position *models.Position,
	probe *Probe,
) (positionFormCallback *PositionFormCallback) {
	positionFormCallback = new(PositionFormCallback)
	positionFormCallback.probe = probe
	positionFormCallback.position = position

	positionFormCallback.CreationMode = (position == nil)

	return
}

type PositionFormCallback struct {
	position *models.Position

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (positionFormCallback *PositionFormCallback) OnSave() {

	log.Println("PositionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	positionFormCallback.probe.formStage.Checkout()

	if positionFormCallback.position == nil {
		positionFormCallback.position = new(models.Position).Stage(positionFormCallback.probe.stageOfInterest)
	}
	position_ := positionFormCallback.position
	_ = position_

	// get the formGroup
	formGroup := positionFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "X":
			FormDivBasicFieldToField(&(position_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(position_.Y), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(position_.Name), formDiv)
		}
	}

	positionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Position](
		positionFormCallback.probe,
	)
	positionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if positionFormCallback.CreationMode {
		positionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PositionFormCallback(
				nil,
				positionFormCallback.probe,
			),
		}).Stage(positionFormCallback.probe.formStage)
		position := new(models.Position)
		FillUpForm(position, newFormGroup, positionFormCallback.probe)
		positionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(positionFormCallback.probe)
}
func __gong__New__UmlStateFormCallback(
	umlstate *models.UmlState,
	probe *Probe,
) (umlstateFormCallback *UmlStateFormCallback) {
	umlstateFormCallback = new(UmlStateFormCallback)
	umlstateFormCallback.probe = probe
	umlstateFormCallback.umlstate = umlstate

	umlstateFormCallback.CreationMode = (umlstate == nil)

	return
}

type UmlStateFormCallback struct {
	umlstate *models.UmlState

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (umlstateFormCallback *UmlStateFormCallback) OnSave() {

	log.Println("UmlStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlstateFormCallback.probe.formStage.Checkout()

	if umlstateFormCallback.umlstate == nil {
		umlstateFormCallback.umlstate = new(models.UmlState).Stage(umlstateFormCallback.probe.stageOfInterest)
	}
	umlstate_ := umlstateFormCallback.umlstate
	_ = umlstate_

	// get the formGroup
	formGroup := umlstateFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlstate_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(umlstate_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(umlstate_.Y), formDiv)
		case "Umlsc:States":
			// we need to retrieve the field owner before the change
			var pastUmlscOwner *models.Umlsc
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Umlsc"
			rf.Fieldname = "States"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				umlstateFormCallback.probe.stageOfInterest,
				umlstateFormCallback.probe.backRepoOfInterest,
				umlstate_,
				&rf)

			if reverseFieldOwner != nil {
				pastUmlscOwner = reverseFieldOwner.(*models.Umlsc)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastUmlscOwner != nil {
					idx := slices.Index(pastUmlscOwner.States, umlstate_)
					pastUmlscOwner.States = slices.Delete(pastUmlscOwner.States, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _umlsc := range *models.GetGongstructInstancesSet[models.Umlsc](umlstateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _umlsc.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newUmlscOwner := _umlsc // we have a match
						if pastUmlscOwner != nil {
							if newUmlscOwner != pastUmlscOwner {
								idx := slices.Index(pastUmlscOwner.States, umlstate_)
								pastUmlscOwner.States = slices.Delete(pastUmlscOwner.States, idx, idx+1)
								newUmlscOwner.States = append(newUmlscOwner.States, umlstate_)
							}
						} else {
							newUmlscOwner.States = append(newUmlscOwner.States, umlstate_)
						}
					}
				}
			}
		}
	}

	umlstateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.UmlState](
		umlstateFormCallback.probe,
	)
	umlstateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if umlstateFormCallback.CreationMode {
		umlstateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__UmlStateFormCallback(
				nil,
				umlstateFormCallback.probe,
			),
		}).Stage(umlstateFormCallback.probe.formStage)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, newFormGroup, umlstateFormCallback.probe)
		umlstateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(umlstateFormCallback.probe)
}
func __gong__New__UmlscFormCallback(
	umlsc *models.Umlsc,
	probe *Probe,
) (umlscFormCallback *UmlscFormCallback) {
	umlscFormCallback = new(UmlscFormCallback)
	umlscFormCallback.probe = probe
	umlscFormCallback.umlsc = umlsc

	umlscFormCallback.CreationMode = (umlsc == nil)

	return
}

type UmlscFormCallback struct {
	umlsc *models.Umlsc

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (umlscFormCallback *UmlscFormCallback) OnSave() {

	log.Println("UmlscFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlscFormCallback.probe.formStage.Checkout()

	if umlscFormCallback.umlsc == nil {
		umlscFormCallback.umlsc = new(models.Umlsc).Stage(umlscFormCallback.probe.stageOfInterest)
	}
	umlsc_ := umlscFormCallback.umlsc
	_ = umlsc_

	// get the formGroup
	formGroup := umlscFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlsc_.Name), formDiv)
		case "Activestate":
			FormDivBasicFieldToField(&(umlsc_.Activestate), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(umlsc_.IsInDrawMode), formDiv)
		case "DiagramPackage:Umlscs":
			// we need to retrieve the field owner before the change
			var pastDiagramPackageOwner *models.DiagramPackage
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Umlscs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				umlscFormCallback.probe.stageOfInterest,
				umlscFormCallback.probe.backRepoOfInterest,
				umlsc_,
				&rf)

			if reverseFieldOwner != nil {
				pastDiagramPackageOwner = reverseFieldOwner.(*models.DiagramPackage)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDiagramPackageOwner != nil {
					idx := slices.Index(pastDiagramPackageOwner.Umlscs, umlsc_)
					pastDiagramPackageOwner.Umlscs = slices.Delete(pastDiagramPackageOwner.Umlscs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](umlscFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _diagrampackage.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDiagramPackageOwner := _diagrampackage // we have a match
						if pastDiagramPackageOwner != nil {
							if newDiagramPackageOwner != pastDiagramPackageOwner {
								idx := slices.Index(pastDiagramPackageOwner.Umlscs, umlsc_)
								pastDiagramPackageOwner.Umlscs = slices.Delete(pastDiagramPackageOwner.Umlscs, idx, idx+1)
								newDiagramPackageOwner.Umlscs = append(newDiagramPackageOwner.Umlscs, umlsc_)
							}
						} else {
							newDiagramPackageOwner.Umlscs = append(newDiagramPackageOwner.Umlscs, umlsc_)
						}
					}
				}
			}
		}
	}

	umlscFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Umlsc](
		umlscFormCallback.probe,
	)
	umlscFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if umlscFormCallback.CreationMode {
		umlscFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__UmlscFormCallback(
				nil,
				umlscFormCallback.probe,
			),
		}).Stage(umlscFormCallback.probe.formStage)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, newFormGroup, umlscFormCallback.probe)
		umlscFormCallback.probe.formStage.Commit()
	}

	fillUpTree(umlscFormCallback.probe)
}
func __gong__New__VerticeFormCallback(
	vertice *models.Vertice,
	probe *Probe,
) (verticeFormCallback *VerticeFormCallback) {
	verticeFormCallback = new(VerticeFormCallback)
	verticeFormCallback.probe = probe
	verticeFormCallback.vertice = vertice

	verticeFormCallback.CreationMode = (vertice == nil)

	return
}

type VerticeFormCallback struct {
	vertice *models.Vertice

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (verticeFormCallback *VerticeFormCallback) OnSave() {

	log.Println("VerticeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticeFormCallback.probe.formStage.Checkout()

	if verticeFormCallback.vertice == nil {
		verticeFormCallback.vertice = new(models.Vertice).Stage(verticeFormCallback.probe.stageOfInterest)
	}
	vertice_ := verticeFormCallback.vertice
	_ = vertice_

	// get the formGroup
	formGroup := verticeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "X":
			FormDivBasicFieldToField(&(vertice_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(vertice_.Y), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(vertice_.Name), formDiv)
		}
	}

	verticeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Vertice](
		verticeFormCallback.probe,
	)
	verticeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if verticeFormCallback.CreationMode {
		verticeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__VerticeFormCallback(
				nil,
				verticeFormCallback.probe,
			),
		}).Stage(verticeFormCallback.probe.formStage)
		vertice := new(models.Vertice)
		FillUpForm(vertice, newFormGroup, verticeFormCallback.probe)
		verticeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(verticeFormCallback.probe)
}
