package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Loginer struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		// 注意使⽤只读上下⽂
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
	c.JSON(http.StatusOK, gin.H{"data": "成功响应的数据"})
	//c.Redirect(http.StatusMovedPermanently, "http://www.fudebeisi.cn/")
	//c.HTML(http.StatusOK, "login.tmpl", gin.H{"headName": "登录", "tableName": "请输入账号密码"})
}

func DoLogin(c *gin.Context) {
	var loginer Loginer
	c.Bind(&loginer)
	if err := c.Bind(&loginer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	if loginer.Username != "rocky" || loginer.Password != "666" {

		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	c.Writer.Header().Set("token", "666666")
	c.JSON(http.StatusOK, loginer)
}
