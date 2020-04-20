package image

import (
	"errors"
	"math"
	"regexp"
	"strconv"

	"github.com/appditto/natricon/color"
	"github.com/appditto/natricon/rand"
)

// Constants
var MinSaturation float64 = 0.2 // Minimum allowed saturation
var MinBrightness float64 = 0.4 // Minimum allowed brightness

// Accessories - represents accessories for natricon
type Accessories struct {
	BodyColor     color.RGB
	HairColor     color.RGB
	BodyAsset     Asset
	HairAsset     Asset
	MouthAsset    Asset
	EyeAsset      Asset
	BackHairAsset *Asset
}

// Hex string regex
var hexRegex = regexp.MustCompile("^[0-9a-fA-F]+$")

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
	bodyColorHSV.V = math.Max(MinBrightness, bodyColorHSV.V)
	bodyColorHSV.S = math.Max(MinSaturation, bodyColorHSV.S)
	accessories.BodyColor = bodyColorHSV.ToRGB()

	// Get hair color using next 6 bits
	accessories.HairColor, err = GetHairColor(accessories.BodyColor, hash[6:12], hash[12:16], hash[16:20])

	// Get body and hair illustrations
	accessories.BodyAsset, err = GetBodyAsset(hash[20:26])
	accessories.HairAsset, err = GetHairAsset(hash[26:32], &accessories.BodyAsset)
	accessories.BackHairAsset = GetBackHairAsset(accessories.HairAsset)

	// Get mouth and eyes
	targetSex := Neutral
	if accessories.BodyAsset.Sex != Neutral {
		targetSex = accessories.BodyAsset.Sex
	} else if accessories.HairAsset.Sex != Neutral {
		targetSex = accessories.HairAsset.Sex
	}
	accessories.MouthAsset, err = GetMouthAsset(hash[32:40], targetSex)
	if targetSex == Neutral && accessories.MouthAsset.Sex != Neutral {
		targetSex = accessories.MouthAsset.Sex
	}
	accessories.EyeAsset, err = GetEyeAsset(hash[40:48], targetSex)

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
	r := rand.Init()
	r.Seed(uint32(randSeed))
	shiftedHue = float64(r.Int31n(270-90)+90) + bodyColorHSV.H

	// If > 360, subtract
	if shiftedHue > 360 {
		shiftedHue = shiftedHue - 360
	}

	// Generate random shift between 0..20 for saturation
	randSeed, err = strconv.ParseInt(sEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r.Seed(uint32(randSeed))
	// Adjust saturation by -20 to + 40
	randNum := float64(r.Int31n(121) - 20)
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
	r.Seed(uint32(randSeed))
	// Adjust brightness by -20 to + 40
	randNum = float64(r.Int31n(121) - 20)
	shiftedBrightness = (bodyColorHSV.V * 100.0) + randNum
	// Cap at 100
	if shiftedBrightness > 100 {
		shiftedBrightness = 100
	}

	hairColorHSV := color.HSV{}
	hairColorHSV.H = shiftedHue
	hairColorHSV.S = math.Max(MinSaturation, shiftedSaturation/100.0)
	hairColorHSV.V = math.Max(MinBrightness, shiftedBrightness/100.0)

	return hairColorHSV.ToRGB(), nil
}

// GetBodyAsset - return body illustration to use with given entropy
func GetBodyAsset(entropy string) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	r := rand.Init()
	r.Seed(uint32(randSeed))
	bodyIndex := r.Int31n(int32(GetAssets().GetNBodyAssets()))

	return GetAssets().GetBodyAssets()[bodyIndex], nil
}

// GetHairAsset - return hair illustration to use with given entropy
func GetHairAsset(entropy string, bodyAsset *Asset) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	hairAssetOptions := GetAssets().GetHairAssets(bodyAsset.Sex)

	r := rand.Init()
	r.Seed(uint32(randSeed))
	hairIndex := r.Int31n(int32(len(hairAssetOptions)))

	return hairAssetOptions[hairIndex], nil
}

// GetBackHairAsset - return back hair illustration for a given hair asset
func GetBackHairAsset(hairAsset Asset) *Asset {
	for _, ba := range GetAssets().GetBackHairAssets() {
		if ba.FileName == hairAsset.FileName {
			return &ba
		}
	}
	return nil
}

// GetEyeAsset - return hair illustration to use with given entropy
func GetEyeAsset(entropy string, sex Sex) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	eyeAssetOptions := GetAssets().GetEyeAssets(sex)

	r := rand.Init()
	r.Seed(uint32(randSeed))
	eyeIndex := r.Int31n(int32(len(eyeAssetOptions)))

	return eyeAssetOptions[eyeIndex], nil
}

// GetEyeAsset - return hair illustration to use with given entropy
func GetMouthAsset(entropy string, sex Sex) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	mouthAssetOptions := GetAssets().GetMouthAssets(sex)

	r := rand.Init()
	r.Seed(uint32(randSeed))
	mouthIndex := r.Int31n(int32(len(mouthAssetOptions)))

	return mouthAssetOptions[mouthIndex], nil
}
