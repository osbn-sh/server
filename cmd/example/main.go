package main

import (
	"fmt"
	"ostadbun/database"
	"ostadbun/repository/postgres/academicRepository"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	dbconf := database.New()
	f := academicRepository.New(dbconf)

	a, _ := f.UserCountLesson(30)

	fmt.Println(a, len(a))
}
