package api

import (
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(c *gin.Context) {
	listCategory := service.CategoryService{}
	err := c.ShouldBind(&listCategory)
	if err == nil {
		res := listCategory.ListCategory(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
