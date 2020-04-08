package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/appditto/natricon/nano"
	"github.com/gin-gonic/gin"
)

func getNatricon(c *gin.Context) {
	address := c.Query("address")
	valid := nano.ValidateAddress(address)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid address")
		return
	}
	sha256 := nano.AddressSha256(address)
	c.String(http.StatusOK, "%s", sha256)
}

func main() {
	// Parse server options
	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	// Setup router
	router := gin.Default()
	router.GET("/natricon", getNatricon)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
