package manager

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	hashedPass := []byte{}
	if body.Password != "" {
		hashedPass, err = bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
	}
	body.Password = string(hashedPass)

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
