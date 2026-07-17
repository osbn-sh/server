package student

import (
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) PassLessonGet(c *fiber.Ctx) error {

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("user id not found").WithKind(richerror.KindInvalid),
			c)
	}

	PassedLessons, errD := h.studentService.Get(userID, c.Context())

	if errD != nil {
		return richerror.Out(errD, c)
	}

	return c.JSON(PassedLessons)

}
