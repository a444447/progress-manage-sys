package myLog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//本包主要是用于初始化日志器与日志示例

var (
	Logger   = logrus.New() //一个项目公用的日志对象
	LogEntry *logrus.Entry
)

func init() {
	logPath := "logs/log"                                         //日志文件存放的位置
	softLink := "logs/latest.log"                                 //最新日志文件的软连接路径(因为我们会按日期分割日志)
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755) //初始化日志文件对象
	if err != nil {
		fmt.Println(err.Error())
	}
	Logger.Out = src

	//日志分离
	Logger.SetLevel(logrus.DebugLevel) // 设置日志的等级
	logWriter, _ := rotatelogs.New(
		logPath+"%Y%m%d.log",                      //设置分割的文件名
		rotatelogs.WithMaxAge(7*24*time.Hour),     //设置最多可以保留多少天的日志，这里设置为保留7天
		rotatelogs.WithRotationTime(24*time.Hour), //设置每天保存一个日志文件
		rotatelogs.WithLinkName(softLink),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter, // info级别使用logWriter写日志
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 格式日志时间
	})
	Logger.AddHook(Hook)
	LogEntry = logrus.NewEntry(Logger)
}
