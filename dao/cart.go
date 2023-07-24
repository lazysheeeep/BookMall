package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDbClient(ctx)}
}

func (dao *CartDao) GetCart(uId uint, bookId uint, bossId uint) (cart model.Cart, err error) {
	err = dao.Model(&model.Cart{}).Where("user_id=? AND book_id=? AND boss_id=?", uId, bookId, bossId).Find(&cart).Error
	return
}

func (dao *CartDao) GetCartById(id uint) (cart model.Cart, err error) {
	err = dao.Model(&model.Cart{}).Where("id=?", id).Find(&cart).Error
	return
}

func (dao *CartDao) Create(cart model.Cart) error {
	err := dao.Model(&model.Cart{}).Create(&cart).Error
	return err
}

func (dao *CartDao) Update(cId uint, cart model.Cart) error {
	err := dao.Model(&model.Cart{}).Where("id=?", cId).Updates(&cart).Error
	return err
}

func (dao *CartDao) Show(uId uint) (cart []model.Cart, err error) {
	err = dao.Model(&model.Cart{}).Where("user_id=?", uId).Find(&cart).Error
	return
}

func (dao *CartDao) Delete(item model.Cart) error {
	err := dao.Model(&model.Cart{}).Delete(&item).Error
	return err
}
