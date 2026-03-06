package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	//dbConf := database.New()

	//sf := academicRepository.New(dbConf)

	//g, r := sf.MajorGet(4)

	//err := notify.Notify("سلام کاظمم")

	//fmt.Println(err)

	g := f{}

	http.ListenAndServe(":8080", g)

}

type f struct {
}

func (e f) ServeHTTP(rs http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq)

	rs.WriteHeader(http.StatusOK)
}
