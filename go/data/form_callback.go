// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewClassdiagramFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	classdiagram *models.Classdiagram,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (classdiagramFormCallback *ClassdiagramFormCallback) {
	classdiagramFormCallback = new(ClassdiagramFormCallback)
	classdiagramFormCallback.stageOfInterest = stageOfInterest
	classdiagramFormCallback.tableStage = tableStage
	classdiagramFormCallback.formStage = formStage
	classdiagramFormCallback.classdiagram = classdiagram
	classdiagramFormCallback.r = r
	classdiagramFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type ClassdiagramFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	classdiagram            *models.Classdiagram
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (classdiagramFormCallback *ClassdiagramFormCallback) OnSave() {

	log.Println("ClassdiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.formStage.Checkout()

	if classdiagramFormCallback.classdiagram == nil {
		classdiagramFormCallback.classdiagram = new(models.Classdiagram).Stage(classdiagramFormCallback.stageOfInterest)
	}
	classdiagram_ := classdiagramFormCallback.classdiagram
	_ = classdiagram_

	// get the formGroup
	formGroup := classdiagramFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(classdiagram_.Name), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInDrawMode), formDiv)
		}
	}

	classdiagramFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Classdiagram](
		classdiagramFormCallback.stageOfInterest,
		classdiagramFormCallback.tableStage,
		classdiagramFormCallback.formStage,
		classdiagramFormCallback.r,
		classdiagramFormCallback.backRepoOfInterest,
	)
	classdiagramFormCallback.tableStage.Commit()
}
func NewDiagramPackageFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	diagrampackage *models.DiagramPackage,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (diagrampackageFormCallback *DiagramPackageFormCallback) {
	diagrampackageFormCallback = new(DiagramPackageFormCallback)
	diagrampackageFormCallback.stageOfInterest = stageOfInterest
	diagrampackageFormCallback.tableStage = tableStage
	diagrampackageFormCallback.formStage = formStage
	diagrampackageFormCallback.diagrampackage = diagrampackage
	diagrampackageFormCallback.r = r
	diagrampackageFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type DiagramPackageFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	diagrampackage            *models.DiagramPackage
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (diagrampackageFormCallback *DiagramPackageFormCallback) OnSave() {

	log.Println("DiagramPackageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagrampackageFormCallback.formStage.Checkout()

	if diagrampackageFormCallback.diagrampackage == nil {
		diagrampackageFormCallback.diagrampackage = new(models.DiagramPackage).Stage(diagrampackageFormCallback.stageOfInterest)
	}
	diagrampackage_ := diagrampackageFormCallback.diagrampackage
	_ = diagrampackage_

	// get the formGroup
	formGroup := diagrampackageFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(diagrampackage_.SelectedClassdiagram), diagrampackageFormCallback.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(diagrampackage_.IsEditable), formDiv)
		case "IsReloaded":
			FormDivBasicFieldToField(&(diagrampackage_.IsReloaded), formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	diagrampackageFormCallback.stageOfInterest.Commit()
	fillUpTable[models.DiagramPackage](
		diagrampackageFormCallback.stageOfInterest,
		diagrampackageFormCallback.tableStage,
		diagrampackageFormCallback.formStage,
		diagrampackageFormCallback.r,
		diagrampackageFormCallback.backRepoOfInterest,
	)
	diagrampackageFormCallback.tableStage.Commit()
}
func NewFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	field *models.Field,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (fieldFormCallback *FieldFormCallback) {
	fieldFormCallback = new(FieldFormCallback)
	fieldFormCallback.stageOfInterest = stageOfInterest
	fieldFormCallback.tableStage = tableStage
	fieldFormCallback.formStage = formStage
	fieldFormCallback.field = field
	fieldFormCallback.r = r
	fieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type FieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	field            *models.Field
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (fieldFormCallback *FieldFormCallback) OnSave() {

	log.Println("FieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fieldFormCallback.formStage.Checkout()

	if fieldFormCallback.field == nil {
		fieldFormCallback.field = new(models.Field).Stage(fieldFormCallback.stageOfInterest)
	}
	field_ := fieldFormCallback.field
	_ = field_

	// get the formGroup
	formGroup := fieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
		}
	}

	fieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Field](
		fieldFormCallback.stageOfInterest,
		fieldFormCallback.tableStage,
		fieldFormCallback.formStage,
		fieldFormCallback.r,
		fieldFormCallback.backRepoOfInterest,
	)
	fieldFormCallback.tableStage.Commit()
}
func NewGongEnumShapeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongenumshape *models.GongEnumShape,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongenumshapeFormCallback *GongEnumShapeFormCallback) {
	gongenumshapeFormCallback = new(GongEnumShapeFormCallback)
	gongenumshapeFormCallback.stageOfInterest = stageOfInterest
	gongenumshapeFormCallback.tableStage = tableStage
	gongenumshapeFormCallback.formStage = formStage
	gongenumshapeFormCallback.gongenumshape = gongenumshape
	gongenumshapeFormCallback.r = r
	gongenumshapeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongEnumShapeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongenumshape            *models.GongEnumShape
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongenumshapeFormCallback *GongEnumShapeFormCallback) OnSave() {

	log.Println("GongEnumShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumshapeFormCallback.formStage.Checkout()

	if gongenumshapeFormCallback.gongenumshape == nil {
		gongenumshapeFormCallback.gongenumshape = new(models.GongEnumShape).Stage(gongenumshapeFormCallback.stageOfInterest)
	}
	gongenumshape_ := gongenumshapeFormCallback.gongenumshape
	_ = gongenumshape_

	// get the formGroup
	formGroup := gongenumshapeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongenumshape_.Position), gongenumshapeFormCallback.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumshape_.Identifier), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(gongenumshape_.Heigth), formDiv)
		}
	}

	gongenumshapeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongEnumShape](
		gongenumshapeFormCallback.stageOfInterest,
		gongenumshapeFormCallback.tableStage,
		gongenumshapeFormCallback.formStage,
		gongenumshapeFormCallback.r,
		gongenumshapeFormCallback.backRepoOfInterest,
	)
	gongenumshapeFormCallback.tableStage.Commit()
}
func NewGongEnumValueEntryFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongenumvalueentry *models.GongEnumValueEntry,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) {
	gongenumvalueentryFormCallback = new(GongEnumValueEntryFormCallback)
	gongenumvalueentryFormCallback.stageOfInterest = stageOfInterest
	gongenumvalueentryFormCallback.tableStage = tableStage
	gongenumvalueentryFormCallback.formStage = formStage
	gongenumvalueentryFormCallback.gongenumvalueentry = gongenumvalueentry
	gongenumvalueentryFormCallback.r = r
	gongenumvalueentryFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongEnumValueEntryFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongenumvalueentry            *models.GongEnumValueEntry
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) OnSave() {

	log.Println("GongEnumValueEntryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueentryFormCallback.formStage.Checkout()

	if gongenumvalueentryFormCallback.gongenumvalueentry == nil {
		gongenumvalueentryFormCallback.gongenumvalueentry = new(models.GongEnumValueEntry).Stage(gongenumvalueentryFormCallback.stageOfInterest)
	}
	gongenumvalueentry_ := gongenumvalueentryFormCallback.gongenumvalueentry
	_ = gongenumvalueentry_

	// get the formGroup
	formGroup := gongenumvalueentryFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Identifier), formDiv)
		}
	}

	gongenumvalueentryFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValueEntry](
		gongenumvalueentryFormCallback.stageOfInterest,
		gongenumvalueentryFormCallback.tableStage,
		gongenumvalueentryFormCallback.formStage,
		gongenumvalueentryFormCallback.r,
		gongenumvalueentryFormCallback.backRepoOfInterest,
	)
	gongenumvalueentryFormCallback.tableStage.Commit()
}
func NewGongStructShapeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongstructshape *models.GongStructShape,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongstructshapeFormCallback *GongStructShapeFormCallback) {
	gongstructshapeFormCallback = new(GongStructShapeFormCallback)
	gongstructshapeFormCallback.stageOfInterest = stageOfInterest
	gongstructshapeFormCallback.tableStage = tableStage
	gongstructshapeFormCallback.formStage = formStage
	gongstructshapeFormCallback.gongstructshape = gongstructshape
	gongstructshapeFormCallback.r = r
	gongstructshapeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongStructShapeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongstructshape            *models.GongStructShape
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongstructshapeFormCallback *GongStructShapeFormCallback) OnSave() {

	log.Println("GongStructShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructshapeFormCallback.formStage.Checkout()

	if gongstructshapeFormCallback.gongstructshape == nil {
		gongstructshapeFormCallback.gongstructshape = new(models.GongStructShape).Stage(gongstructshapeFormCallback.stageOfInterest)
	}
	gongstructshape_ := gongstructshapeFormCallback.gongstructshape
	_ = gongstructshape_

	// get the formGroup
	formGroup := gongstructshapeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstructshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongstructshape_.Position), gongstructshapeFormCallback.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongstructshape_.Identifier), formDiv)
		case "ShowNbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.ShowNbInstances), formDiv)
		case "NbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.NbInstances), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongstructshape_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(gongstructshape_.Heigth), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(gongstructshape_.IsSelected), formDiv)
		}
	}

	gongstructshapeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongStructShape](
		gongstructshapeFormCallback.stageOfInterest,
		gongstructshapeFormCallback.tableStage,
		gongstructshapeFormCallback.formStage,
		gongstructshapeFormCallback.r,
		gongstructshapeFormCallback.backRepoOfInterest,
	)
	gongstructshapeFormCallback.tableStage.Commit()
}
func NewLinkFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	link *models.Link,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.stageOfInterest = stageOfInterest
	linkFormCallback.tableStage = tableStage
	linkFormCallback.formStage = formStage
	linkFormCallback.link = link
	linkFormCallback.r = r
	linkFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type LinkFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	link            *models.Link
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(link_.Middlevertice), linkFormCallback.stageOfInterest, formDiv)
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
		}
	}

	linkFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.stageOfInterest,
		linkFormCallback.tableStage,
		linkFormCallback.formStage,
		linkFormCallback.r,
		linkFormCallback.backRepoOfInterest,
	)
	linkFormCallback.tableStage.Commit()
}
func NewNoteShapeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	noteshape *models.NoteShape,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.stageOfInterest = stageOfInterest
	noteshapeFormCallback.tableStage = tableStage
	noteshapeFormCallback.formStage = formStage
	noteshapeFormCallback.noteshape = noteshape
	noteshapeFormCallback.r = r
	noteshapeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type NoteShapeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	noteshape            *models.NoteShape
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {

	log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	// get the formGroup
	formGroup := noteshapeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
		}
	}

	noteshapeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.NoteShape](
		noteshapeFormCallback.stageOfInterest,
		noteshapeFormCallback.tableStage,
		noteshapeFormCallback.formStage,
		noteshapeFormCallback.r,
		noteshapeFormCallback.backRepoOfInterest,
	)
	noteshapeFormCallback.tableStage.Commit()
}
func NewNoteShapeLinkFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	noteshapelink *models.NoteShapeLink,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (noteshapelinkFormCallback *NoteShapeLinkFormCallback) {
	noteshapelinkFormCallback = new(NoteShapeLinkFormCallback)
	noteshapelinkFormCallback.stageOfInterest = stageOfInterest
	noteshapelinkFormCallback.tableStage = tableStage
	noteshapelinkFormCallback.formStage = formStage
	noteshapelinkFormCallback.noteshapelink = noteshapelink
	noteshapelinkFormCallback.r = r
	noteshapelinkFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type NoteShapeLinkFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	noteshapelink            *models.NoteShapeLink
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (noteshapelinkFormCallback *NoteShapeLinkFormCallback) OnSave() {

	log.Println("NoteShapeLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapelinkFormCallback.formStage.Checkout()

	if noteshapelinkFormCallback.noteshapelink == nil {
		noteshapelinkFormCallback.noteshapelink = new(models.NoteShapeLink).Stage(noteshapelinkFormCallback.stageOfInterest)
	}
	noteshapelink_ := noteshapelinkFormCallback.noteshapelink
	_ = noteshapelink_

	// get the formGroup
	formGroup := noteshapelinkFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshapelink_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(noteshapelink_.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(noteshapelink_.Type), formDiv)
		}
	}

	noteshapelinkFormCallback.stageOfInterest.Commit()
	fillUpTable[models.NoteShapeLink](
		noteshapelinkFormCallback.stageOfInterest,
		noteshapelinkFormCallback.tableStage,
		noteshapelinkFormCallback.formStage,
		noteshapelinkFormCallback.r,
		noteshapelinkFormCallback.backRepoOfInterest,
	)
	noteshapelinkFormCallback.tableStage.Commit()
}
func NewPositionFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	position *models.Position,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (positionFormCallback *PositionFormCallback) {
	positionFormCallback = new(PositionFormCallback)
	positionFormCallback.stageOfInterest = stageOfInterest
	positionFormCallback.tableStage = tableStage
	positionFormCallback.formStage = formStage
	positionFormCallback.position = position
	positionFormCallback.r = r
	positionFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PositionFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	position            *models.Position
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (positionFormCallback *PositionFormCallback) OnSave() {

	log.Println("PositionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	positionFormCallback.formStage.Checkout()

	if positionFormCallback.position == nil {
		positionFormCallback.position = new(models.Position).Stage(positionFormCallback.stageOfInterest)
	}
	position_ := positionFormCallback.position
	_ = position_

	// get the formGroup
	formGroup := positionFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	positionFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Position](
		positionFormCallback.stageOfInterest,
		positionFormCallback.tableStage,
		positionFormCallback.formStage,
		positionFormCallback.r,
		positionFormCallback.backRepoOfInterest,
	)
	positionFormCallback.tableStage.Commit()
}
func NewUmlStateFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	umlstate *models.UmlState,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (umlstateFormCallback *UmlStateFormCallback) {
	umlstateFormCallback = new(UmlStateFormCallback)
	umlstateFormCallback.stageOfInterest = stageOfInterest
	umlstateFormCallback.tableStage = tableStage
	umlstateFormCallback.formStage = formStage
	umlstateFormCallback.umlstate = umlstate
	umlstateFormCallback.r = r
	umlstateFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type UmlStateFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	umlstate            *models.UmlState
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (umlstateFormCallback *UmlStateFormCallback) OnSave() {

	log.Println("UmlStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlstateFormCallback.formStage.Checkout()

	if umlstateFormCallback.umlstate == nil {
		umlstateFormCallback.umlstate = new(models.UmlState).Stage(umlstateFormCallback.stageOfInterest)
	}
	umlstate_ := umlstateFormCallback.umlstate
	_ = umlstate_

	// get the formGroup
	formGroup := umlstateFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlstate_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(umlstate_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(umlstate_.Y), formDiv)
		}
	}

	umlstateFormCallback.stageOfInterest.Commit()
	fillUpTable[models.UmlState](
		umlstateFormCallback.stageOfInterest,
		umlstateFormCallback.tableStage,
		umlstateFormCallback.formStage,
		umlstateFormCallback.r,
		umlstateFormCallback.backRepoOfInterest,
	)
	umlstateFormCallback.tableStage.Commit()
}
func NewUmlscFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	umlsc *models.Umlsc,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (umlscFormCallback *UmlscFormCallback) {
	umlscFormCallback = new(UmlscFormCallback)
	umlscFormCallback.stageOfInterest = stageOfInterest
	umlscFormCallback.tableStage = tableStage
	umlscFormCallback.formStage = formStage
	umlscFormCallback.umlsc = umlsc
	umlscFormCallback.r = r
	umlscFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type UmlscFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	umlsc            *models.Umlsc
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (umlscFormCallback *UmlscFormCallback) OnSave() {

	log.Println("UmlscFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlscFormCallback.formStage.Checkout()

	if umlscFormCallback.umlsc == nil {
		umlscFormCallback.umlsc = new(models.Umlsc).Stage(umlscFormCallback.stageOfInterest)
	}
	umlsc_ := umlscFormCallback.umlsc
	_ = umlsc_

	// get the formGroup
	formGroup := umlscFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlsc_.Name), formDiv)
		case "Activestate":
			FormDivBasicFieldToField(&(umlsc_.Activestate), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(umlsc_.IsInDrawMode), formDiv)
		}
	}

	umlscFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Umlsc](
		umlscFormCallback.stageOfInterest,
		umlscFormCallback.tableStage,
		umlscFormCallback.formStage,
		umlscFormCallback.r,
		umlscFormCallback.backRepoOfInterest,
	)
	umlscFormCallback.tableStage.Commit()
}
func NewVerticeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	vertice *models.Vertice,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (verticeFormCallback *VerticeFormCallback) {
	verticeFormCallback = new(VerticeFormCallback)
	verticeFormCallback.stageOfInterest = stageOfInterest
	verticeFormCallback.tableStage = tableStage
	verticeFormCallback.formStage = formStage
	verticeFormCallback.vertice = vertice
	verticeFormCallback.r = r
	verticeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type VerticeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	vertice            *models.Vertice
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (verticeFormCallback *VerticeFormCallback) OnSave() {

	log.Println("VerticeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticeFormCallback.formStage.Checkout()

	if verticeFormCallback.vertice == nil {
		verticeFormCallback.vertice = new(models.Vertice).Stage(verticeFormCallback.stageOfInterest)
	}
	vertice_ := verticeFormCallback.vertice
	_ = vertice_

	// get the formGroup
	formGroup := verticeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	verticeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Vertice](
		verticeFormCallback.stageOfInterest,
		verticeFormCallback.tableStage,
		verticeFormCallback.formStage,
		verticeFormCallback.r,
		verticeFormCallback.backRepoOfInterest,
	)
	verticeFormCallback.tableStage.Commit()
}