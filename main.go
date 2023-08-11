package main

import (
	"clan-africa/db"
	"clan-africa/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.PopulateFlight()
	db.PopulateFlightOptions()

	router := gin.New()
	router.Use(gin.Logger())

	handlers.FlightRoute(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run("127.0.0.1:6000")

}
