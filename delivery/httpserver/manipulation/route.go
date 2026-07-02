package manipulation

import (
	"fmt"
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	e.Get("/pending", h.GetPending)

	userGroup := e.Group("/manipulation", middlewares.Auth(h.usersvc), middlewares.ManipulationPermission(h.manipulSVC))

	// GetUniversityPending returns all universities with 'pending' status
	// GetUniversity godoc
	// @Summary Checking permission
	// @Tags permission
	// @Produce json
	// @Success 200 {object} any
	// @Router manipulation/permission [get]
	userGroup.Get("/permission", h.BasicPermission)

	userGroup.Post("/lesson", h.addPendingLesson)
	userGroup.Post("/university", h.addPendingUniversity)
	userGroup.Post("/professor", h.addPendingProfessor)
	userGroup.Post("/major", h.addPendingMajor)

	approvementRoute := func(x string) string {
		return fmt.Sprintf("/%s/approvement/:status/:targetID", x)
	}

	userGroup.Post(approvementRoute("lesson"), h.ApprovementLessonPending)
	userGroup.Post(approvementRoute("university"), h.ApprovementUnivPending)
	userGroup.Post(approvementRoute("professor"), h.ApprovementProfPending)
	userGroup.Post(approvementRoute("major"), h.ApprovementMajorPending)

	stabilizeRoute := func(x string) string {
		return fmt.Sprintf("/%s/stabilize/:targetID", x)
	}

	userGroup.Post(stabilizeRoute("lesson"), h.StabilizingLesson)
	userGroup.Post(stabilizeRoute("university"), h.StabilizingUniversity)
	userGroup.Post(stabilizeRoute("professor"), h.StabilizingProfessor)
	userGroup.Post(stabilizeRoute("major"), h.StabilizingMajor)

}
