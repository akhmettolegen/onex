package manager

import (
	"fmt"
	"github.com/akhmettolegen/texert/pkg/helpers"
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"time"
)

func (m *Manager) Upload(file *multipart.FileHeader) (response *models.UploadFileResponse, err error) {
	fmt.Println("3")
	fileExt := helpers.GetFileExt(file.Filename)
	fmt.Println("4")
	randomFileName := helpers.RandStringRunes(16)
	fmt.Println("5")
	objectName := time.Now().Format("20060102") + "_" + randomFileName + "." + fileExt
	fmt.Println("6")

	url, err := helpers.UploadToMinio(m.App.MinIOClient, m.App.Config.Minio.Bucket, objectName, file, fileExt)
	if err != nil {
		return
	}
	fmt.Println("7")

	filepath := "http://" + m.App.Config.Minio.Host + "/" + m.App.Config.Minio.Bucket + "/" + url
	fmt.Println("8")

	response = &models.UploadFileResponse{
		Data: models.File{
			URL:            filepath,
			UploadedUserID: uuid.NewV4(),
		}}

	return
}
