package main

import (
	"github.com/gin-gonic/gin"
	router "github.com/rockyskate/spider/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8090")
}
