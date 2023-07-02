package model

import "gorm.io/gorm"

type FirstCategory struct {
	gorm.Model
	FirstCategory string `gorm:"unique"`
	Text          string
}
