package academic

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/academic")

	// GetUniversityPending returns all universities with 'pending' status
	// GetUniversity godoc
	// @Summary Get university
	// @Tags University
	// @Produce json
	// @Param universityid path string true "University ID"
	// @Success 200 {object} []entity.PendingUniversity
	// @Router academic/university/{universityid} [get]
	userGroup.Get("/university", h.University)

	userGroup.Get("", h.Academics)

}
