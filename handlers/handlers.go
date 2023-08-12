package handlers

import (
	"clan-africa/routes"

	"github.com/gin-gonic/gin"
)


func FlightRoute(incomingRoute *gin.Engine) {
	incomingRoute.GET("/flights", routes.GetFlights())
	incomingRoute.GET("/flights/options", routes.GetFlightOptions())
	incomingRoute.POST("/flights/:id/book", routes.BookFlight())
	incomingRoute.POST("/bookings/:id/confirm-payment", routes.ConfirmBookingPayment())
}