package route

import (
	"BookMall/api"
	"BookMall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousel", api.ListCarousel)
		v1.GET("firstCategory", api.ListCategory)

		//展示商品
		v1.GET("book", api.ListBook)
		//搜索商品
		v1.POST("search", api.SearchBook)

		authed := v1.Group("/") //认证保护
		authed.Use(middleware.JWT())
		{
			//用户模块
			//更新昵称
			authed.PUT("update", api.UserUpdate)
			//上传头像
			authed.POST("upload", api.UserUploadAvatar)

			//绑定手机号码
			//发送验证码
			authed.POST("sendSms", api.SendSms)
			//绑定验证码
			authed.POST("checkSms", api.CheckSms)
			//绑定邮箱 都没写 找个机会去研究一下 小生凡一写的有点奇怪
			authed.GET("email", api.SendEmail)
			authed.POST("email", api.CheckEmail)

			//创建商品
			authed.POST("book", api.CreateBook)

			//余额模块
			//余额充值
			authed.POST("money", api.Recharge)
			//余额展示
			authed.GET("money", api.ShowMoney)

			//收藏夹模块
			//创建收藏夹
			authed.POST("favorite", api.CreateFavorite)
			//展示收藏夹
			authed.GET("favorite", api.ShowFavorite)
			//删除收藏
			authed.DELETE("favorite", api.DeleteFavorite)

			//地址模块
			//创建地址
			authed.POST("address", api.CreateAddress)
			//修改地址
			authed.POST("address/:id", api.UpdateAddress)
			//地址展示
			authed.GET("address", api.ShowAddress)
			//删除地址
			authed.DELETE("address/:id", api.DeleteAddress)

			//购物车模块
			//创建购物车
			authed.POST("cart", api.CreateCart)
			//展示购物车
			authed.GET("cart", api.ShowCart)
			//修改购物车信息
			authed.POST("cart/:id", api.ChangeCart)
			//删除商品
			authed.DELETE("cart/:id", api.DeleteCart)

			//订单模块
			//创建订单
			authed.POST("orders", api.CreateOrder)
			//获取订单详情
			authed.GET("orders/:id", api.ShowOrder)
			//展示订单
			authed.GET("orders", api.GetOrder)
			//修改订单信息
			authed.DELETE("orders/:id", api.DeleteOrder)

			//支付模块
			//订单支付
			authed.POST("payment", api.OrderPay)
		}
	}
	return r
}
