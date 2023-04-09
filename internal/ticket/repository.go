package ticket

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type repository struct {
	db *[]domain.Ticket
}

func NewRepository(db *[]domain.Ticket) Repository {
	fmt.Printf("Dirección en memoria desde Repository INICIADA %p\n", db)
	fmt.Printf("Dirección en memoria desde Repository TERMINADA\n")
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {

	if len(*r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return *r.db, nil
}

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(*r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range *r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetTotalTickets(destination string) (int, error) {
	var totalTickets int
	for _, ticket := range *r.db {
		if ticket.Country == destination {
			totalTickets++
		}
	}

	if totalTickets == 0 {
		return 0, ErrEmptySlice
	}

	return totalTickets, nil

}

func (r *repository) AverageDestination(destination string) (float64, error) {
	var avg float64

	total, err := r.GetTotalTickets(destination)

	if err != nil {
		if errors.Is(ErrEmptySlice, err) {
			return 0, err
		}
		return 0, err
	}

	avg = float64(total) / float64(len(*r.db))

	return avg * 100, nil

}

var (
	ErrEmptySlice = errors.New("No hay elementos disponibles para mostrar")
)
