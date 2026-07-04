package manipulationService

import (
	Activityconstants "ostadbun/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func (m Manipulation) StabilizeLesson(c *fiber.Ctx, lessonID int) error {

	submitterID, err := m.manipulationRepo.StabilizeLesson(lessonID)

	if err == nil {
		m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeLesson)
	}

	return err

}

func (m Manipulation) StabilizeProfessor(c *fiber.Ctx, professorID int) error {

	submitterID, err := m.manipulationRepo.StabilizeProfessor(professorID)
	if err == nil {
		m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeProfessor)
	}

	return err

}

func (m Manipulation) StabilizeMajor(c *fiber.Ctx, majorID int) error {

	submitterID, err := m.manipulationRepo.StabilizeMajor(majorID)

	if err == nil {
		m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeMajor)
	}

	return err
}

func (m Manipulation) StabilizeUniversity(c *fiber.Ctx, universityID int) error {

	submitterID, err := m.manipulationRepo.StabilizeUniversity(universityID)

	if err == nil {
		m.activity.Trigger(c.Context(), int(submitterID), Activityconstants.TriggerStabilizeUniversity)
	}

	return err

}
