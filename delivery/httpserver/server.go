package httpserver

import (
	"fmt"
	"ostadbun/delivery/httpserver/academic"
	homehandler "ostadbun/delivery/httpserver/homeHandler"
	"ostadbun/delivery/httpserver/manipulation"
	"ostadbun/delivery/httpserver/userhandler"
	"ostadbun/pkg/enviroment"
	"ostadbun/service/academicservice"
	"ostadbun/service/activityService"
	"ostadbun/service/manipulationService"

	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	userService    userservice.User
	manipulService manipulationService.Manipulation

	userHandler         userhandler.Handler
	manipulationHandler manipulation.Handler
	academicHandler     academic.Handler
	homeHandler         homehandler.Handler
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
		homeHandler:         homehandler.New(),
	}
}

func (s Server) Serve() {

	//starting
	e := fiber.New()

	//configurations
	e.Use(cors.New(corsConfBuilder()))
	e.Static("/pub", "./public")

	//set handlers
	s.userHandler.SetRoutes(e)
	s.manipulationHandler.SetRoutes(e)
	s.academicHandler.SetRoutes(e)
	s.homeHandler.SetRoutes(e)

	ShowRoutes(e)
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

func ShowRoutes(e *fiber.App) {
	routes := e.Stack()

	if !enviroment.IsProduction() {

		fmt.Println("Registered Routes:")
		for _, stack := range routes {
			for _, route := range stack {
				fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
			}
		}
	}
}
