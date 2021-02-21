package app

import (
	"context"
	"go-app/configs"
	"go-app/repositories/ticketrepo"
	"go-app/services/ticketservice"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-app/controllers"
)

var (
	r = gin.Default()
)

// Run is the App Entry Point
func Run() {

	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI) // use env variables
	// Connect to MongoDB
	mongoDB, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	/*
		====== Setup repositories =======
	*/
	ticketRepo := ticketrepo.NewTicketRepo(mongoDB)
	/*
		====== Setup services ===========
	*/
	ticketService := ticketservice.NewTicketService(ticketRepo)
	/*
		====== Setup controllers ========
	*/
	ticketCtl := controllers.NewTicketController(ticketService)

	/*
		======== Routes ============
	*/

	// API Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my Golang API",
		})
	})

	/*
		===== Ticket Routes =====
	*/
	r.POST("/tickets", ticketCtl.PostTicket)
	r.GET("/tickets", ticketCtl.GetTickets)
	r.GET("/tickets/:id", ticketCtl.GetTicketById)
	r.PUT("/tickets", ticketCtl.UpdateTicketById)
	r.DELETE("/tickets/:id", ticketCtl.DeleteTicketById)
	r.Run()
}
