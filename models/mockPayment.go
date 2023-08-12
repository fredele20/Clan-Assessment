package models

type MockPayment struct {
	Id             int            `json:"id"`
	BookingId      int            `json:"bookingId"`
	Amount         int            `json:"amount"`
	MeansOfPayment MeansOfPayment `json:"meansOfPayment"`
}

type MeansOfPayment string

const (
	Transfer MeansOfPayment = "transfer"
	Cash     MeansOfPayment = "cash"
)


func (m MeansOfPayment) IsValid() bool {
	switch m {
	case Transfer, Cash:
		return true
	default:
		return false
	}
}
