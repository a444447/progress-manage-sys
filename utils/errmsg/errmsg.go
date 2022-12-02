package errmsg

/*
此包主要用来处理错误，定义返回码的含义
*/
const (
	SUCCESS       = 200
	ERROR         = 500
	ErrorValidate = 10000
	// 定义用户模块的错误码, 1000开头
	ErrorPasswordWrong = 1001
	ErrorUserNotExist  = 1002
	ErrorUserNoRight   = 1003
	ErrorTokenNotExist = 1004
	ErrorTokenRuntime  = 1005
	ErrorTokenFmt      = 1006

	// 定义毕业论文模块的错误码, 2000开头

	//定义

)

var codeMsg = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "FAIL",
	ErrorValidate:      "validator不通过",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "不存在的用户名",
	ErrorUserNoRight:   "用户权限不匹配",
	ErrorTokenNotExist: "token不存在",
	ErrorTokenRuntime:  "token过期",
	ErrorTokenFmt:      "token格式错误",
}

// GetErrMsg 实现错误码与对应文本的转换
func GetErrMsg(code int) string {
	return codeMsg[code]
}
