package controller

import (
	"github.com/appditto/natricon/server/db"
	"github.com/gin-gonic/gin"
)

func StatsWorker(statsChan <-chan *gin.Context) {
	// Process stats
	for c := range statsChan {
		// Update unique addresses
		db.GetDB().UpdateStatsAddress(c.Query("address"))
	}
}

// Stats API
func Stats(c *gin.Context) {
	// Get # of unique natricons served
	numServed := db.GetDB().StatsUniqueAddresses()

	// Return response
	c.JSON(200, gin.H{
		"num_unique_natricons_served": numServed,
	})
}
