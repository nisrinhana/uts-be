package helper

import "github.com/gofiber/fiber/v2"

func JSONResponse(c *fiber.Ctx, status int, data interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": status < 400,
		"data":    data,
		"message": message,
	})
}
