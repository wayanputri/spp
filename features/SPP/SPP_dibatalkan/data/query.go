package data

import (
	"database/sql"
	sppdibatalkan "project/features/SPP/SPP_dibatalkan"
)

type SPPData struct {
	db *sql.DB
}

// Select implements sppdibatalkan.SPPDataInterface.
func (repo SPPData) Select() ([]sppdibatalkan.SPPEntity, error) {
	modelSPP := []SPP{
		{Id: 1, TglSPP: "20 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 30000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 2, TglSPP: "10 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 1000000, Status: "Ditolak", Tahapan: "sudah direview"},
		{Id: 3, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 50000000, Status: "Batal", Tahapan: "belum direview"},
		{Id: 4, TglSPP: "23 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 51000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 5, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 32000000, Status: "Dibuat", Tahapan: "belum direview"},
	}
	entitySPP := ListModelToEntity(modelSPP)
	return entitySPP, nil
}

func New(db *sql.DB) sppdibatalkan.SPPDataInterface {
	return SPPData{
		db: db,
	}
}
