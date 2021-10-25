package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

// UnmarshalJSON  为 Time 重写 MarshaJSON 和 UnmarshalJSON 方法，在此方法中实现自定义格式的转换；
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = Time{t1}
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value  为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan  为 Time 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
