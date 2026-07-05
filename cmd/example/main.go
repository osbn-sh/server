package main

import (
	"context"
	"fmt"
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/postgres/activityRepository"
	"ostadbun/repository/redis/redisActivity"
	"ostadbun/service/activityService"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	dbconf := database.New()

	redisClient := redisAdaptor.New()

	g := activityRepository.New(dbconf)

	rds := redisActivity.New(redisClient)
	ttt := activityService.New(g, *rds)

	ctx := context.Background()

	//ttt.Trigger(ctx, 69, Activityconstants.TriggerMakeAdmin)

	//fmt.Print("the user level is: ")

	t := time.Now()
	fmt.Println(ttt.LevelCalculator(ctx, 69))

	fmt.Println(time.Since(t))

}
