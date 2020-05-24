package controller

import "github.com/gin-gonic/gin"

func StatsWorker(statsChan <-chan *gin.Context) {
	// Process stats
	for c := range statsChan {
		print(c.Query("address"))
	}
}
