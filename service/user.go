package service

import (
	"chat-demo/model"
	"chat-demo/serializer"
)

type UserRegisterService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

func (service *UserRegisterService) Register() serializer.Response {
	var user model.User
	code := 200
	count := 0
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count != 0 {
		code = 400
		return serializer.Response{
			Status: code,
			Msg:    "用户名已存在。",
		}
	}
	user = model.User{
		UserName: service.UserName,
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "密码加密出错。",
		}
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "用户创建成功。",
	}
}
