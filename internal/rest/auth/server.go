package auth

import (
	"auth/service/internal/config"
	"auth/service/internal/lib"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	auth   Auth
	appCfg *config.AppConfig
}

func NewAuthApi(appCfg *config.AppConfig, auth Auth) *AuthApi {
	return &AuthApi{
		auth:   auth,
		appCfg: appCfg,
	}
}

type Auth interface {
	Login(ctx context.Context, email, password string, ac *config.AppConfig) (token string, loginErr error)
	Register(ctx context.Context, email, password, user_name string) (id string, registerErr error)
}

func (ap *AuthApi) LoginHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var req lib.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	if err := lib.ValidateLoginRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	token, loginErr := ap.auth.Login(ctx, req.Email, req.Password, ap.appCfg)
	if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": loginErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func (ap *AuthApi) RegisterHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var req lib.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	if err := lib.ValidateRegisterRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	uID, registerErr := ap.auth.Register(ctx, req.Email, req.Password, req.User_name)
	if registerErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": registerErr.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"user_id": uID,
	})
}

func AddAuthHandlers(r *gin.Engine, ap *AuthApi) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is a main get endpoint",
		})
	})

	authGroup := r.Group("/auth")
	authGroup.POST("/login", ap.LoginHandler)
	authGroup.POST("/register", ap.RegisterHandler)
}
