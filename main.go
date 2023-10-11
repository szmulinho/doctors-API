package main

import (
	"fmt"
	"github.com/szmulinho/common/database"
	"github.com/szmulinho/common/utils"
	"github.com/szmulinho/doctors/internal/server"
	"log"
)

func main() {
	fmt.Println("Starting the application...")
	defer fmt.Println("Closing the application")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	ctx, _, wait := utils.Gracefully()

	server.Run(ctx, db)

	wait()
}
