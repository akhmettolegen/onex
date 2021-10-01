package manager

import "github.com/akhmettolegen/onex/pkg/models"

func (m *Manager) SignUp(signUpReq *models.SignUpRequest) (response *models.SignUpResponse, err error){

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