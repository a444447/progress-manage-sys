package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"progress-manage-system/utils/myLog"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now() //先记录开始的时刻
		c.Next()                //next()进入下一层
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := myLog.Logger.WithFields(logrus.Fields{
			"SpendTime": spendTime,
			"path":      url,
			"Method":    method,
			"status":    statusCode,
		})
		if len(c.Errors) > 0 { // 创建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= 500 {
			Log.Error()
		} else if statusCode >= 400 {
			Log.Warn()
		} else {
			Log.Info()
		}
	}
}
