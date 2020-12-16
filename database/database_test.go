package database

import (
	"testing"
)

func TestConfig(t *testing.T) {
	config := dbConfig()
	if config.Host != "localhost" {
		t.Error("DB config did not work.")
	}
}
