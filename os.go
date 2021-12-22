package u

import (
	"path"
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
	return path.Base(fPath)
}

// FilenameNotExt 获取文件名称，不含后缀
func (receiver *osObj) FilenameNotExt(fPath string) string {
	return strings.Trim(fPath, path.Ext(fPath))
}
