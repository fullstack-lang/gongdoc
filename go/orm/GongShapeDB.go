// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongShape_sql sql.NullBool
var dummy_GongShape_time time.Duration
var dummy_GongShape_sort sort.Float64Slice

// GongShapeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongshapeAPI
type GongShapeAPI struct {
	gorm.Model

	models.GongShape

	// encoding of pointers
	GongShapePointersEnconding
}

// GongShapePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongShapePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Position is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	PositionID sql.NullInt64

	// Implementation of a reverse ID for field Classdiagram{}.GongStructShapes []*GongShape
	Classdiagram_GongStructShapesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Classdiagram_GongStructShapesDBID_Index sql.NullInt64
}

// GongShapeDB describes a gongshape in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongshapeDB
type GongShapeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongshapeDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongshapeDB.Identifier
	Identifier_Data sql.NullString

	// Declation for basic field gongshapeDB.ShowNbInstances
	// provide the sql storage for the boolan
	ShowNbInstances_Data sql.NullBool

	// Declation for basic field gongshapeDB.NbInstances
	NbInstances_Data sql.NullInt64

	// Declation for basic field gongshapeDB.Width
	Width_Data sql.NullFloat64

	// Declation for basic field gongshapeDB.Heigth
	Heigth_Data sql.NullFloat64

	// Declation for basic field gongshapeDB.IsSelected
	// provide the sql storage for the boolan
	IsSelected_Data sql.NullBool
	// encoding of pointers
	GongShapePointersEnconding
}

// GongShapeDBs arrays gongshapeDBs
// swagger:response gongshapeDBsResponse
type GongShapeDBs []GongShapeDB

// GongShapeDBResponse provides response
// swagger:response gongshapeDBResponse
type GongShapeDBResponse struct {
	GongShapeDB
}

// GongShapeWOP is a GongShape without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongShapeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Identifier string `xlsx:"2"`

	ShowNbInstances bool `xlsx:"3"`

	NbInstances int `xlsx:"4"`

	Width float64 `xlsx:"5"`

	Heigth float64 `xlsx:"6"`

	IsSelected bool `xlsx:"7"`
	// insertion for WOP pointer fields
}

var GongShape_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Identifier",
	"ShowNbInstances",
	"NbInstances",
	"Width",
	"Heigth",
	"IsSelected",
}

type BackRepoGongShapeStruct struct {
	// stores GongShapeDB according to their gorm ID
	Map_GongShapeDBID_GongShapeDB map[uint]*GongShapeDB

	// stores GongShapeDB ID according to GongShape address
	Map_GongShapePtr_GongShapeDBID map[*models.GongShape]uint

	// stores GongShape according to their gorm ID
	Map_GongShapeDBID_GongShapePtr map[uint]*models.GongShape

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoGongShape *BackRepoGongShapeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoGongShape.stage
	return
}

func (backRepoGongShape *BackRepoGongShapeStruct) GetDB() *gorm.DB {
	return backRepoGongShape.db
}

// GetGongShapeDBFromGongShapePtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongShape *BackRepoGongShapeStruct) GetGongShapeDBFromGongShapePtr(gongshape *models.GongShape) (gongshapeDB *GongShapeDB) {
	id := backRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]
	gongshapeDB = backRepoGongShape.Map_GongShapeDBID_GongShapeDB[id]
	return
}

// BackRepoGongShape.CommitPhaseOne commits all staged instances of GongShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongShape *BackRepoGongShapeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongshape := range stage.GongShapes {
		backRepoGongShape.CommitPhaseOneInstance(gongshape)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongshape := range backRepoGongShape.Map_GongShapeDBID_GongShapePtr {
		if _, ok := stage.GongShapes[gongshape]; !ok {
			backRepoGongShape.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongShape.CommitDeleteInstance commits deletion of GongShape to the BackRepo
func (backRepoGongShape *BackRepoGongShapeStruct) CommitDeleteInstance(id uint) (Error error) {

	gongshape := backRepoGongShape.Map_GongShapeDBID_GongShapePtr[id]

	// gongshape is not staged anymore, remove gongshapeDB
	gongshapeDB := backRepoGongShape.Map_GongShapeDBID_GongShapeDB[id]
	query := backRepoGongShape.db.Unscoped().Delete(&gongshapeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoGongShape.Map_GongShapePtr_GongShapeDBID, gongshape)
	delete(backRepoGongShape.Map_GongShapeDBID_GongShapePtr, id)
	delete(backRepoGongShape.Map_GongShapeDBID_GongShapeDB, id)

	return
}

// BackRepoGongShape.CommitPhaseOneInstance commits gongshape staged instances of GongShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongShape *BackRepoGongShapeStruct) CommitPhaseOneInstance(gongshape *models.GongShape) (Error error) {

	// check if the gongshape is not commited yet
	if _, ok := backRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]; ok {
		return
	}

	// initiate gongshape
	var gongshapeDB GongShapeDB
	gongshapeDB.CopyBasicFieldsFromGongShape(gongshape)

	query := backRepoGongShape.db.Create(&gongshapeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape] = gongshapeDB.ID
	backRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID] = gongshape
	backRepoGongShape.Map_GongShapeDBID_GongShapeDB[gongshapeDB.ID] = &gongshapeDB

	return
}

// BackRepoGongShape.CommitPhaseTwo commits all staged instances of GongShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongShape *BackRepoGongShapeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongshape := range backRepoGongShape.Map_GongShapeDBID_GongShapePtr {
		backRepoGongShape.CommitPhaseTwoInstance(backRepo, idx, gongshape)
	}

	return
}

// BackRepoGongShape.CommitPhaseTwoInstance commits {{structname }} of models.GongShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongShape *BackRepoGongShapeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongshape *models.GongShape) (Error error) {

	// fetch matching gongshapeDB
	if gongshapeDB, ok := backRepoGongShape.Map_GongShapeDBID_GongShapeDB[idx]; ok {

		gongshapeDB.CopyBasicFieldsFromGongShape(gongshape)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value gongshape.Position translates to updating the gongshape.PositionID
		gongshapeDB.PositionID.Valid = true // allow for a 0 value (nil association)
		if gongshape.Position != nil {
			if PositionId, ok := backRepo.BackRepoPosition.Map_PositionPtr_PositionDBID[gongshape.Position]; ok {
				gongshapeDB.PositionID.Int64 = int64(PositionId)
				gongshapeDB.PositionID.Valid = true
			}
		}

		// This loop encodes the slice of pointers gongshape.Fields into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, fieldAssocEnd := range gongshape.Fields {

			// get the back repo instance at the association end
			fieldAssocEnd_DB :=
				backRepo.BackRepoField.GetFieldDBFromFieldPtr(fieldAssocEnd)

			// encode reverse pointer in the association end back repo instance
			fieldAssocEnd_DB.GongShape_FieldsDBID.Int64 = int64(gongshapeDB.ID)
			fieldAssocEnd_DB.GongShape_FieldsDBID.Valid = true
			fieldAssocEnd_DB.GongShape_FieldsDBID_Index.Int64 = int64(idx)
			fieldAssocEnd_DB.GongShape_FieldsDBID_Index.Valid = true
			if q := backRepoGongShape.db.Save(fieldAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		// This loop encodes the slice of pointers gongshape.Links into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, linkAssocEnd := range gongshape.Links {

			// get the back repo instance at the association end
			linkAssocEnd_DB :=
				backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkAssocEnd)

			// encode reverse pointer in the association end back repo instance
			linkAssocEnd_DB.GongShape_LinksDBID.Int64 = int64(gongshapeDB.ID)
			linkAssocEnd_DB.GongShape_LinksDBID.Valid = true
			linkAssocEnd_DB.GongShape_LinksDBID_Index.Int64 = int64(idx)
			linkAssocEnd_DB.GongShape_LinksDBID_Index.Valid = true
			if q := backRepoGongShape.db.Save(linkAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoGongShape.db.Save(&gongshapeDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongShape intance %s", gongshape.Name))
		return err
	}

	return
}

// BackRepoGongShape.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongShape *BackRepoGongShapeStruct) CheckoutPhaseOne() (Error error) {

	gongshapeDBArray := make([]GongShapeDB, 0)
	query := backRepoGongShape.db.Find(&gongshapeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongshapeInstancesToBeRemovedFromTheStage := make(map[*models.GongShape]any)
	for key, value := range backRepoGongShape.stage.GongShapes {
		gongshapeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongshapeDB := range gongshapeDBArray {
		backRepoGongShape.CheckoutPhaseOneInstance(&gongshapeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongshape, ok := backRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]
		if ok {
			delete(gongshapeInstancesToBeRemovedFromTheStage, gongshape)
		}
	}

	// remove from stage and back repo's 3 maps all gongshapes that are not in the checkout
	for gongshape := range gongshapeInstancesToBeRemovedFromTheStage {
		gongshape.Unstage(backRepoGongShape.GetStage())

		// remove instance from the back repo 3 maps
		gongshapeID := backRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]
		delete(backRepoGongShape.Map_GongShapePtr_GongShapeDBID, gongshape)
		delete(backRepoGongShape.Map_GongShapeDBID_GongShapeDB, gongshapeID)
		delete(backRepoGongShape.Map_GongShapeDBID_GongShapePtr, gongshapeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongshapeDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongshapeDB
func (backRepoGongShape *BackRepoGongShapeStruct) CheckoutPhaseOneInstance(gongshapeDB *GongShapeDB) (Error error) {

	gongshape, ok := backRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]
	if !ok {
		gongshape = new(models.GongShape)

		backRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID] = gongshape
		backRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape] = gongshapeDB.ID

		// append model store with the new element
		gongshape.Name = gongshapeDB.Name_Data.String
		gongshape.Stage(backRepoGongShape.GetStage())
	}
	gongshapeDB.CopyBasicFieldsToGongShape(gongshape)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongshape.Stage(backRepoGongShape.GetStage())

	// preserve pointer to gongshapeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongShapeDBID_GongShapeDB)[gongshapeDB hold variable pointers
	gongshapeDB_Data := *gongshapeDB
	preservedPtrToGongShape := &gongshapeDB_Data
	backRepoGongShape.Map_GongShapeDBID_GongShapeDB[gongshapeDB.ID] = preservedPtrToGongShape

	return
}

// BackRepoGongShape.CheckoutPhaseTwo Checkouts all staged instances of GongShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongShape *BackRepoGongShapeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongshapeDB := range backRepoGongShape.Map_GongShapeDBID_GongShapeDB {
		backRepoGongShape.CheckoutPhaseTwoInstance(backRepo, gongshapeDB)
	}
	return
}

// BackRepoGongShape.CheckoutPhaseTwoInstance Checkouts staged instances of GongShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongShape *BackRepoGongShapeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongshapeDB *GongShapeDB) (Error error) {

	gongshape := backRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]
	_ = gongshape // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Position field
	if gongshapeDB.PositionID.Int64 != 0 {
		gongshape.Position = backRepo.BackRepoPosition.Map_PositionDBID_PositionPtr[uint(gongshapeDB.PositionID.Int64)]
	}
	// This loop redeem gongshape.Fields in the stage from the encode in the back repo
	// It parses all FieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongshape.Fields = gongshape.Fields[:0]
	// 2. loop all instances in the type in the association end
	for _, fieldDB_AssocEnd := range backRepo.BackRepoField.Map_FieldDBID_FieldDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if fieldDB_AssocEnd.GongShape_FieldsDBID.Int64 == int64(gongshapeDB.ID) {
			// 4. fetch the associated instance in the stage
			field_AssocEnd := backRepo.BackRepoField.Map_FieldDBID_FieldPtr[fieldDB_AssocEnd.ID]
			// 5. append it the association slice
			gongshape.Fields = append(gongshape.Fields, field_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(gongshape.Fields, func(i, j int) bool {
		fieldDB_i_ID := backRepo.BackRepoField.Map_FieldPtr_FieldDBID[gongshape.Fields[i]]
		fieldDB_j_ID := backRepo.BackRepoField.Map_FieldPtr_FieldDBID[gongshape.Fields[j]]

		fieldDB_i := backRepo.BackRepoField.Map_FieldDBID_FieldDB[fieldDB_i_ID]
		fieldDB_j := backRepo.BackRepoField.Map_FieldDBID_FieldDB[fieldDB_j_ID]

		return fieldDB_i.GongShape_FieldsDBID_Index.Int64 < fieldDB_j.GongShape_FieldsDBID_Index.Int64
	})

	// This loop redeem gongshape.Links in the stage from the encode in the back repo
	// It parses all LinkDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongshape.Links = gongshape.Links[:0]
	// 2. loop all instances in the type in the association end
	for _, linkDB_AssocEnd := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if linkDB_AssocEnd.GongShape_LinksDBID.Int64 == int64(gongshapeDB.ID) {
			// 4. fetch the associated instance in the stage
			link_AssocEnd := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[linkDB_AssocEnd.ID]
			// 5. append it the association slice
			gongshape.Links = append(gongshape.Links, link_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(gongshape.Links, func(i, j int) bool {
		linkDB_i_ID := backRepo.BackRepoLink.Map_LinkPtr_LinkDBID[gongshape.Links[i]]
		linkDB_j_ID := backRepo.BackRepoLink.Map_LinkPtr_LinkDBID[gongshape.Links[j]]

		linkDB_i := backRepo.BackRepoLink.Map_LinkDBID_LinkDB[linkDB_i_ID]
		linkDB_j := backRepo.BackRepoLink.Map_LinkDBID_LinkDB[linkDB_j_ID]

		return linkDB_i.GongShape_LinksDBID_Index.Int64 < linkDB_j.GongShape_LinksDBID_Index.Int64
	})

	return
}

// CommitGongShape allows commit of a single gongshape (if already staged)
func (backRepo *BackRepoStruct) CommitGongShape(gongshape *models.GongShape) {
	backRepo.BackRepoGongShape.CommitPhaseOneInstance(gongshape)
	if id, ok := backRepo.BackRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]; ok {
		backRepo.BackRepoGongShape.CommitPhaseTwoInstance(backRepo, id, gongshape)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongShape allows checkout of a single gongshape (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongShape(gongshape *models.GongShape) {
	// check if the gongshape is staged
	if _, ok := backRepo.BackRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]; ok {

		if id, ok := backRepo.BackRepoGongShape.Map_GongShapePtr_GongShapeDBID[gongshape]; ok {
			var gongshapeDB GongShapeDB
			gongshapeDB.ID = id

			if err := backRepo.BackRepoGongShape.db.First(&gongshapeDB, id).Error; err != nil {
				log.Panicln("CheckoutGongShape : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongShape.CheckoutPhaseOneInstance(&gongshapeDB)
			backRepo.BackRepoGongShape.CheckoutPhaseTwoInstance(backRepo, &gongshapeDB)
		}
	}
}

// CopyBasicFieldsFromGongShape
func (gongshapeDB *GongShapeDB) CopyBasicFieldsFromGongShape(gongshape *models.GongShape) {
	// insertion point for fields commit

	gongshapeDB.Name_Data.String = gongshape.Name
	gongshapeDB.Name_Data.Valid = true

	gongshapeDB.Identifier_Data.String = gongshape.Identifier
	gongshapeDB.Identifier_Data.Valid = true

	gongshapeDB.ShowNbInstances_Data.Bool = gongshape.ShowNbInstances
	gongshapeDB.ShowNbInstances_Data.Valid = true

	gongshapeDB.NbInstances_Data.Int64 = int64(gongshape.NbInstances)
	gongshapeDB.NbInstances_Data.Valid = true

	gongshapeDB.Width_Data.Float64 = gongshape.Width
	gongshapeDB.Width_Data.Valid = true

	gongshapeDB.Heigth_Data.Float64 = gongshape.Heigth
	gongshapeDB.Heigth_Data.Valid = true

	gongshapeDB.IsSelected_Data.Bool = gongshape.IsSelected
	gongshapeDB.IsSelected_Data.Valid = true
}

// CopyBasicFieldsFromGongShapeWOP
func (gongshapeDB *GongShapeDB) CopyBasicFieldsFromGongShapeWOP(gongshape *GongShapeWOP) {
	// insertion point for fields commit

	gongshapeDB.Name_Data.String = gongshape.Name
	gongshapeDB.Name_Data.Valid = true

	gongshapeDB.Identifier_Data.String = gongshape.Identifier
	gongshapeDB.Identifier_Data.Valid = true

	gongshapeDB.ShowNbInstances_Data.Bool = gongshape.ShowNbInstances
	gongshapeDB.ShowNbInstances_Data.Valid = true

	gongshapeDB.NbInstances_Data.Int64 = int64(gongshape.NbInstances)
	gongshapeDB.NbInstances_Data.Valid = true

	gongshapeDB.Width_Data.Float64 = gongshape.Width
	gongshapeDB.Width_Data.Valid = true

	gongshapeDB.Heigth_Data.Float64 = gongshape.Heigth
	gongshapeDB.Heigth_Data.Valid = true

	gongshapeDB.IsSelected_Data.Bool = gongshape.IsSelected
	gongshapeDB.IsSelected_Data.Valid = true
}

// CopyBasicFieldsToGongShape
func (gongshapeDB *GongShapeDB) CopyBasicFieldsToGongShape(gongshape *models.GongShape) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongshape.Name = gongshapeDB.Name_Data.String
	gongshape.Identifier = gongshapeDB.Identifier_Data.String
	gongshape.ShowNbInstances = gongshapeDB.ShowNbInstances_Data.Bool
	gongshape.NbInstances = int(gongshapeDB.NbInstances_Data.Int64)
	gongshape.Width = gongshapeDB.Width_Data.Float64
	gongshape.Heigth = gongshapeDB.Heigth_Data.Float64
	gongshape.IsSelected = gongshapeDB.IsSelected_Data.Bool
}

// CopyBasicFieldsToGongShapeWOP
func (gongshapeDB *GongShapeDB) CopyBasicFieldsToGongShapeWOP(gongshape *GongShapeWOP) {
	gongshape.ID = int(gongshapeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongshape.Name = gongshapeDB.Name_Data.String
	gongshape.Identifier = gongshapeDB.Identifier_Data.String
	gongshape.ShowNbInstances = gongshapeDB.ShowNbInstances_Data.Bool
	gongshape.NbInstances = int(gongshapeDB.NbInstances_Data.Int64)
	gongshape.Width = gongshapeDB.Width_Data.Float64
	gongshape.Heigth = gongshapeDB.Heigth_Data.Float64
	gongshape.IsSelected = gongshapeDB.IsSelected_Data.Bool
}

// Backup generates a json file from a slice of all GongShapeDB instances in the backrepo
func (backRepoGongShape *BackRepoGongShapeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongShapeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongShapeDB, 0)
	for _, gongshapeDB := range backRepoGongShape.Map_GongShapeDBID_GongShapeDB {
		forBackup = append(forBackup, gongshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongShape ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongShape file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongShapeDB instances in the backrepo
func (backRepoGongShape *BackRepoGongShapeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongShapeDB, 0)
	for _, gongshapeDB := range backRepoGongShape.Map_GongShapeDBID_GongShapeDB {
		forBackup = append(forBackup, gongshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongShape")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongShape_Fields, -1)
	for _, gongshapeDB := range forBackup {

		var gongshapeWOP GongShapeWOP
		gongshapeDB.CopyBasicFieldsToGongShapeWOP(&gongshapeWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongshapeWOP, -1)
	}
}

// RestoreXL from the "GongShape" sheet all GongShapeDB instances
func (backRepoGongShape *BackRepoGongShapeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongShapeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongShape"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongShape.rowVisitorGongShape)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongShape *BackRepoGongShapeStruct) rowVisitorGongShape(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongshapeWOP GongShapeWOP
		row.ReadStruct(&gongshapeWOP)

		// add the unmarshalled struct to the stage
		gongshapeDB := new(GongShapeDB)
		gongshapeDB.CopyBasicFieldsFromGongShapeWOP(&gongshapeWOP)

		gongshapeDB_ID_atBackupTime := gongshapeDB.ID
		gongshapeDB.ID = 0
		query := backRepoGongShape.db.Create(gongshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongShape.Map_GongShapeDBID_GongShapeDB[gongshapeDB.ID] = gongshapeDB
		BackRepoGongShapeid_atBckpTime_newID[gongshapeDB_ID_atBackupTime] = gongshapeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongShapeDB.json" in dirPath that stores an array
// of GongShapeDB and stores it in the database
// the map BackRepoGongShapeid_atBckpTime_newID is updated accordingly
func (backRepoGongShape *BackRepoGongShapeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongShapeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongShapeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongShape file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongShapeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongShapeDBID_GongShapeDB
	for _, gongshapeDB := range forRestore {

		gongshapeDB_ID_atBackupTime := gongshapeDB.ID
		gongshapeDB.ID = 0
		query := backRepoGongShape.db.Create(gongshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongShape.Map_GongShapeDBID_GongShapeDB[gongshapeDB.ID] = gongshapeDB
		BackRepoGongShapeid_atBckpTime_newID[gongshapeDB_ID_atBackupTime] = gongshapeDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongShape file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongShape>id_atBckpTime_newID
// to compute new index
func (backRepoGongShape *BackRepoGongShapeStruct) RestorePhaseTwo() {

	for _, gongshapeDB := range backRepoGongShape.Map_GongShapeDBID_GongShapeDB {

		// next line of code is to avert unused variable compilation error
		_ = gongshapeDB

		// insertion point for reindexing pointers encoding
		// reindexing Position field
		if gongshapeDB.PositionID.Int64 != 0 {
			gongshapeDB.PositionID.Int64 = int64(BackRepoPositionid_atBckpTime_newID[uint(gongshapeDB.PositionID.Int64)])
			gongshapeDB.PositionID.Valid = true
		}

		// This reindex gongshape.GongStructShapes
		if gongshapeDB.Classdiagram_GongStructShapesDBID.Int64 != 0 {
			gongshapeDB.Classdiagram_GongStructShapesDBID.Int64 =
				int64(BackRepoClassdiagramid_atBckpTime_newID[uint(gongshapeDB.Classdiagram_GongStructShapesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoGongShape.db.Model(gongshapeDB).Updates(*gongshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongShapeid_atBckpTime_newID map[uint]uint