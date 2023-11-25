package data

import sppditerima "project/features/SPP/SPP_diterima"

type SPP struct {
	Id             int
	TglSPP         string
	UraianSPP      string
	TglDiterima    string
	NilaiUsulan    int
	NilaiDisetujui int
	Tahapan        string
}

func ModelToEntity(spp SPP) sppditerima.SPPEntity {
	return sppditerima.SPPEntity{
		Id:             spp.Id,
		TglSPP:         spp.TglSPP,
		UraianSPP:      spp.UraianSPP,
		TglDiterima:    spp.TglDiterima,
		NilaiUsulan:    spp.NilaiUsulan,
		NilaiDisetujui: spp.NilaiDisetujui,
		Tahapan:        spp.Tahapan,
	}
}

func ListModelToEntity(spp []SPP) []sppditerima.SPPEntity {
	var sppbatal []sppditerima.SPPEntity
	for _, value := range spp {
		sppbatal = append(sppbatal, ModelToEntity(value))
	}
	return sppbatal
}
