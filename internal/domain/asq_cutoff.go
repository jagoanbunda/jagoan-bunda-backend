package domain

import (
	"github.com/google/uuid"
)

// ASQCutoff merepresentasikan cutoff score untuk setiap domain per usia dalam ASQ-3
type ASQCutoff struct {
	ID              uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	QuestionnaireID *uuid.UUID        `gorm:"type:uuid" json:"questionnaire_id"`
	Questionnaire   *ASQQuestionnaire `gorm:"foreignKey:QuestionnaireID" json:"questionnaire"`

	Domain         ASQDomain `gorm:"type:varchar(20);not null" json:"domain"`
	CutoffScore    float64   `gorm:"type:decimal(4,1);not null" json:"cutoff_score"`    // Below this = refer
	MonitoringZone float64   `gorm:"type:decimal(4,1);not null" json:"monitoring_zone"` // Below this but above cutoff = monitor
}

func (ASQCutoff) TableName() string {
	return "asq_cutoffs"
}
