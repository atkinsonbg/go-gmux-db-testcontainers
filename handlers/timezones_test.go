package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	db "github.com/atkinsonbg/go-gmux-db-testcontainers/database"
	"github.com/gorilla/mux"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	packageName := "handlers"
	path, _ := os.Getwd()
	path2 := strings.Replace(path, packageName, "", 1)
	mountFrom := fmt.Sprintf("%s/scripts/init.sql", path2)
	mountTo := "/docker-entrypoint-initdb.d/init.sql"

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:11.6-alpine",
		ExposedPorts: []string{"5432/tcp"},
		BindMounts:   map[string]string{mountFrom: mountTo},
		Env: map[string]string{
			"POSTGRES_DB": "postgresTC",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// Panic and fail since there isn't much we can do if the container doesn't start
		panic(err)
	}
	defer postgresC.Terminate(ctx)

	p, _ := postgresC.MappedPort(ctx, "5432")
	os.Setenv("DBPORT", p.Port())

	fmt.Println("Initing DB")
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
