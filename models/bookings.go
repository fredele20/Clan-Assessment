package models


type Booking struct {
	Id int `json:"id"`
	FlightId int `json:"flightId"`
	Type BookingType `json:"type"`
}

type BookingType string

const (
	FirstClass BookingType = "firstClass"
)

func (b BookingType) IsValid() bool {
	switch b {
	case FirstClass:
		return true
	default:
		return false
	}
}