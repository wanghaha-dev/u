package u

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

const (
	PanicLevel = log.DebugLevel
	FatalLevel = log.FatalLevel
	ErrorLevel = log.ErrorLevel
	WarnLevel = log.WarnLevel
	InfoLevel = log.InfoLevel
	DebugLevel = log.DebugLevel
	TraceLevel = log.TraceLevel
)

var (
	JSONFormatter = &log.JSONFormatter{}
	TextFormatter = &log.TextFormatter{}
)

var logger *log.Logger
var loggerOnce sync.Once

// Log 获取log对象
func Log() *log.Logger {
	loggerOnce.Do(func() {
		logger = log.New()
	})
	return logger
}

// Exist check path exists
func Exist(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}

// LogFile 日志文件
func LogFile(filePath string) (*os.File, error) {
	exist := Exist(filePath)
	if !exist {
		err := os.MkdirAll(filepath.Dir(filePath), 0777)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	return file, nil
}

