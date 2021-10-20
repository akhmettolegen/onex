package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Upload(ctx *gin.Context) {

	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	response, err := h.Manager.Upload(header)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
