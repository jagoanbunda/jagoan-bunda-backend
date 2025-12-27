package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

// Register godoc
// @Summary Register user baru
// @Description Mendaftarkan user baru dengan role parent atau nakes
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Data registrasi user"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (a *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := a.AuthService.Register(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, authResponse)
	return
}

// Login godoc
// @Summary Login user
// @Description Autentikasi user dengan email dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Kredensial login"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := a.AuthService.Login(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authResponse)
	return
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Mendapatkan access token baru menggunakan refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/refresh [post]
func (a *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	authResponse, err := a.AuthService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, authResponse)
}
