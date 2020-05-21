package controller

import (
	"encoding/json"

	"github.com/appditto/natricon/server/model"
	"github.com/appditto/natricon/server/net"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Account donations are checked to
const donationAccount = "nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd"

// Donations at or above this threshold will award "vip" status
const donationThresholdNano = 2.0

// Maximum donation days allowed at a time (cannot buy bigger chunks than this at a time)
const maxVIPDays = 120

type NanoController struct {
	RPCClient *net.RPCClient
}

// Handle callback for donation listener
func (nc NanoController) Callback(c *gin.Context) {
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

// Cron job for checking missed callbacks
func (nc NanoController) CheckMissedCallbacks() {
	if nc.RPCClient == nil {
		return
	}
	historyResponse, err := nc.RPCClient.MakeAccountHistoryRequest(
		donationAccount,
		10,
	)
	if err != nil {
		glog.Fatalf("Error occured checking donation account history %s", err)
		return
	}
	// TODO - implement
	for i := 0; i < len(historyResponse.History); i++ {
		glog.Infof("Found history item %s", historyResponse.History[i].Hash)
	}
}
