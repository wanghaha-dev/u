package u

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

// RandNum 生成随机数
func RandNum(num int) int {
	rand.Seed(time.Now().Unix() + int64(rand.Intn(math.MaxInt)))
	return rand.Intn(num)
}

// RandRangeNum 生成指定范围的随机数
func RandRangeNum(min, max int) int {
	if min >= max {
		return 0
	}
	return max - RandNum(max-min+1) + 1
}

// RandStr 生成随机字符串
func RandStr(length int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strSlice := strings.Split(chars, "")
	strList := make([]string, length)
	for i := 0; i < length; i++ {
		num := RandNum(len(chars))
		strList = append(strList, strSlice[num])
	}
	return strings.Join(strList, "")
}

// Rand32Uid 获取32位长度uid
func Rand32Uid() string {
	return RandStr(32)
}

// RandCheckCode 获取6位验证码
func RandCheckCode() string {
	return ToString(RandRangeNum(100000, 999999))
}
