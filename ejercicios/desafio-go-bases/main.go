package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/service/reader/csv"
)
import "github.com/bootcamp-go/desafio-go-bases/internal/tickets/service/travel"

func main() {
	CsvTicketsReader := csv.NewCSVTicketsReader("tickets.csv")
	NewTravelServiceStruct, _ := travel.NewTravelServiceStruct(*CsvTicketsReader)
	// total number of tickets
	tickets, err := NewTravelServiceStruct.GetTotalTickets("Japan")
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("Total tickets ---> ", tickets)
	// total number of tickets by period
	period := "6"
	ticketsByPeriod, err := NewTravelServiceStruct.GetCountByPeriod(period)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("Total tickets by period %s ----> %d \n", period, ticketsByPeriod)
	// Average destination tickets
	average, err := NewTravelServiceStruct.AverageDestination("Japan")
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("Average tickets for destination Japan  ----> %f%% \n", average*100)
}
