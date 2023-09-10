// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewClassdiagramFormCallback(
	classdiagram *models.Classdiagram,
	playground *Playground,
) (classdiagramFormCallback *ClassdiagramFormCallback) {
	classdiagramFormCallback = new(ClassdiagramFormCallback)
	classdiagramFormCallback.playground = playground
	classdiagramFormCallback.classdiagram = classdiagram
	return
}

type ClassdiagramFormCallback struct {
	classdiagram *models.Classdiagram

	playground *Playground
}

func (classdiagramFormCallback *ClassdiagramFormCallback) OnSave() {

	log.Println("ClassdiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.playground.formStage.Checkout()

	if classdiagramFormCallback.classdiagram == nil {
		classdiagramFormCallback.classdiagram = new(models.Classdiagram).Stage(classdiagramFormCallback.playground.stageOfInterest)
	}
	classdiagram_ := classdiagramFormCallback.classdiagram
	_ = classdiagram_

	// get the formGroup
	formGroup := classdiagramFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(classdiagram_.Name), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInDrawMode), formDiv)
		}
	}

	classdiagramFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Classdiagram](
		classdiagramFormCallback.playground,
	)
	classdiagramFormCallback.playground.tableStage.Commit()
}
func NewDiagramPackageFormCallback(
	diagrampackage *models.DiagramPackage,
	playground *Playground,
) (diagrampackageFormCallback *DiagramPackageFormCallback) {
	diagrampackageFormCallback = new(DiagramPackageFormCallback)
	diagrampackageFormCallback.playground = playground
	diagrampackageFormCallback.diagrampackage = diagrampackage
	return
}

type DiagramPackageFormCallback struct {
	diagrampackage *models.DiagramPackage

	playground *Playground
}

func (diagrampackageFormCallback *DiagramPackageFormCallback) OnSave() {

	log.Println("DiagramPackageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagrampackageFormCallback.playground.formStage.Checkout()

	if diagrampackageFormCallback.diagrampackage == nil {
		diagrampackageFormCallback.diagrampackage = new(models.DiagramPackage).Stage(diagrampackageFormCallback.playground.stageOfInterest)
	}
	diagrampackage_ := diagrampackageFormCallback.diagrampackage
	_ = diagrampackage_

	// get the formGroup
	formGroup := diagrampackageFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(diagrampackage_.SelectedClassdiagram), diagrampackageFormCallback.playground.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(diagrampackage_.IsEditable), formDiv)
		case "IsReloaded":
			FormDivBasicFieldToField(&(diagrampackage_.IsReloaded), formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	diagrampackageFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DiagramPackage](
		diagrampackageFormCallback.playground,
	)
	diagrampackageFormCallback.playground.tableStage.Commit()
}
func NewFieldFormCallback(
	field *models.Field,
	playground *Playground,
) (fieldFormCallback *FieldFormCallback) {
	fieldFormCallback = new(FieldFormCallback)
	fieldFormCallback.playground = playground
	fieldFormCallback.field = field
	return
}

type FieldFormCallback struct {
	field *models.Field

	playground *Playground
}

func (fieldFormCallback *FieldFormCallback) OnSave() {

	log.Println("FieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fieldFormCallback.playground.formStage.Checkout()

	if fieldFormCallback.field == nil {
		fieldFormCallback.field = new(models.Field).Stage(fieldFormCallback.playground.stageOfInterest)
	}
	field_ := fieldFormCallback.field
	_ = field_

	// get the formGroup
	formGroup := fieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	fieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Field](
		fieldFormCallback.playground,
	)
	fieldFormCallback.playground.tableStage.Commit()
}
func NewGongEnumShapeFormCallback(
	gongenumshape *models.GongEnumShape,
	playground *Playground,
) (gongenumshapeFormCallback *GongEnumShapeFormCallback) {
	gongenumshapeFormCallback = new(GongEnumShapeFormCallback)
	gongenumshapeFormCallback.playground = playground
	gongenumshapeFormCallback.gongenumshape = gongenumshape
	return
}

type GongEnumShapeFormCallback struct {
	gongenumshape *models.GongEnumShape

	playground *Playground
}

func (gongenumshapeFormCallback *GongEnumShapeFormCallback) OnSave() {

	log.Println("GongEnumShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumshapeFormCallback.playground.formStage.Checkout()

	if gongenumshapeFormCallback.gongenumshape == nil {
		gongenumshapeFormCallback.gongenumshape = new(models.GongEnumShape).Stage(gongenumshapeFormCallback.playground.stageOfInterest)
	}
	gongenumshape_ := gongenumshapeFormCallback.gongenumshape
	_ = gongenumshape_

	// get the formGroup
	formGroup := gongenumshapeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongenumshape_.Position), gongenumshapeFormCallback.playground.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumshape_.Identifier), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(gongenumshape_.Heigth), formDiv)
		}
	}

	gongenumshapeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongEnumShape](
		gongenumshapeFormCallback.playground,
	)
	gongenumshapeFormCallback.playground.tableStage.Commit()
}
func NewGongEnumValueEntryFormCallback(
	gongenumvalueentry *models.GongEnumValueEntry,
	playground *Playground,
) (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) {
	gongenumvalueentryFormCallback = new(GongEnumValueEntryFormCallback)
	gongenumvalueentryFormCallback.playground = playground
	gongenumvalueentryFormCallback.gongenumvalueentry = gongenumvalueentry
	return
}

type GongEnumValueEntryFormCallback struct {
	gongenumvalueentry *models.GongEnumValueEntry

	playground *Playground
}

func (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) OnSave() {

	log.Println("GongEnumValueEntryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueentryFormCallback.playground.formStage.Checkout()

	if gongenumvalueentryFormCallback.gongenumvalueentry == nil {
		gongenumvalueentryFormCallback.gongenumvalueentry = new(models.GongEnumValueEntry).Stage(gongenumvalueentryFormCallback.playground.stageOfInterest)
	}
	gongenumvalueentry_ := gongenumvalueentryFormCallback.gongenumvalueentry
	_ = gongenumvalueentry_

	// get the formGroup
	formGroup := gongenumvalueentryFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Identifier), formDiv)
		}
	}

	gongenumvalueentryFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValueEntry](
		gongenumvalueentryFormCallback.playground,
	)
	gongenumvalueentryFormCallback.playground.tableStage.Commit()
}
func NewGongStructShapeFormCallback(
	gongstructshape *models.GongStructShape,
	playground *Playground,
) (gongstructshapeFormCallback *GongStructShapeFormCallback) {
	gongstructshapeFormCallback = new(GongStructShapeFormCallback)
	gongstructshapeFormCallback.playground = playground
	gongstructshapeFormCallback.gongstructshape = gongstructshape
	return
}

type GongStructShapeFormCallback struct {
	gongstructshape *models.GongStructShape

	playground *Playground
}

func (gongstructshapeFormCallback *GongStructShapeFormCallback) OnSave() {

	log.Println("GongStructShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructshapeFormCallback.playground.formStage.Checkout()

	if gongstructshapeFormCallback.gongstructshape == nil {
		gongstructshapeFormCallback.gongstructshape = new(models.GongStructShape).Stage(gongstructshapeFormCallback.playground.stageOfInterest)
	}
	gongstructshape_ := gongstructshapeFormCallback.gongstructshape
	_ = gongstructshape_

	// get the formGroup
	formGroup := gongstructshapeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstructshape_.Name), formDiv)
		case "Position":
			FormDivSelectFieldToField(&(gongstructshape_.Position), gongstructshapeFormCallback.playground.stageOfInterest, formDiv)
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

	gongstructshapeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongStructShape](
		gongstructshapeFormCallback.playground,
	)
	gongstructshapeFormCallback.playground.tableStage.Commit()
}
func NewLinkFormCallback(
	link *models.Link,
	playground *Playground,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.playground = playground
	linkFormCallback.link = link
	return
}

type LinkFormCallback struct {
	link *models.Link

	playground *Playground
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.playground.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.playground.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(link_.Middlevertice), linkFormCallback.playground.stageOfInterest, formDiv)
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

	linkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.playground,
	)
	linkFormCallback.playground.tableStage.Commit()
}
func NewNoteShapeFormCallback(
	noteshape *models.NoteShape,
	playground *Playground,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.playground = playground
	noteshapeFormCallback.noteshape = noteshape
	return
}

type NoteShapeFormCallback struct {
	noteshape *models.NoteShape

	playground *Playground
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {

	log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.playground.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.playground.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	// get the formGroup
	formGroup := noteshapeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	noteshapeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.NoteShape](
		noteshapeFormCallback.playground,
	)
	noteshapeFormCallback.playground.tableStage.Commit()
}
func NewNoteShapeLinkFormCallback(
	noteshapelink *models.NoteShapeLink,
	playground *Playground,
) (noteshapelinkFormCallback *NoteShapeLinkFormCallback) {
	noteshapelinkFormCallback = new(NoteShapeLinkFormCallback)
	noteshapelinkFormCallback.playground = playground
	noteshapelinkFormCallback.noteshapelink = noteshapelink
	return
}

type NoteShapeLinkFormCallback struct {
	noteshapelink *models.NoteShapeLink

	playground *Playground
}

func (noteshapelinkFormCallback *NoteShapeLinkFormCallback) OnSave() {

	log.Println("NoteShapeLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapelinkFormCallback.playground.formStage.Checkout()

	if noteshapelinkFormCallback.noteshapelink == nil {
		noteshapelinkFormCallback.noteshapelink = new(models.NoteShapeLink).Stage(noteshapelinkFormCallback.playground.stageOfInterest)
	}
	noteshapelink_ := noteshapelinkFormCallback.noteshapelink
	_ = noteshapelink_

	// get the formGroup
	formGroup := noteshapelinkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	noteshapelinkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.NoteShapeLink](
		noteshapelinkFormCallback.playground,
	)
	noteshapelinkFormCallback.playground.tableStage.Commit()
}
func NewPositionFormCallback(
	position *models.Position,
	playground *Playground,
) (positionFormCallback *PositionFormCallback) {
	positionFormCallback = new(PositionFormCallback)
	positionFormCallback.playground = playground
	positionFormCallback.position = position
	return
}

type PositionFormCallback struct {
	position *models.Position

	playground *Playground
}

func (positionFormCallback *PositionFormCallback) OnSave() {

	log.Println("PositionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	positionFormCallback.playground.formStage.Checkout()

	if positionFormCallback.position == nil {
		positionFormCallback.position = new(models.Position).Stage(positionFormCallback.playground.stageOfInterest)
	}
	position_ := positionFormCallback.position
	_ = position_

	// get the formGroup
	formGroup := positionFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	positionFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Position](
		positionFormCallback.playground,
	)
	positionFormCallback.playground.tableStage.Commit()
}
func NewUmlStateFormCallback(
	umlstate *models.UmlState,
	playground *Playground,
) (umlstateFormCallback *UmlStateFormCallback) {
	umlstateFormCallback = new(UmlStateFormCallback)
	umlstateFormCallback.playground = playground
	umlstateFormCallback.umlstate = umlstate
	return
}

type UmlStateFormCallback struct {
	umlstate *models.UmlState

	playground *Playground
}

func (umlstateFormCallback *UmlStateFormCallback) OnSave() {

	log.Println("UmlStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlstateFormCallback.playground.formStage.Checkout()

	if umlstateFormCallback.umlstate == nil {
		umlstateFormCallback.umlstate = new(models.UmlState).Stage(umlstateFormCallback.playground.stageOfInterest)
	}
	umlstate_ := umlstateFormCallback.umlstate
	_ = umlstate_

	// get the formGroup
	formGroup := umlstateFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	umlstateFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.UmlState](
		umlstateFormCallback.playground,
	)
	umlstateFormCallback.playground.tableStage.Commit()
}
func NewUmlscFormCallback(
	umlsc *models.Umlsc,
	playground *Playground,
) (umlscFormCallback *UmlscFormCallback) {
	umlscFormCallback = new(UmlscFormCallback)
	umlscFormCallback.playground = playground
	umlscFormCallback.umlsc = umlsc
	return
}

type UmlscFormCallback struct {
	umlsc *models.Umlsc

	playground *Playground
}

func (umlscFormCallback *UmlscFormCallback) OnSave() {

	log.Println("UmlscFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlscFormCallback.playground.formStage.Checkout()

	if umlscFormCallback.umlsc == nil {
		umlscFormCallback.umlsc = new(models.Umlsc).Stage(umlscFormCallback.playground.stageOfInterest)
	}
	umlsc_ := umlscFormCallback.umlsc
	_ = umlsc_

	// get the formGroup
	formGroup := umlscFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	umlscFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Umlsc](
		umlscFormCallback.playground,
	)
	umlscFormCallback.playground.tableStage.Commit()
}
func NewVerticeFormCallback(
	vertice *models.Vertice,
	playground *Playground,
) (verticeFormCallback *VerticeFormCallback) {
	verticeFormCallback = new(VerticeFormCallback)
	verticeFormCallback.playground = playground
	verticeFormCallback.vertice = vertice
	return
}

type VerticeFormCallback struct {
	vertice *models.Vertice

	playground *Playground
}

func (verticeFormCallback *VerticeFormCallback) OnSave() {

	log.Println("VerticeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticeFormCallback.playground.formStage.Checkout()

	if verticeFormCallback.vertice == nil {
		verticeFormCallback.vertice = new(models.Vertice).Stage(verticeFormCallback.playground.stageOfInterest)
	}
	vertice_ := verticeFormCallback.vertice
	_ = vertice_

	// get the formGroup
	formGroup := verticeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	verticeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Vertice](
		verticeFormCallback.playground,
	)
	verticeFormCallback.playground.tableStage.Commit()
}
