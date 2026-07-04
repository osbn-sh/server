package userhandler

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {

	userGroup := e.Group("/user")

	userGroup.Get("/logout", h.logout)

	userGroup.Get("/oauth", h.redirector)

	userGroup.Get("/oauth/callback/:provider", h.acceptor)

	userGroup.Get("/switch-permission/:userid", middlewares.Auth(h.userSvc), middlewares.IsAdmin(h.userSvc), h.switchPermission)

	userGroup.Get("/level/:userid", h.Level)

	userGroup.Get("/ow", middlewares.Auth(h.userSvc), middlewares.IsAdmin(h.userSvc), test)

	//new group

	userGroup.Get("/is_login", middlewares.Auth(h.userSvc), h.IsLogin)
	userGroup.Post("/register", h.Register)

}

func test(c *fiber.Ctx) error {

	return c.SendString("you access here yeay")
}
