package academic

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/academic")
	userGroup.Get("", h.Search)

	userGroup.Get("/university/:id", h.UniversityGet)
	userGroup.Get("/professor/:id", h.ProfessorGet)
	userGroup.Get("/major/:id", h.MajorGet)
	userGroup.Get("/lesson/:id", h.LessonGet)

}
