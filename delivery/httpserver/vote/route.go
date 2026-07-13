package votehandler

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {

	userGroup := e.Group("/vote", middlewares.Auth(h.userSvc))

	userGroup.Post("/", h.AddRate)
	userGroup.Delete("/:rate_id", h.DeleteRate)
	userGroup.Get("/:entity/:slug", h.GetRate)
	userGroup.Patch("/:rate_id", h.UpdateRate)

	adminGroup := e.Group("/option", middlewares.Auth(h.userSvc), middlewares.IsAdmin(h.userSvc))

	adminGroup.Post("/", h.AddOption)
	adminGroup.Delete("/:option_id", h.DeleteOption)
	adminGroup.Patch("/:option_id", h.UpdateOption)
}
