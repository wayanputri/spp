package data

import sppdibuat "project/features/SPP/SPP_dibuat"

type SPP struct {
	Id        int
	TglSPP    string
	UraianSPP string
	NilaiSPP  int
	Status    string
	Tahapan   string
}

func ModelToEntity(spp SPP) sppdibuat.SPPEntity {
	return sppdibuat.SPPEntity{
		Id:        spp.Id,
		TglSPP:    spp.TglSPP,
		UraianSPP: spp.UraianSPP,
		NilaiSPP:  spp.NilaiSPP,
		Status:    spp.Status,
		Tahapan:   spp.Tahapan,
	}
}

func ListModelToEntity(spp []SPP) []sppdibuat.SPPEntity {
	var sppbatal []sppdibuat.SPPEntity
	for _, value := range spp {
		sppbatal = append(sppbatal, ModelToEntity(value))
	}
	return sppbatal
}
