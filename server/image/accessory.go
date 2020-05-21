package image

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/appditto/natricon/server/color"
	"github.com/appditto/natricon/server/rand"
)

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
	BadgeAsset        *Asset
	OutlineColor      color.RGB
}

// Hex string regex
const hexRegexStr = "^[0-9a-fA-F]+$"

var hexRegex = regexp.MustCompile(hexRegexStr)

// GetAccessoriesForHash - Return Accessories object based on 64-character hex string
func GetAccessoriesForHash(hash string, withBadge bool, outline bool, outlineColor *color.RGB) (Accessories, error) {
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

	// Get badge
	if withBadge {
		accessories.BadgeAsset = GetBadgeAsset(accessories.BodyAsset)
	}

	// Get mouth and eyes
	targetSex := Neutral
	if accessories.BodyAsset.Sex != Neutral {
		targetSex = accessories.BodyAsset.Sex
	} else if accessories.HairAsset.Sex != Neutral {
		targetSex = accessories.HairAsset.Sex
	}
	accessories.MouthAsset, err = GetMouthAsset(hash[46:55], targetSex, accessories.BodyColor.PerceivedBrightness())
	if targetSex == Neutral && accessories.MouthAsset.Sex != Neutral {
		targetSex = accessories.MouthAsset.Sex
	}
	accessories.EyeAsset, err = GetEyeAsset(hash[55:64], targetSex, accessories.BodyColor.PerceivedBrightness())

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

// GetBadgeAsset - return badge asset for a particular body
func GetBadgeAsset(bodyAsset Asset) *Asset {
	identifier, _ := strconv.Atoi(strings.Split(bodyAsset.FileName, "_")[0])
	searchStr := fmt.Sprintf("b%d", identifier)
	for _, v := range GetAssets().GetBadgeAssets() {
		if strings.Contains(v.FileName, searchStr) {
			return &v
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
func GetEyeAsset(entropy string, sex Sex, luminosity float64) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	eyeAssetOptions := GetAssets().GetEyeAssets(sex, luminosity)

	r := rand.Init()
	r.Seed(uint32(randSeed))
	eyeIndex := r.Int31n(int32(len(eyeAssetOptions)))

	return eyeAssetOptions[eyeIndex], nil
}

// GetEyeAsset - return hair illustration to use with given entropy
func GetMouthAsset(entropy string, sex Sex, luminosity float64) (Asset, error) {
	// Get detemrinistic RNG
	randSeed, err := strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return Asset{}, err
	}

	mouthAssetOptions := GetAssets().GetMouthAssets(sex, luminosity)

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
