package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"

	"github.com/globalsign/mgo/bson"
)

type Course struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Duration    int           `json:"duration" bson:"duration"`
	MainVideoURL    string `json:"main_video_url" bson:"main_video_url"`

	Videos      Videos        `json:"videos" bson:"videos"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type Courses []Course

type Video struct {
	ID          int `json:"id" bson:"id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	URL string        `json:"url" bson:"url"`

	UploadedAt time.Time `json:"uploaded_at" bson:"uploaded_at"`
}

type Videos []Video

func (c Course) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Duration, validation.Required),
		validation.Field(&c.Description, validation.Required),)
}