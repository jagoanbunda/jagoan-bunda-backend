package dto

import (
	"time"

	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
)

type ChildResponse struct{
	Name string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender domain.Gender	`json:"gender"`
	NIK string `json:"nik"`
	BirthWeight float32	`json:"birth_weight"`
	BirthHeight float32	`json:"birth_day"`
}


