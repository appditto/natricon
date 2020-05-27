package main

import (
	"flag"
	"fmt"

	"github.com/appditto/natricon/server/controller"
	"github.com/appditto/natricon/server/net"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jasonlvhit/gocron"
	_ "go.uber.org/automaxprocs"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "https://natricon.com"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

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
	wsUrl := flag.String("nano-ws-url", "", "Nano WS Url to use for tracking donation account")
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
	router.Use(CorsMiddleware())

	// Setup socket IO server
	sio, _ := socketio.NewServer(nil)
	sio.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("bcast")
		s.Emit("connected", s.ID())
		return nil
	})
	go sio.Serve()
	defer sio.Close()
	router.GET("/socket.io/*any", gin.WrapH(sio))
	router.POST("/socket.io/*any", gin.WrapH(sio))

	// Setup channel for stats processing job
	statsChan := make(chan *gin.Context, 100)

	// Setup natricon controller
	natriconController := controller.NatriconController{
		Seed:         seed,
		StatsChannel: &statsChan,
	}
	// Setup nano controller
	nanoController := controller.NanoController{
		RPCClient:       rpcClient,
		SIOServer:       sio,
		DonationAccount: utils.GetEnv("DONATION_ACCOUNT", ""),
	}

	// V1 API
	router.GET("/api/v1/nano", natriconController.GetNano)
	// Stats
	router.GET("/api/v1/nano/stats", controller.Stats)
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

	// Start Nano WS client
	if *wsUrl != "" {
		go net.StartNanoWSClient(*wsUrl, utils.GetEnv("DONATION_ACCOUNT", ""), nanoController.Callback)
	}

	// Start stats worker
	go controller.StatsWorker(statsChan)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
