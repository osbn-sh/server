package student

import (
	"ostadbun/param/studentparam"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) PassLessonAdd(c *fiber.Ctx) error {

	var data studentparam.StudentPassDetail

	err := c.BodyParser(&data)

	if err != nil {
		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("اطلاعات ورودی اشتباه است").WithKind(richerror.KindInvalid),
			c)
	}

	if data.LessonID < 1 || data.ProfessorID < 1 {

		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("lesson_id or professor_id required").WithKind(richerror.KindInvalid),
			c)
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("user id not found").WithKind(richerror.KindInvalid),
			c)
	}

	user, err := h.userService.GetByID(userID)

	if user.UniversityId == nil || *user.UniversityId < 1 {
		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("the user not register university").WithKind(richerror.KindInvalid),
			c)

	}

	if user.MajorId == nil || *user.MajorId < 1 {
		return richerror.Out(
			richerror.New("PassAdd,delivery").WithMessage("the user not register major"),
			c)

	}

	data.UniversityID = *user.UniversityId
	data.MajorID = *user.MajorId

	errDOING := h.studentService.Add(userID, data)

	if errDOING != nil {
		return richerror.Out(errDOING, c)
	}

	return c.SendString("success")

}
