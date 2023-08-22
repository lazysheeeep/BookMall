package e

var (
	Success       = 200
	Error         = 500
	InvalidParams = 400
	ErrorDao      = 600

	//用户模块
	ErrorUserHasExist          = 50001
	ErrorFailEncryption        = 50002
	ErrorFailCreateUser        = 50003
	ErrorUserNotExist          = 50004
	ErrorPassword              = 50005
	ErrorAuthToken             = 50006
	ErrorAuthCheckTokenTimeOut = 50007
	ErrorUpLoadAvatarToStatic  = 50008

	//书本模块
	ErrorUploadBookToStatic = 60001
	ErrorOperationType      = 60002
	ErrorBookNotExist       = 60003

	//收藏夹模块
	ErrorCreateDaoFavorite = 70001
	ErrorNoneFavorite      = 70002

	//购物车模块
	ErrorChangeCart = 80001

	//订单模块
	ErrorOrderNone = 90001
	ErrorPayTime   = 90002

	//账户模块
	ErrorMoneyNotEnough = 10001

	//验证码模块
	ErrorPhoneNum      = 20001
	ErrorApi           = 20002
	ErrorCode          = 20003
	ErrorCheckCodeTime = 20004

	//邮箱绑定模块
	ErrorSendEmail = 30001
)
