package dto

import "github.com/jagoanbunda/jagoanbunda-backend/internal/domain"

type RegisterRequest struct {
	Name     string          `json:"name" binding:"required"`
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=8"`
	Role     domain.UserRole `json:"role" binding:"required,oneof=parent nakes"`
	Phone    *string         `json:"phone" binding:"required"`
	NIK      *string         `json:"NIK"`
	Address  *string         `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Message      *string `json:"message,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
