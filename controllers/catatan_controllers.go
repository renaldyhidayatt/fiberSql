package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/dto"
	"github.com/renaldyhidayatt/pencobaan/repository"
)

func GetAllCatatan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	sumber, err := repository.GetAllCatatan(db)

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

func GetCatatanID(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()
	id := c.Params("id")

	if err != nil {
		return nil
	}

	// ider, err := strconv.Atoi(id)

	catatan, err := repository.GetCatatanID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "catatan not found",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"catatan": catatan,
	})
}

func CreateCatatan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return nil
	}

	catatan := &dto.Catatan{}

	if err := c.BodyParser(catatan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := repository.CreateCatatan(catatan, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"catatan": catatan,
	})

}

func UpdateCatatan(c *fiber.Ctx) error {
	catatan := &dto.Catatan{}

	db, err := config.InitialDatabase()
	id := c.Params("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.BodyParser(catatan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedCatatan, err := repository.GetCatatanID(id, db)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "catatan with whit ID NOT Found",
		})
	}

	validate := validator.New()

	if err := validate.Struct(catatan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Error Validate",
		})
	}

	if err := repository.UpdateCatatan(foundedCatatan.ID, catatan, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteCatatan(c *fiber.Ctx) error {
	db, err := config.InitialDatabase()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := c.Params("id")

	if err = repository.DeleteCatatan(id, db); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
