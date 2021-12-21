package u

// GetError 获取错误
func GetError(err error) {
	if err != nil {
		Log().Errorln(err)
	}
}

// GetErrorCallback 获取错误并执行回调函数
func GetErrorCallback(err error, function func(err error)) {
	if err != nil {
		function(err)
		Log().Errorln(err)
	}
}
