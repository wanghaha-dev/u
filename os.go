package u

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type osObj struct{}

var _os *osObj
var _osOnce sync.Once

func OS() *osObj{
	_osOnce.Do(func() {
		_os = new(osObj)
	})
	return _os
}

// Filename 获取文件名称，包含后缀
func (receiver *osObj) Filename(fPath string) string {
	return filepath.Base(fPath)
}

// FilenameNotExt 获取文件名称，不含后缀
func (receiver *osObj) FilenameNotExt(fPath string) string {
	return strings.Trim(fPath, path.Ext(fPath))
}

// ExtNotDot 获取文件后缀不带"."
func (receiver *osObj) ExtNotDot(fPath string) string {
	return strings.Trim(path.Ext(fPath), ".")
}

// CheckCreateDirs 判断文件夹是否存在，不能存在则创建
func (receiver *osObj)CheckCreateDirs(fPath string) error {
	_, err := os.Stat(fPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(fPath), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
