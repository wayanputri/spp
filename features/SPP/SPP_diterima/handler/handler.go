package handler

import (
	sppditerima "project/features/SPP/SPP_diterima"
	"project/helpers"

	"github.com/gofiber/fiber/v2"
)

type SPPHandler struct {
	sppHandler sppditerima.SPPServiceInterface
}

func (h *SPPHandler) Get(c *fiber.Ctx) error {
	data, err := h.sppHandler.Get()
	if err != nil {
		return helpers.R(nil, 500, "failed get spp diterima", c)
	}
	response := ListEntityToResponse(data)
	return helpers.R(response, 200, "get success all spp diterima", c)
}

func New(handler sppditerima.SPPServiceInterface) *SPPHandler {
	return &SPPHandler{
		sppHandler: handler,
	}
}
