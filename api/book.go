package api

import (
	"BookMall/pkg/util"
	"BookMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBook(c *gin.Context) {
	form, _ := c.MultipartForm()
	file := form.File["file"]
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	createBook := service.BookService{}
	err = c.ShouldBind(&createBook)
	if err == nil && claims != nil {
		res := createBook.Create(c.Request.Context(), file, claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
