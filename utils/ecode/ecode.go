package ecode

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	codes = map[int]struct{}{}
)

// New Error
func New(code int, msg string) *Error {
	return add(code, msg)
}

// add for only inner error
func add(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", code))
	}
	codes[code] = struct{}{}
	return &Error{
		code: code, message: msg,
	}
}

type Errors interface {
	// error以字符串形式返回code
	Error() string
	// 得到error code
	Code() int
	// 得到message
	Message() string
	// 得到error的detail，目前我们返回nil
	Details() []interface{}
	// for compatiable
	Equal(error) bool
	// Reload Message
	Reload(string) Error
}

type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Reload(message string) Error {
	e.message = message
	return e
}

func (e Error) Details() []interface{} { return nil }

func (e Error) Equal(err error) bool { return Equal(err, e) }

func Equal(err error, e Error) bool {
	return Cause(err).Code() == e.Code()
}

/*
Cause:调用errors的Cause,主要用于自定义的Error与原生的error之间进行判定
Equal:首先调用Cause,如果err不是Error类型，那么会根据String()返回一个code=500的Error
String: 如果不是Error类型，保留err.Error()的信息，放在一个code=500的Error中

*/

func Cause(err error) Errors {
	if err == nil {
		return Ok
	}
	if ec, ok := errors.Cause(err).(Errors); ok {
		return ec
	}
	return String(err.Error())

}

func String(s string) *Error {
	if s == "" {
		return Ok
	}
	return &Error{
		code: 500, message: s,
	}
}
