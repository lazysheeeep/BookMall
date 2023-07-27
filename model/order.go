package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId      uint `gorm:"not null"`
	BookId      uint `gorm:"not null"`
	BossId      uint `gorm:"not null"`
	AddressId   uint `gorm:"not null"`
	Num         uint
	OrderNum    uint64 //订单编号
	State       uint   //1未支付 2已支付
	Money       float64
	ExpiredTime float64
}
