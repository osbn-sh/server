package main

import (
	"fmt"
	"ostadbun/database"
	"ostadbun/repository/postgres/academicRepository"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dbConf := database.New()

	sf := academicRepository.New(dbConf)

	g, r := sf.MajorGet(4)

	fmt.Println(g, r)

}
