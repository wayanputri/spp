package data

import (
	"database/sql"
	sppditolak "project/features/SPP/SPP_ditolak"
)

type SPPData struct {
	db *sql.DB
}

// Select implements sppditolak.SPPDataInterface.
func (repo SPPData) Select() ([]sppditolak.SPPEntity, error) {
	modelSPP := []SPP{
		{Id: 1, TglSPP: "20 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 3000000, StatusPerbaikan: "blm diperbaiki", Tahapan: "sudah direview"},
		{Id: 2, TglSPP: "10 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 3000000, StatusPerbaikan: "blm diperbaiki", Tahapan: "sudah direview"},
		{Id: 3, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 3000000, StatusPerbaikan: "blm diperbaiki", Tahapan: "belum direview"},
		{Id: 4, TglSPP: "23 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 3000000, StatusPerbaikan: "blm diperbaiki", Tahapan: "sudah direview"},
		{Id: 5, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglTolak: "21 Januari 2023", NilaiSPP: 3000000, StatusPerbaikan: "blm diperbaiki", Tahapan: "belum direview"},
	}
	entity := ListModelToEntity(modelSPP)
	return entity, nil
}

func New(db *sql.DB) sppditolak.SPPDataInterface {
	return SPPData{
		db: db,
	}
}
