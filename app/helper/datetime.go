package helper

import "time"

// IsNowInTimeRange 当前时间是否在指定范围内 **/
// 参数为时间字符串 格式为 "时:分:秒"
func IsNowInTimeRange(startTimeStr, endTimeStr string) bool {
	// 当前时间
	now := time.Now()
	// 当前时间转换为 "年-月-日" 的格式
	format := now.Format("2006-01-02")
	// 转换为time类型需要的格式
	layout := "2006-01-02 15:04:05"
	// 将开始时间拼接 年-月-日 转换为time类型
	startTime, _ := time.ParseInLocation(layout, format+" "+startTimeStr, time.Local)
	// 将结束时间拼接 年-月-日 转换为time类型
	endTime, _ := time.ParseInLocation(layout, format+" "+endTimeStr, time.Local)
	// 使用time的Before和After方法，判断当前时间是否在参数的时间范围
	return now.Before(endTime) && now.After(startTime)
}

// 获取当天每个小时的编排
func GetTodayHourList() (list []int) {
	hour := 24 // time.Now().Hour()
	for i := 0; i < hour; i++ {
		list = append(list, i)
	}

	return list
}

// GetDaysDateList 根据开始时间戳与天数获取时间段内每天的日期列表 /**
func GetDaysDateList(startTime int64, days int) (list []string) {
	for i := 0; i < days; i++ {
		list = append(list, time.Unix(startTime+int64(i*60*60*24), 0).Format("2006-01-02"))
	}

	return list
}
