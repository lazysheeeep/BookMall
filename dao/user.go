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
