package student

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/student")

	userGroup.Post("/pass", middlewares.Auth(h.userService), h.PassLessonAdd)

	userGroup.Delete("/pass/:id", middlewares.Auth(h.userService), h.PassLessonDelete)

}
