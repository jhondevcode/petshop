package config

import (
	"log"
	"os"
	"strconv"
)

func ServerPort() int {
	defaultPort := 8080
	envPort := os.Getenv("SERVER_PORT")
	if envPort != "" {
		port, err := strconv.Atoi(envPort)
		if err != nil {
			log.Fatalln("Setting default port")
			return defaultPort
		} else {
			return port
		}
	} else {
		return defaultPort
	}
}
