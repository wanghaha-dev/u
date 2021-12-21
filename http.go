package u

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Http 获取http请求对象
func Http() *request {
	return &request{}
}

// Params URL参数别名
type Params map[string]interface{}

// Body Body参数别名
type Body map[string]interface{}

// Headers 请求头别名
type Headers map[string]interface{}

// Cookies cookies
type Cookies []http.Cookie

// Request 请求结构体
type request struct {}

// Response 响应结构体
type response struct {
	Resp []byte
	Error error
}

func (resp *response) String() string {
	return resp.ToString()
}

// Bind 将响应的结果解析到结构体上
func (resp *response) Bind(StructData interface{}) error {
	err := json.Unmarshal(resp.Resp, StructData)
	if err != nil {
		return err
	}
	return nil
}

// ToString 将响应结果转为string
func (resp *response) ToString() string {
	return string(resp.Resp)
}

// GetJsonToMap 将json格式的相应内容转为map类型
func (resp *response) GetJsonToMap(key string) gjson.Result {
	return JsonBytesGet(resp.Resp, key)
}

// Get 普通GET请求
func (r *request) Get(URL string) *response {
	return r.CustomRequest(URL, "GET")
}

// Post 普通POST请求
func (r *request) Post(URL string) *response {
	return r.CustomRequest(URL, "POST")
}

// GetRequest GET请求支持自定义参数和请求头
func (r *request) GetRequest(URL string, v ...interface{}) *response {
	return r.CustomRequest(URL, "GET", v...)
}

// PostRequest POST请求支持自定义参数和请求头
func (r *request) PostRequest(URL string, v ...interface{}) *response {
	return r.CustomRequest(URL, "POST", v...)
}

// PutRequest PUT请求支持自定义参数和请求头
func (r *request) PutRequest(URL string, v ...interface{}) *response {
	return r.CustomRequest(URL, "PUT", v...)
}

// DeleteRequest DELETE请求支持自定义参数和请求头
func (r *request) DeleteRequest(URL string, v ...interface{}) *response {
	return r.CustomRequest(URL, "DELETE", v...)
}

// PatchRequest PATCH请求支持自定义参数和请求头
func (r *request) PatchRequest(URL string, v ...interface{}) *response {
	return r.CustomRequest(URL, "PATCH", v...)
}

// CustomRequest 自定义请求支持自定义参数和请求头
func (r *request) CustomRequest(URL string, METHOD string,  v ...interface{}) *response {
	urlObj := url.Values{}
	payload := strings.NewReader("")
	request, err := http.NewRequest(METHOD, URL, nil)
	if err != nil {
		return &response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToString(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()
		case Body:
			// 添加 body 请求参数
			request.Header.Set("Content-type", "application/x-www-form-urlencoded")
			for key, val := range item.(Body) {
				urlObj.Add(key, ToString(val))
			}
		case Headers:
			// 设置请求头
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToString(val))
			}
		case time.Duration:
			// 设置超时时间
			http.DefaultClient.Timeout = item.(time.Duration)
		case Cookies:
			// 设置cookie
			for _, cookie := range item.(Cookies) {
				request.AddCookie(&cookie)
			}
		}
	}

	// 重新赋值
	payload = strings.NewReader(urlObj.Encode())
	request.Body = ioutil.NopCloser(payload)

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response{
			Resp:  nil,
			Error: err,
		}
	}

	return &response{
		Resp:  data,
		Error: nil,
	}
}

