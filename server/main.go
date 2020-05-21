package main

import (
	"flag"
	"fmt"

	"github.com/appditto/natricon/server/controller"
	"github.com/appditto/natricon/server/net"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jasonlvhit/gocron"
)

/* TODO - remove me
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
}*/

func main() {
	// Get seed from env
	seed := utils.GetEnv("NATRICON_SEED", "1234567890")
	// Parse server options
	/* TODO - remove me
	loadFiles := flag.Bool("load-files", false, "Print assets as GO arrays")
	testBodyDist := flag.Bool("test-bd", false, "Test body distribution")
	testHairDist := flag.Bool("test-hd", false, "Test hair distribution")*/

	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	rpcUrl := flag.String("rpc-url", "", "Optional URL to use for nano RPC Client")
	flag.Parse()

	// TODO - remove me
	/*
		if *loadFiles {
			LoadAssetsToArray()
			return
		} else if *testBodyDist {
			testBodyDistribution()
			return
		} else if *testHairDist {
			testHairDistribution()
			return
		}*/

	var rpcClient *net.RPCClient
	if *rpcUrl != "" {
		glog.Infof("RPC Client configured at %s", *rpcUrl)
		rpcClient = &net.RPCClient{Url: *rpcUrl}
	}

	// Setup router
	router := gin.Default()
	router.Use(cors.Default())

	// Setup natricon controller
	natriconController := controller.NatriconController{
		Seed: seed,
	}
	// Setup nano controller
	nanoController := controller.NanoController{
		RPCClient: rpcClient,
	}

	// V1 API
	router.GET("/api/v1/nano", natriconController.GetNano)
	// Donation callback
	router.POST("/api/nanocallback", nanoController.Callback)
	// For testing
	router.GET("/api/natricon", natriconController.GetNatricon)
	router.GET("/api/random", natriconController.GetRandom)
	router.GET("/api/randomsvg", natriconController.GetRandomSvg)

	// Setup cron job to check for missed donations
	if !gin.IsDebugging() {
		go func() {
			gocron.Every(10).Minutes().Do(nanoController.CheckMissedCallbacks)
			<-gocron.Start()
		}()
	}

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
