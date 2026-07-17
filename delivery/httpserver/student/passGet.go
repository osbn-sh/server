package student

import (
	"ostadbun/param/studentparam"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

// PassLessonAdd @Summary		Register a passed lesson
// @Description	Register a lesson that the authenticated student has already passed.
// @Tags			Student
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		studentparam.StudentPassDetail	true	"Passed lesson information"
// @Success		200		{string}	string	"success"
// @Failure		400		{object}	richerror.UserReport
// @Failure		401		{object}	richerror.UserReport
// @Failure		500		{object}	richerror.UserReport
// @Router			/student/pass [post]
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

	errDOING := h.studentService.Add(userID, data, c.Context())

	if errDOING != nil {
		return richerror.Out(errDOING, c)
	}

	return c.SendString("success")

}
