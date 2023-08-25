package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/pkg/util"
	"BookMall/serializer"
	"context"
	"fmt"
	"strconv"
)

type MoneyService struct {
	Money string `json:"money" form:"money"`
	Key   string `json:"key" form:"key"`
}

func (service *MoneyService) Recharge(ctx context.Context, uId uint) serializer.Response {
	var user model.User
	var err error

	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	if service.Key == "" || len(service.Key) != 16 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    "秘钥长度不足",
		}
	}
	util.Encrypt.SetKey(service.Key)
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	userMoney, _ := strconv.ParseFloat(moneyStr, 64)
	tempMoney, _ := strconv.ParseFloat(service.Money, 64)
	money := userMoney + tempMoney
	lastMoney := fmt.Sprintf("%f", money)
	user.Money = util.Encrypt.AesEncoding(lastMoney)

	err = userDao.UpdateUser(uId, user)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *MoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	var user model.User
	var err error

	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)

	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMoney(user, service.Key),
	}
}
