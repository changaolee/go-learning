package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	// 注册一个 v1 的路由组
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {

			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
