package models

import uuid "github.com/satori/go.uuid"

type File struct {
	URL				string		`json:"url"`
	UploadedUserID	uuid.UUID	`json:"uploadedUserId"`
}

type UploadFileResponse struct {
	Data File `json:"data"`
}
