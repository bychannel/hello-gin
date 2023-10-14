package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 定义路由和处理程序
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "current date:" + time.Now().Format(time.RFC3339),
		})
	})

	// 运行应用程序
	r.Run()
}
