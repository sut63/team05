package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/groupofage"
)

// GroupOfAgeController struct
type GroupOfAgeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateGroupOfAge handles POST requests for adding groupofage entities
// @Summary Create groupofage
// @Description Create groupofage
// @ID create-groupofage
// @Accept   json
// @Produce  json
// @Param groupofage body ent.GroupOfAge true "GroupOfAge entity"
// @Success 200 {object} ent.GroupOfAge
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /groupofages [post]
func (ctl *GroupOfAgeController) CreateGroupOfAge(c *gin.Context) {
	obj := ent.GroupOfAge{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "groupofage binding failed",
		})
		return
	}

	goa, err := ctl.client.GroupOfAge.
		Create().
		SetGroupOfAgeName(obj.GroupOfAgeName).
		SetGroupOfAgeAge(obj.GroupOfAgeAge).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving groupofage failed",
		})
		return
	}

	c.JSON(200, goa)
}

// GetGroupOfAge handles GET requests to retrieve a groupofage entity
// @Summary Get a groupofage entity by ID
// @Description get groupofage by ID
// @ID get-groupofage
// @Produce  json
// @Param id path int true "GroupOfAge ID"
// @Success 200 {object} ent.GroupOfAge
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /groupofages/{id} [get]
func (ctl *GroupOfAgeController) GetGroupOfAge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	goa, err := ctl.client.GroupOfAge.
		Query().
		Where(groupofage.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, goa)
}

// ListGroupOfAge handles request to get a list of groupofage entities
// @Summary List groupofage entities
// @Description list groupofage entities
// @ID list-groupofage
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.GroupOfAge
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /groupofages [get]
func (ctl *GroupOfAgeController) ListGroupOfAge(c *gin.Context) {
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

	groupofages, err := ctl.client.GroupOfAge.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, groupofages)
}

// DeleteGroupOfAge handles DELETE requests to delete a groupofage entity
// @Summary Delete a groupofage entity by ID
// @Description get groupofage by ID
// @ID delete-groupofage
// @Produce  json
// @Param id path int true "GroupOfAge ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /groupofages/{id} [delete]
func (ctl *GroupOfAgeController) DeleteGroupOfAge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.GroupOfAge.
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

// UpdateGroupOfAge handles PUT requests to update a groupofage entity
// @Summary Update a groupofage entity by ID
// @Description update groupofage by ID
// @ID update-groupofage
// @Accept   json
// @Produce  json
// @Param id path int true "GroupOfAge ID"
// @Param groupofage body ent.GroupOfAge true "GroupOfAge entity"
// @Success 200 {object} ent.GroupOfAge
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /groupofages/{id} [put]
func (ctl *GroupOfAgeController) UpdateGroupOfAge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.GroupOfAge{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "groupofage binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	goa, err := ctl.client.GroupOfAge.
		UpdateOneID(int(id)).
		SetGroupOfAgeName(obj.GroupOfAgeName).
		SetGroupOfAgeAge(obj.GroupOfAgeAge).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update groupofage failed",
		})
		return
	}

	c.JSON(200, goa)
}

// NewGroupOfAgeController creates and registers handles for the oder controller
func NewGroupOfAgeController(router gin.IRouter, client *ent.Client) *GroupOfAgeController {
	goac := &GroupOfAgeController{
		client: client,
		router: router,
	}

	goac.register()

	return goac

}

func (ctl *GroupOfAgeController) register() {
	groupofages := ctl.router.Group("/groupofages")

	groupofages.POST("", ctl.CreateGroupOfAge)
	groupofages.GET(":id", ctl.GetGroupOfAge)
	groupofages.GET("", ctl.ListGroupOfAge)
	groupofages.DELETE(":id", ctl.DeleteGroupOfAge)
	groupofages.PUT(":id", ctl.UpdateGroupOfAge)

}
