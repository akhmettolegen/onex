package manager

import (
	"errors"
	"github.com/akhmettolegen/onex/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

func (m *Manager) SignUp(signUpReq *models.SignUpRequest) (response *models.SignUpResponse, err error){

	if !phoneRegex.MatchString(signUpReq.Phone) {
		return nil, errors.New("invalid phone number")
	}

	// check if user exists
	_, err = m.App.DB.GetUserByPhone(signUpReq.Phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == nil {
		err = errors.New("user already exists")
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

	createToken := models.AccessToken{
		UserID:    createUser.ID,
		TTL:       m.App.Config.Token.TTL,
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

func (m *Manager) SignIn(signInReq *models.SignInRequest) (response *models.SignInResponse, err error){

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
