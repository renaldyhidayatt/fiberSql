package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
	"github.com/renaldyhidayatt/pencobaan/utils"
)

func CreateUsers(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return err
	}

	user := &dto.Users{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash

	if err := repository.CreateUsers(user, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	newUser := dto.NewUsers{
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})

}
