package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
)

func GetAllPengeluaran(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	pengeluaran, err := repository.GetAllPengeluaran(db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "pengeluaran not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":       false,
		"pengeluaran": pengeluaran,
	})
}

func GetPengeluaranID(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	id := c.Params("id")

	if err != nil {
		panic(err)
	}

	pengeluaran, err := repository.GetPengeluaranID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "uang not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":       false,
		"pengeluaran": pengeluaran,
	})
}

func CreatePengeluaran(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return nil
	}

	pengeluaran := &dto.Pengeluaran{}

	if err := c.BodyParser(pengeluaran); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.CreatePengeluaran(pengeluaran, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"pengeluaran": pengeluaran,
	})
}

func UpdatePengeluaran(c *fiber.Ctx) error {
	pengeluaran := &dto.Pengeluaran{}

	db, err := config.InitialDatabase()

	id := c.Params("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.BodyParser(pengeluaran); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.UpdatePengeluaran(id, pengeluaran, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeletePengeluaran(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := c.Params("id")

	if err = repository.DeletePengeluaran(id, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
