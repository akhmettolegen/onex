package handlers

import (
	"github.com/akhmettolegen/onex/pkg/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) Upload(ctx *gin.Context) {

	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	fileExt := helpers.GetFileExt(header.Filename)
	randomFileName := helpers.RandStringRunes(16)
	objectName := time.Now().Format("20060102") + "_" + randomFileName + "." + fileExt

	url, err := helpers.UploadToMinio(h.App.MinIOClient, h.App.Config.Minio.Bucket, objectName, header, fileExt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, url)
}
