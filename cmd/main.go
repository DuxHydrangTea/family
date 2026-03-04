package main

import (
	"family/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	application := app.NewApplication()

	if err := application.RunApp(); err != nil {
		panic(err)
	}
}
