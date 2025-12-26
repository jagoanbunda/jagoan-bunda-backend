package dto

import (
	"time"

	"github.com/google/uuid"
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

type CreateAnthropometryRequest struct{
	ChildID uuid.UUID
	UserID uuid.UUID
	AnthropometryResponse
}
