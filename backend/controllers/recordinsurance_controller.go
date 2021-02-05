package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/amountpaid"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/recordinsurance"
)

// RecordinsuranceController defines the struct for the recordinsurance controller
type RecordinsuranceController struct {
	client *ent.Client
	router gin.IRouter
}

// Recordinsurance defines the struct for the recordinsurance
type Recordinsurance struct {
	ProductID              int
	MemberID               int
	HospitalID             int
	OfficerID              int
	AmountpaidID           int
	NumberOfDaysOfTreat    int
	RecordinsuranceContact string
	RecordinsuranceTime    string
	RecordinsuranceAddress string
}

// CreateRecordinsurance handles POST requests for adding recordinsurance entities
// @Summary Create recordinsurance
// @Description Create recordinsurance
// @ID create-recordinsurance
// @Accept   json
// @Produce  json
// @Param Recordinsurance body Recordinsurance true "Recordinsurance entity"
// @Success 200 {object} ent.Recordinsurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /recordinsurances [post]
func (ctl *RecordinsuranceController) CreateRecordinsurance(c *gin.Context) {
	obj := Recordinsurance{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "recordinsurance failed",
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
			"error": "member diagnostic not found",
		})
		return
	}

	h, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(obj.HospitalID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "hospital diagnostic  not found",
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

	ap, err := ctl.client.Amountpaid.
		Query().
		Where(amountpaid.IDEQ(int(obj.AmountpaidID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "amountpaid not found",
		})
		return
	}
	timet := time.Now().Local()

	rin, err := ctl.client.Recordinsurance.
		Create().
		SetProduct(p).
		SetMember(m).
		SetHospital(h).
		SetOfficer(of).
		SetAmountpaid(ap).
		SetNumberOfDaysOfTreat(obj.NumberOfDaysOfTreat).
		SetRecordinsuranceContact(obj.RecordinsuranceContact).
		SetRecordinsuranceTime(timet).
		SetRecordinsuranceAddress(obj.RecordinsuranceAddress).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   rin,
	})
}

// ListRecordinsurance handles request to get a list of recordinsurance entities
// @Summary List recordinsurance entities
// @Description list recordinsurance entities
// @ID list-recordinsurance
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Recordinsurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /recordinsurances [get]
func (ctl *RecordinsuranceController) ListRecordinsurance(c *gin.Context) {
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

	recordinsurances, err := ctl.client.Recordinsurance.
		Query().
		WithProduct().
		WithMember().
		WithHospital().
		WithOfficer().
		WithAmountpaid().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, recordinsurances)
}

// GetRecordinsurance handles GET requests to retrieve a recordinsurance entity
// @Summary Get a recordinsurance entity by ID
// @Description get recordinsurance by ID
// @ID get-recordinsurance
// @Produce  json
// @Param id path int true "Recordinsurance ID"
// @Success 200 {object} ent.Recordinsurance
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /recordinsurances/{id} [get]
func (ctl *RecordinsuranceController) GetRecordinsurance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Recordinsurance.
		Query().
		Where(recordinsurance.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, p)
}

// DeleteRecordinsurance handles DELETE requests to delete a recordinsurance entity
// @Summary Delete a recordinsurance entity by ID
// @Description get recordinsurance by ID
// @ID delete-recordinsurance
// @Produce  json
// @Param id path int true "Recordinsurance ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /recordinsurances/{id} [delete]
func (ctl *RecordinsuranceController) DeleteRecordinsurance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Recordinsurance.
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

// NewRecordinsuranceController creates and registers handles for the recordinsurance controller
func NewRecordinsuranceController(router gin.IRouter, client *ent.Client) *RecordinsuranceController {
	ric := &RecordinsuranceController{
		client: client,
		router: router,
	}
	ric.register()
	return ric
}

// InitRecordinsuranceController registers routes to the main engine
func (ctl *RecordinsuranceController) register() {
	recordinsurances := ctl.router.Group("/recordinsurances")

	recordinsurances.GET("", ctl.ListRecordinsurance)
	recordinsurances.GET(":id", ctl.GetRecordinsurance)
	recordinsurances.POST("", ctl.CreateRecordinsurance)
	recordinsurances.DELETE(":id", ctl.DeleteRecordinsurance)
}
