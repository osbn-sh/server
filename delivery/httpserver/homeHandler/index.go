package homehandler

import (
	renderertempl "ostadbun/pkg/rendererTempl"
	viewindex "ostadbun/view/index"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) IndexPage(c *fiber.Ctx) error {

	return renderertempl.HTML(c, viewindex.Index("0.0.1", "8.4.0", "https://github.com/osbn-sh/app", "https://github.com/osbn-sh/server"))
}
