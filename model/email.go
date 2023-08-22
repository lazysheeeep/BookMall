package model

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	UserId     uint
	Email      string
	Code       string
	ExpireTime int64
}
