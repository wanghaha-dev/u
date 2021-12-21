package u

import (
	"encoding/xml"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

type router struct {
	*httprouter.Router
}

var _server *router

// Server 获取server对象
func Server() *router {
	_server = &router{
		httprouter.New(),
	}
	//http.Handle("/", _server)
	return _server
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
}

// Handle server handle
func Handle(path string, method string, f func(c *Context)) {
	_server.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := new(Context)
		ctx.Response = writer
		ctx.Request = request
		ctx.Params = params
		f(ctx)
	})
}

// GET method GET
func (receiver *router) GET(path string, f func(c *Context)) {
	Handle(path, http.MethodGet, f)
}

// POST method POST
func (receiver *router) POST(path string, f func(c *Context)) {
	Handle(path, http.MethodPost, f)
}

// PUT method PUT
func (receiver *router) PUT(path string, f func(c *Context)) {
	Handle(path, http.MethodPut, f)
}

// DELETE method DELETE
func (receiver *router) DELETE(path string, f func(c *Context)) {
	Handle(path, http.MethodDelete, f)
}

// HEAD method HEAD
func (receiver *router) HEAD(path string, f func(c *Context)) {
	Handle(path, http.MethodHead, f)
}

// OPTIONS method options
func (receiver *router) OPTIONS(path string, f func(c *Context)) {
	Handle(path, http.MethodOptions, f)
}

// Run 运行
func (receiver *router) Run(addr string) error {
	NewColor(FgLightWhite).Println("#######################################################")
	NewColor(LightRed).Printf("uWeb listen %s start time: %s\n", addr, Time().Now().DateTime())
	NewColor(FgLightWhite).Println("#######################################################")
	return http.ListenAndServe(addr, _server)
}

// JSON json响应
func (receiver *Context) JSON(data interface{}) {
	receiver.Response.WriteHeader(200)
	receiver.Response.Header().Set("Content-Type", "application/json")

	switch data.(type) {
	case string:
		_, _ = receiver.Response.Write(ToBytes(data))
	default:
		_, _ = receiver.Response.Write(ToBytes(Json(data)))
	}
}

// XML xml响应
func (receiver *Context) XML(data interface{}) {
	receiver.Response.WriteHeader(200)
	receiver.Response.Header().Set("Content-Type", "application/xml")
	bytes, _ := xml.Marshal(data)
	_, _ = receiver.Response.Write(bytes)
}

// String 字符串响应
func (receiver *Context) String(data string) {
	receiver.Response.WriteHeader(200)
	_, _ = receiver.Response.Write([]byte(data))
}

// GetParam 获取路由参数
func (receiver *Context) GetParam(key string) string {
	return receiver.Params.ByName(key)
}

// GetVar 获取查询参数
func (receiver *Context) GetVar(key string) string {
	return receiver.Request.URL.Query().Get(key)
}

// Bind 绑定参数
func (receiver *Context) Bind(obj interface{}) {
	dataMap := receiver.Request.URL.Query()

	v := reflect.ValueOf(obj)
	for i := 0; i < v.Elem().NumField(); i++ {
		t := v.Type()
		fieldValue := dataMap.Get(t.Elem().Field(i).Tag.Get("field"))

		switch v.Elem().Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.Elem().Field(i).SetInt(ToInt64(fieldValue))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v.Elem().Field(i).SetUint(ToUint64(fieldValue))
		case reflect.Float32, reflect.Float64:
			v.Elem().Field(i).SetFloat(ToFloat64(fieldValue))
		case reflect.Bool:
			v.Elem().Field(i).SetBool(ToBool(fieldValue))
		case reflect.String:
			v.Elem().Field(i).SetString(fieldValue)
		default:
			v.Elem().Field(i).SetString(fieldValue)
		}
	}
}
