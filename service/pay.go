package service

import (
	"BookMall/dao"
	"BookMall/pkg/e"
	"BookMall/pkg/util"
	"BookMall/serializer"
	"context"
	"fmt"
	"strconv"
	"time"
)

type OrderPay struct {
	OrderId  uint   `json:"order_id" form:"order_id"`
	Money    uint   `json:"money" form:"money"`
	OrderNo  string `json:"order_no" form:"order_no"`
	BookId   uint   `json:"book_id" form:"book_id"`
	PayTime  string `json:"pay_time" form:"pay_time"`
	Sign     string `json:"sign" form:"sign"`
	BossId   uint   `json:"boss_id" form:"boss_id"`
	BossName string `json:"boss_name" form:"boss_name"`
	Num      uint   `json:"num" form:"num"`
	Key      string `json:"key" form:"key"`
}

func (service *OrderPay) Pay(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderId(service.OrderId)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//检查订单超时时间
	if order.ExpiredTime < float64(time.Now().Unix()) {
		code = e.ErrorPayTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}

	money := order.Money * float64(service.Num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByUserId(uId)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//对钱解密，减去钱，然后重新加密存入
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat-money < 0.0 {
		code := e.ErrorMoneyNotEnough
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(userMoney)
	err = userDao.UpdateUser(uId, user)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//boss余额增加对应金额
	boss, _ := userDao.GetUserByUserId(service.BossId)
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	bossMoney := fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(bossMoney)
	err = userDao.UpdateUser(service.BossId, boss)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//商品数量减少
	bookDao := dao.NewBookDao(ctx)
	book, _ := bookDao.GetBookById(service.BookId)
	book.Num = book.Num - int(service.Num)
	err = bookDao.UpdateBook(service.BookId, book)

	//删除订单
	order.State = 1
	err = orderDao.DeleteOrder(order)

	//在这里小生凡一创造了一个temp对象并且写入了数据库，但是我暂时并没有看懂这一步操作
	//如果后面反应过来了他在干嘛请在后面添加
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
