package data

import sppditolak "project/features/SPP/SPP_ditolak"

type SPP struct {
	Id              int
	TglSPP          string
	UraianSPP       string
	TglTolak        string
	NilaiSPP        int
	StatusPerbaikan string
	Tahapan         string
}

func ModelToEntity(spp SPP) sppditolak.SPPEntity {
	return sppditolak.SPPEntity{
		Id:              spp.Id,
		TglSPP:          spp.TglSPP,
		UraianSPP:       spp.UraianSPP,
		TglTolak:        spp.TglTolak,
		NilaiSPP:        spp.NilaiSPP,
		StatusPerbaikan: spp.StatusPerbaikan,
		Tahapan:         spp.Tahapan,
	}
}

func ListModelToEntity(spp []SPP) []sppditolak.SPPEntity {
	var sppbatal []sppditolak.SPPEntity
	for _, value := range spp {
		sppbatal = append(sppbatal, ModelToEntity(value))
	}
	return sppbatal
}
