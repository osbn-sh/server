package main

import (
	"fmt"
	"ostadbun/database"
	"ostadbun/repository/postgres/academicRepository"
	academicservice "ostadbun/service/academicService"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	dbconf := database.New()

	g := academicRepository.New(dbconf)

	t := academicservice.New(*g)

	a, b := t.MultiDepend(21, "professor")

	fmt.Println(a, b)

}
