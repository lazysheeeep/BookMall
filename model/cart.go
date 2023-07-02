package model

type Cart struct {
	UserId uint `gorm:"not null"`
	BookId uint `gorm:"not null"`
	BossId uint `gorm:"not null"`
	Num    uint `gorm:"not null"`
	MaxNum uint `gorm:"not null"`
	Check  bool
}
