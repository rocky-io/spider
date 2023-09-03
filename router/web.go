package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rockyskate/spider/controllers"
	"github.com/rockyskate/spider/middlewares"
)

func InitWeb(r *gin.Engine) {
	r.GET("/", controllers.Index)
	r.GET("/login", controllers.Login)
	r.POST("/dologin", controllers.DoLogin)
	user := r.Group("/").Use(middlewares.Auth)
	user.GET("user/show", controllers.Index)
}
