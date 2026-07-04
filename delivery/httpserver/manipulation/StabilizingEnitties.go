package manipulation

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) StabilizingLesson(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeLesson(c, tID)

	if errSvc != nil {
		return errSvc
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingProfessor(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeProfessor(tID)

	if errSvc != nil {
		return errSvc
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingUniversity(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeUniversity(tID)

	if errSvc != nil {
		return errSvc
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingMajor(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeMajor(tID)

	if errSvc != nil {
		return errSvc
	}

	return c.SendString("wow it done!")
}

func GiveTargetID(c *fiber.Ctx) (int, error) {

	tid := c.Params("targetID")
	targetID, errTid := strconv.Atoi(tid)

	if errTid != nil {
		return 0, fmt.Errorf("id is invalid, should be a number")
	}

	return targetID, nil
}
