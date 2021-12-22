package u

import (
	"github.com/disintegration/imaging"
)

// ResizePreserveAspectRatio 保持纵横比调整图片
func ResizePreserveAspectRatio(width int, sourceImagePath string, savePath string) error {
	src, err := imaging.Open(sourceImagePath)
	if err != nil {
		return err
	}

	dstImage := imaging.Resize(src, width, 0, imaging.Lanczos)
	err = imaging.Save(dstImage, savePath)
	if err != nil {
		return err
	}
	return nil
}
