package model

import "encoding/json"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

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
