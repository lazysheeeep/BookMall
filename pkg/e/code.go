package e

var (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	//用户模块
	ErrorUserHasExist          = 50001
	ErrorFailEncryption        = 50002
	ErrorFailCreateUser        = 50003
	ErrorUserNotExist          = 50004
	ErrorPassword              = 50005
	ErrorAuthToken             = 50006
	ErrorAuthCheckTokenTimeOut = 50007
)
