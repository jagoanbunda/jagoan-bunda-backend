package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
	// "github.com/jagoanbunda/jagoanbunda-backend/internal/utils"
)

type AnthropometryService interface {
	GetRecordFromChildID(ctx context.Context, childID uuid.UUID) ([]dto.AnthropometryResponse, error)
	CreateRecordWithChildID(ctx context.Context, request *dto.CreateAnthropometryRequest) (*dto.AnthropometryResponse, error)
}

type anthropometryService struct {
	repository repository.AnthropometryRepository
}

// CreateRecordWithChildID implements [AnthropometryService].
func (a *anthropometryService) CreateRecordWithChildID(ctx context.Context, request *dto.CreateAnthropometryRequest) (*dto.AnthropometryResponse, error) {
	newRecord := &domain.Anthropometry{
		ChildID: request.ChildID,
		UserID: request.UserID,
		Weight: request.Weight,
		Height: request.Height,
		HeadCircumference: request.HeadCircumference,
		ZScoreBBU: request.ZScoreBBU,
		ZScoreTBU: request.ZScoreTBU,
		ZScoreBBTB: request.ZScoreBBTB,
		StatusBBU: request.StatusBBU,
		StatusTBU: request.StatusTBU,
		StatusBBTB: request.StatusBBTB,
	}

	if err := a.repository.Create(ctx, newRecord); err != nil {
		return nil, err
	}
	ageMonths := utils.CalculateAgeInMonths(newRecord.Child.Birthday)

	response := &dto.AnthropometryResponse{
		Date: newRecord.CreatedAt,
		Weight : newRecord.Weight,
		Height: newRecord.Height,
		HeadCircumference: newRecord.HeadCircumference,
		AgeMonths: ageMonths,
		ZScoreBBU: newRecord.ZScoreBBU,
		ZScoreTBU: newRecord.ZScoreTBU,
		ZScoreBBTB: newRecord.ZScoreBBTB,
		StatusBBU: newRecord.StatusBBU,
		StatusTBU: newRecord.StatusTBU,
		StatusBBTB: newRecord.StatusBBTB,
	}

	return response, nil

}

// GetRecord implements [AnthropometryService].
func (a *anthropometryService) GetRecordFromChildID(ctx context.Context, childID uuid.UUID) ([]dto.AnthropometryResponse, error) {
	var response []dto.AnthropometryResponse
	records, err := a.repository.GetFromChildID(ctx, childID)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		ageMonths := utils.CalculateAgeInMonths(record.Child.Birthday)
		response = append(response, dto.AnthropometryResponse{
			Date:              record.CreatedAt,
			Weight:            record.Weight,
			Height:            record.Height,
			HeadCircumference: record.HeadCircumference,
			AgeMonths:         ageMonths,
			ZScoreBBU:         record.ZScoreBBU,
			ZScoreTBU:         record.ZScoreTBU,
			ZScoreBBTB:        record.ZScoreBBTB,
			StatusBBU:         record.StatusBBU,
			StatusTBU:         record.StatusTBU,
			StatusBBTB:        record.StatusBBTB,
		})
	}

	return response, nil
}

func NewAnthropometryService(repository repository.AnthropometryRepository) AnthropometryService {
	return &anthropometryService{repository: repository}
}
