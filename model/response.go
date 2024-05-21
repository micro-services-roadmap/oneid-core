package model

import (
	"encoding/json"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Res(code int, data interface{}, msg string) *Response {
	return &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}
func Ok() *Response {
	return Res(SUCCESS, map[string]interface{}{}, "Success")
}

func OkWithMessage(message string) *Response {
	return Res(SUCCESS, map[string]interface{}{}, message)
}

func OkWithData(data interface{}) *Response {
	return Res(SUCCESS, data, "Success")
}

func OkWithDetailed(data interface{}, message string) *Response {
	return Res(SUCCESS, data, message)
}

func Failed() *Response {
	return Res(ERROR, map[string]interface{}{}, "Fail")
}

func FailWithMessage(message string) *Response {
	return Res(ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(data interface{}, message string) *Response {
	return Res(ERROR, data, message)
}

func FailWithError(data error) *Response {
	return FailWithDetailed(data.Error(), "Failed")
}

func Result(code int, data interface{}, msg string) []byte {
	r := &Response{
		code,
		data,
		msg,
	}

	ser, _ := json.Marshal(r)
	return ser
}

func Fail(data interface{}) []byte {
	return Result(ERROR, data, "操作失败")
}

func Success(data interface{}) []byte {
	return Result(SUCCESS, data, "操作成功")
}
