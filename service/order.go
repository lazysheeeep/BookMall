package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	UserId    uint    `json:"user_id" form:"user_id"`
	BookId    uint    `json:"book_id" form:"book_id"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	Num       uint    `json:"num" form:"num"`
	Money     float64 `json:"money" form:"money"`
	OrderNum  uint64  `json:"order_num" form:"order_num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	PageNum   int     `json:"page_num" form:"page_num"` //订单编号
	PageSize  int     `json:"page_size" form:"page_size"`
	State     uint    `json:"type" form:"type"`
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order model.Order
	var err error

	code := e.Success

	order = model.Order{
		UserId:    uId,
		BookId:    service.BookId,
		BossId:    service.BossId,
		AddressId: service.AddressId,
		Num:       service.Num,
		State:     1,
		Money:     service.Money,
	}

	//生成订单号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	bookNum := strconv.Itoa(int(service.BookId))
	userNum := strconv.Itoa(int(uId))
	number += bookNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	order.OrderNum = orderNum

	expiredTime := float64(time.Now().Unix()) + 20*time.Minute.Seconds()
	order.ExpiredTime = expiredTime

	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.Create(order)

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

func (service *OrderService) Get(ctx context.Context, uId uint) serializer.Response {
	var orders []model.Order
	var err error

	code := e.Success

	if service.PageSize == 0 {
		service.PageSize = 5
	}

	orderDao := dao.NewOrderDao(ctx)
	orders, err = orderDao.GetOrders(uId, service.PageNum, service.PageSize)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	if len(orders) == 0 {
		code = e.ErrorOrderNone
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(len(orders)))
}

func (service *OrderService) Show(ctx context.Context, aId string) serializer.Response {
	orderId, _ := strconv.Atoi(aId)
	var order model.Order
	var err error

	code := e.Success

	orderDao := dao.NewOrderDao(ctx)
	order, err = orderDao.GetOrderId(uint(orderId))

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(order.AddressId)

	bookDao := dao.NewBookDao(ctx)
	book, err := bookDao.GetBookById(order.BookId)

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
		Data:   serializer.BuildOrder(book, address, order),
	}
}

func (service *OrderService) Delete(ctx context.Context, aId string) serializer.Response {
	var order model.Order
	var err error

	code := e.Success

	id, err := strconv.Atoi(aId)

	orderDao := dao.NewOrderDao(ctx)
	order, err = orderDao.GetOrderId(uint(id))
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	err = orderDao.DeleteOrder(order)
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
