package manipulation

import (
	"fmt"
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	notify "ostadbun/pkg/bale/notif"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) EditPendingMajor(c *fiber.Ctx) error {

	fmt.Println("is this??🫟")

	idString := c.Params("id")

	idINT, err := strconv.Atoi(idString)

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	id := int64(idINT)
	var acceptData manipulationParam.PendingMajor

	er := c.BodyParser(&acceptData)

	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "error on parsing request",
			"details": er,
		})
	}

	data := entity.PendingMajor{
		Name:               acceptData.Name,
		NameEnglish:        acceptData.NameEnglish,
		DescriptionEnglish: acceptData.DescriptionEnglish,
		Description:        acceptData.Description,
		SubmittedBy:        int64(userId),
		Action:             "update",
		TargetId:           &id,
	}

	go func() {
		if err := notify.NotifyNewMajor(data); err != nil {
			//TODO log here
		}
	}()

	return richerror.Out(h.manipulSVC.EditPendingMajor(data, userId), c)

}
