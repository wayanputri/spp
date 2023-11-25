package handler

import (
	sppditolak "project/features/SPP/SPP_ditolak"
	"project/helpers"

	"github.com/gofiber/fiber/v2"
)

type SPPHandler struct {
	sppHandler sppditolak.SPPServiceInterface
}

func (h *SPPHandler) Get(c *fiber.Ctx) error {
	data, err := h.sppHandler.Get()
	if err != nil {
		return helpers.R(nil, 500, "failed get all SPP ditolak", c)
	}
	response := ListEntityToResponse(data)
	return helpers.R(response, 200, "get all success SPP ditolak", c)
}

func New(handler sppditolak.SPPServiceInterface) *SPPHandler {
	return &SPPHandler{
		sppHandler: handler,
	}
}
