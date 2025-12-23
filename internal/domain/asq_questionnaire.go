package domain

import (
	"gorm.io/gorm"
)

type ASQQuestionnaire struct{
	gorm.Model
	AgeMonths int `gorm:"type:integer;not null;unique" json:"age_months"`
	AgeRangeMin int `gorm:"type:integer;not null;unique" json:"age_range_min"`
	AgeRangeMax int `gorm:"type:integer;not null;unique" json:"age_range_max"`
	Version string `gorm:"type:varchar(20);default:'1.0'" json:"version"`
	IsActive bool `gorm:"type:boolean;default=true" json:"is_active"`
}
