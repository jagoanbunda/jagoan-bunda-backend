package domain

import (
	"github.com/google/uuid"
)

// ASQDomain adalah enum untuk domain yang diukur dalam ASQ-3
type ASQDomain string

const (
	ASQDomainCommunication  ASQDomain = "communication"
	ASQDomainGrossMotor     ASQDomain = "gross_motor"
	ASQDomainFineMotor      ASQDomain = "fine_motor"
	ASQDomainProblemSolving ASQDomain = "problem_solving"
	ASQDomainPersonalSocial ASQDomain = "personal_social"
)

// ASQQuestion merepresentasikan pertanyaan dalam kuesioner ASQ-3
type ASQQuestion struct {
	ID              uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	QuestionnaireID *uuid.UUID        `gorm:"type:uuid" json:"questionnaire_id"`
	Questionnaire   *ASQQuestionnaire `gorm:"foreignKey:QuestionnaireID" json:"questionnaire"`

	Domain          ASQDomain `gorm:"type:varchar(20);not null" json:"domain"`
	QuestionNumber  int       `gorm:"type:integer;not null" json:"question_number"` // 1-6 within domain
	QuestionText    string    `gorm:"type:text;not null" json:"question_text"`
	QuestionTextID  *string   `gorm:"type:text" json:"question_text_id"` // Indonesian translation
	IllustrationURL *string   `gorm:"type:varchar(500)" json:"illustration_url"`
	HowToCheck      *string   `gorm:"type:text" json:"how_to_check"` // Instructions for parent
	SortOrder       int       `gorm:"type:integer;default:0" json:"sort_order"`
}

func (ASQQuestion) TableName() string {
	return "asq_questions"
}
