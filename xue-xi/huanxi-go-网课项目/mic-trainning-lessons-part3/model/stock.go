package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID       int32 `gorm:"primary_key;type:int"`
	CreateAt *time.Time
	UpdateAt *time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

type Stock struct {
	BaseModel
	ProductId int32 `gorm:"type:int;index"`
	Num int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"`
}
