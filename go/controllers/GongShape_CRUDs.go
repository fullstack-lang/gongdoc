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
var __GongShape__dummysDeclaration__ models.GongShape
var __GongShape_time__dummyDeclaration time.Duration

// An GongShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongShape updateGongShape deleteGongShape
type GongShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongShape updateGongShape
type GongShapeInput struct {
	// The GongShape to submit or modify
	// in: body
	GongShape *orm.GongShapeAPI
}

// GetGongShapes
//
// swagger:route GET /gongshapes gongshapes getGongShapes
//
// # Get all gongshapes
//
// Responses:
// default: genericError
//
//	200: gongshapeDBResponse
func (controller *Controller) GetGongShapes(c *gin.Context) {

	// source slice
	var gongshapeDBs []orm.GongShapeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetGongShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongShape.GetDB()

	query := db.Find(&gongshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongshapeAPIs := make([]orm.GongShapeAPI, 0)

	// for each gongshape, update fields from the database nullable fields
	for idx := range gongshapeDBs {
		gongshapeDB := &gongshapeDBs[idx]
		_ = gongshapeDB
		var gongshapeAPI orm.GongShapeAPI

		// insertion point for updating fields
		gongshapeAPI.ID = gongshapeDB.ID
		gongshapeDB.CopyBasicFieldsToGongShape(&gongshapeAPI.GongShape)
		gongshapeAPI.GongShapePointersEnconding = gongshapeDB.GongShapePointersEnconding
		gongshapeAPIs = append(gongshapeAPIs, gongshapeAPI)
	}

	c.JSON(http.StatusOK, gongshapeAPIs)
}

// PostGongShape
//
// swagger:route POST /gongshapes gongshapes postGongShape
//
// Creates a gongshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("PostGongShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongShape.GetDB()

	// Validate input
	var input orm.GongShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongshape
	gongshapeDB := orm.GongShapeDB{}
	gongshapeDB.GongShapePointersEnconding = input.GongShapePointersEnconding
	gongshapeDB.CopyBasicFieldsFromGongShape(&input.GongShape)

	query := db.Create(&gongshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongShape.CheckoutPhaseOneInstance(&gongshapeDB)
	gongshape := backRepo.BackRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]

	if gongshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongshapeDB)
}

// GetGongShape
//
// swagger:route GET /gongshapes/{ID} gongshapes getGongShape
//
// Gets the details for a gongshape.
//
// Responses:
// default: genericError
//
//	200: gongshapeDBResponse
func (controller *Controller) GetGongShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetGongShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongShape.GetDB()

	// Get gongshapeDB in DB
	var gongshapeDB orm.GongShapeDB
	if err := db.First(&gongshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongshapeAPI orm.GongShapeAPI
	gongshapeAPI.ID = gongshapeDB.ID
	gongshapeAPI.GongShapePointersEnconding = gongshapeDB.GongShapePointersEnconding
	gongshapeDB.CopyBasicFieldsToGongShape(&gongshapeAPI.GongShape)

	c.JSON(http.StatusOK, gongshapeAPI)
}

// UpdateGongShape
//
// swagger:route PATCH /gongshapes/{ID} gongshapes updateGongShape
//
// # Update a gongshape
//
// Responses:
// default: genericError
//
//	200: gongshapeDBResponse
func (controller *Controller) UpdateGongShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("UpdateGongShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongShape.GetDB()

	// Validate input
	var input orm.GongShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongshapeDB orm.GongShapeDB

	// fetch the gongshape
	query := db.First(&gongshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongshapeDB.CopyBasicFieldsFromGongShape(&input.GongShape)
	gongshapeDB.GongShapePointersEnconding = input.GongShapePointersEnconding

	query = db.Model(&gongshapeDB).Updates(gongshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongshapeNew := new(models.GongShape)
	gongshapeDB.CopyBasicFieldsToGongShape(gongshapeNew)

	// get stage instance from DB instance, and call callback function
	gongshapeOld := backRepo.BackRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]
	if gongshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongshapeOld, gongshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongshapeDB
	c.JSON(http.StatusOK, gongshapeDB)
}

// DeleteGongShape
//
// swagger:route DELETE /gongshapes/{ID} gongshapes deleteGongShape
//
// # Delete a gongshape
//
// default: genericError
//
//	200: gongshapeDBResponse
func (controller *Controller) DeleteGongShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("DeleteGongShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongShape.GetDB()

	// Get model if exist
	var gongshapeDB orm.GongShapeDB
	if err := db.First(&gongshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongshapeDeleted := new(models.GongShape)
	gongshapeDB.CopyBasicFieldsToGongShape(gongshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongshapeStaged := backRepo.BackRepoGongShape.Map_GongShapeDBID_GongShapePtr[gongshapeDB.ID]
	if gongshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongshapeStaged, gongshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}