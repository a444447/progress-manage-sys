package baseTime

import (
	"database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
)

//baseTime包是用来实现我们自己的时间格式的包
//需要实现的方法包括:
// 1.UnmarshalJSON:将json传入的字符串转换我们的baseTime结构体
// 2.MarshalJSON：将我们的baseTime结构体转换为json,这样我们才能以json字符串发送我们的结构体
// 3.Value()与Scan():实现这两个方法才能从数据库中读入与写入
//   Value返回的数据是我们要写入数据库的，而Scan中，参数v是我们从数据库读入的

type BaseTime time.Time

func (t *BaseTime) UnmarshalJSON(data []byte) error {
	//data是我们前端发来的时间字符串
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	//要去除接受的str末尾的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = BaseTime(t1)
	return err
}

func (t BaseTime) MarshalJSON() ([]byte, error) {
	formatTime := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatTime), nil
}

func (t *BaseTime) Scan(v interface{}) error {
	//在Value()中，我们以string的形式存储数据, 所以如果正确的话，v也会是字符串，我们需要将字符串v转换为自定义时间类型
	switch vt := v.(type) {
	case string:
		//字符串转换为time.Time
		tTime, _ := time.Parse("2006-01-02 15:04:05", vt)
		*t = BaseTime(tTime)
	default:
		return errors.New("错误的类型")
	}
	return nil
}

func (t BaseTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *BaseTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
