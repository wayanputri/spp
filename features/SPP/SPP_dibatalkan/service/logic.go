package service

import sppdibatalkan "project/features/SPP/SPP_dibatalkan"

type SPPService struct {
	sppService sppdibatalkan.SPPDataInterface
}

// Get implements sppdibatalkan.SPPServiceInterface.
func (service SPPService) Get() ([]sppdibatalkan.SPPEntity, error) {
	data, err := service.sppService.Select()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(service sppdibatalkan.SPPDataInterface) sppdibatalkan.SPPServiceInterface {
	return SPPService{
		sppService: service,
	}
}
