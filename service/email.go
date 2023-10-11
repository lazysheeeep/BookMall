package service

import (
	"BookMall/cache"
	"BookMall/config"
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
	"fmt"
	"gopkg.in/mail.v2"
	"math/rand"
	"time"
)

type EmailService struct {
	Email string `json:"email" form:"email"`
}

type EmailCheckService struct {
	Email string `json:"email" form:"email"`
	Code  string `json:"code" form:"code"`
}

func (service *EmailService) Send(ctx context.Context, uId uint) serializer.Response {
	var err error
	code := e.Success

	//生成验证码
	messageCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().Unix())).Int31n(1000000))

	m := mail.NewMessage()
	m.SetHeader("From", config.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "邮箱验证")
	m.SetBody("text/html", "验证码："+messageCode)

	d := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	//在此处开始将生成的code写入mysql，现在使用redis代替
	err = cache.AddEmailCode(ctx, uId, service.Email, messageCode)
	if err != nil {
		code = e.ErrorRedis
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *EmailCheckService) Check(ctx context.Context, uId uint) serializer.Response {
	var user model.User
	var err error

	code := e.Success

	result := cache.GetEmailCode(ctx, uId, service.Email)
	//验证码超时的情况
	if result == "" {
		code = e.ErrorCheckCodeTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//验证码错误的情况
	if service.Code != result {
		code = e.ErrorCode
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)
	user.Email = service.Email
	err = userDao.UpdateUser(uId, user)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
