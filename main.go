package main

import (
	"log"
	"net/http"

	"github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
	"github.com/atkinsonbg/go-gmux-proper-unit-testing/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/timezones", handlers.ListTimezonesHandler).Methods("GET")
	r.HandleFunc("/timezones/{identifier}", handlers.GetTimezoneHandler).Methods("GET")
	r.HandleFunc("/timezones", handlers.InsertTimezoneHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", r))
}
