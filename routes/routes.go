package routes

import (
	"clan-africa/db"
	"clan-africa/models"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFlights() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var filter models.FilterFlight
		fmt.Println(filter)
		if err := ctx.BindQuery(&filter); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error binding json"})
			return
		}

		flights, err := db.GetFlights(&filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting flights"})
			return
		}

		ctx.JSON(http.StatusOK, flights)
	}
}

func GetFlightOptions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		flightOptions, err := db.GetFlightOptions()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something failed! could not get available flight options"})
			return
		}

		ctx.JSON(http.StatusOK, flightOptions)
	}
}

func BookFlight() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var booking models.Booking
		if err := ctx.BindJSON(&booking); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		flightId := ctx.Param("id")
		flightIdInt, _ := strconv.Atoi(flightId)

		bookFlight, err := db.BookFlight(flightIdInt, &booking)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		location := url.URL{Path: "https://google.com"}

		ctx.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
		ctx.JSON(http.StatusCreated, bookFlight)
	}
}

func ConfirmBookingPayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payment models.MockPayment
		if err := ctx.BindJSON(&payment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		bookingId := ctx.Param("id")
		bookingIdInt, _ := strconv.Atoi(bookingId)

		confirmPayment, err := db.ConfirmBooking(bookingIdInt, &payment)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, confirmPayment)
	}
}