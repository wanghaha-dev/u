package u

import (
	"github.com/spf13/viper"
	"sync"
)

type config struct {
	*viper.Viper
}

var _config *config
var _configOnce sync.Once

// Cfg 获取config对象
func Cfg() *config {
	_configOnce.Do(func() {
		_config = &config{viper.New()}
	})
	return _config
}
