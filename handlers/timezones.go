package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
	"github.com/gorilla/mux"
)

// ListTimezonesHandler lists all the timezones in the database
func ListTimezonesHandler(w http.ResponseWriter, r *http.Request) {
	timezones, err := database.GetAllTimezones()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	results, _ := json.Marshal(timezones)

	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// GetTimezoneHandler gets a single timezone from the database based on the identifier
func GetTimezoneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timezone, err := database.GetTimezone(vars["identifier"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	results, _ := json.Marshal(timezone)

	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// InsertTimezoneHandler inserts a single timezone into the database
func InsertTimezoneHandler(w http.ResponseWriter, r *http.Request) {
	var timezone database.Timezone
	err := json.NewDecoder(r.Body).Decode(&timezone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = database.InsertTimezone(timezone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
