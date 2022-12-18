package ecode

import "fmt"

// 定义响应的数据结构
type ResponseJSON struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(err error, args ...interface{}) *ResponseJSON {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}
	fmt.Printf("%v\n", len(args))

	ec := Cause(err)
	return &ResponseJSON{
		Code:    ec.Code(),
		Message: ec.Message(),
		Data:    data,
	}
}
