package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vp2306/fund-forge/internal/db"
	"github.com/vp2306/fund-forge/internal/routes"
)

func main() {

	//load env vars
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found (using system env vars)")
    }

	//connect to db
	databse, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer databse.Close()
	log.Println("Succesfully connected to the database")
	

	//start router
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	log.Println("listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))


}