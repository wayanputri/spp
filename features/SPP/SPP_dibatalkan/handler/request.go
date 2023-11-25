package handler

type SPPRequest struct {
	Id        int    `json:"id"`
	TglSPP    string `json:"tanggal_SPP"`
	UraianSPP string `json:"uraian_SPP"`
	TglTolak  string `json:"tanggal_tolak"`
	NilaiSPP  int    `json:"nilai_SPP"`
	Status    string `json:"status_perbaikan"`
	Tahapan   string `json:"tahapan"`
}
