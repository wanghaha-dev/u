package u

import (
	"io"
	"io/fs"
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

// CopyFile 复制文件或文件夹
func (receiver *osObj)CopyFile(src, dst string) error {
	srcData, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcData.IsDir() {
		_, err := os.Stat(dst)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dst, 0777)
			if err != nil {
				return err
			}
		}

		err = filepath.Walk(src, func(srcFilePath string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			dstFilename := path.Join(dst,
				path.Join(
					strings.TrimLeft(srcFilePath, src),
					filepath.Base(srcFilePath),
				),
			)
			err = receiver.copyFile(srcFilePath, dstFilename)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	} else {
		err := receiver.copyFile(src, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// copyFile 复制文件
func (receiver *osObj) copyFile(src, dst string) error {
	err := OS().CheckCreateDirs(dst)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer dstFile.Close()
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}
