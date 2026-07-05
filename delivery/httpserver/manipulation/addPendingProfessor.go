package manipulation

import (
	"net/http"
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	notify "ostadbun/pkg/bale/notif"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) addPendingProfessor(c *fiber.Ctx) error {

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return richerror.Out(
			richerror.New("addpendingprofessor.delivery").WithMessage("user not found").WithKind(richerror.KindInvalid),
			c)
	}

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
		Action:             "create",
	}

	go func() {
		if err := notify.NotifyNewProfessor(data); err != nil {
			//TODO log here
		}
	}()

	errSvc := h.manipulSVC.AddPendingProfessor(data, userId)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}
	return c.Status(http.StatusOK).SendString("success")

}
