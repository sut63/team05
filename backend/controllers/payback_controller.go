package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

// PaybackController defines the struct for the payback controller
type PaybackController struct {
	client *ent.Client
	router gin.IRouter
}

// Payback defines the struct for the payback
type Payback struct {
	ProductID     int
	MemberID      int
	BankID        int
	OfficerID     int
	Accountnumber string
	Transfertime  string
}

// CreatePayback handles POST requests for adding payback entities
// @Summary Create payback
// @Description Create payback
// @ID create-payback
// @Accept   json
// @Produce  json
// @Param payback body Payback true "Payback entity"
// @Success 200 {object} ent.Payback
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paybacks [post]
func (ctl *PaybackController) CreatePayback(c *gin.Context) {
	obj := Payback{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payback binding failed",
		})
		return
	}

	p, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.ProductID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "product not found",
		})
		return
	}

	m, err := ctl.client.Member.
		Query().
		Where(member.IDEQ(int(obj.MemberID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "member diagnostic  not found",
		})
		return
	}

	b, err := ctl.client.Bank.
		Query().
		Where(bank.IDEQ(int(obj.BankID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statuscar diagnostic  not found",
		})
		return
	}

	of, err := ctl.client.Officer.
		Query().
		Where(officer.IDEQ(int(obj.OfficerID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "officer not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Transfertime)

	in, err := ctl.client.Payback.
		Create().
		SetProduct(p).
		SetMember(m).
		SetBank(b).
		SetOfficer(of).
		SetAccountnumber(obj.Accountnumber).
		SetTransfertime(time).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, in)
}

// ListPayback handles request to get a list of payback entities
// @Summary List payback entities
// @Description list payback entities
// @ID list-payback
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payback
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paybacks [get]
func (ctl *PaybackController) ListPayback(c *gin.Context) {
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

	paybacks, err := ctl.client.Payback.
		Query().
		WithProduct().
		WithMember().
		WithBank().
		WithOfficer().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paybacks)
}

// DeletePayback handles DELETE requests to delete a payback entity
// @Summary Delete a payback entity by ID
// @Description get payback by ID
// @ID delete-payback
// @Produce  json
// @Param id path int true "Payback ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paybacks/{id} [delete]
func (ctl *PaybackController) DeletePayback(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Payback.
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

// NewPaybackController creates and registers handles for the payback controller
func NewPaybackController(router gin.IRouter, client *ent.Client) *PaybackController {
	ic := &PaybackController{
		client: client,
		router: router,
	}
	ic.register()
	return ic
}

// InitPaybackController registers routes to the main engine
func (ctl *PaybackController) register() {
	paybacks := ctl.router.Group("/paybacks")

	paybacks.GET("", ctl.ListPayback)
	paybacks.POST("", ctl.CreatePayback)
	paybacks.DELETE(":id", ctl.DeletePayback)
}
