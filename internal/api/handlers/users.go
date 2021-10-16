package handlers

import (
	"github.com/akhmettolegen/onex/pkg/helpers"
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (h *Handler) GetUsers(ctx *gin.Context) {
	var query helpers.RequestQuery
	err := ctx.Bind(&query)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	page, size := helpers.ParsePagination(query)

	response, err := h.Manager.GetUsers(page, size)
	if err != nil {
		ctx.JSON(400, gin.H{"message":err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) GetUserByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response, err := h.Manager.GetUserByID(id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var createUserReq models.UserCreateRequest
	if err := ctx.ShouldBindJSON(&createUserReq); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	response, err := h.Manager.CreateUser(createUserReq)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)

}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = h.Manager.DeleteUser(id)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user deleted"})
}