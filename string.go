package u

import (
	"regexp"
	"strings"
)

// SplitSpace 以空格分割字符串-正则方式
func SplitSpace(s string) []string {
	compile, _ := regexp.Compile(`\s+`)
	return compile.Split(s, -1)
}

// Join 组合字符串
func Join(s []string) string {
	return strings.Join(s, "")
}
