package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/pkg/util"
	"BookMall/serializer"
	"context"
	"mime/multipart"
)

type UserService struct {
	UserName string `json:"user_name" form:"user_name"`
	NickName string `json:"nick_name" form:"nick_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	code := e.Success
	var user *model.User
	if service.Key == "" || len(service.Key) != 16 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    "秘钥长度不足",
		}
	}

	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.UserExistOrNotByName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	if exist {
		code = e.ErrorUserHasExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	user = &model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("0"), //初始余额为0
	}

	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//创建用户
	if err := userDao.CreateUser(user); err != nil {
		code = e.ErrorFailCreateUser
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

func (service *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	//查找用户
	user, exist, err := userDao.UserExistOrNotByName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	if !exist {
		code = e.ErrorUserNotExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	//校验密码
	flag := user.CheckPassword(service.Password)
	if !flag {
		code = e.ErrorPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	//签发token
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "token签发失败",
			Err:    err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenDate{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}

func (service *UserService) Update(ctx context.Context, ID uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	//根据userName得到user,claims是在登录是生成的,所以id肯定是存在的
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(ID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	//更新user的nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUser(ID, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (service *UserService) Upload(ctx context.Context, file multipart.File, ID uint) serializer.Response {
	code := e.Success
	var user *model.User
	var err error

	//找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(ID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	//保存图片到本地
	filePath, err := UploadAvatarToLocalStatic(file, ID, user.UserName)
	if err != nil {
		code = e.ErrorUpLoadAvatarToStatic
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	user.Avatar = filePath
	err = userDao.UpdateUser(ID, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
