package repository

import (
	"context"
	"fmt"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type FoodRepository interface {
	Create(ctx context.Context, food *domain.Food) error
	Get(ctx context.Context) ([]domain.Food, error)
	Update(ctx context.Context, food *domain.Food) error
	Delete(ctx context.Context, food *domain.Food) error
}

type foodRepository struct {
	db *gorm.DB
}

// Create implements [FoodRepository].
func (f *foodRepository) Create(ctx context.Context, food *domain.Food) error {
	return f.db.WithContext(ctx).Create(&food).Error
}

// Delete implements [FoodRepository].
func (f *foodRepository) Delete(ctx context.Context, food *domain.Food) error {
	return f.db.WithContext(ctx).Delete(&food).Error
}

// Get implements [FoodRepository].
func (f *foodRepository) Get(ctx context.Context) ([]domain.Food, error) {
	var records []domain.Food
	if err := f.db.WithContext(ctx).
						Find(&records).Error ; err != nil {
							return nil, fmt.Errorf("REPOSITORY ERROR : %w", err)
						}
	return records, nil
}

// Update implements [FoodRepository].
func (f *foodRepository) Update(ctx context.Context, food *domain.Food) error {
	return f.db.WithContext(ctx).Save(&food).Error
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepository{db: db}
}
