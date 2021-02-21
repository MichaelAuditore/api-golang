package ticketrepo

import (
	"context"
	"go-app/domain/ticket"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo Interface
type Repo interface {
	CreateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error)
	GetAllTickets() ([]ticket.Ticket, error)
	GetTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error)
	UpdateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error)
	DeleteTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error)
}

type ticketRepo struct {
	db *mongo.Client
}

// NewTicketRepo will instantiate Ticket Repository
func NewTicketRepo(db *mongo.Client) Repo {
	return &ticketRepo{
		db: db,
	}
}

func (t *ticketRepo) CreateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // releases resources if CreateTicket completes before timeout elapses
	collection := t.db.Database("tickets-db").Collection("tickets")
	ticket.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, *ticket)

	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t *ticketRepo) GetAllTickets() ([]ticket.Ticket, error) {
	var tickets []ticket.Ticket
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := t.db.Database("tickets-db").Collection("tickets")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var ticket ticket.Ticket
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (t *ticketRepo) GetTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error) {
	var ticket ticket.Ticket
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // releases resources if CreateTicket completes before timeout elapses
	collection := t.db.Database("tickets-db").Collection("tickets")
	err := collection.FindOne(ctx, bson.M{"_id": ticketId}).Decode(&ticket)

	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (t *ticketRepo) UpdateTicket(ticket *ticket.Ticket) (*ticket.Ticket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // releases resources if UpdateTicket completes before timeout elapses
	collection := t.db.Database("tickets-db").Collection("tickets")
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(ctx, bson.D{{"_id", ticket.ID}}, bson.D{{"$set", ticket}}, &opts).Decode(&ticket)

	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t *ticketRepo) DeleteTicketById(ticketId *primitive.ObjectID) (*ticket.Ticket, error) {
	var ticket ticket.Ticket
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel() // releases resources if CreateTicket completes before timeout elapses
	collection := t.db.Database("tickets-db").Collection("tickets")
	err := collection.FindOneAndDelete(ctx, bson.M{"_id": ticketId}).Decode(&ticket)

	if err != nil {
		return nil, err
	}
	return &ticket, nil
}
