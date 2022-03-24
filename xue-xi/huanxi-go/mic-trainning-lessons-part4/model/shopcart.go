package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID int32 `gorm:"primary_key"`
	CreateAt *time.Time `gorm:"column:add_time"`
	UpdateAt *time.Time `gorm:"column:update_time"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

type ShopCart struct {
	BaseModel
	AccountId int32 `gorm:"type:int;index"`
	ProductId int32 `gorm:"type:int;index"`
	Num int `gorm:"int"`
	Checked bool
}