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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if data.LessonID < 1 || data.ProfessorID < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "lesson_id or professor_id required")
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	user, err := h.userService.GetByID(userID)

	if user.UniversityId == nil || *user.UniversityId < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "the user not register university ")
	}

	if user.MajorId == nil || *user.MajorId < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "the user not register major ")
	}

	data.UniversityID = *user.UniversityId
	data.MajorID = *user.MajorId

	errDOING := h.studentService.Add(userID, data)

	if errDOING != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDOING.(richerror.RichError).Unwrap().Error(),
		})
	}
	return c.SendString("success")

	//	who are you
	//	get university lesson and

}
