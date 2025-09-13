package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vp2306/fund-forge/internal/db"
	"github.com/vp2306/fund-forge/internal/handlers"
	"github.com/vp2306/fund-forge/internal/repositories"
	"github.com/vp2306/fund-forge/internal/routes"
	"github.com/vp2306/fund-forge/internal/services"
)

func main() {

	//load env vars
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env vars)")
	}

	//connect to db
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to the database")

	//wire dependencies
	repo := repositories.NewETFRepository(db)
	service := services.NewETFService(repo)
	handler := handlers.NewETFHandler(service)

	//start router
	r := chi.NewRouter()
	routes.RegisterETFRoutes(r, handler)
	log.Println("listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
