package image

import (
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type ImageFormat string

func ConvertSvgToBinary(svgData []byte, format ImageFormat, size uint) ([]byte, error) {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	pixelWand := imagick.NewPixelWand()
	if format == "jpg" || format == "jpeg" {
		pixelWand.SetColor("#FFFFFF")
	} else {
		pixelWand.SetColor("none")
	}
	mw.SetGravity(imagick.GRAVITY_CENTER)
	mw.SetBackgroundColor(pixelWand)
	//mw.SetResolution(float64(size), float64(size))
	err := mw.ReadImageBlob(svgData)
	if err != nil {
		return nil, err
	}
	mw.SetImageFormat(strings.ToUpper(string(format)))
	return mw.GetImageBlob(), nil
}
