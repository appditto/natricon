package image

import (
	"errors"
	"math"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/appditto/natricon/color"
)

// Accessories - represents accessories for natricon
type Accessories struct {
	BodyColor  color.RGB
	HairColor  color.RGB
	BodyAsset  string
	HairAsset  string
	MouthAsset string
	EyeAsset   string
}

// Hex string regex
var hexRegex = regexp.MustCompile("^[0-9a-fA-F]+$")

// Constants
var minSaturation float64 = 0.2
var minBrightness float64 = 0.4

// GetAccessoriesForHash - Return Accessories object based on 64-character hex string
func GetAccessoriesForHash(hash string) (Accessories, error) {
	var err error
	if len(hash) != 64 {
		return Accessories{}, errors.New("Invalid hash")
	}
	// Validate is a hex string
	if !hexRegex.MatchString(hash) {
		return Accessories{}, errors.New("Invalid hash")
	}

	// Create empty Accessories object
	var accessories = Accessories{}
	// Body color is first 6 digits as hex string
	bodyColorHex := hash[0:6]

	accessories.BodyColor, err = color.HTMLToRGB(bodyColorHex)
	if err != nil {
		return Accessories{}, err
	}
	// Enforce min saturation and brightness
	bodyColorHSV := accessories.BodyColor.ToHSV()
	bodyColorHSV.V = math.Max(minBrightness, bodyColorHSV.V)
	bodyColorHSV.S = math.Max(minSaturation, bodyColorHSV.S)
	accessories.BodyColor = bodyColorHSV.ToRGB()

	// Get hair color using next 6 bits
	accessories.HairColor, err = GetHairColor(accessories.BodyColor, hash[6:12], hash[12:16], hash[16:20])

	// Get body and hair illustrations
	accessories.BodyAsset, err = GetBodyAsset(hash[20:26])
	accessories.HairAsset, err = GetHairAsset(hash[26:32])

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
	// Adjust saturation by -20 to + 40
	randNum := float64(r.Intn(121) - 20)
	shiftedSaturation = (bodyColorHSV.S * 100.0) + randNum
	// Cap at 100
	if shiftedSaturation > 100 {
		shiftedSaturation = 100
	}

	// Generate random shift between 0..20 for brightness
	randSeed, err = strconv.ParseInt(bEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.New(rand.NewSource(randSeed))
	// Adjust brightness by -20 to + 40
	randNum = float64(r.Intn(121) - 20)
	shiftedBrightness = (bodyColorHSV.V * 100.0) + randNum
	// Cap at 100
	if shiftedBrightness > 100 {
		shiftedBrightness = 100
	}

	hairColorHSV := color.HSV{}
	hairColorHSV.H = shiftedHue
	hairColorHSV.S = math.Max(minSaturation, shiftedSaturation/100.0)
	hairColorHSV.V = math.Max(minBrightness, shiftedBrightness/100.0)

	return hairColorHSV.ToRGB(), nil
}

// GetBodyAsset - return body illustration to use with given entropy
func GetBodyAsset(entropy string) (string, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return "", err
	}

	r := rand.New(rand.NewSource(randSeed))
	bodyIndex := r.Intn(len(BodyIllustrations))

	return GetIllustrationPath(BodyIllustrations[bodyIndex], Body), nil
}

// GetHairAsset - return hair illustration to use with given entropy
func GetHairAsset(entropy string) (string, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return "", err
	}

	r := rand.New(rand.NewSource(randSeed))
	hairIndex := r.Intn(len(HairIllustrations))

	return GetIllustrationPath(HairIllustrations[hairIndex], Hair), nil
}
