package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vp2306/fund-forge/config"
	"github.com/vp2306/fund-forge/internal/handlers"
	"github.com/vp2306/fund-forge/internal/routes"
)

func main() {

	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	// new handler 
	handler := handlers.NewHandlers()
	
	// set up http server
	mux := http.NewServeMux()

	// setup routes
	routes.Routes(mux, handler)

	// server instance
	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr: serverAddr,
		Handler: mux,
	}

	fmt.Printf("Server is up and running on port %s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil{
		log.Fatalf("Server failed to start %v", err)
	}

	
}