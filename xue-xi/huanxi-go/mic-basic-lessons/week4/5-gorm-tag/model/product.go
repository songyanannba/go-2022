package model

type Food struct {
	FoodId uint `gorm:"primarykey"`
	Name   uint `gorm:"column:food_name;type:varchar(64);index:idx_food_name,unique"`
}
