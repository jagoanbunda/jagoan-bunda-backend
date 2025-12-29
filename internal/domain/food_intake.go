package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MealTime string

const (
	Breakfast MealTime = "breakfast"
	Lunch     MealTime = "lunch"
	Dinner    MealTime = "dinner"
)


type FoodIntake struct {
	gorm.Model
	ChildID *uuid.UUID `gorm:"type:uuid" json:"child_id"`
	Child   *Child     `gorm:"foreignKey:ChildID" json:"child"`
	UserID  *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User    *User      `gorm:"foreignKey:UserID" json:"user"`

	MealTime          MealTime   `gorm:"type:varchar(20);not null;default:'breakfast'"`
	Foods             []*Food 		`gorm:"many2many:food_intake_foods;" json:"foods"`
	TotalEnergy       float64    `gorm:"type:decimal(8, 2)" json:"total_energy"`
	TotalProtein      float64    `gorm:"type:decimal(8,2)" json:"total_protein"`
	TotalFat          float64    `gorm:"type:decimal(8,2)" json:"total_fat"`
	TotalCarbohydrate float64    `gorm:"type:decimal(8,2)" json:"total_carbohydrate"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (f FoodIntake) TableName() string {
	return "food_intakes"
}
