package main

import (
	"github.com/gorilla/mux"
	"go-home-content/controller"
	"go-home-content/database"
	"go-home-content/repository"
	"go-home-content/service"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	database.ConnectDatabase()

	// Initialize layers
	repo := &repository.SectionRepository{DB: database.DB}
	service := &services.SectionService{Repo: repo}
	restController := &controller.SectionController{Service: service}

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/api/section", restController.CreateSection).Methods("POST")
	router.HandleFunc("/api/sections", restController.GetSections).Methods("GET")
	router.HandleFunc("/api/section/{id}", restController.GetSectionById).Methods("GET")

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
