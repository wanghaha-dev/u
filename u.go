package u

import (
	"fmt"
	"github.com/go-basic/uuid"
	"log"
	"strings"
	"time"
)

type Map map[string]interface{}

// Author contact author
func Author() {
	fmt.Println("author: 王哈哈")
	fmt.Println("contact: 31931727@qq.com")
}

// HasExists 判断元素是否存在与某列表
func HasExists(dataList interface{}, data interface{}) bool {
	result := false

	switch dataList.(type) {
	case []string:
		for _, item := range dataList.([]string) {
			if item == data {
				result = true
			}
		}
	case []int:
		for _, item := range dataList.([]int) {
			if item == data {
				result = true
			}
		}
	case []int8:
		for _, item := range dataList.([]int8) {
			if item == data {
				result = true
			}
		}
	case []int16:
		for _, item := range dataList.([]int16) {
			if item == data {
				result = true
			}
		}
	case []int32:
		for _, item := range dataList.([]int32) {
			if item == data {
				result = true
			}
		}
	case []int64:
		for _, item := range dataList.([]int64) {
			if item == data {
				result = true
			}
		}
	case []uint:
		for _, item := range dataList.([]int) {
			if item == data {
				result = true
			}
		}
	case []uint8:
		for _, item := range dataList.([]int8) {
			if item == data {
				result = true
			}
		}
	case []uint16:
		for _, item := range dataList.([]int16) {
			if item == data {
				result = true
			}
		}
	case []uint32:
		for _, item := range dataList.([]int32) {
			if item == data {
				result = true
			}
		}
	case []uint64:
		for _, item := range dataList.([]int64) {
			if item == data {
				result = true
			}
		}

	case []float32:
		for _, item := range dataList.([]float32) {
			if item == data {
				result = true
			}
		}
	case []float64:
		for _, item := range dataList.([]float64) {
			if item == data {
				result = true
			}
		}
	}
	return result
}

// UUID 获取UUID
func UUID() string {
	s := uuid.New()
	return s
}

// UUIDStr 获取UUID字符串, 不含"-"
func UUIDStr() string {
	s := uuid.New()
	return strings.ReplaceAll(s, "-", "")
}

// UUIDStrUpper 获取UUID字符串, 不含"-", 转大写
func UUIDStrUpper() string {
	s := uuid.New()
	return strings.ToUpper(strings.ReplaceAll(s, "-", ""))
}

// Foreach 循环执行
func Foreach(n int, callback func()) {
	for i := 0; i < n; i++ {
		callback()
	}
}

// separator 分隔符
const separator = "#############################################"

// Debug 调试输出
func Debug(data ...interface{}) {
	log.Println(separator)
	log.Println(data...)
	log.Println(separator)
}

// DebugYellow 调试输出-黄色
func DebugYellow(data ...interface{}) {
	PrintlnYellow(Time().DateTime(), separator)
	PrintYellow(Time().DateTime(), " ")
	PrintlnYellow(data...)
	PrintlnYellow(Time().DateTime(), separator)
}

// DebugBlue 调试输出-蓝色
func DebugBlue(data ...interface{}) {
	PrintlnBlue(Time().DateTime(), separator)
	PrintBlue(Time().DateTime(), " ")
	PrintlnBlue(data...)
	PrintlnBlue(Time().DateTime(), separator)
}

// DebugGreen 调试输出-绿色
func DebugGreen(data ...interface{}) {
	PrintlnGreen(Time().DateTime(), separator)
	PrintGreen(Time().DateTime(), " ")
	PrintlnGreen(data...)
	PrintlnGreen(Time().DateTime(), separator)
}

// DebugGray 调试输出-灰色
func DebugGray(data ...interface{}) {
	PrintlnGray(Time().DateTime(), separator)
	PrintGray(Time().DateTime(), " ")
	PrintlnGray(data...)
	PrintlnGray(Time().DateTime(), separator)
}

// DebugCyan 调试输出-青色
func DebugCyan(data ...interface{}) {
	PrintlnCyan(Time().DateTime(), separator)
	PrintCyan(Time().DateTime(), " ")
	PrintlnCyan(data...)
	PrintlnCyan(Time().DateTime(), separator)
}

// Delay 计算耗时
func Delay(callback func()) {
	start := time.Now().UnixNano()

	callback()

	fmt.Println("========================= 共耗时 ===========================")

	LightRed.Print(time.Now().Unix() - start / 1000 / 1000 / 1000)
	fmt.Println(" Unix")

	LightRed.Print(time.Now().UnixMilli() - start / 1000 / 1000)
	fmt.Println(" UnixMilli")

	LightRed.Print(time.Now().UnixMicro() - start / 1000)
	fmt.Println(" UnixMicro")

	LightRed.Print(time.Now().UnixNano() - start)
	fmt.Println(" UnixNano")
	fmt.Println("===========================================================")
}


// Version 输出版本
func Version() {
	fmt.Println("v0.0.47")
	fmt.Println("Last update time: 2021-12-23 17:42:00")
}
