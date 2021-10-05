package handlers

import (
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetToken(ctx *gin.Context) {
	var signInReq models.SignInRequest
	if err := ctx.ShouldBindJSON(&signInReq); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response, err := h.Manager.GetToken(&signInReq)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
