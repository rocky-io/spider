package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rockyskate/spider/middlewares"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.Cors())
	InitWeb(r)
}
