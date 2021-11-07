package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) GetToken(signInReq *models.SignInRequest) (response *models.SignInResponse, err error){

	user, err := m.App.DB.GetUserByPhone(signInReq.Phone)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInReq.Password)); err != nil {
		return nil, err
	}

	createToken := models.AccessToken{
		UserID:    user.ID,
		TTL:       m.App.Config.Token.TTL,
	}

	err = m.App.DB.CreateToken(&createToken)
	if err != nil {
		return
	}

	response = &models.SignInResponse{
		Token:  createToken.Token,
		UserID: createToken.UserID,
		TTL:    createToken.TTL,
	}

	return
}
