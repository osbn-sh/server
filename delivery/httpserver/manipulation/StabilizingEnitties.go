package manipulation

import (
	"fmt"
	"ostadbun/pkg/richerror"
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
		return richerror.Out(errSvc, c)
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingProfessor(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeProfessor(c, tID)

	if errSvc != nil {
		return richerror.Out(errSvc, c)
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingUniversity(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeUniversity(c, tID)

	if errSvc != nil {
		return richerror.Out(errSvc, c)
	}

	return c.SendString("wow it done!")
}

func (h Handler) StabilizingMajor(c *fiber.Ctx) error {

	tID, err := GiveTargetID(c)

	if err != nil {
		return err
	}

	errSvc := h.manipulSVC.StabilizeMajor(c, tID)

	if errSvc != nil {
		return richerror.Out(errSvc, c)
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
