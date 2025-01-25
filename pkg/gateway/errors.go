package gateway

import "github.com/gofiber/fiber/v2"

var (
	ErrNoIp                = fiber.NewError(400, "No IP address provided")
	ErrFailedToCalculateIP = fiber.NewError(500, "Failed to determine belonging IP address")
	ErrFailedToPass        = fiber.NewError(401, "Failed to pass gateway")
)
