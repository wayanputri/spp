package service

import (
	sppditolak "project/features/SPP/SPP_ditolak"
)

type SPPService struct {
	sppService sppditolak.SPPDataInterface
}

// Get implements sppditolak.SPPServiceInterface.
func (service SPPService) Get() ([]sppditolak.SPPEntity, error) {
	data, err := service.sppService.Select()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(service sppditolak.SPPDataInterface) sppditolak.SPPServiceInterface {
	return SPPService{
		sppService: service,
	}
}
