package repository

import (
	"context"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type ChildRepository interface {
	GetFromChildID(ctx context.Context, childID string) (*domain.Child, error)
}

type childRepository struct {
	db *gorm.DB
}

// GetFromChildID implements [ChildRepository].
func (c *childRepository) GetFromChildID(ctx context.Context, childID string)(*domain.Child, error) {
	var child domain.Child
	if err := c.db.WithContext(ctx).Where("id = ? ", childID).First(&child).Error ; err != nil{
		return nil, err
	}
	return &child, nil
}

func NewChildRepository(db *gorm.DB) ChildRepository {
	return &childRepository{db: db}
}
