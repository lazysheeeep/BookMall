package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
)

type AddressService struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Province string `json:"province" form:"province"`
	City     string `json:"city" form:"city"`
	Area     string `json:"area" form:"area"`
	Street   string `json:"street" form:"street"`
	Detail   string `json:"detail" form:"detail"`
	Phone    string `json:"phone" form:"phone"`
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

func (service *AddressService) Update(ctx context.Context, uId uint) serializer.Response {
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
	address, err = addressDao.GetAddressById(service.Id)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	address.Name = service.Name
	address.Province = service.Province
	address.City = service.City
	address.Area = service.Area
	address.Street = service.Street
	address.Detail = service.Detail
	address.Phone = service.Phone

	err = addressDao.Update(address, service.Id)
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

func (service *AddressService) Delete(ctx context.Context) serializer.Response {
	code := e.Success
	var err error
	var address model.Address

	addressDao := dao.NewAddressDao(ctx)
	address, _ = addressDao.GetAddressById(service.Id)
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
