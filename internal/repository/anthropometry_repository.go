package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"gorm.io/gorm"
)

type AnthropometryRepository interface {
	Get(ctx context.Context) ([]domain.Anthropometry, error)
	GetByIDWithChildID(ctx context.Context, anthropometryID string, childID uuid.UUID) (*domain.Anthropometry, error)
	GetFromChildID(ctx context.Context, childID uuid.UUID) ([]domain.Anthropometry, error)
	Create(ctx context.Context, anthropometry *domain.Anthropometry) error
	GetLatestRecordFromChildID(ctx context.Context, childID uuid.UUID) (*domain.Anthropometry, error)
	UpdateWithChildID(ctx context.Context, anthropometry *domain.Anthropometry) error
	GetChildBirthdayFromChildID(ctx context.Context, childID uuid.UUID) (*time.Time, error)
	Delete(ctx context.Context, anthropometry *domain.Anthropometry) error
}

type anthropometryRepository struct {
	db *gorm.DB
}

// Delete implements [AnthropometryRepository].
func (a *anthropometryRepository) Delete(ctx context.Context, anthropometry *domain.Anthropometry) error {
	return a.db.WithContext(ctx).Delete(&anthropometry).Error
}

// GetChildBirthdayFromChildID implements [AnthropometryRepository].
func (a *anthropometryRepository) GetChildBirthdayFromChildID(ctx context.Context, childID uuid.UUID) (*time.Time, error) {
	var record domain.Anthropometry
	if err := a.db.WithContext(ctx).Where("child_id = ? ", childID).Preload("Child").First(&record).Error; err != nil {
		return nil, err
	}
	if record.Child == nil {
		return nil, fmt.Errorf("child record not found")
	}

	return &record.Child.Birthday, nil
}

// Update implements [AnthropometryRepository].
func (a *anthropometryRepository) UpdateWithChildID(ctx context.Context, anthropometry *domain.Anthropometry) error {
	return a.db.WithContext(ctx).
		Where("child_id = ? ", &anthropometry.ChildID).
		Select("Weight", "Height", "HeadCircumference", "ZScoreBBU", "ZScoreTBU", "ZScoreBBTB", "StatusBBU", "StatusTBU", "StatusBBTB"). // -> to explicitly update the selected columns
		Updates(anthropometry).
		Error
}

// GetByIDWithChildID implements [AnthropometryRepository].
func (a *anthropometryRepository) GetByIDWithChildID(ctx context.Context, anthropometryID string, childID uuid.UUID) (*domain.Anthropometry, error) {
	var record domain.Anthropometry
	if err := a.db.WithContext(ctx).Preload("Child").
		Where("child_id = ?", childID).
		Where("id = ? ", anthropometryID).
		First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// GetLatestRecordFromChildID implements [AnthropometryRepository].
func (a *anthropometryRepository) GetLatestRecordFromChildID(ctx context.Context, childID uuid.UUID) (*domain.Anthropometry, error) {
	var record domain.Anthropometry
	if err := a.db.WithContext(ctx).Preload("Child").Where("child_id = ?", childID).Last(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// Create implements [AnthropometryRepository].
func (a *anthropometryRepository) Create(ctx context.Context, anthropometry *domain.Anthropometry) error {
	return a.db.WithContext(ctx).Create(anthropometry).Error
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
		Preload("Child").
		Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("cannot find anthropometry records : %v", err)
	}
	return resp, nil
}

func NewAnthropometryRepository(db *gorm.DB) AnthropometryRepository {
	return &anthropometryRepository{db: db}
}
