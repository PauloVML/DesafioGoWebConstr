package ticket

import (
	"context"
	"testing"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

var cxt = context.Background()

var tickets = []domain.Ticket{
	{
		Id:      "1",
		Name:    "Tait Mc Caughan",
		Email:   "tmc0@scribd.com",
		Country: "Finland",
		Time:    "17:11",
		Price:   785.00,
	},
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

var ticketsByDestination = []domain.Ticket{
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

type stubRepo struct {
	db *DbMock
}

type DbMock struct {
	db  []domain.Ticket
	spy bool
	err error
}

func NewRepositoryTest(dbM *DbMock) Repository {
	return &stubRepo{dbM}
}

func (r *stubRepo) GetAll() ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}
	return tickets, nil
}

func (r *stubRepo) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}

func (r *stubRepo) GetTotalTickets(destination string) (int, error) {
	r.db.spy = true
	var total int

	for _, data := range r.db.db {
		if data.Country == destination {
			total++
		}
	}

	return total, nil

}

func (r *stubRepo) AverageDestination(destination string) (float64, error) {
	r.db.spy = true
	var avg float64
	var total float64

	for _, ticket := range r.db.db {
		if ticket.Country == destination {
			total++
		}
	}

	avg = float64(len(r.db.db)) / total

	return avg * 100, nil

}

func TestGetTicketByDestination(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := NewRepositoryTest(dbMock)
	service := NewServiceImpl(repo)

	tkts, err := service.GetTotalTickets("China")

	assert.Nil(t, err)
	assert.True(t, dbMock.spy)
	assert.Equal(t, len(ticketsByDestination), tkts)
}

func TestGetTotalTickets(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := NewRepositoryTest(dbMock)
	service := NewServiceImpl(repo)

	avr, err := service.AverageDestination("China")

	assert.Nil(t, err)
	assert.NotNil(t, avr)
	assert.True(t, dbMock.spy)
}
