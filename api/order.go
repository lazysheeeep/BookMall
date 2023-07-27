package api

import (
	"BookMall/pkg/util"
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	createOrder := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrder); err == nil && claim != nil {
		res := createOrder.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
