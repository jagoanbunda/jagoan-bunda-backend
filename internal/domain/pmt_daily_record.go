package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PMTDailyRecord struct {
	gorm.Model

	UserID    *uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User      *User       `gorm:"foreignKey:UserID" json:"user"`
	ProgramID *uint       `gorm:"index"`
	Program   *PMTProgram `gorm:"foreignKey:ProgramID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Date     time.Time `gorm:"type:date;not null" json:"date"`
	Consumed bool      `gorm:"type:bool;not null"`
	Notes    *string   `gorm:"type:text"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

