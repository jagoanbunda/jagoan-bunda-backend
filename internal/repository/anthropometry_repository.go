package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type AnthropometryRepository interface {
	Get(ctx context.Context) ([]domain.Anthropometry, error)
	GetFromChildID(ctx context.Context, childID uuid.UUID) ([]domain.Anthropometry, error)
}

type anthropometryRepository struct {
	db *gorm.DB
}

// Get implements [AnthropometryRepository].
func (a *anthropometryRepository) Get(ctx context.Context) ([]domain.Anthropometry, error) {
	var records []domain.Anthropometry
	if err := a.db.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetFromChildID implements [AnthropometryRepository].
func (a *anthropometryRepository) GetFromChildID(ctx context.Context, childID uuid.UUID) ([]domain.Anthropometry, error) {
	var resp []domain.Anthropometry
	if err := a.db.WithContext(ctx).Where("child_id = ?", childID).
		Select("created_at", "weight", "height", "head_circumference", "z_score_bbu", "z_score_tbu", "z_score_bbtb", "status_bbu", "status_tbu", "status_bbtb").
		Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("cannot find anthropometry records : %v", err)
	}
	return resp, nil
}

func NewAnthropometryRepository(db *gorm.DB) AnthropometryRepository {
	return &anthropometryRepository{db: db}
}
