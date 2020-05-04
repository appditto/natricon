package image

import (
	"fmt"

	magick "github.com/bbedward/magick"
)

type ImageFormat string

func convertSvgToBinary(svgData []byte, format ImageFormat, size int) ([]byte, error) {
	image, err := magick.NewFromBlob(svgData, "svg")
	if err != nil {
		return nil, err
	}
	defer image.Destroy()
	image.Resize(fmt.Sprintf("%dx%d", size, size))
	return image.ToBlob(string(format))
}
