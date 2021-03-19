package main

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"softwareDB/controller"
	_ "softwareDB/docs"
)

// @title Software-Database API
// @version 1.0
// @description This is the API for the TEKO Software Database Project
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/software", controller.CreateSoftware).Methods("POST")
	// Read
	router.HandleFunc("/software/{name}", controller.GetSoftwareByName).Methods("GET")
	// Read-all
	router.HandleFunc("/software", controller.ListSoftware).Methods("GET")
	// Update
	router.HandleFunc("/software/{id}", controller.UpdateSoftware).Methods("PUT")
	// Delete
	router.HandleFunc("/software/{id}", controller.DeleteSoftwareById).Methods("DELETE")
	router.HandleFunc("/shutdown", controller.Shutdown).Methods("GET")

	// Swagger
	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}



