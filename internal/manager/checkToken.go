package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) CheckToken(tokenReq uuid.UUID) (response *models.CheckTokenResponse, err error) {
	accessToken, err := m.App.DB.GetAccessToken(tokenReq)
	if err != nil && err == models.AccessTokenExpiredError {
		return nil, err
	} else if err != nil && err == models.AccessTokenNotFoundError {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	_, err = m.App.DB.GetUserByID(accessToken.UserID)
	if err != nil {
		return nil, err
	}
	return &models.CheckTokenResponse{
		Token:     accessToken.Token,
		UserID:    accessToken.UserID,
	}, nil
}
