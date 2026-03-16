package main

import (
	"ostadbun/database"
	"ostadbun/repository/postgres/academicRepository"
	"ostadbun/service/academicservice"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	dbconf := database.New()
	f := academicRepository.New(dbconf)
	g := academicservice.New(*f)

	g.MultiDepond(6)

}
