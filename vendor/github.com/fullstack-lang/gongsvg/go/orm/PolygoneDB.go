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

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Polygone_sql sql.NullBool
var dummy_Polygone_time time.Duration
var dummy_Polygone_sort sort.Float64Slice

// PolygoneAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model polygoneAPI
type PolygoneAPI struct {
	gorm.Model

	models.Polygone

	// encoding of pointers
	PolygonePointersEnconding
}

// PolygonePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type PolygonePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Layer{}.Polygones []*Polygone
	Layer_PolygonesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Layer_PolygonesDBID_Index sql.NullInt64
}

// PolygoneDB describes a polygone in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model polygoneDB
type PolygoneDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field polygoneDB.Name
	Name_Data sql.NullString

	// Declation for basic field polygoneDB.Points
	Points_Data sql.NullString

	// Declation for basic field polygoneDB.Color
	Color_Data sql.NullString

	// Declation for basic field polygoneDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field polygoneDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field polygoneDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field polygoneDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field polygoneDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field polygoneDB.Transform
	Transform_Data sql.NullString
	// encoding of pointers
	PolygonePointersEnconding
}

// PolygoneDBs arrays polygoneDBs
// swagger:response polygoneDBsResponse
type PolygoneDBs []PolygoneDB

// PolygoneDBResponse provides response
// swagger:response polygoneDBResponse
type PolygoneDBResponse struct {
	PolygoneDB
}

// PolygoneWOP is a Polygone without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type PolygoneWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Points string `xlsx:"2"`

	Color string `xlsx:"3"`

	FillOpacity float64 `xlsx:"4"`

	Stroke string `xlsx:"5"`

	StrokeWidth float64 `xlsx:"6"`

	StrokeDashArray string `xlsx:"7"`

	StrokeDashArrayWhenSelected string `xlsx:"8"`

	Transform string `xlsx:"9"`
	// insertion for WOP pointer fields
}

var Polygone_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Points",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoPolygoneStruct struct {
	// stores PolygoneDB according to their gorm ID
	Map_PolygoneDBID_PolygoneDB map[uint]*PolygoneDB

	// stores PolygoneDB ID according to Polygone address
	Map_PolygonePtr_PolygoneDBID map[*models.Polygone]uint

	// stores Polygone according to their gorm ID
	Map_PolygoneDBID_PolygonePtr map[uint]*models.Polygone

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoPolygone *BackRepoPolygoneStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoPolygone.stage
	return
}

func (backRepoPolygone *BackRepoPolygoneStruct) GetDB() *gorm.DB {
	return backRepoPolygone.db
}

// GetPolygoneDBFromPolygonePtr is a handy function to access the back repo instance from the stage instance
func (backRepoPolygone *BackRepoPolygoneStruct) GetPolygoneDBFromPolygonePtr(polygone *models.Polygone) (polygoneDB *PolygoneDB) {
	id := backRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]
	polygoneDB = backRepoPolygone.Map_PolygoneDBID_PolygoneDB[id]
	return
}

// BackRepoPolygone.CommitPhaseOne commits all staged instances of Polygone to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPolygone *BackRepoPolygoneStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for polygone := range stage.Polygones {
		backRepoPolygone.CommitPhaseOneInstance(polygone)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, polygone := range backRepoPolygone.Map_PolygoneDBID_PolygonePtr {
		if _, ok := stage.Polygones[polygone]; !ok {
			backRepoPolygone.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoPolygone.CommitDeleteInstance commits deletion of Polygone to the BackRepo
func (backRepoPolygone *BackRepoPolygoneStruct) CommitDeleteInstance(id uint) (Error error) {

	polygone := backRepoPolygone.Map_PolygoneDBID_PolygonePtr[id]

	// polygone is not staged anymore, remove polygoneDB
	polygoneDB := backRepoPolygone.Map_PolygoneDBID_PolygoneDB[id]
	query := backRepoPolygone.db.Unscoped().Delete(&polygoneDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoPolygone.Map_PolygonePtr_PolygoneDBID, polygone)
	delete(backRepoPolygone.Map_PolygoneDBID_PolygonePtr, id)
	delete(backRepoPolygone.Map_PolygoneDBID_PolygoneDB, id)

	return
}

// BackRepoPolygone.CommitPhaseOneInstance commits polygone staged instances of Polygone to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPolygone *BackRepoPolygoneStruct) CommitPhaseOneInstance(polygone *models.Polygone) (Error error) {

	// check if the polygone is not commited yet
	if _, ok := backRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]; ok {
		return
	}

	// initiate polygone
	var polygoneDB PolygoneDB
	polygoneDB.CopyBasicFieldsFromPolygone(polygone)

	query := backRepoPolygone.db.Create(&polygoneDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone] = polygoneDB.ID
	backRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID] = polygone
	backRepoPolygone.Map_PolygoneDBID_PolygoneDB[polygoneDB.ID] = &polygoneDB

	return
}

// BackRepoPolygone.CommitPhaseTwo commits all staged instances of Polygone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolygone *BackRepoPolygoneStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, polygone := range backRepoPolygone.Map_PolygoneDBID_PolygonePtr {
		backRepoPolygone.CommitPhaseTwoInstance(backRepo, idx, polygone)
	}

	return
}

// BackRepoPolygone.CommitPhaseTwoInstance commits {{structname }} of models.Polygone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolygone *BackRepoPolygoneStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, polygone *models.Polygone) (Error error) {

	// fetch matching polygoneDB
	if polygoneDB, ok := backRepoPolygone.Map_PolygoneDBID_PolygoneDB[idx]; ok {

		polygoneDB.CopyBasicFieldsFromPolygone(polygone)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers polygone.Animates into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, animateAssocEnd := range polygone.Animates {

			// get the back repo instance at the association end
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)

			// encode reverse pointer in the association end back repo instance
			animateAssocEnd_DB.Polygone_AnimatesDBID.Int64 = int64(polygoneDB.ID)
			animateAssocEnd_DB.Polygone_AnimatesDBID.Valid = true
			animateAssocEnd_DB.Polygone_AnimatesDBID_Index.Int64 = int64(idx)
			animateAssocEnd_DB.Polygone_AnimatesDBID_Index.Valid = true
			if q := backRepoPolygone.db.Save(animateAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoPolygone.db.Save(&polygoneDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Polygone intance %s", polygone.Name))
		return err
	}

	return
}

// BackRepoPolygone.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoPolygone *BackRepoPolygoneStruct) CheckoutPhaseOne() (Error error) {

	polygoneDBArray := make([]PolygoneDB, 0)
	query := backRepoPolygone.db.Find(&polygoneDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	polygoneInstancesToBeRemovedFromTheStage := make(map[*models.Polygone]any)
	for key, value := range backRepoPolygone.stage.Polygones {
		polygoneInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, polygoneDB := range polygoneDBArray {
		backRepoPolygone.CheckoutPhaseOneInstance(&polygoneDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		polygone, ok := backRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
		if ok {
			delete(polygoneInstancesToBeRemovedFromTheStage, polygone)
		}
	}

	// remove from stage and back repo's 3 maps all polygones that are not in the checkout
	for polygone := range polygoneInstancesToBeRemovedFromTheStage {
		polygone.Unstage(backRepoPolygone.GetStage())

		// remove instance from the back repo 3 maps
		polygoneID := backRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]
		delete(backRepoPolygone.Map_PolygonePtr_PolygoneDBID, polygone)
		delete(backRepoPolygone.Map_PolygoneDBID_PolygoneDB, polygoneID)
		delete(backRepoPolygone.Map_PolygoneDBID_PolygonePtr, polygoneID)
	}

	return
}

// CheckoutPhaseOneInstance takes a polygoneDB that has been found in the DB, updates the backRepo and stages the
// models version of the polygoneDB
func (backRepoPolygone *BackRepoPolygoneStruct) CheckoutPhaseOneInstance(polygoneDB *PolygoneDB) (Error error) {

	polygone, ok := backRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
	if !ok {
		polygone = new(models.Polygone)

		backRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID] = polygone
		backRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone] = polygoneDB.ID

		// append model store with the new element
		polygone.Name = polygoneDB.Name_Data.String
		polygone.Stage(backRepoPolygone.GetStage())
	}
	polygoneDB.CopyBasicFieldsToPolygone(polygone)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	polygone.Stage(backRepoPolygone.GetStage())

	// preserve pointer to polygoneDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_PolygoneDBID_PolygoneDB)[polygoneDB hold variable pointers
	polygoneDB_Data := *polygoneDB
	preservedPtrToPolygone := &polygoneDB_Data
	backRepoPolygone.Map_PolygoneDBID_PolygoneDB[polygoneDB.ID] = preservedPtrToPolygone

	return
}

// BackRepoPolygone.CheckoutPhaseTwo Checkouts all staged instances of Polygone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolygone *BackRepoPolygoneStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, polygoneDB := range backRepoPolygone.Map_PolygoneDBID_PolygoneDB {
		backRepoPolygone.CheckoutPhaseTwoInstance(backRepo, polygoneDB)
	}
	return
}

// BackRepoPolygone.CheckoutPhaseTwoInstance Checkouts staged instances of Polygone to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPolygone *BackRepoPolygoneStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, polygoneDB *PolygoneDB) (Error error) {

	polygone := backRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
	_ = polygone // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem polygone.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	polygone.Animates = polygone.Animates[:0]
	// 2. loop all instances in the type in the association end
	for _, animateDB_AssocEnd := range backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if animateDB_AssocEnd.Polygone_AnimatesDBID.Int64 == int64(polygoneDB.ID) {
			// 4. fetch the associated instance in the stage
			animate_AssocEnd := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB_AssocEnd.ID]
			// 5. append it the association slice
			polygone.Animates = append(polygone.Animates, animate_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(polygone.Animates, func(i, j int) bool {
		animateDB_i_ID := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[polygone.Animates[i]]
		animateDB_j_ID := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[polygone.Animates[j]]

		animateDB_i := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB_i_ID]
		animateDB_j := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB_j_ID]

		return animateDB_i.Polygone_AnimatesDBID_Index.Int64 < animateDB_j.Polygone_AnimatesDBID_Index.Int64
	})

	return
}

// CommitPolygone allows commit of a single polygone (if already staged)
func (backRepo *BackRepoStruct) CommitPolygone(polygone *models.Polygone) {
	backRepo.BackRepoPolygone.CommitPhaseOneInstance(polygone)
	if id, ok := backRepo.BackRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]; ok {
		backRepo.BackRepoPolygone.CommitPhaseTwoInstance(backRepo, id, polygone)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitPolygone allows checkout of a single polygone (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutPolygone(polygone *models.Polygone) {
	// check if the polygone is staged
	if _, ok := backRepo.BackRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]; ok {

		if id, ok := backRepo.BackRepoPolygone.Map_PolygonePtr_PolygoneDBID[polygone]; ok {
			var polygoneDB PolygoneDB
			polygoneDB.ID = id

			if err := backRepo.BackRepoPolygone.db.First(&polygoneDB, id).Error; err != nil {
				log.Panicln("CheckoutPolygone : Problem with getting object with id:", id)
			}
			backRepo.BackRepoPolygone.CheckoutPhaseOneInstance(&polygoneDB)
			backRepo.BackRepoPolygone.CheckoutPhaseTwoInstance(backRepo, &polygoneDB)
		}
	}
}

// CopyBasicFieldsFromPolygone
func (polygoneDB *PolygoneDB) CopyBasicFieldsFromPolygone(polygone *models.Polygone) {
	// insertion point for fields commit

	polygoneDB.Name_Data.String = polygone.Name
	polygoneDB.Name_Data.Valid = true

	polygoneDB.Points_Data.String = polygone.Points
	polygoneDB.Points_Data.Valid = true

	polygoneDB.Color_Data.String = polygone.Color
	polygoneDB.Color_Data.Valid = true

	polygoneDB.FillOpacity_Data.Float64 = polygone.FillOpacity
	polygoneDB.FillOpacity_Data.Valid = true

	polygoneDB.Stroke_Data.String = polygone.Stroke
	polygoneDB.Stroke_Data.Valid = true

	polygoneDB.StrokeWidth_Data.Float64 = polygone.StrokeWidth
	polygoneDB.StrokeWidth_Data.Valid = true

	polygoneDB.StrokeDashArray_Data.String = polygone.StrokeDashArray
	polygoneDB.StrokeDashArray_Data.Valid = true

	polygoneDB.StrokeDashArrayWhenSelected_Data.String = polygone.StrokeDashArrayWhenSelected
	polygoneDB.StrokeDashArrayWhenSelected_Data.Valid = true

	polygoneDB.Transform_Data.String = polygone.Transform
	polygoneDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromPolygoneWOP
func (polygoneDB *PolygoneDB) CopyBasicFieldsFromPolygoneWOP(polygone *PolygoneWOP) {
	// insertion point for fields commit

	polygoneDB.Name_Data.String = polygone.Name
	polygoneDB.Name_Data.Valid = true

	polygoneDB.Points_Data.String = polygone.Points
	polygoneDB.Points_Data.Valid = true

	polygoneDB.Color_Data.String = polygone.Color
	polygoneDB.Color_Data.Valid = true

	polygoneDB.FillOpacity_Data.Float64 = polygone.FillOpacity
	polygoneDB.FillOpacity_Data.Valid = true

	polygoneDB.Stroke_Data.String = polygone.Stroke
	polygoneDB.Stroke_Data.Valid = true

	polygoneDB.StrokeWidth_Data.Float64 = polygone.StrokeWidth
	polygoneDB.StrokeWidth_Data.Valid = true

	polygoneDB.StrokeDashArray_Data.String = polygone.StrokeDashArray
	polygoneDB.StrokeDashArray_Data.Valid = true

	polygoneDB.StrokeDashArrayWhenSelected_Data.String = polygone.StrokeDashArrayWhenSelected
	polygoneDB.StrokeDashArrayWhenSelected_Data.Valid = true

	polygoneDB.Transform_Data.String = polygone.Transform
	polygoneDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToPolygone
func (polygoneDB *PolygoneDB) CopyBasicFieldsToPolygone(polygone *models.Polygone) {
	// insertion point for checkout of basic fields (back repo to stage)
	polygone.Name = polygoneDB.Name_Data.String
	polygone.Points = polygoneDB.Points_Data.String
	polygone.Color = polygoneDB.Color_Data.String
	polygone.FillOpacity = polygoneDB.FillOpacity_Data.Float64
	polygone.Stroke = polygoneDB.Stroke_Data.String
	polygone.StrokeWidth = polygoneDB.StrokeWidth_Data.Float64
	polygone.StrokeDashArray = polygoneDB.StrokeDashArray_Data.String
	polygone.StrokeDashArrayWhenSelected = polygoneDB.StrokeDashArrayWhenSelected_Data.String
	polygone.Transform = polygoneDB.Transform_Data.String
}

// CopyBasicFieldsToPolygoneWOP
func (polygoneDB *PolygoneDB) CopyBasicFieldsToPolygoneWOP(polygone *PolygoneWOP) {
	polygone.ID = int(polygoneDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	polygone.Name = polygoneDB.Name_Data.String
	polygone.Points = polygoneDB.Points_Data.String
	polygone.Color = polygoneDB.Color_Data.String
	polygone.FillOpacity = polygoneDB.FillOpacity_Data.Float64
	polygone.Stroke = polygoneDB.Stroke_Data.String
	polygone.StrokeWidth = polygoneDB.StrokeWidth_Data.Float64
	polygone.StrokeDashArray = polygoneDB.StrokeDashArray_Data.String
	polygone.StrokeDashArrayWhenSelected = polygoneDB.StrokeDashArrayWhenSelected_Data.String
	polygone.Transform = polygoneDB.Transform_Data.String
}

// Backup generates a json file from a slice of all PolygoneDB instances in the backrepo
func (backRepoPolygone *BackRepoPolygoneStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "PolygoneDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PolygoneDB, 0)
	for _, polygoneDB := range backRepoPolygone.Map_PolygoneDBID_PolygoneDB {
		forBackup = append(forBackup, polygoneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Polygone ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Polygone file", err.Error())
	}
}

// Backup generates a json file from a slice of all PolygoneDB instances in the backrepo
func (backRepoPolygone *BackRepoPolygoneStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PolygoneDB, 0)
	for _, polygoneDB := range backRepoPolygone.Map_PolygoneDBID_PolygoneDB {
		forBackup = append(forBackup, polygoneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Polygone")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Polygone_Fields, -1)
	for _, polygoneDB := range forBackup {

		var polygoneWOP PolygoneWOP
		polygoneDB.CopyBasicFieldsToPolygoneWOP(&polygoneWOP)

		row := sh.AddRow()
		row.WriteStruct(&polygoneWOP, -1)
	}
}

// RestoreXL from the "Polygone" sheet all PolygoneDB instances
func (backRepoPolygone *BackRepoPolygoneStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoPolygoneid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Polygone"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoPolygone.rowVisitorPolygone)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoPolygone *BackRepoPolygoneStruct) rowVisitorPolygone(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var polygoneWOP PolygoneWOP
		row.ReadStruct(&polygoneWOP)

		// add the unmarshalled struct to the stage
		polygoneDB := new(PolygoneDB)
		polygoneDB.CopyBasicFieldsFromPolygoneWOP(&polygoneWOP)

		polygoneDB_ID_atBackupTime := polygoneDB.ID
		polygoneDB.ID = 0
		query := backRepoPolygone.db.Create(polygoneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoPolygone.Map_PolygoneDBID_PolygoneDB[polygoneDB.ID] = polygoneDB
		BackRepoPolygoneid_atBckpTime_newID[polygoneDB_ID_atBackupTime] = polygoneDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "PolygoneDB.json" in dirPath that stores an array
// of PolygoneDB and stores it in the database
// the map BackRepoPolygoneid_atBckpTime_newID is updated accordingly
func (backRepoPolygone *BackRepoPolygoneStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoPolygoneid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "PolygoneDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Polygone file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*PolygoneDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_PolygoneDBID_PolygoneDB
	for _, polygoneDB := range forRestore {

		polygoneDB_ID_atBackupTime := polygoneDB.ID
		polygoneDB.ID = 0
		query := backRepoPolygone.db.Create(polygoneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoPolygone.Map_PolygoneDBID_PolygoneDB[polygoneDB.ID] = polygoneDB
		BackRepoPolygoneid_atBckpTime_newID[polygoneDB_ID_atBackupTime] = polygoneDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Polygone file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Polygone>id_atBckpTime_newID
// to compute new index
func (backRepoPolygone *BackRepoPolygoneStruct) RestorePhaseTwo() {

	for _, polygoneDB := range backRepoPolygone.Map_PolygoneDBID_PolygoneDB {

		// next line of code is to avert unused variable compilation error
		_ = polygoneDB

		// insertion point for reindexing pointers encoding
		// This reindex polygone.Polygones
		if polygoneDB.Layer_PolygonesDBID.Int64 != 0 {
			polygoneDB.Layer_PolygonesDBID.Int64 =
				int64(BackRepoLayerid_atBckpTime_newID[uint(polygoneDB.Layer_PolygonesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoPolygone.db.Model(polygoneDB).Updates(*polygoneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoPolygoneid_atBckpTime_newID map[uint]uint