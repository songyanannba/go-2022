package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	FoodId uint `gorm:"primarykey"`
	Name   string `gorm:"column:food_name;type:varchar(64);index:idx_food_name,unique"`
}
