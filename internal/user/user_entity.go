package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ObjectID  primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"-"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
