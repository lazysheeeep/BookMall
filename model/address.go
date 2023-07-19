package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId   uint   `gorm:"not null"`
	Name     string `gorm:"type varchar(20) not null"`
	Province string `gorm:"not null"`
	City     string `gorm:"not null"`
	Area     string `gorm:"not null"`
	Street   string `gorm:"street"`
	Detail   string `gorm:"type varchar(50) not null"`
	Phone    string `gorm:"type varchar(11) not null"`
}
