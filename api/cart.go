package api

import (
	"BookMall/pkg/util"
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	createCart := service.CartService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCart); err == nil && claims != nil {
		res := createCart.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ShowCart(c *gin.Context) {
	showCart := service.CartService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showCart); err == nil && claims != nil {
		res := showCart.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ChangeCart(c *gin.Context) {
	changeCart := service.CartService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&changeCart); err == nil && claims != nil {
		res := changeCart.Change(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func DeleteCart(c *gin.Context) {
	deleteCart := service.DeleteCartService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCart); err == nil && claims != nil {
		res := deleteCart.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
