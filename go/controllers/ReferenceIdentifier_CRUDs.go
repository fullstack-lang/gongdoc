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
var __ReferenceIdentifier__dummysDeclaration__ models.ReferenceIdentifier
var __ReferenceIdentifier_time__dummyDeclaration time.Duration

// An ReferenceIdentifierID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getReferenceIdentifier updateReferenceIdentifier deleteReferenceIdentifier
type ReferenceIdentifierID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ReferenceIdentifierInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postReferenceIdentifier updateReferenceIdentifier
type ReferenceIdentifierInput struct {
	// The ReferenceIdentifier to submit or modify
	// in: body
	ReferenceIdentifier *orm.ReferenceIdentifierAPI
}

// GetReferenceIdentifiers
//
// swagger:route GET /referenceidentifiers referenceidentifiers getReferenceIdentifiers
//
// # Get all referenceidentifiers
//
// Responses:
// default: genericError
//
//	200: referenceidentifierDBResponse
func GetReferenceIdentifiers(c *gin.Context) {
	db := orm.BackRepo.BackRepoReferenceIdentifier.GetDB()

	// source slice
	var referenceidentifierDBs []orm.ReferenceIdentifierDB
	query := db.Find(&referenceidentifierDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	referenceidentifierAPIs := make([]orm.ReferenceIdentifierAPI, 0)

	// for each referenceidentifier, update fields from the database nullable fields
	for idx := range referenceidentifierDBs {
		referenceidentifierDB := &referenceidentifierDBs[idx]
		_ = referenceidentifierDB
		var referenceidentifierAPI orm.ReferenceIdentifierAPI

		// insertion point for updating fields
		referenceidentifierAPI.ID = referenceidentifierDB.ID
		referenceidentifierDB.CopyBasicFieldsToReferenceIdentifier(&referenceidentifierAPI.ReferenceIdentifier)
		referenceidentifierAPI.ReferenceIdentifierPointersEnconding = referenceidentifierDB.ReferenceIdentifierPointersEnconding
		referenceidentifierAPIs = append(referenceidentifierAPIs, referenceidentifierAPI)
	}

	c.JSON(http.StatusOK, referenceidentifierAPIs)
}

// PostReferenceIdentifier
//
// swagger:route POST /referenceidentifiers referenceidentifiers postReferenceIdentifier
//
// Creates a referenceidentifier
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostReferenceIdentifier(c *gin.Context) {
	db := orm.BackRepo.BackRepoReferenceIdentifier.GetDB()

	// Validate input
	var input orm.ReferenceIdentifierAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create referenceidentifier
	referenceidentifierDB := orm.ReferenceIdentifierDB{}
	referenceidentifierDB.ReferenceIdentifierPointersEnconding = input.ReferenceIdentifierPointersEnconding
	referenceidentifierDB.CopyBasicFieldsFromReferenceIdentifier(&input.ReferenceIdentifier)

	query := db.Create(&referenceidentifierDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoReferenceIdentifier.CheckoutPhaseOneInstance(&referenceidentifierDB)
	referenceidentifier := (*orm.BackRepo.BackRepoReferenceIdentifier.Map_ReferenceIdentifierDBID_ReferenceIdentifierPtr)[referenceidentifierDB.ID]

	if referenceidentifier != nil {
		models.AfterCreateFromFront(&models.Stage, referenceidentifier)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, referenceidentifierDB)
}

// GetReferenceIdentifier
//
// swagger:route GET /referenceidentifiers/{ID} referenceidentifiers getReferenceIdentifier
//
// Gets the details for a referenceidentifier.
//
// Responses:
// default: genericError
//
//	200: referenceidentifierDBResponse
func GetReferenceIdentifier(c *gin.Context) {
	db := orm.BackRepo.BackRepoReferenceIdentifier.GetDB()

	// Get referenceidentifierDB in DB
	var referenceidentifierDB orm.ReferenceIdentifierDB
	if err := db.First(&referenceidentifierDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var referenceidentifierAPI orm.ReferenceIdentifierAPI
	referenceidentifierAPI.ID = referenceidentifierDB.ID
	referenceidentifierAPI.ReferenceIdentifierPointersEnconding = referenceidentifierDB.ReferenceIdentifierPointersEnconding
	referenceidentifierDB.CopyBasicFieldsToReferenceIdentifier(&referenceidentifierAPI.ReferenceIdentifier)

	c.JSON(http.StatusOK, referenceidentifierAPI)
}

// UpdateReferenceIdentifier
//
// swagger:route PATCH /referenceidentifiers/{ID} referenceidentifiers updateReferenceIdentifier
//
// # Update a referenceidentifier
//
// Responses:
// default: genericError
//
//	200: referenceidentifierDBResponse
func UpdateReferenceIdentifier(c *gin.Context) {
	db := orm.BackRepo.BackRepoReferenceIdentifier.GetDB()

	// Get model if exist
	var referenceidentifierDB orm.ReferenceIdentifierDB

	// fetch the referenceidentifier
	query := db.First(&referenceidentifierDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ReferenceIdentifierAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	referenceidentifierDB.CopyBasicFieldsFromReferenceIdentifier(&input.ReferenceIdentifier)
	referenceidentifierDB.ReferenceIdentifierPointersEnconding = input.ReferenceIdentifierPointersEnconding

	query = db.Model(&referenceidentifierDB).Updates(referenceidentifierDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	referenceidentifierNew := new(models.ReferenceIdentifier)
	referenceidentifierDB.CopyBasicFieldsToReferenceIdentifier(referenceidentifierNew)

	// get stage instance from DB instance, and call callback function
	referenceidentifierOld := (*orm.BackRepo.BackRepoReferenceIdentifier.Map_ReferenceIdentifierDBID_ReferenceIdentifierPtr)[referenceidentifierDB.ID]
	if referenceidentifierOld != nil {
		models.AfterUpdateFromFront(&models.Stage, referenceidentifierOld, referenceidentifierNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the referenceidentifierDB
	c.JSON(http.StatusOK, referenceidentifierDB)
}

// DeleteReferenceIdentifier
//
// swagger:route DELETE /referenceidentifiers/{ID} referenceidentifiers deleteReferenceIdentifier
//
// # Delete a referenceidentifier
//
// default: genericError
//
//	200: referenceidentifierDBResponse
func DeleteReferenceIdentifier(c *gin.Context) {
	db := orm.BackRepo.BackRepoReferenceIdentifier.GetDB()

	// Get model if exist
	var referenceidentifierDB orm.ReferenceIdentifierDB
	if err := db.First(&referenceidentifierDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&referenceidentifierDB)

	// get an instance (not staged) from DB instance, and call callback function
	referenceidentifierDeleted := new(models.ReferenceIdentifier)
	referenceidentifierDB.CopyBasicFieldsToReferenceIdentifier(referenceidentifierDeleted)

	// get stage instance from DB instance, and call callback function
	referenceidentifierStaged := (*orm.BackRepo.BackRepoReferenceIdentifier.Map_ReferenceIdentifierDBID_ReferenceIdentifierPtr)[referenceidentifierDB.ID]
	if referenceidentifierStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, referenceidentifierStaged, referenceidentifierDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}