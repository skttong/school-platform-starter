package main

import (
	"log"

	"github.com/joho/godotenv"
	"school/internal/app"
)

func main() {
	_ = godotenv.Load()
	srv := app.NewServer()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
