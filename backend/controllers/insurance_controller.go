package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

// InsuranceController defines the struct for the insurance controller
type InsuranceController struct {
	client *ent.Client
	router gin.IRouter
}

// Insurance defines the struct for the insurance
type Insurance struct {
	ProductID             int
	MemberID              int
	HospitalID            int
	OfficerID             int
	InsuranceAddress      string
	InsuranceInsurer      string
	InsuranceTimeBuy      string
	InsuranceTimeFirstpay string
}

// CreateInsurance handles POST requests for adding insurance entities
// @Summary Create insurance
// @Description Create insurance
// @ID create-insurance
// @Accept   json
// @Produce  json
// @Param insurance body Insurance true "Insurance entity"
// @Success 200 {object} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances [post]
func (ctl *InsuranceController) CreateInsurance(c *gin.Context) {
	obj := Insurance{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "insurance binding failed",
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

	h, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(obj.HospitalID))).
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
	timeb, err := time.Parse(time.RFC3339, obj.InsuranceTimeBuy)
	timep, err := time.Parse(time.RFC3339, obj.InsuranceTimeFirstpay)

	in, err := ctl.client.Insurance.
		Create().
		SetProduct(p).
		SetMember(m).
		SetHospital(h).
		SetOfficer(of).
		SetInsuranceAddress(obj.InsuranceAddress).
		SetInsuranceInsurer(obj.InsuranceInsurer).
		SetInsuranceTimeBuy(timeb).
		SetInsuranceTimeFirstpay(timep).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, in)
}

// ListInsurance handles request to get a list of insurance entities
// @Summary List insurance entities
// @Description list insurance entities
// @ID list-insurance
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances [get]
func (ctl *InsuranceController) ListInsurance(c *gin.Context) {
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

	insurances, err := ctl.client.Insurance.
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

	c.JSON(200, insurances)
}

// DeleteInsurance handles DELETE requests to delete a insurance entity
// @Summary Delete a insurance entity by ID
// @Description get insurance by ID
// @ID delete-insurance
// @Produce  json
// @Param id path int true "Insurance ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances/{id} [delete]
func (ctl *InsuranceController) DeleteInsurance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Insurance.
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

// NewInsuranceController creates and registers handles for the insurance controller
func NewInsuranceController(router gin.IRouter, client *ent.Client) *InsuranceController {
	ic := &InsuranceController{
		client: client,
		router: router,
	}
	ic.register()
	return ic
}

// InitInsuranceController registers routes to the main engine
func (ctl *InsuranceController) register() {
	insurances := ctl.router.Group("/insurances")

	insurances.GET("", ctl.ListInsurance)
	insurances.POST("", ctl.CreateInsurance)
	insurances.DELETE(":id", ctl.DeleteInsurance)
}
