package main

import (
	"fmt"
	"ostadbun/database"
	"ostadbun/repository/postgres/voteRepository"

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

	gf := voteRepository.New(dbconf)

	//err := gf.AddOption(voteparam.Option{
	//	Name:   "good",
	//	Weight: 3,
	//})

	//err := gf.AddRate(69, voteparam.Vote{
	//	OptionID: 2,
	//	Rate:     10,
	//	Target:   "professor",
	//	TargetID: 35,
	//})

	//fmt.Println(gf.ChangeRate(25, 2))

	//fmt.Println(gf.ChangeOption(2, voteparam.Option{Name: "ظاهر آراسته", Weight: 2}))

	fmt.Println(gf.CalcVotesByOption("professor", 35))

}
