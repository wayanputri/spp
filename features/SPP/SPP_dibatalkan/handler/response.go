package handler

import sppdibatalkan "project/features/SPP/SPP_dibatalkan"

type SPPResponse struct {
	Id        int    `json:"id"`
	TglSPP    string `json:"tanggal_SPP,omitempty"`
	UraianSPP string `json:"uraian_SPP,omitempty"`
	TglTolak  string `json:"tanggal_tolak,omitempty"`
	NilaiSPP  int    `json:"nilai_SPP,omitempty"`
	Status    string `json:"status_perbaikan,omitempty"`
	Tahapan   string `json:"tahapan,omitempty"`
}

func EntityToRespoonse(spp sppdibatalkan.SPPEntity) SPPResponse {
	return SPPResponse{
		Id:        spp.Id,
		TglSPP:    spp.TglSPP,
		UraianSPP: spp.UraianSPP,
		TglTolak:  spp.TglTolak,
		NilaiSPP:  spp.NilaiSPP,
		Status:    spp.Status,
		Tahapan:   spp.Tahapan,
	}
}

func ListEntityToResponse(spp []sppdibatalkan.SPPEntity) []SPPResponse {
	var sppresponse []SPPResponse
	for _, value := range spp {
		sppresponse = append(sppresponse, EntityToRespoonse(value))
	}
	return sppresponse
}
