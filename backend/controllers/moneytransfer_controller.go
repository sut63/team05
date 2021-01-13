package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/moneytransfer"
)

// MoneytransferController defines the struct for the Moneytransfer controller
type MoneytransferController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMoneytransfer handles POST requests for adding moneytransfer entities
// @Summary Create moneytransfer
// @Description Create moneytransfer
// @ID create-moneyTransfer
// @Accept   json
// @Produce  json
// @Param moneytransfer body ent.Moneytransfer true "Moneytransfer entity"
// @Success 200 {object} ent.Moneytransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneytransfers [post]
func (ctl *MoneytransferController) CreateMoneytransfer(c *gin.Context) {
	obj := ent.Moneytransfer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "moneytransfer binding failed",
		})
		return
	}

	m, err := ctl.client.Moneytransfer.
		Create().
		SetMoneytransferType(obj.MoneytransferType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetMoneytransfer handles GET requests to retrieve a moneytransfer entity
// @Summary Get a moneytransfer entity by ID
// @Description get moneytransfer by ID
// @ID get-moneytransfer
// @Produce  json
// @Param id path int true "Moneytransfer ID"
// @Success 200 {object} ent.Moneytransfer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneytransfers/{id} [get]
func (ctl *MoneytransferController) GetMoneytransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Moneytransfer.
		Query().
		Where(moneytransfer.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMoneytransfer handles request to get a list of moneytransfer entities
// @Summary List moneytransfer entities
// @Description list moneytransfer entities
// @ID list-moneytransfer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Moneytransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneytransfers [get]
func (ctl *MoneytransferController) ListMoneytransfer(c *gin.Context) {
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

	moneytransfers, err := ctl.client.Moneytransfer.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, moneytransfers)
}

// DeleteMoneytransfer handles DELETE requests to delete a moneytransfer entity
// @Summary Delete a moneytransfer entity by ID
// @Description get moneytransfer by ID
// @ID delete-moneytransfer
// @Produce  json
// @Param id path int true "Moneytransfer ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneytransfers/{id} [delete]
func (ctl *MoneytransferController) DeleteMoneytransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Moneytransfer.
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

// UpdateMoneytransfer handles PUT requests to update a moneytransfer entity
// @Summary Update a moneytransfer entity by ID
// @Description update moneytransfer by ID
// @ID update-moneytransfer
// @Accept   json
// @Produce  json
// @Param id path int true "moneytransfer ID"
// @Param moneytransfer body ent.Moneytransfer true "Moneytransfer entity"
// @Success 200 {object} ent.Moneytransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneytransfers/{id} [put]
func (ctl *MoneytransferController) UpdateMoneytransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Moneytransfer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Moneytransfer binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.Moneytransfer.
		UpdateOneID(int(id)).
		SetMoneytransferType(obj.MoneytransferType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewMoneytransferController creates and registers handles for the MoneytransferController
func NewMoneytransferController(router gin.IRouter, client *ent.Client) *MoneytransferController {
	mc := &MoneytransferController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMoneytransferController registers routes to the main engine
func (ctl *MoneytransferController) register() {
	moneytransfers := ctl.router.Group("/moneytransfers")
	moneytransfers.GET("", ctl.ListMoneytransfer)
	// CRUD
	moneytransfers.POST("", ctl.CreateMoneytransfer)
	moneytransfers.GET(":id", ctl.GetMoneytransfer)
	moneytransfers.PUT(":id", ctl.UpdateMoneytransfer)
	moneytransfers.DELETE(":id", ctl.DeleteMoneytransfer)
}
