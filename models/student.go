package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Student struct {
	ID bson.ObjectId `json:"_id" bson:"_id,omitempty"`

	PersonalDetails PersonalDetails `json:"personal_details" bson:"personal_details"`
	ContactDetails  ContactDetails  `json:"contact_details" bson:"contact_details"`

	Admitted  bool      `json:"admitted" bson:"admitted"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	SchoolID bson.ObjectId `json:"school_id" bson:"school_id"`
}

type PersonalDetails struct {
	Firstname   string    `json:"firstname" bson:"firstname"`
	MiddleName  string    `json:"middlename" bson:"middlename"`
	Lastname    string    `json:"lastname" bson:"lastname"`
	StudentID   string    `json:"student_id" bson:"student_id"`
	DOB         time.Time `json:"dob" bson:"dob"`
	Nationality string    `json:"nationality" bson:"nationality"`
	Gender      string    `json:"gender" bson:"gender"`
}

type ContactDetails struct {
	AddressLine1 string `json:"address_line_1" bson:"address_line_1"`
	AddressLine2 string `json:"address_line_2" bson:"address_line_2"`
	City         string `json:"city" bson:"city"`
	State        string `json:"state" bson:"state"`
	Country      string `json:"country" bson:"country"`

	PinCode     string `json:"pin_code" bson:"pin_code"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`

	Email string `json:"email" bson:"email"`
}

type Students []Student

func (st Student) Validate() error {
	return validation.ValidateStruct(&st,
		validation.Field(&st.PersonalDetails),
		validation.Field(&st.ContactDetails),
		validation.Field(&st.SchoolID, validation.Required))
}

func (pd PersonalDetails) Validate() error {
	return validation.ValidateStruct(&pd,
		validation.Field(&pd.Gender,
			validation.Required.Error("Gender is required")),

		validation.Field(&pd.Firstname, validation.Required),
		validation.Field(&pd.Lastname, validation.Required),
	)
}

func (cd ContactDetails) Validate() error {
	return validation.ValidateStruct(&cd,
		validation.Field(&cd.AddressLine1, validation.Required),
		validation.Field(&cd.AddressLine2, validation.Required),
		validation.Field(&cd.City, validation.Required),
		validation.Field(&cd.Country, validation.Required),
		validation.Field(&cd.Email, validation.Required),
	)
}
