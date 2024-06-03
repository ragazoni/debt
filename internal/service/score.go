package service

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v3"
	"github.com/ragazoni/debt/internal/db"
	"github.com/ragazoni/debt/internal/models"
)

func getScore(c fiber.Ctx) error {
	var document []models.Debt

	cursor, err := db.Find("debts", &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not calculate score",
		})
	}

	defer cursor.Close(context.Background())

	var total float64
	var count int

	for cursor.Next(context.Background()) {
		var debt models.Debt
		cursor.Decode(&debt)
		total += debt.Amount
		count++
	}

	if count == 0 {
		return c.JSON(fiber.Map{
			"score": 1000,
		})
	}

	average := total / float64(count)
	score := math.Sqrt(average) * 10

	return c.JSON(fiber.Map{
		"score": 1000 - score,
	})
}
