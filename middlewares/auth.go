package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func Auth(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization != "666" {
		c.AbortWithStatusJSON(http.StatusForbidden, "请登录")
	}
}
