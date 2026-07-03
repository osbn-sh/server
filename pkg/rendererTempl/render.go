package renderertempl

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func HTML(c *fiber.Ctx, component templ.Component) error {
	var buf bytes.Buffer

	if err := component.Render(c.Context(), &buf); err != nil {
		return err
	}

	c.Type("html")
	return c.Send(buf.Bytes())
}
