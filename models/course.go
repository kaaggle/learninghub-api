package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Course struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Videos      Videos        `json:"videos" bson:"videos"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type Courses []Course

type Video struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`

	UploadedAt time.Time `json:"uploaded_at" bson:"uploaded_at"`
}

type Videos []Video
