## 目前的模块问题
+ 下载模块：基本上传、下载、删除、更新功能完成，但是细节没有优化
+ 任务模块: 任务模块的deadline在数据库中的存储形式没有完成/
+ 其他: 
  + 由于golang会对所有的变量默认有零值，因此即使我们gorm设置为可以为null
但是依然在插入的时候会有零值(比如，设置attatchment->int 可以为空，但是
实际上，会附上0，这样的话就会触发外键约束)
  + 日志模块
  + validator重构 
  + 定义gorm的scope实现分页功能