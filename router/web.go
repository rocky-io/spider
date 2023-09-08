package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rockyskate/spider/controllers"
	"github.com/rockyskate/spider/middlewares"
)

func InitWeb(r *gin.Engine) {
	r.Any("/", controllers.Index)
	r.GET("/login", controllers.Login)
	r.POST("/dologin", controllers.DoLogin)
	user := r.Group("/user", middlewares.Auth)
	user.Any("show", controllers.UserShow)
	user.GET("gorm", controllers.UserAdd)
	user.GET("prototest", controllers.Prototest)

}
