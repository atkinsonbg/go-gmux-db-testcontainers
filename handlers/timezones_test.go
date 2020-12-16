package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	db "github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	db.InitDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestListTimezonesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/timezones", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTimezonesHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var timezones []db.Timezone
	err = json.NewDecoder(rr.Body).Decode(&timezones)
	if err != nil {
		t.Error(err.Error())
		t.Error("Error retreiving list of timezones.")
	}

	if len(timezones) == 0 {
		t.Error("Error retreiving list of timezones.")
	}
}

func TestGetTimezoneHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/timezones/est", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/timezones/{identifier}", GetTimezoneHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var timezone db.Timezone
	err = json.NewDecoder(rr.Body).Decode(&timezone)
	if err != nil {
		t.Error(err.Error())
		t.Error("Error retreiving specific of timezone.")
	}

	if timezone.Name != "eastern" {
		t.Error("Error retreiving specific of timezone.")
	}
}

func TestInsertTimezoneHandler(t *testing.T) {
	var data = []byte(`
	{
		"name": "xyz",
    	"timeoffset": -10,
    	"identifier": "xyz"
	}`)

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", "/timezones", b)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InsertTimezoneHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}
