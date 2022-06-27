package u

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

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

// PoolGoroutine 携程池
func PoolGoroutine(poolNum int, countNum int, function func()) {
	defer ants.Release()
	var wg sync.WaitGroup

	pool, _ := ants.NewPool(poolNum)

	task := func() {
		function()
		wg.Done()
	}

	for i:=0;i<countNum;i++ {
		wg.Add(1)
		_ = pool.Submit(task)
	}

	wg.Wait()
}
