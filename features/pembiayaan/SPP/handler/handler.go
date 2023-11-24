package handler

import (
	"pembiayaan/features/pembiayaan/SPP/data"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	modelSPP := []data.SPP{
		{Id: 1, TglSPP: "20 Januari 2023", UraianSPP: "spp", NilaiSPP: 30000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 2, TglSPP: "10 Januari 2023", UraianSPP: "spp", NilaiSPP: 1000000, Status: "Ditolak", Tahapan: "sudah direview"},
		{Id: 3, TglSPP: "11 Januari 2023", UraianSPP: "spp", NilaiSPP: 50000000, Status: "Batal", Tahapan: "belum direview"},
		{Id: 4, TglSPP: "23 Januari 2023", UraianSPP: "spp", NilaiSPP: 51000000, Status: "Diterima", Tahapan: "sudah direview"},
		{Id: 5, TglSPP: "11 Januari 2023", UraianSPP: "spp", NilaiSPP: 32000000, Status: "Dibuat", Tahapan: "belum direview"},
	}

	return c.JSON(map[string]any{
		"data":    modelSPP,
		"status":  200,
		"message": "success get data",
	})
}

func Create(c *fiber.Ctx) error {
	var request data.SPP
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(map[string]any{
			"message": "input not valid",
			"Status":  fiber.ErrBadRequest,
		})
	}
	id := 5
	request.Id = id
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "success create data",
		"data":    request,
	})

}
