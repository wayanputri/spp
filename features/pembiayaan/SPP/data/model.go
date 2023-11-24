package data

type SPP struct {
	Id        int    `json:"id"`
	TglSPP    string `json:"tanggal_SPP"`
	UraianSPP string `json:"Uraian_SPP"`
	NilaiSPP  int    `json:"nilai_SPP"`
	Status    string `json:"status"`
	Tahapan   string `json:"tahapan"`
}
