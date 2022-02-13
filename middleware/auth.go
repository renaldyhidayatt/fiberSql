package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
)

func Proctected() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})

}
