package serverAddress

import "os"

func Addr() (string, bool) {

	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		return "localhost:3000", false
	}

	return serverAddress, true
}
