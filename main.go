package main

import (
	"fmt"
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/postgres/academicRepository"
	"ostadbun/repository/postgres/activityRepository"
	"ostadbun/repository/postgres/manipulationRepository"
	"ostadbun/repository/postgres/userRepository"
	"ostadbun/repository/redis/redisActivity"
	"ostadbun/repository/redis/redisOauth"
	"ostadbun/repository/redis/redisUser"
	"ostadbun/service/academicservice"
	"ostadbun/service/activityService"

	"ostadbun/service/manipulationService"

	"github.com/joho/godotenv"

	"ostadbun/delivery/httpserver"
	"ostadbun/service/oauthservice"
	"ostadbun/service/userservice"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file resume and load from local env")
	}
	dbConf := database.New()

	redisClient := redisAdaptor.New()

	//oauth
	OauthRds := redisOauth.New(redisClient)
	oauth := oauthservice.NewOAuthService(OauthRds)

	//activity
	activeRds := redisActivity.New(redisClient)
	activeRepo := activityRepository.New(dbConf)
	activeSvc := activityService.New(activeRepo, activeRds)

	//user
	userRds := redisUser.New(redisClient)
	userRepo := userRepository.New(dbConf)
	userSvc := userservice.New(*oauth, activeSvc, userRds, userRepo)

	//manipulation
	maniRepo := manipulationRepository.New(dbConf)
	maniSVC := manipulationService.New(activeSvc, *maniRepo)

	//academic
	academicRepo := academicRepository.New(dbConf)
	acaSVC := academicservice.New(*academicRepo)

	//engine
	server := httpserver.New(userSvc, activeSvc, maniSVC, acaSVC)

	fmt.Println("listening on events ...")
	server.Serve()

}
