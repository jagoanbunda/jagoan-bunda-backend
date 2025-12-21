package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.AuthResponse, error)
	RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest) (*dto.AuthResponse, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func (s *authService) RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	claims, err := utils.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token : %v", err)
	}

	user, err := s.userRepository.FindByID(ctx, claims.UserID.String())
	if err != nil {
		return nil, fmt.Errorf("user not found : %w", err)
	}

	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error on generating access token : %v", err)
	}

	return &dto.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: &request.RefreshToken,
	}, nil
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepository.FindByEmail(ctx, request.Email)
	if err != nil || user == nil {
		return nil, fmt.Errorf(
			"cannot find user : %w",
			err,
		)
	}
	//fmt.Printf(`DEBUG USER : %+v\n DEBUG REQUEST : %+v\n`, user, request)

	if err = utils.CheckPassword(user.PasswordHash, request.Password); err != nil {
		return nil, fmt.Errorf("credentials is not found : %w", err)
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error on generating access token : %w", err)
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error on generating refresh token : %w", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: utils.StringToPtr(refreshToken),
	}, nil

}

func (s *authService) Register(ctx context.Context, request *dto.RegisterRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepository.FindByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("Error on finding email : %v", err.Error()))
	}
	if user != nil {
		return nil, errors.New("email sudah dipakai")
	}
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to hash password : %v", err.Error()))
	}
	newUser := &domain.User{
		Name:         request.Name,
		Email:        request.Email,
		Phone:        request.Phone,
		PasswordHash: hashedPassword,
		Role:         request.Role,
		IsVerified:   false,
		Address:      request.Address,
		NIK:          request.NIK,
	}
	err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("User gagal dibuat : %v", err.Error()))
	}

	accessToken, err := utils.GenerateAccessToken(newUser.ID, newUser.Email, newUser.Role)
	if err != nil {
		return nil, fmt.Errorf("error on generating access token : %w", err)
	}
	refreshToken, err := utils.GenerateRefreshToken(newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("error on generating refresh token : %w", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: utils.StringToPtr(refreshToken),
	}, nil
}

func (s *authService) NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}
