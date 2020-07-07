package controller

import (
	"fmt"
	"os"
	"path"

	"github.com/appditto/natricon/server/db"
	"github.com/appditto/natricon/server/image"
	"github.com/appditto/natricon/server/spc"
	"github.com/appditto/natricon/server/utils"
	"github.com/gin-gonic/gin"
)

// Go routine for processing stats messages
func StatsWorker(statsChan <-chan *gin.Context) {
	// Process stats
	for c := range statsChan {
		// Update unique addresses
		db.GetDB().UpdateStatsAddress(c.Query("address"))
		// Update daily
		db.GetDB().UpdateStatsDate(c.Query("address"))
		// Update by service
		if c.Query("svc") != "" {
			db.GetDB().UpdateStatsByService(c.Query("svc"), c.Query("address"))
		}
	}
}

// Stats API
func Stats(c *gin.Context) {
	// Get # of unique natricons served
	numServed := db.GetDB().StatsUniqueAddresses()
	numServedTotal := db.GetDB().StatsTotal()
	svcStats := db.GetDB().ServiceStats()
	daily := db.GetDB().DailyStats()

	// Return response
	c.JSON(200, gin.H{
		"unique_served": numServed,
		"total_served":  numServedTotal,
		"services":      svcStats,
		"daily":         daily,
	})
}

// For generating CSV documents for algorithm analysis
func TestBodyDistribution(seed string) {
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
	ret := "h,s,b,pb\n"
	lt20 := 0
	lt40 := 0
	lt60 := 0
	lt80 := 0
	lt100 := 0
	for i := 0; i < 10000; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.AddressSha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256, spc.BTNone, false, nil)
		ret += fmt.Sprintf("%f,%f,%f,%f\n", accessories.BodyColor.ToHSB().H, accessories.BodyColor.ToHSB().S*100.0, accessories.BodyColor.ToHSB().B*100.0, accessories.BodyColor.PerceivedBrightness())
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

func TestHairDistribution(seed string) {
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
	ret := "h,s,b,pb\n"
	lt20 := 0
	lt40 := 0
	lt60 := 0
	lt80 := 0
	lt100 := 0
	for i := 0; i < 10000; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.AddressSha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256, spc.BTNone, false, nil)
		ret += fmt.Sprintf("%f,%f,%f,%f\n", accessories.HairColor.ToHSB().H, accessories.HairColor.ToHSB().S*100.0, accessories.HairColor.ToHSB().B*100.0, accessories.HairColor.PerceivedBrightness())
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
