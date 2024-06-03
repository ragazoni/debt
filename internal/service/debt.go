package service

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/ragazoni/debt/internal/db"
	"github.com/ragazoni/debt/internal/models"
)

func addDebt(c fiber.Ctx) error {
	var debts models.Debt

	if err := c.Bind().Body(&debts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	debts.DueDate = time.Now()
	_, err := db.Insert("debts", debts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create debt",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Debt created successfully",
	})
}

func listDebts(c fiber.Ctx) error {
	var document []models.Debt

	err := db.FindByUser("debts", &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not list debts",
		})
	}
	return c.JSON(document)
}
