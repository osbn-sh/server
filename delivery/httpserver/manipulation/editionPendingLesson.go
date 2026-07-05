package manipulation

import (
	"fmt"
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) EditPendingLesson(c *fiber.Ctx) error {

	fmt.Println("is this??🫟")

	idString := c.Params("id")

	idINT, err := strconv.Atoi(idString)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id not a number",
		})
	}
	id := int64(idINT)

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
		Action:             "update",
		TargetId:           &id,
	}

	go func() {
		//if err := notify.NotifyNewLesson(data); err != nil {
		//	//TODO log here
		//}
	}()

	rs := h.manipulSVC.EditPendingLesson(data, userId)

	if rs != nil {
		return richerror.Out(rs, c)
	}
	return rs
}
