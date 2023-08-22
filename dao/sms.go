package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type SmsDao struct {
	*gorm.DB
}

func NewSmsDao(ctx context.Context) SmsDao {
	return SmsDao{NewDbClient(ctx)}
}

func (dao *SmsDao) CreateSms(sms model.SmsCode) error {
	return dao.Model(&model.SmsCode{}).Create(&sms).Error
}

func (dao *SmsDao) GetSmsByCode(code string) (sms model.SmsCode, err error) {
	err = dao.Model(&model.SmsCode{}).Where("code=?", code).Find(&sms).Error
	return
}

func (dao *SmsDao) UpdateSms(sId uint, sms model.SmsCode) error {
	return dao.Model(&model.SmsCode{}).Where("id=?", sId).Updates(&sms).Error
}
