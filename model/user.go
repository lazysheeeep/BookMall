package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	NickName       string
	Email          string
	Status         string
	Avatar         string
	Money          string
}
