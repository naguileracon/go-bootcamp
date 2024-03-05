package reader

import "github.com/bootcamp-go/desafio-go-bases/internal/tickets/structure"

type Reader interface {
	Read() ([]structure.Ticket, error)
}
