package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDbClient(ctx)}
}

func (dao *UserDao) UserExistOrNotByName(username string) (user model.User, exist bool, err error) {
	var count int64
	err = dao.Model(&model.User{}).Where("user_name=?", username).Find(&user).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user model.User) error {
	err := dao.Model(&model.User{}).Create(&user).Error
	return err
}

func (dao *UserDao) GetUserByUserId(ID uint) (user model.User, err error) {
	err = dao.Model(&model.User{}).Where("id=?", ID).Find(&user).Error
	if err != nil {
		return
	}
	return user, nil
}

func (dao *UserDao) UpdateUser(uId uint, user model.User) error {
	return dao.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}
