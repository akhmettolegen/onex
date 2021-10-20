package manager

import (
	"github.com/akhmettolegen/onex/pkg/helpers"
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"time"
)

func (m *Manager) Upload(file *multipart.FileHeader) (response *models.UploadFileResponse, err error) {

	fileExt := helpers.GetFileExt(file.Filename)
	randomFileName := helpers.RandStringRunes(16)
	objectName := time.Now().Format("20060102") + "_" + randomFileName + "." + fileExt

	url, err := helpers.UploadToMinio(m.App.MinIOClient, m.App.Config.Minio.Bucket, objectName, file, fileExt)

	filepath := "http://" + m.App.Config.Minio.Host + "/" + m.App.Config.Minio.Bucket + "/" + url

	if err != nil {
		return
	}
	response = &models.UploadFileResponse{
		Data: models.File{
			URL:            filepath,
			UploadedUserID: uuid.NewV4(),
		}}

	return
}
