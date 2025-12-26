// Package service is used to serve the Service in Service-Repository pattern, where bussiness logic applied with using the repository
package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type UserService interface {
	Get(ctx context.Context, uuid string) (*dto.UserResponse, error)
	UpdateProfile(ctx context.Context, userID string, req *dto.UpdateUserRequest, profilePicture *multipart.FileHeader) (*dto.UserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func (u *userService) Get(ctx context.Context, uuid string) (*dto.UserResponse, error) {
	user, err := u.repository.FindByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	data := &dto.UserResponse{
		UUID:            &user.ID,
		Name:            user.Name,
		ProfilePicture:  user.ProfilePicture,
		Email:           user.Email,
		IsVerified:      user.IsVerified,
		Address:         user.Address,
		NIK:             user.NIK,
		AssignedParents: user.AssignedParents,
	}

	return data, nil
}

func (u *userService) UpdateProfile(ctx context.Context, userID string, req *dto.UpdateUserRequest, profilePicture *multipart.FileHeader) (*dto.UserResponse, error) {
	// Get existing user
	user, err := u.repository.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Update fields if provided
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Address != nil {
		user.Address = req.Address
	}
	if req.NIK != nil {
		user.NIK = req.NIK
	}

	// Handle profile picture upload if provided
	if profilePicture != nil {
		// Validate file type
		if !utils.IsAllowedImageType(profilePicture.Filename) {
			return nil, fmt.Errorf("invalid file type. Allowed types: jpg, jpeg, png, gif, webp")
		}

		// Validate file size (max 5MB)
		maxSize := utils.GetMaxFileSizeMB() * 1024 * 1024
		if profilePicture.Size > maxSize {
			return nil, fmt.Errorf("file size exceeds maximum limit of %dMB", utils.GetMaxFileSizeMB())
		}

		// Delete old profile picture if exists
		if user.ProfilePicture != nil && *user.ProfilePicture != "" {
			oldFilePath := filepath.Join(utils.GetUploadDir(), "profile_pictures", filepath.Base(*user.ProfilePicture))
			_ = os.Remove(oldFilePath) // Ignore error if file doesn't exist
		}

		// Create upload directory
		uploadDir := filepath.Join(utils.GetUploadDir(), "profile_pictures")
		if err := utils.EnsureDir(uploadDir); err != nil {
			return nil, fmt.Errorf("failed to create upload directory: %w", err)
		}

		// Generate unique filename
		uniqueFilename := utils.GenerateUniqueFilename(profilePicture.Filename)
		filePath := filepath.Join(uploadDir, uniqueFilename)

		// Open source file
		src, err := profilePicture.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open uploaded file: %w", err)
		}
		defer src.Close()

		// Create destination file
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create destination file: %w", err)
		}
		defer dst.Close()

		// Copy file content
		if _, err := io.Copy(dst, src); err != nil {
			return nil, fmt.Errorf("failed to save file: %w", err)
		}

		// Set profile picture URL
		relativePath := fmt.Sprintf("/uploads/profile_pictures/%s", uniqueFilename)
		user.ProfilePicture = &relativePath
	}

	// Save updated user
	if err := u.repository.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Return updated user response
	return &dto.UserResponse{
		UUID:            &user.ID,
		Name:            user.Name,
		ProfilePicture:  user.ProfilePicture,
		Email:           user.Email,
		IsVerified:      user.IsVerified,
		Address:         user.Address,
		NIK:             user.NIK,
		AssignedParents: user.AssignedParents,
	}, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repository: repo}
}
