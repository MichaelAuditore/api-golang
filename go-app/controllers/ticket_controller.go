package controllers

import (
	"time"

	"go-app/domain/ticket"
	"go-app/services/ticketservice"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// TicketOutput represents HTTP Response Body structure
type TicketOutput struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	CreatedAt time.Time `json:"createdDate" bson:"createdDate"`
	UpdatedAt time.Time `json:"lastUpdate" bson:"lastUpdate"`
	Status bool `json:"status"`
}

// TicketInput represents postTicket body format
type TicketInput struct {
	Username string `json:"username" bson:"username"`
	Status bool `json:"status"`
}

// TicketUpdate represents postTicket body format
type TicketUpdate struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Status bool `json:"status"`
}


// TicketController interface
type TicketController interface {
	PostTicket(*gin.Context)
	GetTickets(*gin.Context)
	GetTicketById(*gin.Context)
	UpdateTicketById(*gin.Context)
	DeleteTicketById(*gin.Context)
}

type ticketController struct {
	ts ticketservice.TicketService
}

// NewTicketController instantiates Ticket Controller
func NewTicketController(ts ticketservice.TicketService) TicketController {
	return &ticketController{ts: ts}
}

func (ctl *ticketController) PostTicket(c *gin.Context) {
	// Read user input
	var ticketInput TicketInput
	if err := c.ShouldBindJSON(&ticketInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	t := ctl.inputToTicket(ticketInput)
	// Create ticket
	// If an Error Occurs while creating return the error
	if _, err := ctl.ts.CreateTicket(&t); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// If ticket is successfully created return a structured Response
	ticketOutput := ctl.mapToTicketOutput(&t)
	HTTPRes(c, http.StatusOK, "Ticket Created", ticketOutput)
}

// GetTickets function to get All tickets created and saved in MongoDB
func (ctl *ticketController) GetTickets(c *gin.Context) {
	tickets, err := ctl.ts.GetAllTickets()
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	HTTPRes(c, http.StatusOK, "All Tickets", tickets)
}

// GetTicketById function to get a specific ticket
func (ctl *ticketController) GetTicketById(c *gin.Context) {
	ticketId, err := primitive.ObjectIDFromHex(c.Param("id"))
	
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ticket, err := ctl.ts.GetTicketById(&ticketId)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	ticketOutput := ctl.mapToTicketOutput(ticket)
	HTTPRes(c, http.StatusOK, "Ticket by Id", ticketOutput)
}

// UpdateTicketById function to update a specific ticket
func (ctl *ticketController) UpdateTicketById(c *gin.Context) {
	var ticketUpdate TicketUpdate
	if err := c.ShouldBindJSON(&ticketUpdate); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	t := ctl.inputToTicketUpdate(ticketUpdate)

	// Update ticket
	// If an Error Occurs while creating return the error
	if _, err := ctl.ts.UpdateTicket(&t); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// If ticket is successfully created return a structured Response
	ticketOutput := ctl.mapToTicketOutput(&t)
	HTTPRes(c, http.StatusOK, "Ticket Updated", ticketOutput)

}

// DeleteTicketById function to delete a specific ticket
func (ctl *ticketController) DeleteTicketById(c *gin.Context) {
	ticketId, err := primitive.ObjectIDFromHex(c.Param("id"))
	
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ticket, err := ctl.ts.DeleteTicketById(&ticketId)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	ticketOutput := ctl.mapToTicketOutput(ticket)
	HTTPRes(c, http.StatusOK, "Deleted Succeed", ticketOutput)
}

// Private Methods
func (ctl *ticketController) inputToTicket(input TicketInput) ticket.Ticket {
	return ticket.Ticket{
		Username:  input.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: input.Status,
	}
}

func (ctl *ticketController) inputToTicketUpdate(input TicketUpdate) ticket.Ticket {
	return ticket.Ticket{
		ID: input.ID,
		Username:  input.Username,
		UpdatedAt: time.Now(),
		Status: input.Status,
	}
}


func (ctl *ticketController) mapToTicketOutput(t *ticket.Ticket) *TicketOutput {
	return &TicketOutput{
		ID:  t.ID,
		Username: t.Username,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Status: t.Status,
	}
}