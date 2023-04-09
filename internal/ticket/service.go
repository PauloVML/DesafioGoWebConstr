package ticket

import "github.com/bootcamp-go/desafio-go-web/internal/domain"

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type serviceImpl struct {
	storage Repository
}

// Constructor
func NewServiceImpl(storage Repository) Service {
	return &serviceImpl{
		storage,
	}
}

func (s *serviceImpl) GetAll() ([]domain.Ticket, error) {
	return s.storage.GetAll()
}

func (s *serviceImpl) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	return s.storage.GetTicketByDestination(destination)
}

func (s *serviceImpl) GetTotalTickets(destination string) (int, error) {
	return s.storage.GetTotalTickets(destination)
}

func (s *serviceImpl) AverageDestination(destination string) (float64, error) {
	return s.storage.AverageDestination(destination)
}
