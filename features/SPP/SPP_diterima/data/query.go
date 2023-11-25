package data

import (
	"database/sql"
	sppditerima "project/features/SPP/SPP_diterima"
)

type SPPData struct {
	db *sql.DB
}

// Select implements sppditerima.SPPDataInterface.
func (repo SPPData) Select() ([]sppditerima.SPPEntity, error) {
	modelSPP := []SPP{
		{Id: 1, TglSPP: "20 Januari 2023", UraianSPP: "spp", TglDiterima: "25 Januari 2023", NilaiUsulan: 6000000, NilaiDisetujui: 5000000, Tahapan: "sudah direview"},
		{Id: 2, TglSPP: "10 Januari 2023", UraianSPP: "spp", TglDiterima: "25 Januari 2023", NilaiUsulan: 6000000, NilaiDisetujui: 5000000, Tahapan: "sudah direview"},
		{Id: 3, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglDiterima: "25 Januari 2023", NilaiUsulan: 6000000, NilaiDisetujui: 5000000, Tahapan: "belum direview"},
		{Id: 4, TglSPP: "23 Januari 2023", UraianSPP: "spp", TglDiterima: "25 Januari 2023", NilaiUsulan: 6000000, NilaiDisetujui: 5000000, Tahapan: "sudah direview"},
		{Id: 5, TglSPP: "11 Januari 2023", UraianSPP: "spp", TglDiterima: "25 Januari 2023", NilaiUsulan: 6000000, NilaiDisetujui: 5000000, Tahapan: "belum direview"},
	}
	entity := ListModelToEntity(modelSPP)
	return entity, nil
}

func New(db *sql.DB) sppditerima.SPPDataInterface {
	return SPPData{
		db: db,
	}
}
