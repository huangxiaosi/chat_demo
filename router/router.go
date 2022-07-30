package router

import (
	"chat-demo/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//conf.Init()
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger()) //异常恢复和日志 中间件。
	v1 := r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "SUCCESS")
		})
		v1.POST("user/register", api.UserRegister)
	}

	return r

}
