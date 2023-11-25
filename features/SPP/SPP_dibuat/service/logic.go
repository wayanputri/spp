package service

import (
	sppdibuat "project/features/SPP/SPP_dibuat"
)

type SPPService struct {
	sppService sppdibuat.SPPDataInterface
}

// Get implements sppdibuat.SPPServiceInterface.
func (service SPPService) Get() ([]sppdibuat.SPPEntity, error) {
	data, err := service.sppService.Select()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(service sppdibuat.SPPDataInterface) sppdibuat.SPPServiceInterface {
	return SPPService{
		sppService: service,
	}
}
