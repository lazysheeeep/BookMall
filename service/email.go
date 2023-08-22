package service

import (
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
	Code string `json:"code" form:"code"`
}

func (service *EmailService) Send(ctx context.Context, uId uint) serializer.Response {
	var err error
	var email model.Email
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

	email.UserId = uId
	email.Code = messageCode
	email.Email = service.Email
	email.ExpireTime = time.Now().Unix() + 120
	emailDao := dao.NewEmailDao(ctx)
	err = emailDao.CreateEmail(email)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *EmailCheckService) Check(ctx context.Context, uId uint) serializer.Response {
	var user model.User
	var email model.Email
	var nilEmail model.Email
	var err error

	code := e.Success

	emailDao := dao.NewEmailDao(ctx)
	email, err = emailDao.GetEmailByCodeAndId(service.Code, uId)

	if email == nilEmail {
		code = e.ErrorCode
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if email.ExpireTime < time.Now().Unix() {
		code = e.ErrorCheckCodeTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)
	user.Email = email.Email
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
