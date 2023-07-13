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

		authed := v1.Group("/") //认证保护
		authed.Use(middleware.JWT())
		{
			//更新昵称
			authed.PUT("user/update", api.UserUpdate)
			//上传头像
			authed.POST("user/upload", api.UserUploadAvatar)
			//绑定手机号码
			//绑定邮箱 都没写 找个机会去研究一下 小生凡一写的有点奇怪

			//创建商品
			authed.POST("user/book", api.CreateBook)
		}
	}
	return r
}
