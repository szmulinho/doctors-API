package main

import (
	"github.com/szmulinho/doctors/cmd/server"
	"github.com/szmulinho/doctors/internal/database"
)

func main() {

	database.Connect()

	server.Run()
}
