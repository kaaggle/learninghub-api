package models

type School struct {
	Approved bool   `json:"approved"`
	Username string `json:"school_username" bson:"school_username"`
	User
}
