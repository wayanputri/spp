package sppditolak

type SPPEntity struct {
	Id              int
	TglSPP          string
	UraianSPP       string
	TglTolak        string
	NilaiSPP        int
	StatusPerbaikan string
	Tahapan         string
}

type SPPDataInterface interface {
	Select() ([]SPPEntity, error)
}

type SPPServiceInterface interface {
	Get() ([]SPPEntity, error)
}
