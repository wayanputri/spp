package handler

import sppdibuat "project/features/SPP/SPP_dibuat"

type SPPResponse struct {
	Id        int    `json:"id"`
	TglSPP    string `json:"tanggal_SPP,omitempty"`
	UraianSPP string `json:"uraian_SPP,omitempty"`
	NilaiSPP  int    `json:"nilai_SPP,omitempty"`
	Status    string `json:"status_perbaikan,omitempty"`
	Tahapan   string `json:"tahapan,omitempty"`
}

func EntityToRespoonse(spp sppdibuat.SPPEntity) SPPResponse {
	return SPPResponse{
		Id:        spp.Id,
		TglSPP:    spp.TglSPP,
		UraianSPP: spp.UraianSPP,
		NilaiSPP:  spp.NilaiSPP,
		Status:    spp.Status,
		Tahapan:   spp.Tahapan,
	}
}
func ListEntityToResponse(spp []sppdibuat.SPPEntity) []SPPResponse {
	var sppresponse []SPPResponse
	for _, value := range spp {
		sppresponse = append(sppresponse, EntityToRespoonse(value))
	}
	return sppresponse
}
