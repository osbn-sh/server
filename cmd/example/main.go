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

	//redisClient := redisAdaptor.New()

	//g := activityRepository.New(dbconf)

	//rds := redisActivity.New(redisClient)

	//ttt.Trigger(ctx, 69, Activityconstants.TriggerMakeAdmin)

	//fmt.Print("the user level is: ")

	acaRepo := academicRepository.New(dbconf)

	err := acaRepo.AddPendingLessonPreRequisites(41, 36)
	fmt.Println(err)
	//fmt.Println(gf.ChangeRate(25, 2))

	//fmt.Println(gf.ChangeOption(2, voteparam.Option{Name: "ظاهر آراسته", Weight: 2}))

}
