package main

import (
	"fmt"
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
	c, d := g.MultiDepend(22, "university")

	//c, d := g.UniversityGet(22)
	fmt.Println(c, d)
}
