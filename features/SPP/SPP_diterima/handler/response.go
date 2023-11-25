package handler

import sppditerima "project/features/SPP/SPP_diterima"

type SPPResponse struct {
	Id             int    `json:"id"`
	TglSPP         string `json:"tgl_spp,omitempty"`
	UraianSPP      string `json:"uraian_spp,omitempty"`
	TglDiterima    string `json:"tgl_diterima,omitempty"`
	NilaiUsulan    int    `json:"nilai_usulan,omitempty"`
	NilaiDisetujui int    `json:"nilai disetujui,omitempty"`
	Tahapan        string `json:"tahapan,omitempty"`
}

func EntityToRespoonse(spp sppditerima.SPPEntity) SPPResponse {
	return SPPResponse{
		Id:             spp.Id,
		TglSPP:         spp.TglSPP,
		UraianSPP:      spp.UraianSPP,
		TglDiterima:    spp.TglDiterima,
		NilaiUsulan:    spp.NilaiUsulan,
		NilaiDisetujui: spp.NilaiDisetujui,
		Tahapan:        spp.Tahapan,
	}
}
func ListEntityToResponse(spp []sppditerima.SPPEntity) []SPPResponse {
	var sppresponse []SPPResponse
	for _, value := range spp {
		sppresponse = append(sppresponse, EntityToRespoonse(value))
	}
	return sppresponse
}
