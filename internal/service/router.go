package service

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ragazoni/debt/internal/middleware"
)

func SetRouter(r fiber.Router) {
	debts := r.Group("/debts")

	debts.Use("/", middleware.AuthRequired)
	debts.Post("/", addDebt)
	debts.Get("/", listDebts)
	debts.Get("/score", getScore)

}
