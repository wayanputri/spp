package sppditerima

type SPPEntity struct {
	Id             int
	TglSPP         string
	UraianSPP      string
	TglDiterima    string
	NilaiUsulan    int
	NilaiDisetujui int
	Tahapan        string
}

type SPPDataInterface interface {
	Select() ([]SPPEntity, error)
}
type SPPServiceInterface interface {
	Get() ([]SPPEntity, error)
}
