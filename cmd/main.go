package main

import "family/internal/app"

func main() {
	application := app.NewApplication()
	
	if err := application.RunApp(); err != nil {
		panic(err)
	}
}
