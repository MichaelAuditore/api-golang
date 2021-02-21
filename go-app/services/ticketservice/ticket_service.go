package ticketservice

import (
	"go-app/domain/ticket"
	"go-app/repositories/ticketrepo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TicketService interface
type TicketService interface {
	CreateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error)
	GetAllTickets() ([]ticket.Ticket, error)
	GetTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error)
	UpdateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error)
	DeleteTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error)
}

type ticketService struct {
	Repo ticketrepo.Repo
}

// NewTicketService will instantiate User Service
func NewTicketService(
	repo ticketrepo.Repo,
) TicketService {

	return &ticketService{
		Repo: repo,
	}
}

func (ts *ticketService) CreateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error) {
	return ts.Repo.CreateTicket(ticket)
}

func (ts *ticketService) GetAllTickets() ([]ticket.Ticket, error) {
	return ts.Repo.GetAllTickets()
}

func (ts *ticketService) GetTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error) {
	return ts.Repo.GetTicketById(ticketId)
}

func (ts *ticketService) UpdateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error) {
	return ts.Repo.UpdateTicket(ticket)
}

func (ts *ticketService) DeleteTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error) {
	return ts.Repo.DeleteTicketById(ticketId)
}
