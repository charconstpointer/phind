package main

import (
	"net/http"
	"tickle/pkg/coordy"

	"github.com/gin-gonic/gin"
)

type TripRequest struct {
	C [2]float32 `json:"coords"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/phind", func(c *gin.Context) {
		var coords TripRequest
		if err := c.ShouldBindJSON(&coords); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := coordy.GetCoords(coords.C)
		c.JSON(http.StatusOK, gin.H{"phind": res})
	})
	r.Run()
}
