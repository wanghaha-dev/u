package u

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type cronObj struct {
	*cron.Cron
}

var _cron *cronObj
var _cronOnce sync.Once

// Cron 获取cron对象
func Cron() *cronObj {
	_cronOnce.Do(func() {
		_cron = &cronObj{cron.New(cron.WithSeconds())}
	})
	return _cron
}
