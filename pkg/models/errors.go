package models

import "errors"

var (
	UserNotFoundError = errors.New("user not found")
	UserAlreadyExistsError = errors.New("user already exists")
	AccessTokenNotFoundError = errors.New("access token not found")
	UnauthorizedError = errors.New("unauthorized")
	AccessTokenExpiredError = errors.New("access token expired")
	InvalidPhoneNumberError =  errors.New("invalid phone number")
)

