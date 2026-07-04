package main

import (
	"context"
	"fmt"
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/postgres/activityRepository"
	"ostadbun/repository/redis/redisActivity"
	"ostadbun/service/activityService"

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

	a, b := ttt.LevelCalculator(ctx, 69)

	fmt.Println(a, b)

}
