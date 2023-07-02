package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	Avatar         string
}
