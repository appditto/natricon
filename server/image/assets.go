package image

import (
	"fmt"
	"os"
	"path"
)

type IllustrationType string

const (
	Body  IllustrationType = "Body"
	Hair  IllustrationType = "Hair"
	Mouth IllustrationType = "Mouth"
	Eye   IllustrationType = "Eye"
)

var BodyIllustrations = [...]string{
	"NarrowRound-10.svg",
	"NarrowRound-11.svg",
	"NarrowRound-12.svg",
	"NarrowRound-13.svg",
	"NarrowRound-14.svg",
	"NarrowRound-15.svg",
	"NarrowRound-5.svg",
	"NarrowRound-6.svg",
	"NarrowRound-7.svg",
	"NarrowRound-8.svg",
	"NarrowRound-9.svg",
	"Square-100.svg",
	"Square-104.svg",
	"Square-108.svg",
	"Square-112.svg",
	"Square-116.svg",
	"Square-120.svg",
	"Square-124.svg",
	"Square-128.svg",
	"Square-48.svg",
	"Square-52.svg",
	"Square-56.svg",
	"Square-60.svg",
	"Square-64.svg",
	"Square-68.svg",
	"Square-72.svg",
	"Square-76.svg",
	"Square-80.svg",
	"Square-84.svg",
	"Square-88.svg",
	"Square-92.svg",
	"Square-96.svg",
}

var HairIllustrations = [...]string{
	"Bubble-1.svg",
	"Slick.svg",
	"Weird.svg",
}

var EyeIllustrations = [...]string{
	"Eyeglasses-1.svg",
	"Eyeglasses-2.svg",
	"Eyeglasses-3.svg",
	"Eyeglasses-4.svg",
	"Eyes-1.svg",
}

var MouthIllustrations = [...]string{
	"Mustache-Slick.svg",
	"Smile-Bigger.svg",
	"Smile-Simple.svg",
	"Smile-Teeth.svg",
}

// GetIllustrationPath - get full path of image
func GetIllustrationPath(illustration string, iType IllustrationType) string {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	fPath := path.Join(wd, "assets", "illustrations", string(iType), illustration)
	if _, err := os.Stat(fPath); err != nil {
		// File does not exist
		panic(fmt.Sprintf("File %s does not exist", fPath))
	}
	return fPath
}
