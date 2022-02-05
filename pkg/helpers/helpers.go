package helpers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"math/rand"
	"mime/multipart"
	"strconv"
	"strings"
)

// GetFileExt - returns file extension or "jpeg" if not found
func GetFileExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	nameLength := len(filename)
	if dotIndex == -1 {
		return "jpeg"
	}
	return filename[dotIndex+1:nameLength]
}

func ParsePaginationFromQuery(ctx *gin.Context) (int, int) {
	page := 1
	size := 15

	qPage := ctx.Query("page")
	if len(qPage) > 0 {
		page, _ = strconv.Atoi(qPage)
	}

	qSize := ctx.Query("size")
	if len(qSize) > 0 {
		size, _ = strconv.Atoi(qSize)
	}

	return page, size
}

func ParseTypeFromQuery(ctx *gin.Context) string {
	fileType := ""
	qType := ctx.Query("type")
	if len(qType) > 0 {
		fileType = qType
	}

	return fileType
}

func ParseForGalleryQuery(ctx *gin.Context) string {
	forGallery := "false"
	qType := ctx.Query("forGallery")
	if len(qType) > 0 {
		forGallery = qType
	}

	return forGallery
}

func ParseSkipSaveFromQuery(ctx *gin.Context) bool {
	skipSave, _ := strconv.ParseBool(ctx.Query("skipSave"))
	return skipSave
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func UploadToMinio(
	minioClient *minio.Client,
	bucketName string,
	objectName string,
	file *multipart.FileHeader,
	contentType string) (uploadUrl string, err error) {
	fmt.Println("6.1")

	var src multipart.File
	src, err = file.Open()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("6.2")
	defer src.Close()

	fmt.Println("6.3")
	fmt.Println("bucketName: ", bucketName, "objectName: ", objectName, "src: ", src, "contentType: ", contentType)

	var uploadInfo minio.UploadInfo
	uploadInfo, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("6.4")

	uploadUrl = uploadInfo.Key
	return
}

