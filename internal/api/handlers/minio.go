package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Upload godoc
// @Tags CDN
// @Summary Upload file to CDN
// @ID upload
// @Security ApiKeyAuth
// @Accept x-www-form-urlencoded
// @Param file	formData file true "File"
// @Produce json
// @Success 200 {object} models.UploadFileResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /files/upload [post]
func (h *Handler) Upload(ctx *gin.Context) {
	fmt.Println("1")
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("2")
	response, err := h.Manager.Upload(header)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
