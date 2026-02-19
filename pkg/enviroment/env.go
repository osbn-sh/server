package enviroment

import (
	"os"
	"strconv"
)

func IsProduction() bool {
	env := os.Getenv("ENVIROMENT")

	envState, err := strconv.Atoi(env)

	if err != nil {
		panic("can not detect environment")
	}
	return !(envState == 0)
}
