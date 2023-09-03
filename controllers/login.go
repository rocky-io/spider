package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "Login")
}

func DoLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "DoLogin")
}
