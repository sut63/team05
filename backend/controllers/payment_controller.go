package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/payment"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/moneytransfer"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/insurance"
)

// PaymentController defines the struct for the payment controller
type PaymentController struct {
	client *ent.Client
	router gin.IRouter
}

// Payment defines the struct for the payment
type Payment struct {
	ProductID             int
	MemberID              int
    MoneytransferID       int
	InsuranceID           int
	BankID                int
	AccountName         string
	AccountNumber       string
	TransferTime        string

}

// CreatePayment handles POST requests for adding payment entities
// @Summary Create payment
// @Description Create payment
// @ID create-payment
// @Accept   json
// @Produce  json
// @Param payment body Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [post]
func (ctl *PaymentController) CreatePayment(c *gin.Context) {
	obj := Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment binding failed",
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

	mn, err := ctl.client.MoneyTransfer.
		Query().
		Where(moneytransfer.IDEQ(int(obj.MoneytransferID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "moneytransfer diagnostic  not found",
		})
		return
	}

	in, err := ctl.client.Insurance.
		Query().
		Where(insurance.IDEQ(int(obj.InsuranceID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "insurance not found",
		})
		return
	}
	
	b, err := ctl.client.Bank.
		Query().
		Where(bank.IDEQ(int(obj.BankID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bank not found",
		})
		return
	}



	time, err := time.Parse(time.RFC3339, obj.TransferTime)
	

	py, err := ctl.client.Payment.
		Create().
		SetProduct(p).
		SetMember(m).
		SetMoneyTransfer(mn).
		SetInsurance(in).
		SetBank(b).
		SetAccountName(obj.AccountName).
		SetAccountNumber(obj.AccountNumber).
		SetTransferTime(time).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, py)
}

// ListPayment handles request to get a list of payment entities
// @Summary List payment entities
// @Description list payment entities
// @ID list-payment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [get]
func (ctl *PaymentController) ListPayment(c *gin.Context) {
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

	payments, err := ctl.client.Payment.
		Query().
		WithProduct().
		WithMember().
		WithMoneyTransfer().
		WithInsurance().
		WithBank().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, payments)
}

// GetPayment handles GET requests to retrieve a payment entity
// @Summary Get a payment entity by ID
// @Description get payment by ID
// @ID get-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments/{id} [get]
func (ctl *PaymentController) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Payment.
		Query().
		WithProduct().
		Where(payment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// DeletePayment handles DELETE requests to delete a payment entity
// @Summary Delete a payment entity by ID
// @Description get payment by ID
// @ID delete-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments/{id} [delete]
func (ctl *PaymentController) DeletePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Payment.
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

// NewPaymentController creates and registers handles for the payment controller
func NewPaymentController(router gin.IRouter, client *ent.Client) *PaymentController {
	ic := &PaymentController{
		client: client,
		router: router,
	}
	ic.register()
	return ic
}

// InitPaymentController registers routes to the main engine
func (ctl *PaymentController) register() {
	payments := ctl.router.Group("/payments")

	payments.GET("", ctl.ListPayment)
	payments.GET(":id", ctl.GetPayment)
	payments.POST("", ctl.CreatePayment)
	payments.DELETE(":id", ctl.DeletePayment)
}