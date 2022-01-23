package main

import (
	"github.com/joho/godotenv"
	"github.com/myl7/bibak/internal"
	"github.com/myl7/bibak/internal/config"
	"github.com/myl7/bibak/internal/db"
	"log"
)

func main() {
	_ = godotenv.Load()

	err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.LoadDB()
	if err != nil {
		log.Fatalln(err)
	}

	internal.Loop()
}
