package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	router := gin.Default()

	// 定义一个GET请求的路由规则
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	router.Run()
}
