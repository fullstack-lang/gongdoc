// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongEnumValueEntry__dummysDeclaration__ models.GongEnumValueEntry
var __GongEnumValueEntry_time__dummyDeclaration time.Duration

// An GongEnumValueEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnumValueEntry updateGongEnumValueEntry deleteGongEnumValueEntry
type GongEnumValueEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumValueEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnumValueEntry updateGongEnumValueEntry
type GongEnumValueEntryInput struct {
	// The GongEnumValueEntry to submit or modify
	// in: body
	GongEnumValueEntry *orm.GongEnumValueEntryAPI
}

// GetGongEnumValueEntrys
//
// swagger:route GET /gongenumvalueentrys gongenumvalueentrys getGongEnumValueEntrys
//
// # Get all gongenumvalueentrys
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func GetGongEnumValueEntrys(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongEnumValueEntry.GetDB()

	// source slice
	var gongenumvalueentryDBs []orm.GongEnumValueEntryDB

	// type Values map[string][]string
	values := c.Request.URL.Query()
	if len(values) == 1 {
		value := values["stack"]
		if len(value) == 1 {
			// we have a single parameter
			// we assume it is the stack
			stackParam := value[0]
			log.Println("GET all params", stackParam)
		}
	}

	query := db.Find(&gongenumvalueentryDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumvalueentryAPIs := make([]orm.GongEnumValueEntryAPI, 0)

	// for each gongenumvalueentry, update fields from the database nullable fields
	for idx := range gongenumvalueentryDBs {
		gongenumvalueentryDB := &gongenumvalueentryDBs[idx]
		_ = gongenumvalueentryDB
		var gongenumvalueentryAPI orm.GongEnumValueEntryAPI

		// insertion point for updating fields
		gongenumvalueentryAPI.ID = gongenumvalueentryDB.ID
		gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(&gongenumvalueentryAPI.GongEnumValueEntry)
		gongenumvalueentryAPI.GongEnumValueEntryPointersEnconding = gongenumvalueentryDB.GongEnumValueEntryPointersEnconding
		gongenumvalueentryAPIs = append(gongenumvalueentryAPIs, gongenumvalueentryAPI)
	}

	c.JSON(http.StatusOK, gongenumvalueentryAPIs)
}

// PostGongEnumValueEntry
//
// swagger:route POST /gongenumvalueentrys gongenumvalueentrys postGongEnumValueEntry
//
// Creates a gongenumvalueentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostGongEnumValueEntry(c *gin.Context) {

	// Validate input
	var input orm.GongEnumValueEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenumvalueentry
	gongenumvalueentryDB := orm.GongEnumValueEntryDB{}
	gongenumvalueentryDB.GongEnumValueEntryPointersEnconding = input.GongEnumValueEntryPointersEnconding
	gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry(&input.GongEnumValueEntry)

	db := orm.BackRepo.BackRepoGongEnumValueEntry.GetDB()
	query := db.Create(&gongenumvalueentryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoGongEnumValueEntry.CheckoutPhaseOneInstance(&gongenumvalueentryDB)
	gongenumvalueentry := (*orm.BackRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]

	if gongenumvalueentry != nil {
		models.AfterCreateFromFront(&models.Stage, gongenumvalueentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumvalueentryDB)
}

// GetGongEnumValueEntry
//
// swagger:route GET /gongenumvalueentrys/{ID} gongenumvalueentrys getGongEnumValueEntry
//
// Gets the details for a gongenumvalueentry.
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func GetGongEnumValueEntry(c *gin.Context) {

	// type Values map[string][]string
	values := c.Request.URL.Query()
	if len(values) == 1 {
		value := values["stack"]
		if len(value) == 1 {
			// we have a single parameter
			// we assume it is the stack
			stackParam := value[0]
			log.Println("GET params", stackParam)
		}
	}

	db := orm.BackRepo.BackRepoGongEnumValueEntry.GetDB()

	// Get gongenumvalueentryDB in DB
	var gongenumvalueentryDB orm.GongEnumValueEntryDB
	if err := db.First(&gongenumvalueentryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumvalueentryAPI orm.GongEnumValueEntryAPI
	gongenumvalueentryAPI.ID = gongenumvalueentryDB.ID
	gongenumvalueentryAPI.GongEnumValueEntryPointersEnconding = gongenumvalueentryDB.GongEnumValueEntryPointersEnconding
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(&gongenumvalueentryAPI.GongEnumValueEntry)

	c.JSON(http.StatusOK, gongenumvalueentryAPI)
}

// UpdateGongEnumValueEntry
//
// swagger:route PATCH /gongenumvalueentrys/{ID} gongenumvalueentrys updateGongEnumValueEntry
//
// # Update a gongenumvalueentry
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func UpdateGongEnumValueEntry(c *gin.Context) {

	// Validate input
	var input orm.GongEnumValueEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	db := orm.BackRepo.BackRepoGongEnumValueEntry.GetDB()

	// Get model if exist
	var gongenumvalueentryDB orm.GongEnumValueEntryDB

	// fetch the gongenumvalueentry
	query := db.First(&gongenumvalueentryDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry(&input.GongEnumValueEntry)
	gongenumvalueentryDB.GongEnumValueEntryPointersEnconding = input.GongEnumValueEntryPointersEnconding

	query = db.Model(&gongenumvalueentryDB).Updates(gongenumvalueentryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueentryNew := new(models.GongEnumValueEntry)
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentryNew)

	// get stage instance from DB instance, and call callback function
	gongenumvalueentryOld := (*orm.BackRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]
	if gongenumvalueentryOld != nil {
		models.AfterUpdateFromFront(&models.Stage, gongenumvalueentryOld, gongenumvalueentryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumvalueentryDB
	c.JSON(http.StatusOK, gongenumvalueentryDB)
}

// DeleteGongEnumValueEntry
//
// swagger:route DELETE /gongenumvalueentrys/{ID} gongenumvalueentrys deleteGongEnumValueEntry
//
// # Delete a gongenumvalueentry
//
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func DeleteGongEnumValueEntry(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongEnumValueEntry.GetDB()

	// Get model if exist
	var gongenumvalueentryDB orm.GongEnumValueEntryDB
	if err := db.First(&gongenumvalueentryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongenumvalueentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueentryDeleted := new(models.GongEnumValueEntry)
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentryDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumvalueentryStaged := (*orm.BackRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr)[gongenumvalueentryDB.ID]
	if gongenumvalueentryStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, gongenumvalueentryStaged, gongenumvalueentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}