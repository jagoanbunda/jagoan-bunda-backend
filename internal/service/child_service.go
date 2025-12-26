package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
)

type ChildService interface {
	GetChildByIDWithAccess(ctx context.Context, childID uuid.UUID, userID uuid.UUID, role domain.UserRole) ([]dto.ChildResponse, error)
	Create(ctx context.Context, request *dto.CreateChildRequest) (*dto.ChildResponse, error)
	GetChildWithAccess(ctx context.Context, userID uuid.UUID, role domain.UserRole) ([]dto.ChildResponse, error)
}
type childService struct {
	repository repository.ChildRepository
}

// GetChildWithAccess implements [ChildService].
func (c *childService) GetChildWithAccess(ctx context.Context, userID uuid.UUID, role domain.UserRole) ([]dto.ChildResponse, error){
	var responses []dto.ChildResponse
	var records []domain.Child
	var err error
	switch role {
		case domain.RoleNakes:
			records, err = c.repository.GetBySupervisorID(ctx, userID.String())
		case domain.RoleParent:
			records, err = c.repository.GetByParentID(ctx, userID.String())
	}

	if err != nil {
		return nil, err
	}

	for _, record := range records{
		responses = append(responses, dto.ChildResponse{
			Name: record.Name,
			Birthday: record.Birthday,
			Gender: record.Gender,
			NIK: record.NIK,
			BirthWeight: record.BirthWeight,
			BirthHeight: record.BirthHeight,
		})
	}
	return responses, nil

}

// Create implements [ChildService].
func (c *childService) Create(ctx context.Context, request *dto.CreateChildRequest) (*dto.ChildResponse, error) {

	newChild := domain.Child{
		Name:        request.Name,
		Birthday:    request.Birthday,
		Gender:      request.Gender,
		NIK:         request.NIK,
		BirthWeight: request.BirthWeight,
		BirthHeight: request.BirthHeight,
		ParentID:    request.ParentID,
	}

	if err := c.repository.Create(ctx, &newChild); err != nil {
		return nil, err
	}

	response := &dto.ChildResponse{
		Name:        newChild.Name,
		Birthday:    newChild.Birthday,
		NIK:         newChild.NIK,
		BirthWeight: newChild.BirthWeight,
		BirthHeight: newChild.BirthHeight,
	}
	return response, nil
}

// GetChildByIDWithAccess implements [ChildService].
func (c *childService) GetChildByIDWithAccess(ctx context.Context, childID uuid.UUID, userID uuid.UUID, role domain.UserRole) ([]dto.ChildResponse, error) {
	var childResponses []dto.ChildResponse
	children, err := c.repository.GetFromChildID(ctx, childID.String())
	if err != nil {
		return nil, err
	}

	for _, record := range children {
		if role == domain.RoleParent && record.ParentID != userID {
			continue
		} else if role == domain.RoleNakes && record.Parent.SupervisorID != &userID {
			continue
		}

		childResponses = append(childResponses, dto.ChildResponse{
			Name:        record.Name,
			Birthday:    record.Birthday,
			Gender:      record.Gender,
			NIK:         record.NIK,
			BirthWeight: record.BirthWeight,
			BirthHeight: record.BirthHeight,
		})
	}

	return childResponses, nil
}

func NewChildService(repository repository.ChildRepository) ChildService {
	return &childService{repository: repository}
}
