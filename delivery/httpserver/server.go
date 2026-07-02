package httpserver

import (
	"fmt"
	"ostadbun/delivery/httpserver/academic"
	"ostadbun/delivery/httpserver/manipulation"
	"ostadbun/delivery/httpserver/userhandler"
	docstempl "ostadbun/docs"

	"ostadbun/pkg/enviroment"
	renderertempl "ostadbun/pkg/rendererTempl"
	viewindex "ostadbun/view/index"

	"ostadbun/service/academicservice"
	"ostadbun/service/activityService"
	"ostadbun/service/manipulationService"

	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/swaggo/fiber-swagger"
)

type Server struct {
	userService    userservice.User
	manipulService manipulationService.Manipulation

	userHandler         userhandler.Handler
	manipulationHandler manipulation.Handler
	academicHandler     academic.Handler
}

func New(
	userService userservice.User,
	activity activityService.Activity,
	manipulService manipulationService.Manipulation,
	academicService academicservice.Service,

) Server {
	return Server{
		userService:         userService,
		manipulService:      manipulService,
		userHandler:         userhandler.New(userService, activity),
		manipulationHandler: manipulation.New(manipulService, userService),
		academicHandler:     academic.New(academicService),
	}
}

func (s Server) Serve() {

	e := fiber.New()

	e.Use(cors.New(corsConfBuilder()))

	s.userHandler.SetRoutes(e)
	s.manipulationHandler.SetRoutes(e)
	s.academicHandler.SetRoutes(e)

	e.Static("/pub", "./public")

	routes := e.Stack()

	if false {

		fmt.Println("Registered Routes:")
		for _, stack := range routes {
			for _, route := range stack {
				fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
			}
		}
	}

	e.Get("/", func(c *fiber.Ctx) error {

		return renderertempl.HTML(c, viewindex.Index("0.0.1", "8.4.0", "https://github.com/osbn-sh/app", "https://github.com/osbn-sh/server"))
	})

	e.Get("/swagger/*", fiberSwagger.WrapHandler)

	e.Get("/doc", func(c *fiber.Ctx) error {

		Url := fmt.Sprintf("http://%s/openapi.json", c.Hostname())

		return renderertempl.HTML(c, docstempl.Docs(Url))

	},
	)

	e.Get("/openapi.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	log.Fatal(e.Listen(":3000"))

}

func corsConfBuilder() cors.Config {

	if enviroment.IsProduction() {

		return cors.Config{
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowOrigins:     "https://ostadbun.tech,https://api.ostadbun.tech,https://app.ostadbun.tech",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,PUT,DELETE",
		}
	} else {
		return cors.Config{
			AllowOrigins: "*",
		}
	}
}
