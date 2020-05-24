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

	_ "go.uber.org/automaxprocs"
)

func main() {
	// Get seed from env
	seed := utils.GetEnv("NATRICON_SEED", "1234567890")
	// Parse server options
	loadFiles := flag.Bool("load-files", false, "Print assets as GO arrays")
	testBodyDist := flag.Bool("test-bd", false, "Test body distribution")
	testHairDist := flag.Bool("test-hd", false, "Test hair distribution")

	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	rpcUrl := flag.String("rpc-url", "", "Optional URL to use for nano RPC Client")
	flag.Parse()

	if *loadFiles {
		LoadAssetsToArray()
		return
	}

	if *testBodyDist {
		controller.TestBodyDistribution(seed)
		return
	} else if *testHairDist {
		controller.TestHairDistribution(seed)
		return
	}

	var rpcClient *net.RPCClient
	if *rpcUrl != "" {
		glog.Infof("RPC Client configured at %s", *rpcUrl)
		rpcClient = &net.RPCClient{Url: *rpcUrl}
	}

	// Setup router
	router := gin.Default()
	router.Use(cors.Default())

	// Setup channel for stats processing job
	statsChan := make(chan *gin.Context, 100)

	// Setup natricon controller
	natriconController := controller.NatriconController{
		Seed:         seed,
		StatsChannel: &statsChan,
	}
	// Setup nano controller
	nanoController := controller.NanoController{
		RPCClient: rpcClient,
	}

	// V1 API
	router.GET("/api/v1/nano", natriconController.GetNano)
	// Stats
	router.GET("/api/v1/stats", controller.Stats)
	// Donation callback
	router.POST("/api/nanocallback", nanoController.Callback)
	// For testing
	router.GET("/api/natricon", natriconController.GetNatricon)
	router.GET("/api/random", natriconController.GetRandom)
	router.GET("/api/randomsvg", natriconController.GetRandomSvg)

	// Setup cron jobs
	if !gin.IsDebugging() {
		go func() {
			// Checking missed donations
			gocron.Every(10).Minutes().Do(nanoController.CheckMissedCallbacks)
			// Updating principal rep requirement
			gocron.Every(30).Minutes().Do(nanoController.UpdatePrincipalWeight)
			// Update principal reps, this is heavier so dont do it so often
			gocron.Every(30).Minutes().Do(nanoController.UpdatePrincipalReps)
			<-gocron.Start()
		}()
	}

	// Start stats worker
	go controller.StatsWorker(statsChan)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
