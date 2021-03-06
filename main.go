package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"babelist.ph/api/models/babe"
	"babelist.ph/api/config/database"
)


func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func initRequestHandler() {

	http.HandleFunc("/", home)

	// creates a new instance of a mux router
    router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/babes", babe.New).Methods("POST")
	router.HandleFunc("/api/v1/babes/{id}", babe.Delete).Methods("DELETE")
	router.HandleFunc("/api/v1/babes", babe.GetAll)
	router.HandleFunc("/api/v1/babes/{id}", babe.Find)
	router.HandleFunc("/api/v1/babes/{id}", babe.Update).Methods("PUT")

	// set the headers
	 c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })

    handler := c.Handler(router)

	// start the webserver
    log.Fatal(http.ListenAndServe(":8000", handler ))

}

func main() {
	// connect to the database
	database.Connect()

	// initialize routes
	initRequestHandler()
}