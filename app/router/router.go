package router

import (
	"database/sql"
	dataSPPBatal "project/features/SPP/SPP_dibatalkan/data"
	handlerSPPBatal "project/features/SPP/SPP_dibatalkan/handler"
	serviceSPPBatal "project/features/SPP/SPP_dibatalkan/service"

	dataSPPDibuat "project/features/SPP/SPP_dibuat/data"
	handlerSPPDibuat "project/features/SPP/SPP_dibuat/handler"
	serviceSPPDibuat "project/features/SPP/SPP_dibuat/service"

	dataSPPDiterima "project/features/SPP/SPP_diterima/data"
	handlerSPPDiterima "project/features/SPP/SPP_diterima/handler"
	serviceSPPDiterima "project/features/SPP/SPP_diterima/service"

	dataSPPDitolak "project/features/SPP/SPP_ditolak/data"
	handlerSPPDitolak "project/features/SPP/SPP_ditolak/handler"
	serviceSPPDitolak "project/features/SPP/SPP_ditolak/service"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(c *fiber.App, db *sql.DB) {
	DSPPB := dataSPPBatal.New(db)
	SSPPB := serviceSPPBatal.New(DSPPB)
	HSPPB := handlerSPPBatal.New(SSPPB)

	c.Get("/api/spp/batal", HSPPB.Get)

	DSPPD := dataSPPDibuat.New(db)
	SSPPD := serviceSPPDibuat.New(DSPPD)
	HSPPD := handlerSPPDibuat.New(SSPPD)
	c.Get("/api/spp/buat", HSPPD.Get)

	DSPPT := dataSPPDiterima.New(db)
	SSPPT := serviceSPPDiterima.New(DSPPT)
	HSPPT := handlerSPPDiterima.New(SSPPT)
	c.Get("/api/spp/terima", HSPPT.Get)

	DSPPTT := dataSPPDitolak.New(db)
	SSPPTT := serviceSPPDitolak.New(DSPPTT)
	HSPPTT := handlerSPPDitolak.New(SSPPTT)
	c.Get("/api/spp/tolak", HSPPTT.Get)
}
