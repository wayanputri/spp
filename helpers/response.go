package helpers

import "github.com/gofiber/fiber/v2"

type Response struct {
	// value message success and error message
	Message string `json:"message"`
	// response to code http 200, 404, 500, etc
	Code int `json:"code"`
	// standart status message success and failed
	Success bool `json:"success"`
	// send data format json, if request success.
	// or request data error send message failed data
	Data any `json:"data,omitempty"`
}

func R(data any, status int, message string, c *fiber.Ctx) error {

	response := Response{
		Message: message,
		Code:    status,
		Success: true,
		Data:    data,
	}
	if status >= 400 {
		response.Success = false
	}
	return c.Status(status).JSON(response)
}
