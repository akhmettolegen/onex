package models

import uuid "github.com/satori/go.uuid"

type UserInfo struct {
	UserID     uuid.UUID
}

var UserInfoKey = "user-info"