package model

import "gorm.io/gorm"

type SecondCategory struct {
	gorm.Model
	SecondCategory string `gorm:"unique"`
	Text           string
}
