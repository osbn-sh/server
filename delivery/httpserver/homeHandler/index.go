package homehandler

import (
	renderertempl "ostadbun/pkg/rendererTempl"
	viewindex "ostadbun/view/index"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) IndexPage(c *fiber.Ctx) error {

	clientVersion, errClient := h.GithubCheckingVersionService.GetClientVersion(c.Context())
	serverVersion, errServer := h.GithubCheckingVersionService.GetServerVersion(c.Context())

	if errClient != nil {
		return errClient
	}

	if errServer != nil {
		return errServer
	}
	return renderertempl.HTML(c, viewindex.Index(clientVersion, serverVersion, "https://github.com/osbn-sh/app", "https://github.com/osbn-sh/server"))
}
