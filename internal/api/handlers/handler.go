package handlers

import (
	"errors"
	"fmt"
	"github.com/akhmettolegen/onex/internal/manager"
	"github.com/akhmettolegen/onex/pkg/application"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// Handler model
type Handler struct {
	App application.Application
	Manager *manager.Manager
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	manager, _ := manager.Get(&app)

	return &Handler{
		App:     app,
		Manager: manager,
	}
}

func (h *Handler) CheckChannelToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("err")
		// 401 400 or other status code
		c.AbortWithStatusJSON(400, gin.H{"message": errors.New("empty authorization header").Error()})
		return
	}
	splittedAuthHeader := strings.Split(authHeader, " ")
	if len(splittedAuthHeader) != 2 || strings.ToLower(splittedAuthHeader[0]) != "bearer" {
		// 401 400 or other status code;
		c.AbortWithStatusJSON(400, gin.H{"message": errors.New("invalid token").Error()})
		return
	}

	token, err := uuid.FromString(splittedAuthHeader[1])
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": errors.New("provided token is not uuid").Error()})
		return
	}

	tokenData, err := h.Manager.CheckToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": err.Error()})
		return
	}

	c.Set("token", token)
	c.Set("user_id", tokenData.UserID.String())
	c.Next()
}