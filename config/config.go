package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {

	err := godotenv.Load("stack.env")
	if err != nil {
		fmt.Print("Error loading stack.env file")
	}
	return os.Getenv(key)
}
