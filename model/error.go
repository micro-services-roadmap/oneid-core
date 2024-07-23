package model

const defaultCode = 001_001_001

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"err"`
}

func (e *CodeError) Raw() error {
	return e.Err
}

func (e *CodeError) Error() string {
	if len(e.Msg) != 0 {
		return e.Msg
	}

	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Msg
}

func (e *CodeError) Resp() *Response {
	return &Response{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func NewRawError(code int, err error) error {
	return &CodeError{Code: code, Msg: "", Err: err}
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg, Err: nil}
}

func NewError(code int, msg string, err error) error {
	return &CodeError{Code: code, Msg: msg, Err: err}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}
