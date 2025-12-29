package repository

import "gorm.io/gorm"

type FoodIntakeRepository interface {

}

type foodIntakeRepository struct {
	db *gorm.DB
}

func NewFoodIntakeRepository(db *gorm.DB) FoodIntakeRepository {
	return &foodIntakeRepository{db : db}
}


