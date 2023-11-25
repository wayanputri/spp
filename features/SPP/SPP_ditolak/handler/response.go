package handler

import sppditolak "project/features/SPP/SPP_ditolak"

type SPPResponse struct {
	Id              int    `json:"id,omitempty"`
	TglSPP          string `json:"tgl_spp,omitempty"`
	UraianSPP       string `json:"uraian_spp,omitempty"`
	TglTolak        string `json:"tgl_tolak,omitempty"`
	NilaiSPP        int    `json:"nilai_spp,omitempty"`
	StatusPerbaikan string `json:"status_perbaikan,omitempty"`
	Tahapan         string `json:"tahapan,omitempty"`
}

func EntityToRespoonse(spp sppditolak.SPPEntity) SPPResponse {
	return SPPResponse{
		Id:              spp.Id,
		TglSPP:          spp.TglSPP,
		UraianSPP:       spp.UraianSPP,
		TglTolak:        spp.TglTolak,
		NilaiSPP:        spp.NilaiSPP,
		StatusPerbaikan: spp.StatusPerbaikan,
		Tahapan:         spp.Tahapan,
	}
}
func ListEntityToResponse(spp []sppditolak.SPPEntity) []SPPResponse {
	var sppresponse []SPPResponse
	for _, value := range spp {
		sppresponse = append(sppresponse, EntityToRespoonse(value))
	}
	return sppresponse
}
