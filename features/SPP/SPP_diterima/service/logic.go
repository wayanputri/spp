package service

import (
	sppditerima "project/features/SPP/SPP_diterima"
)

type SPPService struct {
	sppService sppditerima.SPPDataInterface
}

// Get implements sppditerima.SPPServiceInterface.
func (service SPPService) Get() ([]sppditerima.SPPEntity, error) {
	data, err := service.sppService.Select()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(service sppditerima.SPPDataInterface) sppditerima.SPPServiceInterface {
	return SPPService{
		sppService: service,
	}
}
