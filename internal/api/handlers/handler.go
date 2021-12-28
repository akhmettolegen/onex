package handlers

import (
	"errors"
	"fmt"
	"github.com/akhmettolegen/texert/internal/manager"
	"github.com/akhmettolegen/texert/pkg/application"
	"github.com/akhmettolegen/texert/pkg/models"
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

func (h *Handler) CORSMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Accept, Referer, User-Agent")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}
	ctx.Next()
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

func (h *Handler) FetchMobileUserInfo(ctx *gin.Context) {
	var userInfo models.UserInfo

	userID, isOk := ctx.Get("user_id")
	if !isOk {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errors.New("incorrect userID received").Error()})
		return
	}

	res, err := uuid.FromString(fmt.Sprintf("%v", userID))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errors.New("provided token is not uuid").Error()})
		return
	}

	userInfo.UserID = res

	_, isOk = ctx.Get("token")
	if !isOk {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errors.New("incorrect token received").Error()})
		return
	}

	ctx.Set(models.UserInfoKey, &userInfo)
	ctx.Next()
}