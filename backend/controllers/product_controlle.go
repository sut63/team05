package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/gender"
	"github.com/sut63/team05/ent/groupofage"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

//ProductController struct
type ProductController struct {
	client *ent.Client
	router gin.IRouter
}

//Product struct
type Product struct {
	ProductName          string
	ProductPrice         int
	ProductTime          int
	ProductPaymentOfYear float64

	GenderID   int
	GroupOfAge int
	Officer    int
}

// CreateProduct handles POST requests for adding product entities
// @Summary Create product
// @Description Create product
// @ID create-product
// @Accept   json
// @Produce  json
// @Param product body Product true "Product entity"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products [post]
func (ctl *ProductController) CreateProduct(c *gin.Context) {
	obj := Product{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "product binding failed",
		})
		return
	}

	gd, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.GenderID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	goa, err := ctl.client.GroupOfAge.
		Query().
		Where(groupofage.IDEQ(int(obj.GroupOfAge))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "groupofage not found",
		})
		return
	}

	ofc, err := ctl.client.Officer.
		Query().
		Where(officer.IDEQ(int(obj.Officer))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "officer not found",
		})
		return
	}

	p, err := ctl.client.Product.
		Create().
		SetProductName(obj.ProductName).
		SetProductPrice(obj.ProductPrice).
		SetProductTime(obj.ProductTime).
		SetProductPaymentOfYear(obj.ProductPaymentOfYear).
		SetProductGender(gd).
		SetProductGroupage(goa).
		SetProductOfficer(ofc).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// DeleteProduct handles DELETE requests to delete a product entity
// @Summary Delete a product entity by ID
// @Description get product by ID
// @ID delete-product
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/{id} [delete]
func (ctl *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Product.
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

// GetProduct handles GET requests to retrieve a product entity
// @Summary Get a product entity by ID
// @Description get product by ID
// @ID get-product
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/{id} [get]
func (ctl *ProductController) GetProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, p)
}

// UpdateProduct handles PUT requests to update a Product entity
// @Summary Update a product entity by ID
// @Description update product by ID
// @ID update-product
// @Accept   json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body ent.Product true "Product entity"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/{id} [put]
func (ctl *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Product{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "product binding failed",
		})
		return
	}
	obj.ID = int(id)
	p, err := ctl.client.Product.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, p)
}

// ListProduct handles request to get a list of product entities
// @Summary List product entities
// @Description list product entities
// @ID list-product
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products [get]
func (ctl *ProductController) ListProduct(c *gin.Context) {
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

	products, err := ctl.client.Product.
		Query().
		WithProductGender().
		WithProductGroupage().
		WithProductOfficer().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

// NewProductController creates and registers handles for the product controller
func NewProductController(router gin.IRouter, client *ent.Client) *ProductController {
	pc := &ProductController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *ProductController) register() {
	products := ctl.router.Group("/products")
	products.GET("", ctl.ListProduct)
	// CRUD
	products.POST("", ctl.CreateProduct) //
	products.GET(":id", ctl.GetProduct)
	products.PUT(":id", ctl.UpdateProduct)
	products.DELETE(":id", ctl.DeleteProduct) //
}
