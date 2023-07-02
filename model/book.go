package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN           string
	Name           string
	Author         string
	Publisher      string
	Info           string
	ImgPath        string
	Price          string
	DiscountPrice  string
	OnSale         bool
	Num            int
	FirstCategory  string //采用二级分类
	SecondCategory string
	BossId         uint
	BossName       string
	BossAvatar     string
}
