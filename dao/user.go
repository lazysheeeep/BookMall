package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type userDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *userDao {
	return &userDao{NewDbClient(ctx)}
}

func (dao *userDao) UserExistOrNotByName(username string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.Model(&model.User{}).Where("user_name=?", username).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *userDao) CreateUser(user *model.User) error {
	err := dao.Model(&model.User{}).Create(&user).Error
	return err
}

func (dao *userDao) GetUserByUserId(ID uint) (user *model.User, err error) {
	err = dao.Model(&model.User{}).Where("id=?", ID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *userDao) UpdateUser(uId uint, user *model.User) error {
	return dao.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}
