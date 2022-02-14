package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func Manage(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.html", nil)
}
func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie") // 获取Cookie
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	} else if cookie != "test" {
		c.Abort()
		return
	}

	c.Next()
}
