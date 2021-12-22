package u

import (
	"github.com/disintegration/imaging"
	"os"
	"path"
	"sync"
)

type img struct {}

var _img *img
var _imgOnce sync.Once

// Image 获取img对象
func Image() *img {
	_imgOnce.Do(func() {
		_img = &img{}
	})
	return _img
}

// ResizePreserveAspectRatio 保持纵横比调整图片
func (receiver *img)ResizePreserveAspectRatio(width int, sourceImagePath string, savePath string) error {
	src, err := imaging.Open(sourceImagePath)
	if err != nil {
		return err
	}

	dstImage := imaging.Resize(src, width, 0, imaging.Lanczos)
	_, err = os.Stat(savePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(savePath), 0777)
		if err != nil {
			return err
		}
	}

	err = imaging.Save(dstImage, savePath)
	if err != nil {
		return err
	}
	return nil
}

// Resize 调整图片
func (receiver *img) Resize(sourceImagePath string, width, height int, savePath string) error {
	src, err := imaging.Open(sourceImagePath)
	if err != nil {
		return err
	}

	dstImage := imaging.Resize(src, width, height, imaging.Lanczos)
	_, err = os.Stat(savePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(savePath), 0777)
		if err != nil {
			return err
		}
	}

	err = imaging.Save(dstImage, savePath)
	if err != nil {
		return err
	}
	return nil
}
