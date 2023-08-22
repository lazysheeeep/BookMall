package service

import (
	"BookMall/config"
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"math/rand"
	"time"
)

type SmsService struct {
	Phone string `json:"phone" form:"phone"`
}

type SmsCheckService struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

func (service *SmsService) Send(ctx context.Context) serializer.Response {
	var smsMode model.SmsCode
	var err error

	code := e.Success

	if len(service.Phone) != 11 {
		code = e.ErrorPhoneNum
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	smsCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	credential := common.NewCredential(config.SecretId, config.SecretKey)
	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()

	/* SDK默认使用POST方法。
	 * 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求 */
	cpf.HttpProfile.ReqMethod = "POST"

	/* SDK有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	// cpf.HttpProfile.ReqTimeout = 5

	/* 指定接入地域域名，默认就近地域接入域名为 sms.tencentcloudapi.com ，也支持指定地域域名访问，例如广州地域的域名为 sms.ap-guangzhou.tencentcloudapi.com */
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	/* SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段 */
	cpf.SignMethod = "HmacSHA1"

	/* 实例化要请求产品(以sms为例)的client对象
	 * 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，支持的地域列表参考 https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8 */
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	var templateParams []string
	templateParams = append(templateParams, smsCode)
	templateParams = append(templateParams, "2")

	request := sms.NewSendSmsRequest()
	//*短信应用ID
	request.SmsSdkAppId = common.StringPtr(config.SdkAppId)
	//*短信签名内容
	request.SignName = common.StringPtr(config.SignName)
	//*模板ID
	request.TemplateId = common.StringPtr(config.TemplateId)
	//*模板参数
	request.TemplateParamSet = common.StringPtrs(templateParams)
	//*下发手机号码
	request.PhoneNumberSet = common.StringPtrs([]string{service.Phone})

	//*通过client对象调用想要访问的接口，需要传入对象
	response, err := client.SendSms(request)

	//*处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		code = e.ErrorApi
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//非SDK异常
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	smsMode.Phone = service.Phone
	smsMode.Code = smsCode
	smsMode.ExpireTime = time.Now().Unix() + 120
	smsMode.State = 0

	smsDao := dao.NewSmsDao(ctx)
	err = smsDao.CreateSms(smsMode)
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
		Data:   response.Response,
	}
}

func (service *SmsCheckService) Check(ctx context.Context, uId uint) serializer.Response {
	var smsMode model.SmsCode
	var user model.User
	var err error

	code := e.Success
	nilSms := model.SmsCode{}

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)

	smsDao := dao.NewSmsDao(ctx)
	smsMode, err = smsDao.GetSmsByCode(service.Code)

	if smsMode == nilSms {
		code = e.ErrorCode
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if time.Now().Unix() > smsMode.ExpireTime {
		code = e.ErrorCheckCodeTime
		smsMode.State = 1
		err = smsDao.UpdateSms(smsMode.ID, smsMode)
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

	user.Mobile = service.Phone
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
