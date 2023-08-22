package model

import "gorm.io/gorm"

type SmsCode struct {
	gorm.Model
	Phone      string `gorm:"type varchar(11)"`
	UserId     uint
	Code       string `gorm:"type varchar(6)"`
	ExpireTime int64
}
