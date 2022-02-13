package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
)

func GetAllPemasukan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	pemasukan, err := repository.GetAllPemasukan(db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "pemasukan not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":     false,
		"pemasukan": pemasukan,
	})
}

func GetPemasukanID(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	id := c.Params("id")

	if err != nil {
		panic(err)
	}

	pemasukan, err := repository.GetPemasukanID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "pemasukan not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":     false,
		"pemasukan": pemasukan,
	})
}

func CreatePemasukan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return nil
	}

	pemasukan := &dto.Pemasukan{}

	if err := c.BodyParser(pemasukan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.CreatePemasukan(pemasukan, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"pemasukan": pemasukan,
	})
}

func UpdatePemasukan(c *fiber.Ctx) error {
	pemasukan := &dto.Pemasukan{}

	db, err := config.InitialDatabase()

	id := c.Params("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.BodyParser(pemasukan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.UpdatePemasukan(id, pemasukan, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeletePemasukan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := c.Params("id")

	if err = repository.DeletePemasukan(id, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
