package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
)

type AnthropometryResponse struct{
	Date time.Time `json:"time"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	HeadCircumference float32 `json:"head_circumference"`
	AgeMonths int `json:"age_months"`
	ZScoreBBU         float32  `json:"zscore_bbu"`
	ZScoreTBU         float32  `json:"zscore_tbu"`
	ZScoreBBTB        float32  `json:"zscore_bbtb"`
	StatusBBU         float32  `json:"status_bbu"`
	StatusTBU         float32  `json:"status_tbu"`
	StatusBBTB        float32  `json:"status_bbtb"`
}

func NewAnthropometryResponse(record *domain.Anthropometry, ageMonths int) ( *AnthropometryResponse) {
	return &AnthropometryResponse{
		Date: record.CreatedAt,
		Weight: record.Weight,
		Height: record.Height,
		HeadCircumference: record.HeadCircumference,
		AgeMonths: ageMonths,
		ZScoreBBU: record.ZScoreBBU,
		ZScoreTBU: record.ZScoreTBU,
		ZScoreBBTB: record.ZScoreBBTB,
		StatusBBU: record.StatusBBU,
		StatusTBU : record.StatusTBU,
		StatusBBTB: record.StatusBBTB,
	}
}

type CreateAnthropometryRequest struct{
	ChildID uuid.UUID
	UserID uuid.UUID
	AnthropometryResponse
}
type UpdateAnthropometryRequest struct{
	ID uint
	ChildID uuid.UUID
	AnthropometryResponse
}

type DeleteAnthropometryRequest struct{
	ID uint
	ChildID uuid.UUID
}
