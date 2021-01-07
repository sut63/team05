package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team05/ent"
	"github.com/sut63/team05/ent/member"
)

// MemberController defines the struct for the member controller
type MemberController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMember handles POST requests for adding member entities
// @Summary Create member
// @Description Create member
// @ID create-member
// @Accept   json
// @Produce  json
// @Param Member body ent.Member true "Member entity"
// @Success 200 {object} ent.Member
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members [post]
func (ctl *MemberController) CreateMember(c *gin.Context) {
	obj := ent.Member{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "member binding failed",
		})
		return
	}

	m, err := ctl.client.Member.
		Create().
		SetMemberEmail(obj.MemberEmail).
		SetMemberName(obj.MemberName).
		SetMemberPassword(obj.MemberPassword).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetMember handles GET requests to retrieve a member entity
// @Summary Get a member entity by ID
// @Description get member by ID
// @ID get-member
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} ent.Member
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members/{id} [get]
func (ctl *MemberController) GetMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Member.
		Query().
		Where(member.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMember handles request to get a list of member entities
// @Summary List member entities
// @Description list member entities
// @ID list-member
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Member
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members [get]
func (ctl *MemberController) ListMember(c *gin.Context) {
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

	members, err := ctl.client.Member.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, members)
}

// DeleteMember handles DELETE requests to delete a member entity
// @Summary Delete a member entity by ID
// @Description get member by ID
// @ID delete-member
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members/{id} [delete]
func (ctl *MemberController) DeleteMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Member.
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

// UpdateMember handles PUT requests to update a member entity
// @Summary Update a member entity by ID
// @Description update member by ID
// @ID update-member
// @Accept   json
// @Produce  json
// @Param id path int true "Member ID"
// @Param member body ent.Member true "Member entity"
// @Success 200 {object} ent.Member
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members/{id} [put]
func (ctl *MemberController) UpdateMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Member{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "member binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.Member.
		UpdateOneID(int(id)).
		SetMemberEmail(obj.MemberEmail).
		SetMemberName(obj.MemberName).
		SetMemberPassword(obj.MemberPassword).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewMemberController creates and registers handles for the MemberController
func NewMemberController(router gin.IRouter, client *ent.Client) *MemberController {
	mc := &MemberController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMemberController registers routes to the main engine
func (ctl *MemberController) register() {
	members := ctl.router.Group("/members")
	members.GET("", ctl.ListMember)
	// CRUD
	members.POST("", ctl.CreateMember)
	members.GET(":id", ctl.GetMember)
	members.PUT(":id", ctl.UpdateMember)
	members.DELETE(":id", ctl.DeleteMember)
}
