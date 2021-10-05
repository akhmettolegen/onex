package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) SignUp(signUpReq *models.SignUpRequest) (response *models.SignUpResponse, err error){

	hashedPass := []byte{}
	if signUpReq.Password != "" {
		hashedPass, err = bcrypt.GenerateFromPassword([]byte(signUpReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
	}
	signUpReq.Password = string(hashedPass)

	createUser := models.User{
		Name:     signUpReq.Name,
		Phone:    signUpReq.Phone,
		Password: signUpReq.Password,
		Channel:  "MOBILE",
	}

	err = m.App.DB.CreateUser(&createUser)
	if err != nil {
		return
	}

	createToken := models.AccessToken{
		UserID:    createUser.ID,
		TTL:       100,
	}

	err = m.App.DB.CreateToken(&createToken)
	if err != nil {
		return
	}

	response = &models.SignUpResponse{
		Token:  createToken.Token,
		UserID: createToken.UserID,
		TTL:    createToken.TTL,
	}

	return
}

