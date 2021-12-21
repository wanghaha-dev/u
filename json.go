package u

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

// JsonGet 获取指定key的值
func JsonGet(str, key string) gjson.Result {
	return gjson.Get(str, key)
}

// JsonBytesGet 获取指定key的值
func JsonBytesGet(data []byte, key string) gjson.Result {
	return gjson.GetBytes(data, key)
}

// JsonGetMany 获取多个key值
func JsonGetMany(str string, key ...string) []gjson.Result {
	return gjson.GetMany(str, key...)
}

// JsonGetManyBytes 获取多个key值
func JsonGetManyBytes(data []byte, key ...string) []gjson.Result {
	return gjson.GetManyBytes(data, key...)
}

// JsonExists 是否存在
func JsonExists(str, key string) bool {
	return gjson.Get(str, key).Exists()
}

// Json 转Json字符串
func Json(val interface{}) string {
	data, err := json.Marshal(val)
	if err != nil {
		return ""
	}
	return string(data)
}

// UnJson json反序列化
func UnJson(data string, val interface{}) error {
	err := json.Unmarshal([]byte(data), val)
	if err != nil {
		return err
	}
	return nil
}
