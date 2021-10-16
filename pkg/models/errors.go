package models

import "errors"

var (
	UserNotFoundError = errors.New("user not found")
	UserLogNotFoundError = errors.New("user log not found")
	ClientNotFoundError = errors.New("client not found")
	AccessTokenNotFoundError = errors.New("access token not found")
	UnauthorizedError = errors.New("unauthorized")
	RefreshTokenNotFoundError = errors.New("refresh token not found")
	AccessTokenExpiredError = errors.New("access token expired")
	RefreshTokenExpiredError = errors.New("refresh token expired")
	ChannelNotFoundError = errors.New("channel not found")
	VerificationTokenNotFoundError = errors.New("verification token not found")
	VerificationTokenExpiredError = errors.New("verification token expired")
	SessionNotFoundError = errors.New("session not found")
)

