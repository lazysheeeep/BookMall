package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
)

type AddressService struct {
	Name     string `json:"name" form:"name"`
	Province string `json:"province" form:"province"`
	City     string `json:"city" form:"city"`
	Area     string `json:"area" form:"area"`
	Street   string `json:"street" form:"street"`
	Detail   string `json:"detail" form:"detail"`
	Phone    string `json:"phone" form:"phone"`
}

type DeleteAddressService struct {
	Id uint `json:"id" form:"id"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address model.Address
	var err error

	code := e.Success

	address = model.Address{
		UserId:   uId,
		Name:     service.Name,
		Province: service.Province,
		City:     service.City,
		Area:     service.Area,
		Street:   service.Street,
		Detail:   service.Detail,
		Phone:    service.Phone,
	}

	addressDao := dao.NewAddressDao(ctx)
	err = addressDao.Create(address)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	var addresses []model.Address
	addresses, err = addressDao.GetAddress(uId)
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
		Data: serializer.ListData{
			List:  serializer.BuildAddresses(addresses),
			Total: uint(len(addresses)),
		},
	}
}

func (service *AddressService) Show(ctx context.Context, uId uint) serializer.Response {
	var addresses []model.Address
	var err error

	code := e.Success

	addressDao := dao.NewAddressDao(ctx)
	addresses, err = addressDao.GetAddress(uId)

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
		Data: serializer.ListData{
			List:  serializer.BuildAddresses(addresses),
			Total: uint(len(addresses)),
		},
	}
}

func (service *DeleteAddressService) Delete(ctx context.Context, id uint) serializer.Response {
	code := e.Success
	var err error
	var address model.Address

	addressDao := dao.NewAddressDao(ctx)
	address, _ = addressDao.GetAddressById(id)
	err = addressDao.Delete(address)

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
