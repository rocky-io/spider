package middlewares

import (
	"fmt"

	"github.com/unrolled/secure"

	"github.com/gin-gonic/gin"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		//3.当接收到请求后中间件执行,封装一个secure
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			//如果出现错误，请不要继续。
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
