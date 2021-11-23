package handlers

import (
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var signUpReq models.SignUpRequest
	if err := ctx.ShouldBindJSON(&signUpReq); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response, err := h.Manager.SignUp(&signUpReq)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var signInReq models.SignInRequest
	if err := ctx.ShouldBindJSON(&signInReq); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response, err := h.Manager.SignIn(&signInReq)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
