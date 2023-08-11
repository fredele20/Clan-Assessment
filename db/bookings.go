package db

import (
	"clan-africa/models"
	_ "embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

func BookFlight(flightId int, payload *models.Booking) (*models.Booking, error) {
	filename := "db/bookings.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	bookings := []models.Booking{}

	json.Unmarshal(file, &bookings)

	var flights []models.Flight
	err = json.Unmarshal([]byte(flightsJson), &flights)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for i := 0; i < len(flights); i++ {
		if flightId == flights[i].Id {
			payload.FlightId = flightId
			payload.Id = len(bookings) + 1
			if payload.Type == "" {
				return nil, errors.New("please select booking option")
			}
			if !payload.Type.IsValid() {
				return nil, errors.New("Invalid booking option")
			}

			bookings = append(bookings, *payload)

			dataBytes, _ := json.MarshalIndent(bookings, "", "")
			_ = ioutil.WriteFile(filename, dataBytes, 0644)

			return payload, nil

		} else {
			return nil, errors.New("No flight with the given ID")
		}
	}
	return payload, nil
}
