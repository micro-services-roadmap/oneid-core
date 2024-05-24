package model

const defaultCode = 001_001_001

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Resp() *Response {
	return &Response{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}
