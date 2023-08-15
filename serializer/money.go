package serializer

import (
	"BookMall/model"
	"BookMall/pkg/util"
)

type MoneyVO struct {
	UserId    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserMoney string `json:"user_money"`
}

func BuildMoney(user model.User, key string) MoneyVO {
	util.Encrypt.SetKey(key)
	return MoneyVO{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money),
	}
}
