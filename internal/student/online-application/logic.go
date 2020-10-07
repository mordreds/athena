package online_application

import (
	"errors"
)

var (
	ValidationErr = errors.New("Validation Error")
)

type onlineApplicationService struct {
	onlineRegRepo OnlineApplicationRepository
}

func (ors *onlineApplicationService) SearchById(id int) (*OnlineApplication, error) {
	//panic("implement me")
	return ors.onlineRegRepo.SearchById(id)
}

func (ors *onlineApplicationService) SearchByUuid(uuid string) (*OnlineApplication, error) {
	//panic("implement me")
	return ors.onlineRegRepo.SearchByUuid(uuid)
}

func (ors *onlineApplicationService) GetAllRecords() ([]*OnlineApplication, error) {
	//panic("implement me")
	return ors.onlineRegRepo.GetAllRecord()
}

func (ors *onlineApplicationService) SaveRecord(postData *OnlineApplication) error {
	panic("implement me")
}

func (ors *onlineApplicationService) UpdateById(id int, newData *OnlineApplication) (*OnlineApplication, error) {
	panic("implement me")
}

func (ors *onlineApplicationService) UpdateByUuid(uuuid string, newData *OnlineApplication) (*OnlineApplication, error) {
	panic("implement me")
}

func (ors *onlineApplicationService) DeleteById(id int) (*OnlineApplication, error) {
	panic("implement me")
}

func (ors *onlineApplicationService) DeleteByUuid(uuid string) (*OnlineApplication, error) {
	panic("implement me")
}

// Instantiating New Online-Registration-Service
func NewOnlineRegService(onlineRegRepo OnlineApplicationRepository) OnlineApplicationService {
	return &onlineApplicationService{
		onlineRegRepo,
	}
}