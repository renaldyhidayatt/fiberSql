package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
)

func GetUangAll(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	uang, err := repository.GetAllUang(db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "uang not found",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"uang":  uang,
	})
}

func GetUangID(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	id := c.Params("id")

	if err != nil {
		panic(nil)
	}

	uang, err := repository.GetUangID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "uang not found",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"uang":  uang,
	})
}

func CreateUang(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return nil
	}

	uang := &dto.Uang{}

	if err := c.BodyParser(uang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.CreateUang(uang, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"uang":  uang,
	})
}

func UpdateUang(c *fiber.Ctx) error {
	uang := &dto.Uang{}

	db, err := config.InitialDatabase()
	id := c.Params("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.BodyParser(uang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// foundedSumber, err := repository.GetSumberID(int8(ider), db)

	if err := repository.UpdateUang(id, uang, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteUang(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := c.Params("id")

	if err = repository.DeleteUang(id, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
