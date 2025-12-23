package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ASQScreeningStatus adalah status dari sesi screening ASQ-3
type ASQScreeningStatus string

const (
	ASQScreeningStatusDraft     ASQScreeningStatus = "draft"
	ASQScreeningStatusCompleted ASQScreeningStatus = "completed"
	ASQScreeningStatusReviewed  ASQScreeningStatus = "reviewed"
)

// ASQScreening merepresentasikan sesi screening ASQ-3 yang diisi oleh orang tua
type ASQScreening struct {
	ID              uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ChildID         *uuid.UUID        `gorm:"type:uuid" json:"child_id"`
	Child           *Child            `gorm:"foreignKey:ChildID" json:"child"`
	QuestionnaireID *uuid.UUID        `gorm:"type:uuid" json:"questionnaire_id"`
	Questionnaire   *ASQQuestionnaire `gorm:"foreignKey:QuestionnaireID" json:"questionnaire"`

	ScreeningDate  time.Time          `gorm:"type:date;not null" json:"screening_date"`
	AgeAtScreening int                `gorm:"type:integer;not null" json:"age_at_screening"` // Child's age in months
	CompletedByID  *uuid.UUID         `gorm:"type:uuid" json:"completed_by_id"`
	CompletedBy    *User              `gorm:"foreignKey:CompletedByID" json:"completed_by"` // Parent who filled
	ReviewedByID   *uuid.UUID         `gorm:"type:uuid" json:"reviewed_by_id"`
	ReviewedBy     *User              `gorm:"foreignKey:ReviewedByID" json:"reviewed_by"` // Nakes who reviewed (optional)
	Status         ASQScreeningStatus `gorm:"type:varchar(20);default:'completed'" json:"status"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ASQScreening) TableName() string {
	return "asq_screenings"
}
