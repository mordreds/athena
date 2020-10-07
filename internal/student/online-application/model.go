package online_application

import "time"

// Required Model Schema
type OnlineApplication struct {
	// Standard Required Fields
	ID int64 `json:"id"`
	Status int8 `json:"status"`
	Imported bool `json:"imported"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Person Details
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

	// Contact Details
	PrimaryPhone string `json:"primary_phone"`
	SecondaryPhone string `json:"secondary_phone"`

	// Work Details
	CompanyName string `json:"company_name"`
	Department string `json:"department"`
	Position string `json:"position"`
}