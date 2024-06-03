package service

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ragazoni/debt/internal/middleware"
)

func SetRouter(r fiber.Router) {
	debts := r.Group("/debts")

	debts.Use("/", middleware.AuthRequired)
	debts.Post("/", AddDebt)
	debts.Get("/", ListDebts)
	debts.Get("/score", GetScore)

}
