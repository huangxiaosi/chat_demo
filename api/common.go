package api

import (
	"chat-demo/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(validator.ValidationErrors); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "参数错误。",
			Error:  fmt.Sprint(err),
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配。",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "参数错误。",
		Error:  fmt.Sprint(err),
	}
}
