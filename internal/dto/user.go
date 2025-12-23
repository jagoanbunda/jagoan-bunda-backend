// Package dto are used to serialize data to object for transfer purposes
package dto

import (
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
)

type UserResponse struct {
	UUID            *uuid.UUID    `json:"uuid"`
	Name            string        `json:"name"`
	ProfilePicture  *string       `json:"profile_picture"`
	Email           string        `json:"email"`
	IsVerified      bool          `json:"is_verified"`
	Address         *string       `json:"address"`
	NIK             *string       `json:"nik"`
	AssignedParents []domain.User `json:"assigned_parents"`
}

// UpdateUserRequest represents the request payload for updating user profile
// All fields are optional - only provided fields will be updated
type UpdateUserRequest struct {
	Name    *string `form:"name"`
	Phone   *string `form:"phone"`
	Address *string `form:"address"`
	NIK     *string `form:"nik"`
}
