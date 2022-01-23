package main

import (
	"github.com/joho/godotenv"
	"github.com/myl7/bibak/internal/config"
	"log"
)

func main() {
	_ = godotenv.Load()

	err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
