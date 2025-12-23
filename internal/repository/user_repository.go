package repository

import (
	"context"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}
