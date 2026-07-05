package manipulation

import (
	"fmt"
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {

	userGroup := e.Group("/manipulation", middlewares.Auth(h.usersvc), middlewares.ManipulationPermission(h.manipulSVC))
	userGroup.Get("/get-all", middlewares.IsAdmin(h.usersvc), h.GetPending)

	userGroup.Get("/my-all", middlewares.Auth(h.usersvc), h.GetMyPending)

	userGroup.Get("/permission", h.BasicPermission)
	userGroup.Post("/lesson", h.addPendingLesson)
	userGroup.Post("/university", h.addPendingUniversity)
	userGroup.Post("/professor", h.addPendingProfessor)
	userGroup.Post("/major", h.addPendingMajor)

	userGroup.Put("/lesson/:id", h.EditPendingLesson)
	userGroup.Put("/university/:id", h.EditPendingUniversity)
	userGroup.Put("/professor/:id", h.EditPendingProfessor)
	userGroup.Put("/major/:id", h.EditPendingMajor)

	approvementRoute := func(x string) string {
		return fmt.Sprintf("/%s/approvement/:status/:targetID", x)
	}

	userGroup.Post(approvementRoute("lesson"), middlewares.IsAdmin(h.usersvc), h.ApprovementLessonPending)
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
