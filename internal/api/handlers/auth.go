package handlers

import (
	"github.com/akhmettolegen/texert/pkg/models"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Tags OAuth
// @Summary SignUp endpoint
// @ID sign-up
// @Accept json
// @Param SignUpRequest	body models.SignUpRequest true "SignUp Request"
// @Produce  json
// @Success 200 {object} models.SignUpResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /auth/sign-up [post]
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

// SignIn godoc
// @Tags OAuth
// @Summary SignIn endpoint
// @ID sign-in
// @Accept json
// @Param SignInRequest	body models.SignInRequest true "SignIn Request"
// @Produce  json
// @Success 200 {object} models.SignInResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /auth/sign-in [post]
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
