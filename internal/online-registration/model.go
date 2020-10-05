package online_registration

// Required Model Schema
type OnlineRegistration struct {
	// Standard Required Fields
	ID int64 `json:"id"`
	Status int8 `json:"status"`
	RecordImported bool `json:"record_imported"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

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


// Required Implementation Details (Actions)
type OnlineRegistrationService interface {
	// Search Record By id/uuid
	Search(uuid string) (*OnlineRegistration, error)
	// Get All Records
	Index() ([]*OnlineRegistration, error)
	// Save Record
	Store(postData *OnlineRegistration) error
	// Edit/Update Record
	Update(uuuid string) (*OnlineRegistration, error)
	// Delete Record
	Delete(uuid string) (*OnlineRegistration, error)
}
