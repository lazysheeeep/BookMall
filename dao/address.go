package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDbClient(ctx)}
}

func (dao *AddressDao) Create(address model.Address) error {
	err := dao.Model(&model.Address{}).Create(&address).Error
	return err
}

func (dao *AddressDao) Update(address model.Address, id uint) error {
	err := dao.Model(&model.Address{}).Where("id=?", id).Updates(&address).Error
	return err
}

func (dao *AddressDao) Delete(address model.Address) error {
	err := dao.Model(&model.Address{}).Delete(&address).Error
	return err
}

func (dao *AddressDao) GetAddressById(id uint) (address model.Address, err error) {
	err = dao.Model(&model.Address{}).Where("id=?", id).Find(&address).Error
	return
}

func (dao *AddressDao) GetAddress(uId uint) (addresses []model.Address, err error) {
	err = dao.Model(&model.Address{}).Where("user_id=?", uId).Order("created_at desc").Find(&addresses).Error
	return
}
