package image

import (
	"errors"
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
	accessories.hairColor, err = GetHairColor(accessories.bodyColor, hash[6:12])

	return accessories, nil
}

// GetHairColor - Get a complementary color with given entropy
func GetHairColor(bodyColor color.RGB, entropy string) (color.RGB, error) {
	var err error
	// Get as HSL color
	bodyColorHSL := bodyColor.ToHSL()

	var randSeed int64
	var shiftedHue float64
	// Want to shift the hue between 90-270
	// Get detemrinistic RNG
	randSeed, err = strconv.ParseInt(entropy, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}

	// Generate random shift between 90...270
	r := rand.New(rand.NewSource(randSeed))
	shiftedHue = float64(r.Intn(270-90)+90) + bodyColorHSL.H

	// If > 360, subtract
	if bodyColorHSL.H+shiftedHue > 360 {
		shiftedHue = 360 - shiftedHue
	}

	hairColorHSL := color.HSL{}
	hairColorHSL.H = shiftedHue
	hairColorHSL.S = bodyColorHSL.S
	hairColorHSL.L = bodyColorHSL.L

	return hairColorHSL.ToRGB(), nil
}
