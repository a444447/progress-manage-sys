package ecode

var (
	Ok = add(200, "ok")

	//服务端错误
	ErrSqlInit   = add(10001, "数据库初始化失败")
	ErrCreate    = add(10002, "数据库添加出错")
	ErrUpdate    = add(10003, "数据库更新出错")
	ErrDelete    = add(10004, "数据库删除出错")
	ErrFind      = add(10005, "数据库查找出错")
	ErrMapData   = add(10006, "解析form-data或者json出错")
	ErrParamData = add(10007, "解析url中的参数失败")

	//用户模块
	ErrUserNotFound  = add(20101, "登陆失败:用户名不存在")
	ErrPasswordWrong = add(20102, "密码错误")
	ErrUserExisted   = add(20103, "添加用户失败:用户已经存在")
	ErrValidFail     = add(20104, "输入格式有误")

	//论文模块
	ErrThesisStuId     = add(20201, "添加失败:不存在库中的学生")
	ErrThesisTeacherId = add(20202, "添加失败: 不存在库中的教师")
)
