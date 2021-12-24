package u

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

type uTime struct {
	time.Time
}

// Time 获取uTime对象
func Time() *uTime {
	return &uTime{time.Now()}
}

// Now 获取当前时间UTime对象
func (receiver *uTime) Now() *uTime {
	receiver.Time = time.Now()
	return receiver
}

// ParseTime 字符串是时间解析为time.Time时间对象
func (receiver *uTime) ParseTime(t string) time.Time {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}
	}
	return parseTime
}

// ParseUTime 字符串是时间解析为UTime时间对象
func (receiver *uTime) ParseUTime(t string) *uTime {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return receiver
	}
	return &uTime{parseTime}
}

// TimestampToTimeObject 时间戳转time.Time对象
func (receiver *uTime) TimestampToTimeObject(t int64) time.Time {
	return receiver.ParseTime(receiver.TimestampToDatetime(t))
}

// TimestampToUTimeObject 时间戳转UTime对象
func (receiver *uTime) TimestampToUTimeObject(t int64) *uTime {
	return &uTime{receiver.ParseTime(receiver.TimestampToDatetime(t))}
}

// DateString 获取日期字符串
func (receiver *uTime) DateString() string {
	return receiver.Format("2006-01-02")
}

// TimeString 获取时间字符串
func (receiver *uTime) TimeString() string {
	return receiver.Format("15:04:05")
}

// DateTimeString 获取日期时间字符串
func (receiver *uTime) DateTimeString() string {
	return receiver.Format("2006-01-02 15:04:05")
}

// AddYear 增加1年
func (receiver *uTime) AddYear() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddYear().Carbon2Time()
	return receiver
}

// AddYears 增加年数
func (receiver *uTime) AddYears(years int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddYears(years).Carbon2Time()
	return receiver
}

// AddMonth 增加一月
func (receiver *uTime) AddMonth() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddMonth().Carbon2Time()
	return receiver
}

// AddMonths 增加月数
func (receiver *uTime) AddMonths(months int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddMonths(months).Carbon2Time()
	return receiver
}

// AddDay 增加1天
func (receiver *uTime) AddDay() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddDay().Carbon2Time()
	return receiver
}

// AddDays 增加天数
func (receiver *uTime) AddDays(days int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddDays(days).Carbon2Time()
	return receiver
}

// AddHour 增加1小时
func (receiver *uTime) AddHour() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddHour().Carbon2Time()
	return receiver
}

// AddHours 增加小时数
func (receiver *uTime) AddHours(hours int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddHours(hours).Carbon2Time()
	return receiver
}

// AddMinute 增加1分钟
func (receiver *uTime) AddMinute() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddMinute().Carbon2Time()
	return receiver
}

// AddMinutes 增加分钟数
func (receiver *uTime) AddMinutes(minutes int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddMinutes(minutes).Carbon2Time()
	return receiver
}

// AddSecond 增加1秒
func (receiver *uTime) AddSecond() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddSecond().Carbon2Time()
	return receiver
}

// AddSeconds 增加秒数
func (receiver *uTime) AddSeconds(seconds int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).AddSeconds(seconds).Carbon2Time()
	return receiver
}

// SubYear 减去1年
func (receiver *uTime) SubYear() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubYear().Carbon2Time()
	return receiver
}

// SubYears 减去年数
func (receiver *uTime) SubYears(years int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubYears(years).Carbon2Time()
	return receiver
}

// SubMonth 减去1月
func (receiver *uTime) SubMonth() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubMonth().Carbon2Time()
	return receiver
}

// SubMonths 减去月数
func (receiver *uTime) SubMonths(months int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubMonths(months).Carbon2Time()
	return receiver
}

// SubDay 减去1天
func (receiver *uTime) SubDay() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubDay().Carbon2Time()
	return receiver
}

// SubDays 减去天数
func (receiver *uTime) SubDays(days int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubDays(days).Carbon2Time()
	return receiver
}

// SubHour 减去1小时
func (receiver *uTime) SubHour() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubHour().Carbon2Time()
	return receiver
}

// SubHours 减去小时数
func (receiver *uTime) SubHours(hours int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubHours(hours).Carbon2Time()
	return receiver
}

// SubMinute 减去1分钟
func (receiver *uTime) SubMinute() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubMinute().Carbon2Time()
	return receiver
}

// SubMinutes 减去分钟数
func (receiver *uTime) SubMinutes(minutes int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubMinutes(minutes).Carbon2Time()
	return receiver
}

// SubSecond 减去1秒钟
func (receiver *uTime) SubSecond() *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubSecond().Carbon2Time()
	return receiver
}

// SubSeconds 减去秒钟数
func (receiver *uTime) SubSeconds(seconds int) *uTime {
	receiver.Time = carbon.CreateFromTimestamp(receiver.Unix()).SubSeconds(seconds).Carbon2Time()
	return receiver
}

// ToTime UTime对象转time.Time对象
func (receiver *uTime) ToTime() time.Time {
	return receiver.Time
}

// Yesterday 获取昨天日期
func (receiver *uTime) Yesterday() *uTime {
	return receiver.SubDay()
}

// Tomorrow 获取明天日期
func (receiver *uTime) Tomorrow() *uTime {
	return receiver.AddDay()
}

// TimeFormat time.Time时间格式化
func (receiver *uTime) TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// Timestamp 获取时间戳-秒
func (receiver *uTime) Timestamp() int64 {
	return receiver.Unix()
}

// TimestampMilliSecond 获取时间戳-毫秒
func (receiver *uTime) TimestampMilliSecond() int64 {
	return receiver.UnixMilli()
}

// TimestampMicroSecond 获取时间戳-微秒
func (receiver *uTime) TimestampMicroSecond() int64 {
	return receiver.UnixMicro()
}

// TimestampNanoSecond 获取时间戳-纳秒
func (receiver *uTime) TimestampNanoSecond() int64 {
	return receiver.UnixNano()
}

// DateTime 获取当前日期时间
func (receiver *uTime) DateTime() string {
	return receiver.Format("2006-01-02 15:04:05")
}

// Date 获取日期
func (receiver *uTime) Date() string {
	return receiver.Format("2006-01-02")
}

// TimestampToDatetime 时间戳转格式化日期时间
func (receiver *uTime) TimestampToDatetime(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02 15:04:05")
}

// DatetimeToTimestamp 日期时间转时间戳
func (receiver *uTime) DatetimeToTimestamp(datetime string) int64 {
	parse, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0
	}
	return parse.Unix()
}

// Year get Year
func (receiver *uTime) Year() int {
	return receiver.Time.Year()
}

// Month get Month
func (receiver *uTime) Month() int {
	return int(receiver.Time.Month())
}

// Day get Day
func (receiver *uTime) Day() int {
	return receiver.Time.Day()
}

// Hour get Hour
func (receiver *uTime) Hour() int {
	return receiver.Time.Hour()
}

// Minute get Minute
func (receiver *uTime) Minute() int {
	return receiver.Time.Minute()
}

// Second get Second
func (receiver *uTime) Second() int {
	return receiver.Time.Second()
}
