package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	BookId    uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderNum  uint64
	State     uint
	Money     float64
}
