package csv

import (
	"bufio"
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/structure"
	"io"
	"os"
	"strings"
	"time"
)

func NewCSVTicketsReader(filePath string) *CSVTicketsReader {
	return &CSVTicketsReader{filePath: filePath}
}

type CSVTicketsReader struct {
	Tickets  []structure.Ticket
	filePath string
}

func (ct *CSVTicketsReader) Read() (tickets []structure.Ticket, error error) {
	// opening file
	file, err := os.OpenFile(ct.filePath, os.O_RDONLY, 0644)
	if err != nil {
		error = fmt.Errorf("error opening file: %w", err)
		return
	}
	defer file.Close()

	// reading file
	rd := bufio.NewReader(file)

	for {
		data, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if data != "" {
					ticket, _ := ct.parseCsvToTicket(data)
					tickets = append(tickets, ticket)
				}
				break
			}
			error = fmt.Errorf("error reading file: %w", err)
			return
		}
		// building tickets
		ticket, _ := ct.parseCsvToTicket(data)
		tickets = append(tickets, ticket)
	}
	return
}

func (ct *CSVTicketsReader) parseCsvToTicket(data string) (ticket structure.Ticket, error error) {
	fields := strings.Split(data, ",")
	flightTime, err := time.Parse("15:04", strings.TrimSpace(fields[4]))
	if err != nil {
		error = fmt.Errorf("error parsing flight time: %w", err)
	}
	ticket = structure.Ticket{
		ID:                 strings.TrimSpace(fields[0]),
		Name:               strings.TrimSpace(fields[1]),
		Email:              strings.TrimSpace(fields[2]),
		DestinationCountry: strings.TrimSpace(fields[3]),
		FlightTime:         flightTime,
		Price:              strings.TrimSpace(fields[5]),
	}
	return
}
