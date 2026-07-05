package manipulation

import (
	"fmt"
	"net/http"
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	notify "ostadbun/pkg/bale/notif"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) EditPendingProfessor(c *fiber.Ctx) error {

	fmt.Println("is this??🫟")

	idString := c.Params("id")

	idINT, err := strconv.Atoi(idString)

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return richerror.Out(
			richerror.New("addpendingprofessor.delivery").WithMessage("user not found").WithKind(richerror.KindInvalid),
			c)
	}
	id := int64(idINT)
	var acceptData manipulationParam.PendingProfessor

	er := c.BodyParser(&acceptData)

	if er != nil {

		return richerror.Out(
			richerror.New("addpendingprofessor.delivery").WithMessage("error on parsing request").WithKind(richerror.KindInvalid).WithErr(er),
			c)
	}

	data := entity.PendingProfessor{
		Name:               acceptData.Name,
		NameEnglish:        acceptData.NameEnglish,
		DescriptionEnglish: acceptData.DescriptionEnglish,
		Description:        acceptData.Description,
		ImageUrl:           &acceptData.ImageUrl,
		EducationHistory:   acceptData.EducationHistory,
		SubmittedBy:        int64(userId),
		Action:             "update",
		TargetId:           &id,
	}

	go func() {
		if err := notify.NotifyNewProfessor(data); err != nil {
			//TODO log here
		}
	}()

	errSvc := h.manipulSVC.EditPendingProfessor(data, userId)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}
	return c.Status(http.StatusOK).SendString("success")

}
