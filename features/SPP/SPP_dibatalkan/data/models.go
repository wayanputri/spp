package data

import sppdibatalkan "project/features/SPP/SPP_dibatalkan"

type SPP struct {
	Id        int
	TglSPP    string
	UraianSPP string
	TglTolak  string
	NilaiSPP  int
	Status    string
	Tahapan   string
}

func ModelToEntity(spp SPP) sppdibatalkan.SPPEntity {
	return sppdibatalkan.SPPEntity{
		Id:        spp.Id,
		TglSPP:    spp.TglSPP,
		UraianSPP: spp.UraianSPP,
		TglTolak:  spp.TglTolak,
		NilaiSPP:  spp.NilaiSPP,
		Status:    spp.Status,
		Tahapan:   spp.Tahapan,
	}
}

func ListModelToEntity(spp []SPP) []sppdibatalkan.SPPEntity {
	var sppbatal []sppdibatalkan.SPPEntity
	for _, value := range spp {
		sppbatal = append(sppbatal, ModelToEntity(value))
	}
	return sppbatal
}
