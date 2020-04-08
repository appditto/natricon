package image

import (
	"errors"
	"math"
	"math/rand"
	"strconv"

	"github.com/appditto/natricon/color"
)

// Accessories - represents accessories for natricon
type Accessories struct {
	bodyColor color.RGB
	hairColor color.RGB
}

// GetAccessoriesForHash - Return Accessories object based on 64-character hex string
func GetAccessoriesForHash(hash string) (Accessories, error) {
	var err error
	if len(hash) != 64 {
		return Accessories{}, errors.New("Invalid hash")
	}
	// Validate is a hex string
	_, err = strconv.ParseUint(hash, 16, 64)
	if err != nil {
		return Accessories{}, err
	}

	// Create empty Accessories object
	var accessories = Accessories{}
	// Body color is first 6 digits as hex string
	bodyColorHex := hash[0:6]

	accessories.bodyColor, err = color.HTMLToRGB(bodyColorHex)
	if err != nil {
		return Accessories{}, err
	}

	// Get hair color using next 6 bits
	accessories.hairColor, err = GetHairColor(accessories.bodyColor, hash[6:12], hash[12:16], hash[16:20])

	return accessories, nil
}

// GetHairColor - Get a complementary color with given entropy
func GetHairColor(bodyColor color.RGB, hEntropy string, sEntropy string, bEntropy string) (color.RGB, error) {
	var err error
	// Get as HSV color
	bodyColorHSV := bodyColor.ToHSV()

	var randSeed int64
	var shiftedHue float64
	var shiftedSaturation float64
	var shiftedBrightness float64
	// Want to shift the hue between 90-270
	// Get detemrinistic RNG
	randSeed, err = strconv.ParseInt(hEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}

	// Generate random shift between 90...270
	r := rand.New(rand.NewSource(randSeed))
	shiftedHue = float64(r.Intn(270-90)+90) + bodyColorHSV.H

	// If > 360, subtract
	if shiftedHue > 360 {
		shiftedHue = shiftedHue - 360
	}

	// Generate random shift between 0..20 for saturation
	randSeed, err = strconv.ParseInt(sEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.New(rand.NewSource(randSeed))
	shiftedSaturation = float64(r.Intn(21)) + bodyColorHSV.S

	// Generate random shift between 0..20 for brightness
	randSeed, err = strconv.ParseInt(bEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.New(rand.NewSource(randSeed))
	shiftedBrightness = float64(r.Intn(21)) + bodyColorHSV.V

	// If > 100, subtract
	if shiftedSaturation > 100 {
		shiftedSaturation = 100 - shiftedSaturation
	}
	if shiftedBrightness > 100 {
		shiftedBrightness = 100 - shiftedBrightness
	}

	hairColorHSV := color.HSV{}
	hairColorHSV.H = shiftedHue
	hairColorHSV.S = math.Max(10, shiftedSaturation)
	hairColorHSV.V = math.Max(10, shiftedBrightness)

	return hairColorHSV.ToRGB(), nil
}
