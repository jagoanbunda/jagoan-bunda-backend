package domain

import (
	"github.com/google/uuid"
)

// ASQRecommendation merepresentasikan rekomendasi stimulasi untuk domain yang memerlukan dukungan
type ASQRecommendation struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Domain              ASQDomain `gorm:"type:varchar(20);not null" json:"domain"`
	AgeMonthsMin        int       `gorm:"type:integer;not null" json:"age_months_min"`
	AgeMonthsMax        int       `gorm:"type:integer;not null" json:"age_months_max"`
	ActivityTitle       string    `gorm:"type:varchar(255);not null" json:"activity_title"`
	ActivityDescription string    `gorm:"type:text;not null" json:"activity_description"`
	VideoURL            *string   `gorm:"type:varchar(500)" json:"video_url"`
	SortOrder           int       `gorm:"type:integer;default:0" json:"sort_order"`
}

func (ASQRecommendation) TableName() string {
	return "asq_recommendations"
}
