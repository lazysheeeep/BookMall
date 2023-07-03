package e

var MsgFlags = map[int]string{
	Success:             "ok",
	Error:               "fail",
	InvalidParams:       "参数错误",
	ErrorUserHasExist:   "用户名已经存在",
	ErrorFailEncryption: "密码加密失败",
	ErrorFailCreateUser: "创建用户失败",
	ErrorUserNotExist:   "用户名不存在",
	ErrorPassword:       "用户密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
