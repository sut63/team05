package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/bank"
)

// BankController defines the struct for the Bank controller
type BankController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBank handles POST requests for adding Bank entities
// @Summary Create Bank
// @Description Create Bank
// @ID create-Bank
// @Accept   json
// @Produce  json
// @Param Bank body ent.Bank true "Bank entity"
// @Success 200 {object} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks [post]
func (ctl *BankController) CreateBank(c *gin.Context) {
	obj := ent.Bank{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bank binding failed",
		})
		return
	}

	m, err := ctl.client.Bank.
		Create().
		SetBankType(obj.BankType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetBank handles GET requests to retrieve a Bank entity
// @Summary Get a Bank entity by ID
// @Description get Bank by ID
// @ID get-Bank
// @Produce  json
// @Param id path int true "Bank ID"
// @Success 200 {object} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks/{id} [get]
func (ctl *BankController) GetBank(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Bank.
		Query().
		Where(bank.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListBank handles request to get a list of Bank entities
// @Summary List Bank entities
// @Description list Bank entities
// @ID list-Bank
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks [get]
func (ctl *BankController) ListBank(c *gin.Context) {
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

	banks, err := ctl.client.Bank.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, banks)
}

// DeleteBank handles DELETE requests to delete a Bank entity
// @Summary Delete a Bank entity by ID
// @Description get Bank by ID
// @ID delete-Bank
// @Produce  json
// @Param id path int true "Bank ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks/{id} [delete]
func (ctl *BankController) DeleteBank(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bank.
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

// UpdateBank handles PUT requests to update a Bank entity
// @Summary Update a Bank entity by ID
// @Description update Bank by ID
// @ID update-Bank
// @Accept   json
// @Produce  json
// @Param id path int true "Bank ID"
// @Param Bank body ent.Bank true "Bank entity"
// @Success 200 {object} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks/{id} [put]
func (ctl *BankController) UpdateBank(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Bank{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bank binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.Bank.
		UpdateOneID(int(id)).
		SetBankType(obj.BankType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewBankController creates and registers handles for the BankController
func NewBankController(router gin.IRouter, client *ent.Client) *BankController {
	mc := &BankController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitBankController registers routes to the main engine
func (ctl *BankController) register() {
	banks := ctl.router.Group("/banks")
	banks.GET("", ctl.ListBank)
	// CRUD
	banks.POST("", ctl.CreateBank)
	banks.GET(":id", ctl.GetBank)
	banks.PUT(":id", ctl.UpdateBank)
	banks.DELETE(":id", ctl.DeleteBank)
}
