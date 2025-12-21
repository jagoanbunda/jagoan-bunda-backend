package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KPSPScreening struct {
	gorm.Model

	NakesID *uuid.UUID `gorm:"type:uuid" json:"nakes_id"`
	Nakes   *User      `gorm:"foreignKey:NakesID" json:"nakes"`

	ChildID *uuid.UUID `gorm:"type:uuid" json:"child_id"`
	Child   *Child     `gorm:"foreignKey:ChildID" json:"child"`

	Date    time.Time `gorm:"type:date;not null" json:"date"`
	Answers []string  `gorm:"type:json" json:"answers"`
	Result  *string   `gorm:"type:text" json:"result"`
}

func (KPSPScreening) TableName() string {
	return "kpsp_screenings"
}
