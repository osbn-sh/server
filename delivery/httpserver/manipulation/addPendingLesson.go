package manipulation

import (
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) addPendingLesson(c *fiber.Ctx) error {

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	var acceptData manipulationParam.PendingLesson

	er := c.BodyParser(&acceptData)

	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "error on parsing request",
			"details": er,
		})
	}

	data := entity.PendingLesson{
		Name:               acceptData.Name,
		NameEnglish:        acceptData.NameEnglish,
		DescriptionEnglish: acceptData.DescriptionEnglish,
		Description:        acceptData.Description,
		Term:               acceptData.Term,
		Difficulty:         acceptData.Difficulty,
		SubmittedBy:        int64(userId),
		Action:             "create",
	}

	go func() {
		//if err := notify.NotifyNewLesson(data); err != nil {
		//	//TODO log here
		//}
	}()

	rs := h.manipulSVC.AddPendingLesson(data, userId)

	if rs != nil {
		return richerror.Out(rs, c)
	}
	return rs
}
