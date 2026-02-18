package main

import (
	"fmt"
	"ostadbun/database"
	"ostadbun/repository/postgres/manipulationRepository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file resume and load from local env")
	}
	dbConf := database.New()

	maniRepo := manipulationRepository.New(dbConf)

	fmt.Println("is ok")
	errme := maniRepo.StabilizeUniversity(1)

	if errme != nil {
		fmt.Println("💀")
	} else {
		fmt.Println("👱🏻‍♀️")
	}

	fmt.Println(errme)

}
