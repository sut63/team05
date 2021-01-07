package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/category"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

// inquiryController defines the struct for the inquiry controller
type InquiryController struct {
	client *ent.Client
	router gin.IRouter
}

// inquiry defines the struct for the inquiry
type inquiry struct {
	ProductID           int
	MemberID            int
	CategoryID          int
	OfficerID           int
	InquiryMessages     string
	InquiryTimeMessages string
}

// CreateInquiry handles POST requests for adding inquiry entities
// @Summary Create inquiry
// @Description Create inquiry
// @ID create-inquiry
// @Accept   json
// @Produce  json
// @Param inquiry body inquiry true "Inquiry entity"
// @Success 200 {object} ent.Inquiry
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inquirys [post]
func (ctl *InquiryController) Createinquiry(c *gin.Context) {
	obj := Inquiry{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "inquiry binding failed",
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

	cg, err := ctl.client.Hospital.
		Query().
		Where(category.IDEQ(int(obj.CategoryID))).
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
	timeb, err := time.Parse(time.RFC3339, obj.inquiryTimeBuy)
	timep, err := time.Parse(time.RFC3339, obj.inquiryTimeFirstpay)

	in, err := ctl.client.Inquiry.
		Create().
		SetProduct(p).
		SetMember(m).
		SetCategory(cg).
		SetOfficer(of).
		SetI(obj.inquiryAddress).
		SetinquiryInsurer(obj.inquiryInsurer).
		SetinquiryTimeBuy(timeb).
		SetinquiryTimeFirstpay(timep).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, in)
}

// Listinquiry handles request to get a list of inquiry entities
// @Summary List inquiry entities
// @Description list inquiry entities
// @ID list-inquiry
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.inquiry
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inquirys [get]
func (ctl *inquiryController) Listinquiry(c *gin.Context) {
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

	inquirys, err := ctl.client.inquiry.
		Query().
		WithProduct().
		WithMember().
		WithHospital().
		WithOfficer().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, inquirys)
}

// Deleteinquiry handles DELETE requests to delete a inquiry entity
// @Summary Delete a inquiry entity by ID
// @Description get inquiry by ID
// @ID delete-inquiry
// @Produce  json
// @Param id path int true "inquiry ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inquirys/{id} [delete]
func (ctl *inquiryController) Deleteinquiry(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.inquiry.
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

// NewinquiryController creates and registers handles for the inquiry controller
func NewinquiryController(router gin.IRouter, client *ent.Client) *inquiryController {
	ic := &inquiryController{
		client: client,
		router: router,
	}
	ic.register()
	return ic
}

// InitinquiryController registers routes to the main engine
func (ctl *inquiryController) register() {
	inquirys := ctl.router.Group("/inquirys")

	inquirys.GET("", ctl.Listinquiry)
	inquirys.POST("", ctl.Createinquiry)
	inquirys.DELETE(":id", ctl.Deleteinquiry)
}
