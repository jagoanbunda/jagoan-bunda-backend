package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Child struct {
	ID       uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ParentID uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	Parent   *User      `gorm:"foreignKey:ParentID" json:"parent"`

	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Birthday    time.Time `gorm:"type:date;not null" json:"birthday"`
	Gender      Gender    `gorm:"type:varchar(10);not null" json:"gender"`
	NIK         string    `gorm:"type:varchar(255);not null" json:"nik"`
	BirthWeight float32   `gorm:"type:decimal(5, 2);not null" json:"birth_weight"`
	BirthHeight float32   `gorm:"type:decimal(5, 2);not null" json:"birth_height"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Child) TableName() string {
	return "children"
}
