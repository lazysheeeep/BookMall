package model

import "gorm.io/gorm"

type SmsCode struct {
	gorm.Model
	Phone      string `gorm:"type varchar(11)"`
	Code       string `gorm:"type varchar(6)"`
	ExpireTime int64
	State      int //0未过期 1已过期或已验证
}
