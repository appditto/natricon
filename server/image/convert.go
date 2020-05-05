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
	mw.SetBackgroundColor(pixelWand)
	mw.SetImageUnits(imagick.RESOLUTION_PIXELS_PER_INCH)
	density := 96.0 * float64(size) / float64(DefaultSize)
	mw.SetResolution(density, density)
	err := mw.ReadImageBlob(svgData)
	if err != nil {
		return nil, err
	}
	mw.SetImageFormat(strings.ToUpper(string(format)))
	mw.SetImageCompressionQuality(95)
	return mw.GetImageBlob(), nil
}
