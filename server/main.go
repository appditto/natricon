package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/appditto/natricon/image"
	"github.com/appditto/natricon/nano"
	"github.com/gin-gonic/gin"
)

var seed *string

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

	c.JSON(200, gin.H{
		"bodyColor": accessories.BodyColor.ToHTML(),
		"hairColor": accessories.HairColor.ToHTML(),
		"hash":      sha256,
	})
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

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
