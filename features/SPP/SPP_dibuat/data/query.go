package data

import (
	"database/sql"
	sppdibuat "project/features/SPP/SPP_dibuat"
)

type SPPData struct {
	db *sql.DB
}

// Select implements sppdibuat.SPPDataInterface.
func (repo SPPData) Select() ([]sppdibuat.SPPEntity, error) {
	modelSPP := []SPP{
		{Id: 1, TglSPP: "20 Januari 2023", UraianSPP: "spp", NilaiSPP: 30000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 2, TglSPP: "10 Januari 2023", UraianSPP: "spp", NilaiSPP: 1000000, Status: "Ditolak", Tahapan: "sudah direview"},
		{Id: 3, TglSPP: "11 Januari 2023", UraianSPP: "spp", NilaiSPP: 50000000, Status: "Batal", Tahapan: "belum direview"},
		{Id: 4, TglSPP: "23 Januari 2023", UraianSPP: "spp", NilaiSPP: 51000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 5, TglSPP: "11 Januari 2023", UraianSPP: "spp", NilaiSPP: 32000000, Status: "Dibuat", Tahapan: "belum direview"},
	}
	entity := ListModelToEntity(modelSPP)
	return entity, nil
}

func New(db *sql.DB) sppdibuat.SPPDataInterface {
	return SPPData{
		db: db,
	}
}
