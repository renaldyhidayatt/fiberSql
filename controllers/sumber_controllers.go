package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
)

func GetSumberAll(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	sumber, err := repository.GetAllSumber(db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "sumber not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"sumber": sumber,
	})
}

func GetSumberID(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()
	id := c.Params("id")

	if err != nil {
		return nil
	}

	// ider, err := strconv.Atoi(id)

	sumber, err := repository.GetSumberID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "sumber not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"sumber": sumber,
	})
}

func CreateSumber(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return nil
	}

	sumber := &dto.Sumber{}

	if err := c.BodyParser(sumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.CreateSumber(sumber, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"sumber": sumber,
	})

}

func UpdateSumber(c *fiber.Ctx) error {
	sumber := &dto.Sumber{}

	db, err := config.InitialDatabase()
	id := c.Params("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.BodyParser(sumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// foundedSumber, err := repository.GetSumberID(int8(ider), db)

	if err := repository.UpdateSumber(id, sumber, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteSumber(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := c.Params("id")

	if err = repository.DeleteSumber(id, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
