package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/dto"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/repository"
)

type ChildService interface {
	GetChildWithAccess(ctx context.Context, childID uuid.UUID, userID uuid.UUID, role domain.UserRole) (*dto.ChildResponse, error)
}
type childService struct {
	repository repository.ChildRepository
}

// GetChildWithAccess implements [ChildService].
func (c *childService) GetChildWithAccess(ctx context.Context, childID uuid.UUID, userID uuid.UUID, role domain.UserRole) (*dto.ChildResponse, error) {
	child, err := c.repository.GetFromChildID(ctx, childID.String())
	if err != nil {
		return nil, err
	}

	if role == domain.RoleParent && child.ParentID != userID{
		return nil, err
	} else if role == domain.RoleNakes && child.Parent.SupervisorID != &userID{
		return nil, err
	}

	childResponse := &dto.ChildResponse{
		Name: child.Name,
		Birthday: child.Birthday,
		Gender: child.Gender,
		NIK: child.NIK,
		BirthWeight: child.BirthWeight,
		BirthHeight: child.BirthHeight,
	}

	return childResponse, nil
}

func NewChildService(repository repository.ChildRepository) ChildService {
	return &childService{repository: repository}
}
