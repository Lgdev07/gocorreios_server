package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}
}
