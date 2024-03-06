package travel

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/service/reader/csv"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/structure"
	"strconv"
)

func NewTravelServiceStruct(ticketsCsvReader csv.CSVTicketsReader) (serviceStruct *TravelServiceStruct, err error) {
	tickets, err := ticketsCsvReader.Read()
	if err != nil {
		err = fmt.Errorf("error getting tickets: %w", err)
		return
	}
	serviceStruct = &TravelServiceStruct{tickets: tickets}
	return

}

type TravelServiceStruct struct {
	tickets []structure.Ticket
}

func (ts *TravelServiceStruct) GetTotalTickets(destination string) (totalTickets int, error error) {
	for _, ticket := range ts.tickets {
		if ticket.DestinationCountry == destination {
			totalTickets++
		}
	}
	return
}

func (ts *TravelServiceStruct) GetCountByPeriod(time string) (total int, error error) {
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		error = fmt.Errorf("error converting time to int: %w", err)
		return
	}
	switch {
	case timeInt >= 0 && timeInt <= 6:
		total, error = ts.getPeopleCountByPeriod("earlyMorning")
		if error != nil {
			error = fmt.Errorf("error getting people count by period: %w", error)
			return
		}
	case timeInt >= 7 && timeInt <= 12:
		total, error = ts.getPeopleCountByPeriod("morning")
		if error != nil {
			error = fmt.Errorf("error getting people count by period: %w", error)
			return
		}
	case timeInt >= 13 && timeInt <= 19:
		total, error = ts.getPeopleCountByPeriod("evening")
		if error != nil {
			error = fmt.Errorf("error getting people count by period: %w", error)
			return
		}
	case timeInt >= 20 && timeInt <= 23:
		total, error = ts.getPeopleCountByPeriod("night")
		if error != nil {
			error = fmt.Errorf("error getting people count by period: %w", error)
			return
		}
	default:
		error = fmt.Errorf("error time not valid")
		return
	}
	return
}

func (ts *TravelServiceStruct) getPeopleCountByPeriod(period string) (count int, error error) {
	for _, ticket := range ts.tickets {
		flightTime := ticket.FlightTime
		switch period {
		case "earlyMorning":
			if flightTime.Hour() >= 0 && flightTime.Hour() <= 6 {
				count++
			}
		case "morning":
			if flightTime.Hour() >= 7 && flightTime.Hour() <= 12 {
				count++
			}
		case "evening":
			if flightTime.Hour() >= 13 && flightTime.Hour() <= 19 {
				count++
			}
		case "night":
			if flightTime.Hour() >= 20 && flightTime.Hour() <= 23 {
				count++
			}
		default:
			error = fmt.Errorf("invalid time period")
			return
		}
	}
	return
}

func (ts *TravelServiceStruct) AverageDestination(destination string) (average float64, error error) {
	totalDestinationTickets, err := ts.GetTotalTickets(destination)
	if err != nil {
		error = fmt.Errorf("error getting total destination tickets: %w", err)
		return
	}
	average = float64(totalDestinationTickets) / float64(len(ts.tickets))
	return
}
