package db

import (
	"clan-africa/models"
	"encoding/json"
	_ "embed"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func BookFlight(flightId int, payload *models.Booking) (*models.Booking, error) {
	jsonFile, err := os.Open("db/bookings.json")
	defer jsonFile.Close()

	file, err := ioutil.ReadAll(jsonFile)
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
			if payload.Option == "" {
				return nil, errors.New("please select booking option")
			}
			if !payload.Option.IsValid() {
				return nil, errors.New("Invalid booking option")
			}

			if payload.Option == models.FirstClass {
				payload.Price = models.FirstClassPrice
			} else if payload.Option == models.MiddleClass {
				payload.Price = models.MiddleClassPrice
			} else {
				payload.Price = models.RegularClassPrice
			}

			bookings = append(bookings, *payload)

			dataBytes, _ := json.MarshalIndent(bookings, "", "")
			_ = ioutil.WriteFile(jsonFile.Name(), dataBytes, 0644)

			return payload, nil

		} else {
			return nil, errors.New("No flight with the given ID")
		}
	}
	return payload, nil
}

//go:embed bookings.json
var bookingJson string
func ConfirmBooking(bookingId int, payload *models.MockPayment) (*models.MockPayment, error) {
	jsonFile, err := os.Open("db/mockPayment.json")
	defer jsonFile.Close()

	file, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	payments := []models.MockPayment{}

	json.Unmarshal(file, &payments)

	var bookings []models.Booking
	err = json.Unmarshal([]byte(bookingJson), &bookings)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for i := 0; i < len(bookings); i++ {
		if bookingId == bookings[i].Id {
			payload.BookingId = bookingId
			payload.Amount = int(bookings[i].Price)
			payload.Id = len(payments) + 1
			if payload.MeansOfPayment == "" {
				return nil, errors.New("please select a payment medium")
			}
			if !payload.MeansOfPayment.IsValid() {
				return nil, errors.New("Invalid payment medium")
			}

			payments = append(payments, *payload)

			dataBytes, _ := json.MarshalIndent(payments, "", "")
			_ = ioutil.WriteFile(jsonFile.Name(), dataBytes, 0644)

			return payload, nil

		} else {
			return nil, errors.New("No booking with the given ID")
		}
	}
	return payload, nil
}
