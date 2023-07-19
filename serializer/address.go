package serializer

import (
	"BookMall/model"
)

type AddressVO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Street   string `json:"street"`
	Detail   string `json:"detail"`
	Phone    string `json:"phone"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(address model.Address) AddressVO {
	return AddressVO{
		ID:       address.ID,
		Name:     address.Name,
		Province: address.Province,
		City:     address.City,
		Area:     address.Area,
		Street:   address.Street,
		Detail:   address.Detail,
		Phone:    address.Phone,
		Seen:     false,
		CreateAt: address.CreatedAt.Unix(),
	}
}

func BuildAddresses(addresses []model.Address) (items []AddressVO) {
	for _, item := range addresses {
		tmp := BuildAddress(item)
		items = append(items, tmp)
	}
	return
}
