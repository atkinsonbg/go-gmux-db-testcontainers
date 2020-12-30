package database

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	packageName := "database"
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
	InitDB()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetAllTimezones(t *testing.T) {
	tzones, err := GetAllTimezones()
	if err != nil {
		t.Error("Get All Timezones failed.")
	}

	if len(tzones) == 0 {
		t.Error("Timezones did not return any values.")
	}
}

func TestGetTimezone(t *testing.T) {
	tzone, err := GetTimezone("est")
	if err != nil {
		t.Error("Get a Timezone failed.")
	}

	if tzone.Name != "eastern" {
		t.Error("Timezone did not return correct values.")
	}
}

func TestInsertTimezone(t *testing.T) {
	tzone := Timezone{}
	tzone.Name = "Test"
	tzone.Timeoffset = 10
	tzone.Identifier = "tst"

	rowid, err := InsertTimezone(tzone)
	if err != nil {
		t.Error("Insert a Timezone failed.")
	}

	if rowid != 6 {
		t.Error("Insert timezone failed.")
	}
}
