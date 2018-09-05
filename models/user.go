package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Name      string        `json:"name"`
	UserID    string        `json:"user_id" bson:"user_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	Role string `json:"role" bson:"role"`
}
