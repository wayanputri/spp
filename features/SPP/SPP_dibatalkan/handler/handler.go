package handler

import (
	sppdibatalkan "project/features/SPP/SPP_dibatalkan"
	"project/helpers"

	"github.com/gofiber/fiber/v2"
)

type SPPHandler struct {
	sppHandler sppdibatalkan.SPPServiceInterface
}

func (h *SPPHandler) Get(c *fiber.Ctx) error {
	resp, err := h.sppHandler.Get()
	if err != nil {
		return helpers.R(nil, fiber.StatusInternalServerError, "error get data", c)
	}
	reponse := ListEntityToResponse(resp)
	return helpers.R(reponse, fiber.StatusOK, "success get all data spp dibatalkan", c)

}

func New(handler sppdibatalkan.SPPServiceInterface) *SPPHandler {
	return &SPPHandler{
		sppHandler: handler,
	}
}
