package database

import (
	"log"
)

type Timezone struct {
	ID         int
	Created    string
	Modified   string
	Name       string
	Timeoffset int
	Identifier string
}

// GetAllTimezones returns all timezones from the DB
func GetAllTimezones() ([]Timezone, error) {
	rows, err := db.Query("SELECT * FROM timezones")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var timezones []Timezone

	defer rows.Close()
	for rows.Next() {
		var timezone Timezone
		err := rows.Scan(&timezone.ID, &timezone.Created, &timezone.Modified, &timezone.Name, &timezone.Timeoffset, &timezone.Identifier)
		if err != nil {
			log.Print(err)
		}
		timezones = append(timezones, timezone)
	}

	return timezones, nil
}

// GetTimezone returns a single timezone from the DB for the given identifier
func GetTimezone(identifier string) (Timezone, error) {
	row := db.QueryRow("SELECT * FROM timezones WHERE identifier = $1", identifier)
	var timezone Timezone
	err := row.Scan(&timezone.ID, &timezone.Created, &timezone.Modified, &timezone.Name, &timezone.Timeoffset, &timezone.Identifier)
	if err != nil {
		log.Print(err)
		return timezone, err
	}

	return timezone, nil
}

// InsertTimezone returns a single timezone from the DB for the given identifier
func InsertTimezone(timezone Timezone) (int, error) {
	var rowid int
	err := db.QueryRow("INSERT INTO timezones(name, timeoffset, identifier) VALUES($1, $2, $3) RETURNING id", timezone.Name, timezone.Timeoffset, timezone.Identifier).Scan(&rowid)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	return rowid, nil
}
