package dto

import "github.com/jagoanbunda/jagoanbunda-backend/internal/domain"

type FoodRequest struct{
	ID *uint 				`json:"id"`
	Name          string 	 `json:"name"`
	Category      *string `json:"category"`
	PortionName   string `json:"portion_name"`
	PortionGram   float64 `json:"portion_gram"`
	EnergyKcal    float64 `json:"energy_kcal"`
	ProteinG      float64 `json:"protein_g"`
	FatG          float64 `json:"fat_g"`
	CarbohydrateG float64 `json:"carbohydrate_g"`
	IsActive      bool `json:"is_active"`
}

func NewFoodDomainFromRequest(request *FoodRequest) *domain.Food{
	newFoodDomain := &domain.Food{
		Name : request.Name,
		Category: request.Category,
		PortionName: request.PortionName,
		PortionGram: request.PortionGram,
		EnergyKcal: request.EnergyKcal,
		ProteinG : request.ProteinG,
		FatG : request.FatG,
		CarbohydrateG: request.CarbohydrateG,
		IsActive: request.IsActive,
	}

	if request.ID != nil {
		newFoodDomain.ID = *request.ID
	}
	return newFoodDomain

}


type FoodResponse struct {
	FoodRequest
}
func NewFoodResponseFromDomain(domain *domain.Food) *FoodResponse {
	return &FoodResponse{
		FoodRequest: FoodRequest{
			Name : domain.Name,
			Category: domain.Category,
			PortionName: domain.PortionName,
			PortionGram: domain.PortionGram,
			EnergyKcal: domain.EnergyKcal,
			ProteinG : domain.ProteinG,
			FatG : domain.FatG,
			CarbohydrateG: domain.CarbohydrateG,
			IsActive: domain.IsActive,
		},
	}
}
