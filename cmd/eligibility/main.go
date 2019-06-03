package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gemsorg/eligibility/pkg/database"
	"github.com/joho/godotenv"

	"github.com/gemsorg/eligibility/pkg/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatal("mysql connection error", err)
	}
	defer db.Close()

	s := server.New(db)
	log.Println("info", fmt.Sprintf("Starting service on port 3000"))
	http.Handle("/", s)
	http.ListenAndServe(":3000", nil)
}
