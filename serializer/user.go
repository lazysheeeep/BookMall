package serializer

import (
	"BookMall/config"
	"BookMall/model"
)

type UserVO struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
}

func BuildUser(user model.User) UserVO {
	return UserVO{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   config.Host + config.HttpPort + config.AvatarPath + user.Avatar,
	}
}
