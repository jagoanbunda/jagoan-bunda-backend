package domain

import (
	"gorm.io/gorm"
)

// Food merepresentasikan database makanan dengan informasi nutrisi
type Food struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Category      *string   `gorm:"type:varchar(100)" json:"category"`
	PortionName   string    `gorm:"type:varchar(100);not null" json:"portion_name"`
	PortionGram   float64   `gorm:"type:decimal(8,2);not null" json:"portion_gram"`
	EnergyKcal    float64   `gorm:"type:decimal(8,2);not null" json:"energy_kcal"`
	ProteinG      float64   `gorm:"type:decimal(8,2);not null" json:"protein_g"`
	FatG          float64   `gorm:"type:decimal(8,2);not null" json:"fat_g"`
	CarbohydrateG float64   `gorm:"type:decimal(8,2);not null" json:"carbohydrate_g"`
	IsActive      bool      `gorm:"type:boolean;default:true" json:"is_active"`

	FoodIntakes []*FoodIntake `gorm:"many2many:food_intake_foods;" json:"food_intakes"`
}

func (Food) TableName() string {
	return "foods"
}
