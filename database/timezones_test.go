package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
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
