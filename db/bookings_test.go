package db

import (
	"clan-africa/models"
)

type bookingTestData struct {
	payload *models.Booking
}

// TODO: Something wrong with this test: I will come back to it

// func TestBookFlight(t *testing.T) {
// 	tests := []bookingTestData{
// 		{ payload: &models.Booking{ Id: 8, FlightId: 1, Option: "firstClass", Price: 30000 }, },
// 	}

// 	for _, test := range tests {
// 		fmt.Println(test.payload.FlightId, test.payload)
// 		bookFlight1, err := BookFlight(test.payload.FlightId, test.payload)
// 		require.NoError(t, err)
// 		require.Empty(t, bookFlight1)
// 	}
// }
