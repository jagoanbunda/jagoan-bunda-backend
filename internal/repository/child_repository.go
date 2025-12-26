package repository

import (
	"context"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type ChildRepository interface {
	GetByParentID(ctx context.Context, parentID string) ([]domain.Child, error)
	GetBySupervisorID(ctx context.Context, SupervisorID string) ([]domain.Child, error)

	GetFromChildID(ctx context.Context, childID string) ([]domain.Child, error)
	Create(ctx context.Context, child *domain.Child) error
}

type childRepository struct {
	db *gorm.DB
}

// GetBySupervisorID implements [ChildRepository].
func (c *childRepository) GetBySupervisorID(ctx context.Context, SupervisorID string) ([]domain.Child, error) {
	var children []domain.Child

	if err := c.db.WithContext(ctx).
						Joins("JOIN users ON users.id = children.parent_id").
						Where("users.supervisor_id = ? ", SupervisorID).Find(&children).Error;
		err != nil{
			return nil, err
		}

	return children, nil
}

// GetByParentID implements [ChildRepository].
func (c *childRepository) GetByParentID(ctx context.Context, parentID string) ([]domain.Child, error) {
	var child []domain.Child

	if err := c.db.WithContext(ctx).Where("parent_id = ? ", parentID).Find(&child).Error; err != nil {
		return nil, err
	}

	return child, nil
}

// Create implements [ChildRepository].
func (c *childRepository) Create(ctx context.Context, child *domain.Child) error {
	return c.db.Create(child).Error
}

// GetFromChildID implements [ChildRepository].
func (c *childRepository) GetFromChildID(ctx context.Context, childID string) ([]domain.Child, error) {
	var child []domain.Child
	if err := c.db.WithContext(ctx).Where("id = ? ", childID).Find(&child).Error; err != nil {
		return nil, err
	}
	return child, nil
}

func NewChildRepository(db *gorm.DB) ChildRepository {
	return &childRepository{db: db}
}
