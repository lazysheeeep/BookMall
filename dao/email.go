package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type EmailDao struct {
	*gorm.DB
}

func NewEmailDao(ctx context.Context) EmailDao {
	return EmailDao{NewDbClient(ctx)}
}

func (dao *EmailDao) CreateEmail(email model.Email) error {
	return dao.Model(&model.Email{}).Create(&email).Error
}

func (dao *EmailDao) GetEmailByCodeAndId(code string, uId uint) (email model.Email, err error) {
	err = dao.Model(&model.Email{}).Where("code=? AND user_id=?", code, uId).Find(&email).Error
	return
}
