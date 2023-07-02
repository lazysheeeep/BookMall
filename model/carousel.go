package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	BookId  uint `gorm:"not null"`
	ImgPath string
}
