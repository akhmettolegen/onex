package manager

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
	"time"
)

var phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

func (m *Manager) SignUp(signUpReq *models.SignUpRequest) (response *models.SignUpResponse, err error){

	if !phoneRegex.MatchString(signUpReq.Phone) {
		return nil, models.InvalidPhoneNumberError
	}

	// check if user exists
	_, err = m.App.DB.GetUserByPhone(signUpReq.Phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, models.UserNotFoundError
	} else if err == nil {
		err = models.UserAlreadyExistsError
		return
	}

	channel := models.ChannelMobile
	if signUpReq.Channel != nil {
		channel = *signUpReq.Channel
	}

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
		Channel:  channel,
	}

	err = m.App.DB.CreateUser(&createUser)
	if err != nil {
		return
	}

	token := models.AccessToken{
		UserID:    createUser.ID,
		TTL:       m.App.Config.Token.TTL,
	}

	err = m.App.DB.CreateToken(&token)
	if err != nil {
		return
	}

	expireAt := token.CreatedAt.Add(time.Second * time.Duration(token.TTL))
	response = &models.SignUpResponse{
		Token:  token.Token,
		UserID: token.UserID,
		TTL:    token.TTL,
		ExpiresAt: expireAt,
	}

	return
}

func (m *Manager) SignIn(signInReq *models.SignInRequest) (response *models.SignInResponse, err error){

	user, err := m.App.DB.GetUserByPhone(signInReq.Phone)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInReq.Password)); err != nil {
		return nil, err
	}

	token := models.AccessToken{
		UserID:    user.ID,
		TTL:       m.App.Config.Token.TTL,
	}

	err = m.App.DB.CreateToken(&token)
	if err != nil {
		return
	}

	expireAt := token.CreatedAt.Add(time.Second * time.Duration(token.TTL))
	response = &models.SignInResponse{
		Token:  token.Token,
		UserID: token.UserID,
		TTL:    token.TTL,
		ExpiresAt: expireAt,
	}

	return
}

func (m *Manager) CheckToken(tokenReq uuid.UUID) (response *models.CheckTokenResponse, err error) {
	accessToken, err := m.App.DB.GetAccessToken(tokenReq)
	if err != nil && err == models.AccessTokenExpiredError {
		return nil, err
	} else if err != nil && err == models.AccessTokenNotFoundError {
		return nil, models.UnauthorizedError
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