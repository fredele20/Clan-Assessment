package db

import (
	"clan-africa/models"
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

//go:embed flights.json
var flightsJson string

func PopulateFlight() {
	flights := []models.Flight{
		{
			Id:          1,
			Departure:   "Lagos",
			Destination: "Abuja",
			Date:        time.Date(2023, time.August, 15, 18, 0, 0, 0, time.UTC),
		},

		{
			Id:          2,
			Departure:   "Ilorin",
			Destination: "Port Harcourt",
			Date:        time.Date(2023, time.August, 13, 13, 0, 0, 0, time.UTC),
		},

		{
			Id:          3,
			Departure:   "Lagos",
			Destination: "Port Harcourt",
			Date:        time.Date(2023, time.August, 20, 12, 0, 0, 0, time.UTC),
		},

		{
			Id:          4,
			Departure:   "Ondo",
			Destination: "Lagos",
			Date:        time.Date(2023, time.August, 20, 12, 0, 0, 0, time.UTC),
		},

		{
			Id:          5,
			Departure:   "Lagos",
			Destination: "Abuja",
			Date:        time.Date(2023, time.August, 15, 9, 0, 0, 0, time.UTC),
		},
	}

	file, _ := json.MarshalIndent(flights, "", "")
	_ = ioutil.WriteFile("db/flights.json", file, 0644)
}

func PopulateFlightOptions() {
	options := []models.FlightOptions{
		{ Option: "First Class", Price: 30000 },
		{ Option: "Middle Class", Price: 20000 },
		{ Option: "Regular", Price: 10000 },
	}
	data, _ := json.MarshalIndent(options, "", "")
	_ = ioutil.WriteFile("db/flightOptions.json", data, 0644)
}

func GetFlights(filter *models.FilterFlight) (*models.ListFlight, error) {

	var flights []models.Flight
	err := json.Unmarshal([]byte(flightsJson), &flights)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	if filter.Departure == "" && filter.Destination == "" {
		return &models.ListFlight{
			Data:  flights,
			Count: len(flights),
		}, nil
	}

	var filteredFlight []models.Flight
	for _, flight := range flights {
		if strings.ToLower(flight.Departure) == strings.ToLower(filter.Departure) &&
			strings.ToLower(flight.Destination) == strings.ToLower(filter.Destination) ||
			flight.Date == filter.Date {
			filteredFlight = append(filteredFlight, flight)
		}
	}

	return &models.ListFlight{
		Data:  filteredFlight,
		Count: len(filteredFlight),
	}, nil
}

//go:embed flightOptions.json
var flightOptionsJson string
func GetFlightOptions() (*[]models.FlightOptions, error) {
	options := []models.FlightOptions{}
	err := json.Unmarshal([]byte(flightOptionsJson), &options)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}
	return &options, nil
}
