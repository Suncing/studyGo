package api

import (
	"adminManager/webscoket3/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

/**
*   @author wangchenyang
*   @date 2022/8/17 10:25
*   @description:
 */

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(validator.ValidationErrors); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "参数错误",
			Error:  fmt.Sprint(err),
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "Json类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
