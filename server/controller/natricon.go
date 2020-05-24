package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/appditto/natricon/server/color"
	"github.com/appditto/natricon/server/image"
	"github.com/appditto/natricon/server/magickwand"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-gonic/gin"
)

const minConvertedSize = 100  // Minimum size of PNG/WEBP/JPG converted output
const maxConvertedSize = 1000 // Maximum size of PNG/WEBP/JPG converted output

type NatriconController struct {
	Seed string
}

// Special addresses
type Vanity struct {
	// Optional fields
	hash  string          // Will generate the natricon with specific hash
	badge image.BadgeType // Will generate natricon with specified badge
	// If using any of the below then ALL of them are required
	bodyColor    *color.RGB
	hairColor    *color.RGB
	bodyAssetID  int
	hairAssetID  int
	mouthAssetID int
	eyeAssetID   int
}

var vanities = map[string]*Vanity{
	/* Example to base off of a hash
	"2535ce406f14c289f09e3b471ef9744e36cc0f585b23cfaafcc6412e283dacb4": {
		hash:  "2f2f45946be8ee4f4a9fdc328f2ebb2ba6a163fbf4c8a5c8f5e23d43790ef7d8",
		check: true,
	},*/
	"2535ce406f14c289f09e3b471ef9744e36cc0f585b23cfaafcc6412e283dacb4": {
		bodyColor:    color.HTMLToRGBAlt("#bababa"),
		hairColor:    color.HTMLToRGBAlt("#1378f2"),
		bodyAssetID:  18,
		hairAssetID:  14,
		mouthAssetID: 15,
		eyeAssetID:   12,
		badge:        image.BTDonor,
	},
}

// APIs
// Generate natricon with given nano address
func (nc NatriconController) GetNano(c *gin.Context) {
	address := c.Query("address")
	valid := utils.ValidateAddress(address)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid address")
		return
	}

	var sha256 string
	var badgeType image.BadgeType
	specialNatricon := false
	pubKey := utils.AddressToPub(address)
	vanity := vanities[pubKey]
	if vanity == nil {
		sha256 = utils.PKSha256(pubKey, nc.Seed)
		badgeType = image.GetBadgeSvc().GetBadgeType(pubKey)
	} else {
		badgeType = vanity.badge
		if badgeType == "" {
			badgeType = image.BTNone
		}
		if vanity.bodyAssetID > 0 && vanity.hairAssetID > 0 && vanity.eyeAssetID > 0 && vanity.bodyColor != nil && vanity.hairColor != nil {
			specialNatricon = true
		} else if vanity.hash == "" {
			sha256 = utils.PKSha256(pubKey, nc.Seed)
		} else {
			sha256 = vanity.hash
		}
	}

	if specialNatricon {
		generateSpecialIcon(vanity, badgeType, c)
	} else {
		generateIcon(&sha256, badgeType, c)
	}
}

// TODO - remove us: Testing APIs
func (nc NatriconController) GetRandomSvg(c *gin.Context) {
	var err error

	address := utils.GenerateAddress()
	sha256 := utils.AddressSha256(address, nc.Seed)

	accessories, err := image.GetAccessoriesForHash(sha256, image.BTNone, false, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	bodyHsv := accessories.BodyColor.ToHSB()
	hairHsv := accessories.HairColor.ToHSB()
	deltaHsv := color.HSB{}
	deltaHsv.H = hairHsv.H - bodyHsv.H
	deltaHsv.S = hairHsv.S - bodyHsv.S
	deltaHsv.B = hairHsv.B - bodyHsv.B
	svg, err := image.CombineSVG(accessories)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error occured")
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}

func (nc NatriconController) GetRandom(c *gin.Context) {
	var err error

	address := utils.GenerateAddress()
	sha256 := utils.AddressSha256(utils.AddressToPub(address), nc.Seed)

	accessories, err := image.GetAccessoriesForHash(sha256, image.BTNone, false, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	bodyHsv := accessories.BodyColor.ToHSB()
	hairHsv := accessories.HairColor.ToHSB()
	bodyHsl := accessories.BodyColor.ToHSL()
	hairHsl := accessories.HairColor.ToHSL()
	deltaHsv := color.HSB{}
	deltaHsv.H = hairHsv.H - bodyHsv.H
	deltaHsv.S = hairHsv.S - bodyHsv.S
	deltaHsv.B = hairHsv.B - bodyHsv.B
	svg, err := image.CombineSVG(accessories)
	var svgStr string
	if err != nil {
		svgStr = "Error"
	} else {
		svgStr = string(svg)
	}
	c.JSON(200, gin.H{
		"bodyColor": accessories.BodyColor.ToHTML(false),
		"hairColor": accessories.HairColor.ToHTML(false),
		"hash":      sha256,
		"bodyH":     int16(bodyHsv.H),
		"bodyS":     int16(bodyHsv.S * 100.0),
		"bodyV":     int16(bodyHsv.B * 100.0),
		"bodyHSLH":  int16(bodyHsl.H),
		"bodyHSLS":  int16(bodyHsl.S * 100.0),
		"bodyHSLV":  int16(bodyHsl.L * 100.0),
		"hairH":     int16(hairHsv.H),
		"hairS":     int16(hairHsv.S * 100.0),
		"hairV":     int16(hairHsv.B * 100.0),
		"hairHSLH":  int16(hairHsl.H),
		"hairHSLS":  int16(hairHsl.S * 100.0),
		"haiorHSLV": int16(hairHsl.L * 100.0),
		"deltaH":    int16(deltaHsv.H),
		"deltaS":    int16(deltaHsv.S * 100.0),
		"deltaV":    int16(deltaHsv.B * 100.0),
		"address":   address,
		"svg":       svgStr,
	})
	/*newHTML := strings.Replace(testhtml, "#000", "#"+accessories.HairColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "#FFF", "#"+accessories.BodyColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "address_1", address, -1)
	c.Data(200, "text/html; charset=utf-8", []byte(newHTML))*/
}

func (nc NatriconController) GetNatricon(c *gin.Context) {
	var err error

	address := c.Query("address")
	// valid := utils.BalidateAddress(address)
	// if !valid {
	// c.String(http.StatusBadRequest, "Invalid address")
	// return
	// }
	sha256 := utils.AddressSha256(utils.AddressToPub(address), nc.Seed)

	accessories, err := image.GetAccessoriesForHash(sha256, image.BTNone, false, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	bodyHsv := accessories.BodyColor.ToHSB()
	hairHsv := accessories.HairColor.ToHSB()
	deltaHsv := color.HSB{}
	deltaHsv.H = hairHsv.H - bodyHsv.H
	deltaHsv.S = hairHsv.S - bodyHsv.S
	deltaHsv.B = hairHsv.B - bodyHsv.B
	c.JSON(200, gin.H{
		"bodyColor": accessories.BodyColor.ToHTML(false),
		"hairColor": accessories.HairColor.ToHTML(false),
		"hash":      sha256,
		"bodyH":     int16(bodyHsv.H),
		"bodyS":     int16(bodyHsv.S * 100.0),
		"bodyV":     int16(bodyHsv.B * 100.0),
		"hairH":     int16(hairHsv.H),
		"hairS":     int16(hairHsv.S * 100.0),
		"hairV":     int16(hairHsv.B * 100.0),
		"deltaH":    int16(deltaHsv.H),
		"deltaS":    int16(deltaHsv.S * 100.0),
		"deltaV":    int16(deltaHsv.B * 100.0),
		"address":   address,
	})
}

// Generate natricon with given hash
func generateIcon(hash *string, badgeType image.BadgeType, c *gin.Context) {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		c.String(http.StatusBadRequest, "%s", "Valid formats are 'svg', 'png', or 'webp'")
		return
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			c.String(http.StatusBadRequest, "%s", "Size is required when format is not svg")
			return
		}
		size, err = strconv.Atoi(c.Query("size"))
		if err != nil || size < minConvertedSize || size > maxConvertedSize {
			c.String(http.StatusBadRequest, "%s", fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
			return
		}
	}

	outline := strings.ToLower(c.Query("outline")) == "true"
	// Get outline and outline color info, black is default
	var outlineColor *color.RGB
	if outline {
		if strings.ToLower(c.Query("outline_color")) == "white" {
			outlineColor = &color.RGB{R: 255.0, G: 255.0, B: 255.0}
		} else {
			outlineColor = &color.RGB{R: 0.0, G: 0.0, B: 0.0}
		}
	}

	accessories, err := image.GetAccessoriesForHash(*hash, badgeType, outline, outlineColor)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	bodyHsv := accessories.BodyColor.ToHSB()
	hairHsv := accessories.HairColor.ToHSB()
	deltaHsv := color.HSB{}
	deltaHsv.H = hairHsv.H - bodyHsv.H
	deltaHsv.S = hairHsv.S - bodyHsv.S
	deltaHsv.B = hairHsv.B - bodyHsv.B
	svg, err := image.CombineSVG(accessories)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error occured")
		return
	}
	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = magickwand.ConvertSvgToBinary(svg, magickwand.ImageFormat(format), uint(size))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error occured")
			return
		}
		c.Data(200, fmt.Sprintf("image/%s", format), converted)
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}

// Generate icon for special accounts
func generateSpecialIcon(vanity *Vanity, badgeType image.BadgeType, c *gin.Context) {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		c.String(http.StatusBadRequest, "%s", "Valid formats are 'svg', 'png', or 'webp'")
		return
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			c.String(http.StatusBadRequest, "%s", "Size is required when format is not svg")
			return
		}
		size, err = strconv.Atoi(c.Query("size"))
		if err != nil || size < minConvertedSize || size > maxConvertedSize {
			c.String(http.StatusBadRequest, "%s", fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
			return
		}
	}

	outline := strings.ToLower(c.Query("outline")) == "true"
	// Get outline and outline color info, black is default
	var outlineColor *color.RGB
	if outline {
		if strings.ToLower(c.Query("outline_color")) == "white" {
			outlineColor = &color.RGB{R: 255.0, G: 255.0, B: 255.0}
		} else {
			outlineColor = &color.RGB{R: 0.0, G: 0.0, B: 0.0}
		}
	}

	accessories := image.GetSpecificNatricon(badgeType, outline, outlineColor, vanity.bodyColor, vanity.hairColor, vanity.bodyAssetID, vanity.hairAssetID, vanity.mouthAssetID, vanity.eyeAssetID)
	svg, err := image.CombineSVG(accessories)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error occured")
		return
	}
	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = magickwand.ConvertSvgToBinary(svg, magickwand.ImageFormat(format), uint(size))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error occured")
			return
		}
		c.Data(200, fmt.Sprintf("image/%s", format), converted)
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}
