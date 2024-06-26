package e

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "参数错误",
	ErrorDao:                   "数据库错误",
	ErrorRedis:                 "redis缓存错误",
	ErrorUserHasExist:          "用户名已经存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorFailCreateUser:        "创建用户失败",
	ErrorUserNotExist:          "用户名不存在",
	ErrorPassword:              "用户密码错误",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeOut: "token过期",
	ErrorUpLoadAvatarToStatic:  "上传头像到本地失败",
	ErrorOperationType:         "操作数选择错误",
	ErrorBookNotExist:          "没有找到书本",
	ErrorCreateDaoFavorite:     "创建收藏失败",
	ErrorNoneFavorite:          "收藏夹为空",
	ErrorChangeCart:            "书本数量错误",
	ErrorOrderNone:             "订单数目为零",
	ErrorPayTime:               "订单已超时，请重新下单",
	ErrorMoneyNotEnough:        "用户余额不足",
	ErrorPhoneNum:              "手机号错误，请重新输入",
	ErrorApi:                   "Api接口错误",
	ErrorCode:                  "验证码错误",
	ErrorCheckCodeTime:         "验证码超时，请重新发送验证码",
	ErrorSendEmail:             "发送邮件失败",
	ErrorUploadPictureToQiniu:  "上传图片到七牛云失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
