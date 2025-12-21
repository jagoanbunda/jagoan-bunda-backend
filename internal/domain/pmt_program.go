package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PMTStatus string

const (
	Active    PMTStatus = "active"
	Completed PMTStatus = "completed"
	Cancelled PMTStatus = "cancelled"
)

type PMTProgram struct {
	gorm.Model
	ChildID *uuid.UUID `gorm:"type:uuid" json:"child_id"`
	Child   *Child     `gorm:"foreignKey:ChildID" json:"child"`
	UserID  *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User    *User      `gorm:"foreignKey:UserID" json:"user"`

	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;not null" json:"end_date"`
	Status    PMTStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	Notes     *string   `gorm:"type:text"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (PMTProgram) TableName() string {
	return "pmt_programs"
}
