package handlers

import (
	"clan-africa/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFlights(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	FlightRoute(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/flights", nil)
	router.ServeHTTP(w, req)

	var flights models.Flight
	json.Unmarshal(w.Body.Bytes(), &flights)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, flights)
}

func TestGetFlightOptions(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	FlightRoute(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/flights/options", nil)
	router.ServeHTTP(w, req)

	var flightsOptions models.FlightOptions
	json.Unmarshal(w.Body.Bytes(), &flightsOptions)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, flightsOptions)
}
