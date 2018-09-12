package common

import (
	"fmt"
	"strings"
	"time"
)

var C_format_datetime string = "2006-01-02 15:04:05"
var C_format_datemonth string = "2006-01"
var C_format_date_ymd string = "2006-01-02"

/**
* 判断两个字符串日期相差的天数
 */
func DifferDays(startDate, endDate string) int {
	//startDateStr := "2015-10-07 15:20:10"
	//endDate := time.Now().Format("2006-01-02 15:04:05")
	//将初始日期处理为0时0分0秒
	startT, _ := FormatDateCCT("2006-01-02 15:04:05", startDate)
	startStr := startT.Format("2006-01-02 00:00:00")

	//将日期字符串转成时间格式
	startStrT, _ := FormatDateCCT("2006-01-02 00:00:00", startStr)
	endT, _ := FormatDateCCT("2006-01-02 15:04:05", endDate)
	//将时间格式转成时间戳
	startUnix := startStrT.Unix()
	endUnix := endT.Unix()
	//通过Unix时间戳获取两个日期相差的天数
	//dayNum := (endUnix - startUnix) / 86400
	unixCount := int(endUnix - startUnix)
	var dayNum int

	dayNum = (unixCount) / 86400
	return dayNum
}

/**
* 统一日期格式
 */
func FormatDateString(str, strFormat, dateFormat string, start, length int) string {
	dateStr := Substr(str, start, length)
	d, _ := time.ParseDuration("+12h")
	//endDate := "2015-10-13 15:05:23"
	dateTime, _ := FormatDateCCT(strFormat, dateStr)
	if strings.Contains(str, "PM") {
		dateTime = dateTime.Add(d)
	}
	formatStr := dateTime.Format(dateFormat)
	return formatStr
}

/**
* 计算时差返回字符串
* 字符串类型:   2006-01-02 15:04:05
 */
func TimeDiffString(dateStr, timeDiff string) string {
	res := dateStr
	d, _ := time.ParseDuration(timeDiff)
	dateTime, err := FormatDateCCT("2006-01-02 15:04:05", dateStr)
	if err == nil {
		dateTime = dateTime.Add(d)
		res = dateTime.Format("2006-01-02 15:04:05")
	}
	return res
}

/*
 * 格式化字符串为时间戳
 * @date 待转化为时间戳的字符串
 * @format 转化所需模板 如 20060102
 */
func FormatStr2Unix(date string, format string) (int64, error) {
	loc, _ := time.LoadLocation("Local")                    //重要：获取时区
	theTime, err := time.ParseInLocation(format, date, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0, err
	}
	return theTime.Unix(), nil
}

/*
 * 时间戳转日期
 * @time_unix 时间戳
 * @format 转化所需模板 如 20060102
 */
func FormatUnix2Str(time_unix int64, format string) string {
	dataTimeStr := time.Unix(time_unix, 0).Format(format)
	return dataTimeStr
}

/*
 * 格式化字符串北京时间
 */
func FormatDateCCT(format, date string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")                    //重要：获取时区
	theTime, err := time.ParseInLocation(format, date, loc) //使用模板在对应时区转化为time.time类型
	return theTime, err
}

/*
 * 获取
 * 2016-11-xx 00:00:00
 * 2016-11-xx 23:59:59
 */
func GetSDateAndEDateSJC(sjc int) (string, string) {
	format_date := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)

	add_day := fmt.Sprintf("%dh", sjc)
	th_bet_dt, _ := time.ParseDuration(add_day)
	s_time := now_time.Add(th_bet_dt)
	s_date := s_time.Format(format_date)
	s_date += " 00:00:00"

	e_date := now_time.Format(format_date)
	e_date += " 23:59:59"

	return s_date, e_date
}

/*
 * 获取上周一到周天
 * 2016-11-xx 00:00:00
 * 2016-11-xx 23:59:59
 */
func GetSDateAndEDateSJCmond2sun(sjc int) (string, string) {
	format_date := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// now_time := time.Now().In(loc)
	timeLayout := "2006-01-02 15:04:05"
	monday_date := GetMondayDate() + " 00:00:00"
	theTime, _ := time.ParseInLocation(timeLayout, monday_date, loc)
	add_day := fmt.Sprintf("%dh", sjc)
	th_bet_dt, _ := time.ParseDuration(add_day)
	s_time := theTime.Add(th_bet_dt)
	s_date := s_time.Format(format_date)
	s_date += " 00:00:00"

	e_date := monday_date
	// e_date += " 23:59:59"

	return s_date, e_date
}

/*
获取现在日期时间
*/
func GetNowDatetime(format string) string {
	return time.Now().Format(format)
}

/*
* 改变日期
* @time_date 要改变的日期
* @format 转化所需模板 如 20060102
* @years 增加或减少的年
* @months 增加或减少的月
* @days 增加或减少的日
 */
func ChangeDate(time_date, format string, years int, months int, days int) string {
	time_type, _ := time.Parse(format, time_date)
	new_time_type := time_type.AddDate(years, months, days)
	new_time_date := new_time_type.Format(format)
	return new_time_date
}

/**
* 获取两个时间间隔内的所有日期
* 单位是天
 */
func GetDateDiff(time_date, format string, days int) []string {
	date_arr := []string{}
	days_int := 0
	day_i := 0
	date_str := ""
	if days < 0 {
		days_int = days * (-1)
	}
	for i := 0; i <= days_int; i++ {
		day_i = days - (days_int/days)*i
		date_str = ChangeDate(time_date, format, 0, 0, day_i)
		date_arr = append(date_arr[0:], date_str)
	}
	return date_arr
}

/**
* 获取本周周一的日期(从周一计算本周开始)
 */
func GetMondayDate() string {
	monday_date := ""
	//当前日期
	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)
	now_day := now_time.Format("2006-01-02")
	weekday := now_time.Weekday()
	weekday_str := weekday.String()
	switch weekday_str {
	case "Sunday":
		dt, _ := time.ParseDuration("-144h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Monday":
		monday_date = now_day
	case "Tuesday":
		dt, _ := time.ParseDuration("-24h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Wednesday":
		dt, _ := time.ParseDuration("-48h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Thursday":
		dt, _ := time.ParseDuration("-72h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Friday":
		dt, _ := time.ParseDuration("-96h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Saturday":
		dt, _ := time.ParseDuration("-120h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	}
	return monday_date
}

/**
* 获取本周周一的日期(从周一计算本周开始)
 */
func GetMondayDateByDate(date_str, format string) string {
	monday_date := ""
	//当前日期
	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)
	if len(date_str) > 0 {
		now_time, _ = FormatDateCCT(format, date_str)
	}
	now_day := now_time.Format("2006-01-02")
	weekday := now_time.Weekday() - time.Monday
	weekday_str := weekday.String()
	switch weekday_str {
	case "Sunday":
		dt, _ := time.ParseDuration("-144h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Monday":
		monday_date = now_day
	case "Tuesday":
		dt, _ := time.ParseDuration("-24h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Wednesday":
		dt, _ := time.ParseDuration("-48h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Thursday":
		dt, _ := time.ParseDuration("-72h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Friday":
		dt, _ := time.ParseDuration("-96h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	case "Saturday":
		dt, _ := time.ParseDuration("-120h")
		monday_date = now_time.Add(dt).Format("2006-01-02")
	}
	return monday_date
}

/*
获取本周一的日期
return 2017-04-17
*/
func GetMonday() string {

	format := "2006-01-02"

	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)
	what_day := now_time.Weekday() - time.Monday //现在星期几
	what_day_str := what_day.String()

	now_day := now_time.Format(format) //今天

	monday_date := "" //返回的 周1的日期

	switch what_day_str {

	case "Sunday": //周7
		monday_date = ChangeDateByNow("-144h", format)

	case "Monday": //周1
		monday_date = now_day

	case "Tuesday": //周2
		monday_date = ChangeDateByNow("-24h", format)

	case "Wednesday": //周3
		monday_date = ChangeDateByNow("-48h", format)

	case "Thursday": //周4
		monday_date = ChangeDateByNow("-72h", format)

	case "Friday": //周5
		monday_date = ChangeDateByNow("-96h", format)

	case "Saturday": //周6
		monday_date = ChangeDateByNow("-120h", format)

	}

	return monday_date
}

/*
获取本周一的日期
return 2017-04-17
*/
func GetTodayWeek() int {
	res := 0
	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)
	what_day := now_time.Weekday() //现在星期几
	what_day_str := what_day.String()

	switch what_day_str {
	case "Sunday": //周7
		res = 7
	case "Monday": //周1
		res = 1
	case "Tuesday": //周2
		res = 2
	case "Wednesday": //周3
		res = 3
	case "Thursday": //周4
		res = 4
	case "Friday": //周5
		res = 5
	case "Saturday": //周6
		res = 6
	}

	return res
}

/*
增加或者减少当前时间
@cha_str	= -3h;当前时间减少3小时
@format	返回时间的格式
*/
func ChangeDateByNow(cha_str, format string) string {
	dt, _ := time.ParseDuration(cha_str)
	return time.Now().Add(dt).Format(format)
}

/*
获取今天开始和结束时间的时间戳
*/
func GetTodayUnix() (int64, int64) {

	today_ymd := GetNowDatetime(C_format_date_ymd)

	today_morning := today_ymd + " 00:00:00"
	today_evening := today_ymd + " 23:59:59"

	today_morning_time, _ := FormatStr2Unix(today_morning, C_format_datetime)
	today_evening_time, _ := FormatStr2Unix(today_evening, C_format_datetime)

	return today_morning_time, today_evening_time
}

/*
获取今天开始和结束时间的时间戳字符串
*/
func GetTodayUnixStr() (string, string) {
	today_morning_time, today_evening_time := GetTodayUnix()
	today_morning_time_str := InterfaceToString(today_morning_time)
	today_evening_time_str := InterfaceToString(today_evening_time)

	return today_morning_time_str, today_evening_time_str
}

/*
获取今天的开始和结束时间
*/
func GetTodayDate() (string, string) {
	today_ymd := GetNowDatetime(C_format_date_ymd)

	today_morning := today_ymd + " 00:00:00"
	today_evening := today_ymd + " 23:59:59"

	return today_morning, today_evening
}

/*
获取上周六周日的开始结束时间
*/
func GetMondaySundy() (string, int, int) {
	status := "1"

	loc, _ := time.LoadLocation("Local")
	now_time := time.Now().In(loc)
	what_day := now_time.Weekday() //现在星期几
	what_day_str := what_day.String()

	todayStart := time.Now().Format("2006-01-02 00:00:00")
	format_date := "2006-01-02 15:04:05"
	todayStartStr, _ := FormatStr2Unix(todayStart, format_date) //当天零零点的时间戳

	today_int := int(todayStartStr)

	saturday_star := 0
	sunday_end := 0

	switch what_day_str {

	case "Sunday": //周7
		status = "0"

	case "Monday": //周1
		saturday_star = today_int - (86400 * 2)
		sunday_end = today_int

	case "Tuesday": //周2
		saturday_star = today_int - (86400 * 3)
		sunday_end = today_int - (86400)

	case "Wednesday": //周3
		saturday_star = today_int - (86400 * 4)
		sunday_end = today_int - (86400 * 2)

	case "Thursday": //周4
		saturday_star = today_int - (86400 * 5)
		sunday_end = today_int - (86400 * 3)

	case "Friday": //周5
		saturday_star = today_int - (86400 * 6)
		sunday_end = today_int - (86400 * 4)

	case "Saturday": //周6
		status = "0"
	}

	return status, saturday_star, sunday_end
}

/*
获取周1 zhi 周5的开始结束时间
*/
func GetAllWeek() []int {
	_, _, sunday_end := GetMondaySundy()

	lists_s := []int{}

	for i := 0; i < 5; i++ {
		data := sunday_end + (86400 * i)
		lists_s = append(lists_s, data)
	}

	return lists_s
}
