package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// message := c.PostForm("message")
	// nick := c.DefaultPostForm("nick", "anonymous")
	// c.JSON(200, gin.H{
	// 	"status":  "posted",
	// 	"message": message,
	// 	"nick":    nick,
	// })

	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./upload" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
