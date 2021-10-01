package handlers

import (
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var signUpReq models.SignUpRequest
	if err := ctx.ShouldBindJSON(&signUpReq); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	response, err := h.Manager.SignUp(&signUpReq)
	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.JSON(200, response)
}