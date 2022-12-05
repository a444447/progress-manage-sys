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
	ErrorUserExisted   = 1003
	ErrorUserNoRight   = 1004
	ErrorTokenNotExist = 1005
	ErrorTokenRuntime  = 1006
	ErrorTokenFmt      = 1007
	ErrorLoginParams   = 1008

	// 定义毕业论文模块的错误码, 2000开头

	//定义

)

var codeMsg = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "FAIL",
	ErrorValidate:      "validator不通过",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "不存在的用户名",
	ErrorUserExisted:   "用户已经存在",
	ErrorUserNoRight:   "用户权限不匹配",
	ErrorTokenNotExist: "token不存在",
	ErrorTokenRuntime:  "token过期",
	ErrorTokenFmt:      "token格式错误",
	ErrorLoginParams:   "登陆传递参数有误",
}

// GetErrMsg 实现错误码与对应文本的转换
func GetErrMsg(code int) string {
	return codeMsg[code]
}

func ErrTrans(err error) string {
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func ErrDataTrans(data interface{}, err error) interface{} {
	if err != nil {
		return nil
	}
	return data
}
