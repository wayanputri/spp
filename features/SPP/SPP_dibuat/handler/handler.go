package handler

import (
	sppdibuat "project/features/SPP/SPP_dibuat"
	"project/helpers"

	"github.com/gofiber/fiber/v2"
)

type SPPHandler struct {
	sppHandler sppdibuat.SPPServiceInterface
}

func (h *SPPHandler) Get(c *fiber.Ctx) error {
	data, err := h.sppHandler.Get()
	if err != nil {
		return helpers.R(nil, 500, "failed get data", c)
	}
	response := ListEntityToResponse(data)
	return helpers.R(response, 200, "success get data spp dibuat", c)
}

func New(handler sppdibuat.SPPServiceInterface) *SPPHandler {
	return &SPPHandler{
		sppHandler: handler,
	}
}
