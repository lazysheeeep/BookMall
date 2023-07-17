package api

import (
	"BookMall/pkg/util"
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorite(c *gin.Context) {
	createFavorite := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createFavorite); err == nil && claims != nil {
		res := createFavorite.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ShowFavorite(c *gin.Context) {
	showFavorite := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showFavorite); err == nil && claims != nil {
		res := showFavorite.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavorite := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavorite); err == nil && claims != nil {
		res := deleteFavorite.Delete(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
