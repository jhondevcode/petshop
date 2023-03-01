package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(kill bool) {
	err := godotenv.Load(".env")
	if err != nil {
		if kill {
			log.Panic(err)
		} else {
			log.Println(err)
		}
	}
}
