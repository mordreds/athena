package online_application

type OnlineApplicationRepository interface {
	// Search Record By id/uuid
	SearchById(id int) (*OnlineApplication, error)
	SearchByUuid(uuid string) (*OnlineApplication, error)
	// Get All Records
	GetAllRecord() ([]*OnlineApplication, error)
	// Save Record
	SaveRecord(postData *OnlineApplication) error
	// Edit/Update Record
	UpdateById(id int, newData *OnlineApplication) (*OnlineApplication, error)
	UpdateByUuid(uuuid string, newData *OnlineApplication) (*OnlineApplication, error)
	// Delete Record
	DeleteById(id int) (*OnlineApplication, error)
	DeleteByUuid(uuid string) (*OnlineApplication, error)
}