package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/amountpaid"
)

// AmountpaidController defines the struct for the Amountpaid controller
type AmountpaidController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateAmountpaid handles POST requests for adding Amountpaid entities
// @Summary Create Amountpaid
// @Description Create Amountpaid
// @ID create-Amountpaid
// @Accept   json
// @Produce  json
// @Param Amountpaid body ent.Amountpaid true "Amountpaid entity"
// @Success 200 {object} ent.Amountpaid
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /amountpaids [post]
func (ctl *AmountpaidController) CreateAmountpaid(c *gin.Context) {
	obj := ent.Amountpaid{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "amountpaid failed",
		})
		return
	}

	m, err := ctl.client.Amountpaid.
		Create().
		SetAmountpaidMoney(obj.AmountpaidMoney).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetAmountpaid handles GET requests to retrieve a Amountpaid entity
// @Summary Get a Amountpaid entity by ID
// @Description get Amountpaid by ID
// @ID get-Amountpaid
// @Produce  json
// @Param id path int true "Amountpaid ID"
// @Success 200 {object} ent.Amountpaid
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /amountpaids/{id} [get]
func (ctl *AmountpaidController) GetAmountpaid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ap, err := ctl.client.Amountpaid.
		Query().
		Where(amountpaid.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ap)
}

// ListAmountpaid handles request to get a list of Amountpaid entities
// @Summary List Amountpaid entities
// @Description list Amountpaid entities
// @ID list-Amountpaid
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Amountpaid
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /amountpaids [get]
func (ctl *AmountpaidController) ListAmountpaid(c *gin.Context) {
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

	amountpaids, err := ctl.client.Amountpaid.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, amountpaids)
}

// DeleteAmountpaid handles DELETE requests to delete a Amountpaid entity
// @Summary Delete a Amountpaid entity by ID
// @Description get Amountpaid by ID
// @ID delete-Amountpaid
// @Produce  json
// @Param id path int true "Amountpaid ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /amountpaids/{id} [delete]
func (ctl *AmountpaidController) DeleteAmountpaid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Amountpaid.
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

// UpdateAmountpaid handles PUT requests to update a Amountpaid entity
// @Summary Update a Amountpaid entity by ID
// @Description update Amountpaid by ID
// @ID update-Amountpaid
// @Accept   json
// @Produce  json
// @Param id path int true "Amountpaid ID"
// @Param Amountpaid body ent.Amountpaid true "Amountpaid entity"
// @Success 200 {object} ent.Amountpaid
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /amountpaids/{id} [put]
func (ctl *AmountpaidController) UpdateAmountpaid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Amountpaid{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "amountpaid failed",
		})
		return
	}
	obj.ID = int(id)
	ap, err := ctl.client.Amountpaid.
		UpdateOneID(int(id)).
		SetAmountpaidMoney(obj.AmountpaidMoney).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ap)
}

// NewAmountpaidController creates and registers handles for the amountpaid Controller
func NewAmountpaidController(router gin.IRouter, client *ent.Client) *AmountpaidController {
	apc := &AmountpaidController{
		client: client,
		router: router,
	}
	apc.register()
	return apc
}

// InitAmountpaidController registers routes to the main engine
func (ctl *AmountpaidController) register() {
	amountpaids := ctl.router.Group("/amountpaids")
	amountpaids.GET("", ctl.ListAmountpaid)
	// CRUD
	amountpaids.POST("", ctl.CreateAmountpaid)
	amountpaids.GET(":id", ctl.GetAmountpaid)
	amountpaids.PUT(":id", ctl.UpdateAmountpaid)
	amountpaids.DELETE(":id", ctl.DeleteAmountpaid)
}
