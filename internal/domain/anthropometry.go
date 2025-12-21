package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Anthropometry struct {
	gorm.Model            // <- ID, CreatedAt, UpdatedAt
	ChildID    *uuid.UUID `gorm:"type:uuid;unique" json:"child_id"`
	Child      *Child     `gorm:"foreignKey:ChildID" json:"child"`
	UserID     *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User       *User      `gorm:"foreignKey:UserID" json:"user"`

	Weight            float32  `gorm:"type:decimal(5, 2); not null" json:"weight"`
	Height            float32  `gorm:"type:decimal(5, 2); not null" json:"height"`
	HeadCircumference *float32 `gorm:"decimal(5, 2)" json:"head_circumference"`
	ZScoreBBU         float32  `gorm:"type:decimal(4, 2)" json:"zscore_bbu"`
	ZScoreTBU         float32  `gorm:"type:decimal(4, 2)" json:"zscore_tbu"`
	ZScoreBBTB        float32  `gorm:"type:decimal(4, 2)" json:"zscore_bbtb"`
	StatusBBU         float32  `gorm:"type:decimal(4, 2)" json:"status_bbu"`
	StatusTBU         float32  `gorm:"type:decimal(4, 2)" json:"status_tbu"`
	StatusBBTB        float32  `gorm:"type:decimal(4, 2)" json:"status_bbtb"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Anthropometry) TableName() string {
	return "anthropometries"
}
