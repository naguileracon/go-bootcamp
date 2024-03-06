package structure

import "time"

type Ticket struct {
	ID                 string
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         time.Time
	Price              string
}
