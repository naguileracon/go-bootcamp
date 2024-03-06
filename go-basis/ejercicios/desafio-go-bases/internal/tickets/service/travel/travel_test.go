package travel

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/structure"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("success - get total tickets of an existent country", func(t *testing.T) {
		// arrange - given
		destination := "Japan"
		tickets := []structure.Ticket{{
			DestinationCountry: destination,
		},
			{
				DestinationCountry: destination,
			},
			{
				DestinationCountry: "Colombia",
			}}
		expectedResult := 2
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.GetTotalTickets(destination)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})

	t.Run("success - total tickets of no existent country should be 0", func(t *testing.T) {
		// arrange - given
		destination := "Japan"
		searchedDestination := "Colombia"
		tickets := []structure.Ticket{{
			DestinationCountry: destination,
		},
			{
				DestinationCountry: destination,
			},
		}
		expectedResult := 0
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.GetTotalTickets(searchedDestination)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})

}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("success - get total tickets in the early morning", func(t *testing.T) {
		// arrange - given
		period := "3"
		flightTimeEarlyMorning, err := time.Parse("15:04", "05:20")
		fmt.Println(flightTimeEarlyMorning)
		flightTimeAfternoon, _ := time.Parse("15:04", "16:20")
		tickets := []structure.Ticket{{
			FlightTime: flightTimeEarlyMorning,
		},
			{
				FlightTime: flightTimeEarlyMorning,
			},
			{
				FlightTime: flightTimeAfternoon,
			}}
		expectedResult := 2
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.GetCountByPeriod(period)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})

	t.Run("success - get total tickets in the afternoon", func(t *testing.T) {
		// arrange - given
		period := "16"
		flightTimeEarlyMorning, err := time.Parse("15:04", "05:20")
		fmt.Println(flightTimeEarlyMorning)
		flightTimeAfternoon, _ := time.Parse("15:04", "16:20")
		tickets := []structure.Ticket{{
			FlightTime: flightTimeEarlyMorning,
		},
			{
				FlightTime: flightTimeEarlyMorning,
			},
			{
				FlightTime: flightTimeAfternoon,
			}}
		expectedResult := 1
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.GetCountByPeriod(period)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("success - get average tickets for a destination compared to total tickets", func(t *testing.T) {
		// arrange - given
		destination := "Japan"
		tickets := []structure.Ticket{{
			DestinationCountry: destination,
		},
			{
				DestinationCountry: destination,
			},
			{
				DestinationCountry: "Colombia",
			},
			{
				DestinationCountry: "Turkey",
			},
		}
		expectedResult := float64(2) / float64(len(tickets))
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.AverageDestination(destination)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})

	t.Run("success - average tickets for a no existent destination compared to total tickets should be 0", func(t *testing.T) {
		// arrange - given
		destination := "Japan"
		tickets := []structure.Ticket{
			{
				DestinationCountry: "Colombia",
			},
			{
				DestinationCountry: "Turkey",
			},
		}
		expectedResult := float64(0)
		travelServiceStruct := TravelServiceStruct{tickets}
		// act - when
		totalTickets, err := travelServiceStruct.AverageDestination(destination)
		// assert then
		require.Equal(t, expectedResult, totalTickets)
		require.Nil(t, err)
	})

}
