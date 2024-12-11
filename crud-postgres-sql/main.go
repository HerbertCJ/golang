package main

import (
	"crud-postgres-sql/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var app app.App

	app.Initialize()
}
