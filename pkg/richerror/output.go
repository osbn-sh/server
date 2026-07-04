package richerror

import (
	"errors"
	"ostadbun/pkg/enviroment"

	"github.com/gofiber/fiber/v2"
)

func Out(err error, c *fiber.Ctx) error {
	var re RichError
	if !errors.As(err, &re) {
		return err
	}

	resp := fiber.Map{
		"message": re.Error(),
	}

	if !enviroment.IsProduction() {
		resp["root_cause_message"] = re.RootCause()
		
		resp["developer_operation_chain"] = re.OperationChain()
	}

	return c.Status(re.kind.HTTPStatus()).JSON(resp)
}
