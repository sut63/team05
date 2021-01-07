package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/officer"
)

//OfficerController struct
type OfficerController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateOfficer handles POST requests for adding officer entities
// @Summary Create officer
// @Description Create officer
// @ID create-officer
// @Accept   json
// @Produce  json
// @Param officer body ent.Officer true "Officer entity"
// @Success 200 {object} ent.Officer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officers [post]
func (ctl *OfficerController) CreateOfficer(c *gin.Context) {
	obj := ent.Officer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "officer binding failed",
		})
		return
	}

	ofc, err := ctl.client.Officer.
		Create().
		SetOfficerEmail(obj.OfficerEmail).
		SetOfficerName(obj.OfficerName).
		SetOfficerPassword(obj.OfficerPassword).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving officer failed",
		})
		return
	}

	c.JSON(200, ofc)
}

// GetOfficer handles GET requests to retrieve a officer entity
// @Summary Get a officer entity by ID
// @Description get officer by ID
// @ID get-officer
// @Produce  json
// @Param id path int true "Officer ID"
// @Success 200 {object} ent.Officer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officers/{id} [get]
func (ctl *OfficerController) GetOfficer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ofc, err := ctl.client.Officer.
		Query().
		Where(officer.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ofc)
}

// ListOfficer handles request to get a list of officer entities
// @Summary List officer entities
// @Description list officer entities
// @ID list-officer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Officer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officers [get]
func (ctl *OfficerController) ListOfficer(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	officers, err := ctl.client.Officer.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, officers)
}

// DeleteOfficer handles DELETE requests to delete a officer entity
// @Summary Delete a officer entity by ID
// @Description get officer by ID
// @ID delete-officer
// @Produce  json
// @Param id path int true "Officer ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officers/{id} [delete]
func (ctl *OfficerController) DeleteOfficer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Officer.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateOfficer handles PUT requests to update a officer entity
// @Summary Update a officer entity by ID
// @Description update officer by ID
// @ID update-officer
// @Accept   json
// @Produce  json
// @Param id path int true "Officer ID"
// @Param officer body ent.Officer true "Officer entity"
// @Success 200 {object} ent.Officer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officers/{id} [put]
func (ctl *OfficerController) UpdateOfficer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Officer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "officer binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	ofc, err := ctl.client.Officer.
		UpdateOneID(int(id)).
		SetOfficerEmail(obj.OfficerEmail).
		SetOfficerName(obj.OfficerName).
		SetOfficerPassword(obj.OfficerPassword).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update officer failed",
		})
		return
	}

	c.JSON(200, ofc)
}

// NewOfficerController creates and registers handles for the oder controller
func NewOfficerController(router gin.IRouter, client *ent.Client) *OfficerController {
	ofcc := &OfficerController{
		client: client,
		router: router,
	}

	ofcc.register()

	return ofcc

}

func (ctl *OfficerController) register() {
	officers := ctl.router.Group("/officers")

	officers.POST("", ctl.CreateOfficer)
	officers.GET(":id", ctl.GetOfficer)
	officers.GET("", ctl.ListOfficer)
	officers.DELETE(":id", ctl.DeleteOfficer)
	officers.PUT(":id", ctl.UpdateOfficer)

}
