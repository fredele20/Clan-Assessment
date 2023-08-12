package models

import "time"

type Flight struct {
	Id          int       `json:"id"`
	Departure   string    `json:"departure"`
	Destination string    `json:"destination"`
	Date        time.Time `json:"date"`
}

type FilterFlight struct {
	Departure   string    `form:"departure"`
	Destination string    `form:"destination"`
	Date        time.Time `form:"date"`
}

type ListFlight struct {
	Data  []Flight `json:"data"`
	Count int      `json:"count"`
}

type FlightOptions struct {
	Option string `json:"option"`
	Price  int    `json:"price"`
}
