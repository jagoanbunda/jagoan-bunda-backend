// Package handler are used for gin to handle request, utilizes the service package
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/service"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type UserHandler interface {
	Get(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

// Get godoc
// @Summary Get current user profile
// @Description Mendapatkan profil user yang sedang login
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/me [get]
func (u *userHandler) Get(c *gin.Context) {
	uuid, exist := c.Get("userInfo")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}
	data, ok := uuid.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad claims"})
		return
	}
	userResponse, err := u.userService.Get(c.Request.Context(), data.UserID.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userResponse,
	})
}

// UpdateProfile handles user profile update with optional profile picture upload
// @Summary Update user profile
// @Description Update user profile information and optionally upload a profile picture
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string false "User name"
// @Param phone formData string false "User phone number"
// @Param address formData string false "User address"
// @Param nik formData string false "User NIK"
// @Param profile_picture formData file false "Profile picture (jpg, jpeg, png, gif, webp)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/profile [put]
func (u *userHandler) UpdateProfile(c *gin.Context) {
	// Get user info from context (set by auth middleware)
	userInfo, exists := c.Get("userInfo")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	claims, ok := userInfo.(*utils.AccessTokenClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid token claims",
		})
		return
	}

	// Bind form data
	var req dto.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request data: " + err.Error(),
		})
		return
	}

	// Get optional profile picture file
	file, _ := c.FormFile("profile_picture")

	// Call service to update profile
	userResponse, err := u.userService.UpdateProfile(c.Request.Context(), claims.UserID.String(), &req, file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "profile updated successfully",
		"data":    userResponse,
	})
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}
