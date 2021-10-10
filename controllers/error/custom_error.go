package cerror

import "net/http"

type IError interface {
	Error() string
	ErrorCode() int32
	ErrorMsg() string
	StatusCode() int
}

type TError struct {
	errCode    int32
	statusCode int
	defaultMsg string
}

func (e *TError) Error() string {
	return e.defaultMsg
}

func (e *TError) ErrorCode() int32 {
	return e.errCode
}

func (e *TError) ErrorMsg() string {
	return e.defaultMsg
}

func (e *TError) StatusCode() int {
	return e.statusCode
}

func NewError(errCode int32, msg string, statusCode ...int) IError {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return &TError{
		errCode:    errCode,
		statusCode: code,
		defaultMsg: msg,
	}
}
