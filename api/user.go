package api

import (
	"chat-demo/service"
	"github.com/gin-gonic/gin"
	loggin "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		loggin.Info(err)
	}
}
