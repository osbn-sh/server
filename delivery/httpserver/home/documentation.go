package homehandler

import (
	"fmt"
	docstempl "ostadbun/docs"
	renderertempl "ostadbun/pkg/rendererTempl"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) InitScalarDoc(c *fiber.Ctx) error {

	Url := fmt.Sprintf("http://%s/openapi.json", c.Hostname())

	return renderertempl.HTML(c, docstempl.Docs(Url))

}
