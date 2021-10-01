package models

type Channel string
const (
	ChannelWeb Channel = "WEB"
	ChannelMobile Channel = "MOBILE"
)

type User struct {
	Base
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Channel Channel `json:"channel"`
}

type UsersListResponse struct {
	Data  []User `json:"data"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
}

type UserByIDResponse struct {
	Data    User 	      `json:"data"`
}

type UserCreateRequest struct {
	Name     string                `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Channel string `json:"channel"`
}
