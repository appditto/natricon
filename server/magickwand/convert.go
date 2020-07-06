package magickwand

import (
	"strings"

	"github.com/appditto/natricon/server/image"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type ImageFormat string

func ConvertSvgToBinary(svgData []byte, format ImageFormat, size uint) ([]byte, error) {
	mw := imagick.NewMagickWand()
	mw.SetImageFormat("SVG")
	pixelWand := imagick.NewPixelWand()
	pixelWand.SetColor("none")
	mw.SetBackgroundColor(pixelWand)
	mw.SetImageUnits(imagick.RESOLUTION_PIXELS_PER_INCH)
	density := 96.0 * float64(size) / float64(image.DefaultSize)
	mw.SetResolution(density, density)
	err := mw.ReadImageBlob(svgData)
	if err != nil {
		return nil, err
	}
	mw.SetImageCompression(imagick.COMPRESSION_NO)
	mw.SetImageCompressionQuality(100)
	//mw.SetAntialias(true)
	mw.SetImageFormat(strings.ToUpper(string(format)))
	return mw.GetImageBlob(), nil
}
