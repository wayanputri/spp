package sppdibuat

type SPPEntity struct {
	Id        int
	TglSPP    string
	UraianSPP string
	NilaiSPP  int
	Status    string
	Tahapan   string
}

type SPPDataInterface interface {
	Select() ([]SPPEntity, error)
}
type SPPServiceInterface interface {
	Get() ([]SPPEntity, error)
}
