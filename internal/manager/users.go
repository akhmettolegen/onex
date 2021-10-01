package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) GetUsers(page, size int) (response *models.UsersListResponse, err error) {
	users, total, err := m.App.DB.GetUsers(page, size)
	if err != nil {
		return nil, err
	}

	response =  &models.UsersListResponse{
		Data:  users,
		Page:  page,
		Size:  size,
		Total: total,
	}
	return
}

func (m *Manager) GetUserByID(userID uuid.UUID) (response *models.UserByIDResponse, err error) {
	user, err := m.App.DB.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	response =  &models.UserByIDResponse{
		Data: user,
	}
	return
}

func (m *Manager) CreateUser(body models.UserCreateRequest) (response *models.UserByIDResponse, err error) {
	createUser := models.User{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
		Channel:  models.Channel(body.Channel),
	}

	err = m.App.DB.CreateUser(&createUser)
	if err != nil {
		return
	}

	response = &models.UserByIDResponse{Data: createUser}
	return
}

func (m *Manager) DeleteUser(userID uuid.UUID) (err error) {
	return m.App.DB.DeleteUser(userID)
}
