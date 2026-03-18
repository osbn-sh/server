package main

import (
	"ostadbun/database"

	"ostadbun/repository/postgres/studentRepository"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	dbconf := database.New()

	g := studentRepository.New(dbconf)
	g.RemovePass(6, 21)

}
