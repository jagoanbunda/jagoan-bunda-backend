package domain

import (
	"github.com/google/uuid"
)

// ASQAnswerType adalah enum untuk jenis jawaban dalam ASQ-3
type ASQAnswerType string

const (
	ASQAnswerYes       ASQAnswerType = "yes"
	ASQAnswerSometimes ASQAnswerType = "sometimes"
	ASQAnswerNotYet    ASQAnswerType = "not_yet"
)

// ASQAnswer merepresentasikan jawaban untuk setiap pertanyaan dalam screening ASQ-3
type ASQAnswer struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ScreeningID *uuid.UUID    `gorm:"type:uuid" json:"screening_id"`
	Screening   *ASQScreening `gorm:"foreignKey:ScreeningID" json:"screening"`
	QuestionID  *uuid.UUID    `gorm:"type:uuid" json:"question_id"`
	Question    *ASQQuestion  `gorm:"foreignKey:QuestionID" json:"question"`

	Answer ASQAnswerType `gorm:"type:varchar(20);not null" json:"answer"`
	Score  int           `gorm:"type:integer;not null" json:"score"` // yes=10, sometimes=5, not_yet=0
	Notes  *string       `gorm:"type:text" json:"notes"`
}

func (ASQAnswer) TableName() string {
	return "asq_answers"
}
