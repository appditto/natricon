package image

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/appditto/natricon/server/color"
	"github.com/appditto/natricon/server/rand"
)

// Constants
var MinSaturation float64 = 0.2 // Minimum allowed saturation
var MinBrightness float64 = 0.4 // Minimum allowed brightness

// Accessories - represents accessories for natricon
type Accessories struct {
	BodyColor         color.RGB
	HairColor         color.RGB
	BodyAsset         Asset
	HairAsset         Asset
	MouthAsset        Asset
	EyeAsset          Asset
	BackHairAsset     *Asset
	BodyOutlineAsset  *Asset
	HairOutlineAsset  *Asset
	MouthOutlineAsset *Asset
	OutlineColor      color.RGB
}

// Hex string regex
var hexRegex = regexp.MustCompile("^[0-9a-fA-F]+$")

// GetAccessoriesForHash - Return Accessories object based on 64-character hex string
func GetAccessoriesForHash(hash string, outline bool, outlineColor *color.RGB) (Accessories, error) {
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
	// Body color uses first 12 digits of hash as seed
	accessories.BodyColor, err = GetBodyColor(hash[0:16])
	if err != nil {
		return Accessories{}, err
	}

	// Get hair color
	accessories.HairColor, err = GetHairColor(accessories.BodyColor, hash[16:26], hash[26:30], hash[30:34])

	// Get body and hair illustrations
	accessories.BodyAsset, err = GetBodyAsset(hash[34:40])
	accessories.HairAsset, err = GetHairAsset(hash[40:46], &accessories.BodyAsset)
	accessories.BackHairAsset = GetBackHairAsset(accessories.HairAsset)

	// Get mouth and eyes
	targetSex := Neutral
	if accessories.BodyAsset.Sex != Neutral {
		targetSex = accessories.BodyAsset.Sex
	} else if accessories.HairAsset.Sex != Neutral {
		targetSex = accessories.HairAsset.Sex
	}
	accessories.MouthAsset, err = GetMouthAsset(hash[46:55], targetSex)
	if targetSex == Neutral && accessories.MouthAsset.Sex != Neutral {
		targetSex = accessories.MouthAsset.Sex
	}
	accessories.EyeAsset, err = GetEyeAsset(hash[55:64], targetSex)

	// Get outlines
	if outline {
		accessories.BodyOutlineAsset = GetBodyOutlineAsset(accessories.BodyAsset)
		accessories.HairOutlineAsset = GetHairOutlineAsset(accessories.HairAsset)
		accessories.MouthOutlineAsset = GetMouthOutlineAsset(accessories.MouthAsset)
		if outlineColor != nil {
			accessories.OutlineColor = *outlineColor
		} else {
			accessories.OutlineColor = color.RGB{R: 0, G: 0, B: 0}
		}
	}

	return accessories, nil
}

// GetBodyColor - Get body color with given entropy
func GetBodyColor(entropy string) (color.RGB, error) {
	// Want to generate hue between 0-360
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy[0:4], 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	outHSV := color.HSV{}
	// Generate hue
	r := rand.Init()
	r.Seed(uint32(randSeed))
	outHSV.H = float64(r.Int31n(360))
	// Generate Saturation
	randSeed, err = strconv.ParseInt(entropy[4:8], 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.Init()
	r.Seed(uint32(randSeed))
	minSatInt := int32(MinSaturation * 100)
	outHSV.S = float64(r.Int31n(100-minSatInt)+minSatInt) / 100.0
	// Generate Brightness
	randSeed, err = strconv.ParseInt(entropy[8:12], 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.Init()
	r.Seed(uint32(randSeed))
	minBInt := int32(MinBrightness * 100)
	outHSV.V = float64(r.Int31n(100-minBInt)+minBInt) / 100.0

	print(fmt.Sprintf("Got %f %f %f", outHSV.H, outHSV.S, outHSV.V))

	return outHSV.ToRGB(), nil
}

// GetHairColor - Get a complementary color with given entropy
func GetHairColor(bodyColor color.RGB, hEntropy string, sEntropy string, bEntropy string) (color.RGB, error) {
	var err error
	// Get as HSV color
	bodyColorHSV := bodyColor.ToHSV()
	hairColorHSV := color.HSV{}

	var randSeed int64
	// Want to shift the hue between 90-270
	// Get detemrinistic RNG
	randSeed, err = strconv.ParseInt(hEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}

	// Generate random shift between 90...270
	r := rand.Init()
	r.Seed(uint32(randSeed))
	hairColorHSV.H = float64(r.Int31n(270-90)+90) + bodyColorHSV.H

	// If > 360, subtract
	if hairColorHSV.H > 360 {
		hairColorHSV.H = hairColorHSV.H - 360
	}

	// Generate random saturation between MinimumSaturation - 100
	randSeed, err = strconv.ParseInt(sEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.Init()
	r.Seed(uint32(randSeed))
	minSatInt := int32(MinSaturation * 100)
	hairColorHSV.S = float64(r.Int31n(100-minSatInt)+minSatInt) / 100.0

	// Generate random brightess between MinimumBrightness - 100
	randSeed, err = strconv.ParseInt(bEntropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	r = rand.Init()
	r.Seed(uint32(randSeed))
	minBInt := int32(MinBrightness * 100)
	hairColorHSV.V = float64(r.Int31n(100-minBInt)+minBInt) / 100.0

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

// GetBodyOutlineAsset - return body outline illustration for a given body asset
func GetBodyOutlineAsset(bodyAsset Asset) *Asset {
	for _, ba := range GetAssets().GetBodyOutlineAssets() {
		if ba.FileName == bodyAsset.FileName {
			return &ba
		}
	}
	return nil
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

// GetHairOutlineAsset - return hair outline illustration for a given hair asset
func GetHairOutlineAsset(hairAsset Asset) *Asset {
	for _, ba := range GetAssets().GetHairOutlineAssets() {
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

// GetMouthOutlineAsset - return mouth outline illustration for a given mouth asset
func GetMouthOutlineAsset(mouthAsset Asset) *Asset {
	for _, ba := range GetAssets().GetMouthOutlineAssets() {
		if ba.FileName == mouthAsset.FileName {
			return &ba
		}
	}
	return nil
}
