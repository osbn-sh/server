package main

import (
	"context"
	"fmt"
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/repository/redis/redisGithubVersionChecking"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file resume and load from local env")
	}

	redisClient := redisAdaptor.New()

	t := redisGithubVersionChecking.New(redisClient)

	//a := t.SetClientVersion(context.Background(), "3.3.3")
	y, a := t.GetClientVersion(context.Background())

	fmt.Println(a, y)

}
