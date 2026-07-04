package manipulationService

import (
	"fmt"
	Activityconstants "ostadbun/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func (m Manipulation) StabilizeLesson(c *fiber.Ctx, lessonID int) error {

	submitterID, err := m.manipulationRepo.StabilizeLesson(lessonID)

	if err == nil {
		errTrigger := m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeLesson)

		if errTrigger != nil {
			fmt.Println("⚡️Trigger Error :", errTrigger)
		}
	}

	return err

}

func (m Manipulation) StabilizeProfessor(c *fiber.Ctx, professorID int) error {

	submitterID, err := m.manipulationRepo.StabilizeProfessor(professorID)
	if err == nil {
		errTrigger := m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeProfessor)

		if errTrigger != nil {
			fmt.Println("⚡️Trigger Error :", errTrigger)
		}
	}

	return err

}

func (m Manipulation) StabilizeMajor(c *fiber.Ctx, majorID int) error {

	submitterID, err := m.manipulationRepo.StabilizeMajor(majorID)

	if err == nil {
		errTrigger := m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeMajor)

		if errTrigger != nil {
			fmt.Println("⚡️Trigger Error :", errTrigger)
		}
	}

	return err
}

func (m Manipulation) StabilizeUniversity(c *fiber.Ctx, universityID int) error {

	submitterID, err := m.manipulationRepo.StabilizeUniversity(universityID)

	if err == nil {
		errTrigger := m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeUniversity)

		if errTrigger != nil {
			fmt.Println("⚡️Trigger Error :", errTrigger)
		}
	}

	return err

}
