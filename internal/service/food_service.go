package service

import (
	"context"
	"fmt"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
)

type FoodService interface {
	Create(ctx context.Context, request *dto.FoodRequest) (*dto.FoodResponse, error)
	Get(ctx context.Context) ([]dto.FoodResponse, error)
	Update(ctx context.Context, request *dto.FoodRequest) (*dto.FoodResponse, error)
	Delete(ctx context.Context, request *dto.FoodRequest) error
	Search(ctx context.Context, key string) ([]dto.FoodResponse, error)
}

type foodService struct {
	repository repository.FoodRepository
}

// Search implements [FoodService].
func (f *foodService) Search(ctx context.Context, key string) ([]dto.FoodResponse, error) {
	var response []dto.FoodResponse

	records, err := f.repository.Search(ctx, key)
	if err != nil {
		return nil, err
	}

	for _, record := range records{
		response = append(response, *dto.NewFoodResponseFromDomain(&record))
	}

	return response, nil
}

// Create implements [FoodService].
func (f *foodService) Create(ctx context.Context, request *dto.FoodRequest) (*dto.FoodResponse, error) {
	food := dto.NewFoodDomainFromRequest(request)
	if err := f.repository.Create(ctx, food); err != nil {
		return nil, err
	}
	response := dto.NewFoodResponseFromDomain(food)
	return response, nil
}

// Delete implements [FoodService].
func (f *foodService) Delete(ctx context.Context, request *dto.FoodRequest) error {
	food := dto.NewFoodDomainFromRequest(request)
	if err := f.repository.Delete(ctx, food); err != nil {
		return err
	}
	return nil
}

// Get implements [FoodService].
func (f *foodService) Get(ctx context.Context) ([]dto.FoodResponse, error) {
	var response []dto.FoodResponse

	records, err := f.repository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("SERVICE ERROR : %w", err)
	}

	for _, record := range records {
		response = append(response, *dto.NewFoodResponseFromDomain(&record))
	}

	return response, nil
}

// Update implements [FoodService].
func (f *foodService) Update(ctx context.Context, request *dto.FoodRequest) (*dto.FoodResponse, error) {
	food := dto.NewFoodDomainFromRequest(request)
	if err := f.repository.Update(ctx, food); err != nil {
		return nil, err
	}
	updatedFood := dto.NewFoodResponseFromDomain(food)
	return updatedFood, nil
}

func NewFoodService(repository repository.FoodRepository) FoodService {
	return &foodService{repository: repository}
}
