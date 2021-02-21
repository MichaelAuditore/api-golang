package ticket

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Ticket struct
type Ticket struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	CreatedAt time.Time `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	UpdatedAt time.Time `json:"lastUpdate,omitempty" bson:"lastUpdate,omitempty"`
	Status bool `json:"status"`
}
