package controllers

import (
	"net/http"

	"book_library/data"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, data.Books)

}
