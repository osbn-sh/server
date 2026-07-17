package student

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/student")

	userGroup.Get("/pass", middlewares.Auth(h.userService), h.PassLessonGet)

	userGroup.Post("/pass", middlewares.Auth(h.userService), h.PassLessonAdd)

	userGroup.Delete("/pass/:id", middlewares.Auth(h.userService), h.PassLessonDelete)

}
