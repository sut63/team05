package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/moneytransfer"
)

// MoneyTransferController defines the struct for the MoneyTransfer controller
type MoneyTransferController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMoneyTransfer handles POST requests for adding MoneyTransfer entities
// @Summary Create MoneyTransfer
// @Description Create MoneyTransfer
// @ID create-MoneyTransfer
// @Accept   json
// @Produce  json
// @Param MoneyTransfer body ent.MoneyTransfer true "MoneyTransfer entity"
// @Success 200 {object} ent.MoneyTransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneyTransfers [post]
func (ctl *MoneyTransferController) CreateMoneyTransfer(c *gin.Context) {
	obj := ent.MoneyTransfer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "moneyTransfer binding failed",
		})
		return
	}

	m, err := ctl.client.MoneyTransfer.
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

// GetMoneyTransfer handles GET requests to retrieve a MoneyTransfer entity
// @Summary Get a MoneyTransfer entity by ID
// @Description get MoneyTransfer by ID
// @ID get-MoneyTransfer
// @Produce  json
// @Param id path int true "MoneyTransfer ID"
// @Success 200 {object} ent.MoneyTransfer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneyTransfers/{id} [get]
func (ctl *MoneyTransferController) GetMoneyTransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.MoneyTransfer.
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

// ListMoneyTransfer handles request to get a list of MoneyTransfer entities
// @Summary List MoneyTransfer entities
// @Description list MoneyTransfer entities
// @ID list-MoneyTransfer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MoneyTransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneyTransfers [get]
func (ctl *MoneyTransferController) ListMoneyTransfer(c *gin.Context) {
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

	moneyTransfers, err := ctl.client.MoneyTransfer.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, moneyTransfers)
}

// DeleteMoneyTransfer handles DELETE requests to delete a MoneyTransfer entity
// @Summary Delete a MoneyTransfer entity by ID
// @Description get MoneyTransfer by ID
// @ID delete-MoneyTransfer
// @Produce  json
// @Param id path int true "MoneyTransfer ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneyTransfers/{id} [delete]
func (ctl *MoneyTransferController) DeleteMoneyTransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.MoneyTransfer.
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

// UpdateMoneyTransfer handles PUT requests to update a MoneyTransfer entity
// @Summary Update a MoneyTransfer entity by ID
// @Description update MoneyTransfer by ID
// @ID update-MoneyTransfer
// @Accept   json
// @Produce  json
// @Param id path int true "MoneyTransfer ID"
// @Param MoneyTransfer body ent.MoneyTransfer true "MoneyTransfer entity"
// @Success 200 {object} ent.MoneyTransfer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /moneyTransfers/{id} [put]
func (ctl *MoneyTransferController) UpdateMoneyTransfer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.MoneyTransfer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "MoneyTransfer binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.MoneyTransfer.
		UpdateOneID(int(id)).
		SetMoneytransferType(obj.MoneytransferType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewMoneyTransferController creates and registers handles for the MoneyTransferController
func NewMoneyTransferController(router gin.IRouter, client *ent.Client) *MoneyTransferController {
	mc := &MoneyTransferController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMoneyTransferController registers routes to the main engine
func (ctl *MoneyTransferController) register() {
	moneyTransfers := ctl.router.Group("/moneyTransfers")
	moneyTransfers.GET("", ctl.ListMoneyTransfer)
	// CRUD
	moneyTransfers.POST("", ctl.CreateMoneyTransfer)
	moneyTransfers.GET(":id", ctl.GetMoneyTransfer)
	moneyTransfers.PUT(":id", ctl.UpdateMoneyTransfer)
	moneyTransfers.DELETE(":id", ctl.DeleteMoneyTransfer)
}
