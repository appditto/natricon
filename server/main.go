package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/appditto/natricon/server/color"
	"github.com/appditto/natricon/server/image"
	"github.com/appditto/natricon/server/magickwand"
	"github.com/appditto/natricon/server/model"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var seed string

const minConvertedSize = 100  // Minimum size of PNG/WEBP/JPG converted output
const maxConvertedSize = 1000 // Maximum size of PNG/WEBP/JPG converted output

// Generate natricon with given hash
func generateIcon(hash *string, c *gin.Context) {
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
		if strings.ToLower(c.Query("outlineColor")) == "white" {
			outlineColor = &color.RGB{R: 255.0, G: 255.0, B: 255.0}
		} else {
			outlineColor = &color.RGB{R: 0.0, G: 0.0, B: 0.0}
		}
	}

	accessories, err := image.GetAccessoriesForHash(*hash, outline, outlineColor)
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

// Handle callback for donation listener
func callback(c *gin.Context) {
	var callbackData model.Callback
	err := c.BindJSON(&callbackData)
	if err != nil {
		glog.Errorf("Error processing callback")
		return
	}
	var blockData model.Block
	err = json.Unmarshal([]byte(callbackData.Block), &blockData)
	if err != nil {
		glog.Errorf("Error parsing callback block json")
		return
	}
	glog.Infof("Received callback link %s", blockData.LinkAsAccount)
	amount, err := utils.RawToNano(callbackData.Amount)
	glog.Infof("Received callback amount %f", amount)
}

// Generate natricon with given nano address
func getNano(c *gin.Context) {
	address := c.Query("address")
	valid := utils.ValidateAddress(address)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid address")
		return
	}

	sha256 := utils.AddressSha256(address, seed)

	generateIcon(&sha256, c)
}

// Testing APIs
func getRandomSvg(c *gin.Context) {
	var err error

	address := utils.GenerateAddress()
	sha256 := utils.AddressSha256(address, seed)

	accessories, err := image.GetAccessoriesForHash(sha256, false, nil)
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

func getRandom(c *gin.Context) {
	var err error

	address := utils.GenerateAddress()
	sha256 := utils.AddressSha256(address, seed)

	accessories, err := image.GetAccessoriesForHash(sha256, false, nil)
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

func getNatricon(c *gin.Context) {
	var err error

	address := c.Query("address")
	// valid := utils.BalidateAddress(address)
	// if !valid {
	// c.String(http.StatusBadRequest, "Invalid address")
	// return
	// }
	sha256 := utils.AddressSha256(address, seed)

	accessories, err := image.GetAccessoriesForHash(sha256, false, nil)
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

func testBodyDistribution() {
	wd, _ := os.Getwd()
	output := path.Join(wd, "body_distribution.csv")
	outputF, err := os.Create(output)
	defer outputF.Close()
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	var address string
	var sha256 string
	var accessories image.Accessories
	ret := "h,s,b\n"
	lt20 := 0
	lt40 := 0
	lt60 := 0
	lt80 := 0
	lt100 := 0
	for i := 0; i < 10000; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.AddressSha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256, false, nil)
		ret += fmt.Sprintf("%f,%f,%f\n", accessories.BodyColor.ToHSB().H, accessories.BodyColor.ToHSB().S*100.0, accessories.BodyColor.ToHSB().B*100.0)
		if accessories.BodyColor.ToHSB().S*100.0 < 20 {
			lt20 += 1
		} else if accessories.BodyColor.ToHSB().S*100.0 < 40 {
			lt40 += 1
		} else if accessories.BodyColor.ToHSB().S*100.0 < 60 {
			lt60 += 1
		} else if accessories.BodyColor.ToHSB().S*100.0 < 80 {
			lt80 += 1
		} else {
			lt100 += 1
		}
	}
	outputF.WriteString(ret)
	print(fmt.Sprintf("S 0-20 %d\n", lt20))
	print(fmt.Sprintf("S 20-40 %d\n", lt40))
	print(fmt.Sprintf("S 40-60 %d\n", lt60))
	print(fmt.Sprintf("S 60-80 %d\n", lt80))
	print(fmt.Sprintf("S 80-100 %d\n", lt100))
}

func testHairDistribution() {
	wd, _ := os.Getwd()
	output := path.Join(wd, "hair_distribution.csv")
	outputF, err := os.Create(output)
	defer outputF.Close()
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	var address string
	var sha256 string
	var accessories image.Accessories
	ret := "h,s,b\n"
	lt20 := 0
	lt40 := 0
	lt60 := 0
	lt80 := 0
	lt100 := 0
	for i := 0; i < 10000; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.AddressSha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256, false, nil)
		ret += fmt.Sprintf("%f,%f,%f\n", accessories.HairColor.ToHSB().H, accessories.HairColor.ToHSB().S*100.0, accessories.HairColor.ToHSB().B*100.0)
		if accessories.HairColor.ToHSB().S*100.0 < 20 {
			lt20 += 1
		} else if accessories.HairColor.ToHSB().S*100.0 < 40 {
			lt40 += 1
		} else if accessories.HairColor.ToHSB().S*100.0 < 60 {
			lt60 += 1
		} else if accessories.HairColor.ToHSB().S*100.0 < 80 {
			lt80 += 1
		} else {
			lt100 += 1
		}
	}
	outputF.WriteString(ret)
	print(fmt.Sprintf("S 0-20 %d\n", lt20))
	print(fmt.Sprintf("S 20-40 %d\n", lt40))
	print(fmt.Sprintf("S 40-60 %d\n", lt60))
	print(fmt.Sprintf("S 60-80 %d\n", lt80))
	print(fmt.Sprintf("S 80-100 %d\n", lt100))
}

func main() {
	// Get seed from env
	seed = utils.GetEnv("NATRICON_SEED", "1234567890")
	// Parse server options
	loadFiles := flag.Bool("load-files", false, "Print assets as GO arrays")
	testBodyDist := flag.Bool("test-bd", false, "Test body distribution")
	testHairDist := flag.Bool("test-hd", false, "Test hair distribution")
	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	if *loadFiles {
		LoadAssetsToArray()
		return
	} else if *testBodyDist {
		testBodyDistribution()
		return
	} else if *testHairDist {
		testHairDistribution()
		return
	}

	// Setup router
	router := gin.Default()
	router.Use(cors.Default())
	// V1 API
	router.GET("/api/v1/nano", getNano)
	// Donation callback
	router.POST("/api/nanocallback", callback)
	// For testing
	router.GET("/api/natricon", getNatricon)
	router.GET("/api/random", getRandom)
	router.GET("/api/randomsvg", getRandomSvg)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
