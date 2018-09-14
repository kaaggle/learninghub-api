package models

import (
	"github.com/go-ozzo/ozzo-validation"
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

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Name, validation.Required))
}
