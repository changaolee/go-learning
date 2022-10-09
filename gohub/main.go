package main

import (
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	if err := router.Run(":8000"); err != nil {
		panic(err.Error())
	}
}
