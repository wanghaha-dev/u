package u

// ConcurrentHandle 并发任务
func ConcurrentHandle(dataMapList []map[string]interface{}, callback func(map[string]interface{}) map[string]interface{}) {
	ch := make(chan bool)
	for _, item := range dataMapList {
		go func(dataArg map[string]interface{}) {
			callback(dataArg)

			ch <- true
		}(item)
	}

	Foreach(len(dataMapList), func() {
		<- ch
	})
}
