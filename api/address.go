package api

import (
	"BookMall/pkg/util"
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	createAddress := service.AddressService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil && claims != nil {
		res := createAddress.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateAddress(c *gin.Context) {
	updateAddress := service.AddressService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateAddress); err == nil && claims != nil {
		res := updateAddress.Update(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ShowAddress(c *gin.Context) {
	showAddress := service.AddressService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showAddress); err == nil && claims != nil {
		res := showAddress.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func DeleteAddress(c *gin.Context) {
	deleteAddress := service.AddressService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddress); err == nil && claims != nil {
		res := deleteAddress.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
