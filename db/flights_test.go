package db

import (
	"clan-africa/models"
	"testing"

	"github.com/stretchr/testify/require"
)

type testData struct {
	Filter       *models.FilterFlight
	FlightOutput models.ListFlight
}

func TestGetFlights(t *testing.T) {
	tests := []testData{
		{Filter: &models.FilterFlight{Departure: "lagos", Destination: "abuja"}},
		{Filter: &models.FilterFlight{Departure: "ajah", Destination: "abuja"}},
		{Filter: &models.FilterFlight{Departure: "ilorin", Destination: "sokoto"}},
		{Filter: &models.FilterFlight{Departure: "", Destination: ""}, FlightOutput: models.ListFlight{Count: 5}},
	}

	for i := 0; i < len(tests); i++ {
		output1, err := GetFlights(tests[0].Filter)
		require.NoError(t, err)
		require.NotEmpty(t, output1)
		// require.Len(t, output1)

		output2, err := GetFlights(tests[1].Filter)
		require.NoError(t, err)
		require.Empty(t, output2)

		output3, err := GetFlights(tests[2].Filter)
		require.NoError(t, err)
		require.Empty(t, output3)

		output4, err := GetFlights(tests[3].Filter)
		require.NoError(t, err)
		require.NotEmpty(t, output4)
		require.Equal(t, output4.Count, tests[3].FlightOutput.Count)
	}
}

func TestGetFlightOptions(t *testing.T) {
	result, err := GetFlightOptions()
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
