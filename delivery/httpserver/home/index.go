package homehandler

import (
	renderertempl "ostadbun/pkg/rendererTempl"
	viewindex "ostadbun/view/index"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) IndexPage(c *fiber.Ctx) error {

	clientVersion, errClient := h.GithubCheckingVersionService.GetClientVersion(c.Context())
	serverVersion, errServer := h.GithubCheckingVersionService.GetServerVersion(c.Context())

	if errClient != nil || errServer != nil {
		clientVersion, serverVersion = "1.0.0", "1.0.0"
	}

	return renderertempl.HTML(c, viewindex.Index(clientVersion, serverVersion, "https://github.com/osbn-sh/app", "https://github.com/osbn-sh/server"))
}
