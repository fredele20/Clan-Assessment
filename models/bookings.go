package models

type Booking struct {
	Id       int                `json:"id"`
	FlightId int                `json:"flightId"`
	Option   BookingOption      `json:"option"`
	Price    BookingOptionPrice `json:"price"`
}

type BookingOption string
type BookingOptionPrice int

const (
	FirstClassPrice   BookingOptionPrice = 30000
	MiddleClassPrice  BookingOptionPrice = 20000
	RegularClassPrice BookingOptionPrice = 10000
)

const (
	FirstClass   BookingOption = "firstClass"
	MiddleClass  BookingOption = "secondClass"
	RegularClass BookingOption = "regular"
)

func (b BookingOption) IsValid() bool {
	switch b {
	case FirstClass, MiddleClass, RegularClass:
		return true
	default:
		return false
	}
}
