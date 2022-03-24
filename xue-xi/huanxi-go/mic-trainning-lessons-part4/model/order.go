package model

import "time"

type OrderItem struct {
	BaseModel
	AccountId int32  `gorm:"type:int;index"`
	OrderNo   string `gorm:"type:varchar(64);index""`
	PayType   string `gorm:"type:varchar(16)"` //支付方式
	Status    string `gorm:"type:varchar(16)"` //订单状态
	TradeNo   string `gorm:"type:varchar(64)"` //如 微信支付号
	Addr string `gorm:"type:varchar(64)"`
	Receiver string 	`gorm:"type:varchar(16)"`
	ReceiverModel string `gorm:"type:varchar(11)"`
	PostCode string `gorm:"type:varchar(16)"` //邮编
	OrderAmount float32  //订单总价
	PayTime *time.Time `gorm:"type:datetime"`
}


type OrderProduct struct {
	BaseModel
	OrderId int32 `gorm:"type:int;index"`
	ProductId int32 `gorm:"type:int;index"`
	ProductName string `gorm:"type:varchar(64);index"`
	CoverImage string `gorm:"type:varchar(128)"`
	RealPrice float32
	Num int32 `gorm:"type:int"`
}