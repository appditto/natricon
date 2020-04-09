package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/appditto/natricon/image"
	"github.com/appditto/natricon/nano"
	"github.com/gin-gonic/gin"
)

var seed *string

var testhtml string = `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
.square {
  height: 200px;
  width: 200px;
  background-color: #FFF;
}
.squareTop {
  height: 50px;
  width: 200px;
  background-color: #000;
}
</style>
</head>
<body>
<div>address_1</div>
<div class="squareTop"></div>
<div class="square"></div>
</body>
</html> 
`

func getRandom(c *gin.Context) {
	var err error

	address := nano.GenerateAddress()
	sha256 := nano.AddressSha256(address, *seed)

	accessories, err := image.GetAccessoriesForHash(sha256)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	newHTML := strings.Replace(testhtml, "#000", "#"+accessories.HairColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "#FFF", "#"+accessories.BodyColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "address_1", address, -1)
	c.Data(200, "text/html; charset=utf-8", []byte(newHTML))
}

func getNatricon(c *gin.Context) {
	var err error

	address := c.Query("address")
	valid := nano.ValidateAddress(address)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid address")
		return
	}
	sha256 := nano.AddressSha256(address, *seed)

	accessories, err := image.GetAccessoriesForHash(sha256)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	newHTML := strings.Replace(testhtml, "#000", "#"+accessories.HairColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "#FFF", "#"+accessories.BodyColor.ToHTML(), -1)
	newHTML = strings.Replace(newHTML, "address_1", address, -1)
	c.Data(200, "text/html; charset=utf-8", []byte(newHTML))
	/*
		c.JSON(200, gin.H{
			"bodyColor": accessories.BodyColor.ToHTML(),
			"hairColor": accessories.HairColor.ToHTML(),
			"hash":      sha256,
		})*/
}

func main() {
	// Parse server options
	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	seed = flag.String("seed", "1234567890", "Seed to use for icon generation")
	flag.Parse()

	// Setup router
	router := gin.Default()
	router.GET("/natricon", getNatricon)
	router.GET("/random", getRandom)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
