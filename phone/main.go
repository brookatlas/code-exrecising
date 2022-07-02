package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func initializeDatabase(db sql.DB) {
	statement, err := db.Prepare(
		`
			CREATE TABLE 'phoneNumber' (
				'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
				'number' VARCHAR(32) NOT NULL,
			);
		`,
	)
	checkError(err)

	_, err := statement.Exec()
	checkError(err)

	return true
}

func insertPhoneNumber(db sql.DB, tableName string, fields []string, values []string) {

	query := `
		INSERT INTO phoneNumber(number) values(?, ?, ?)
	`,

	statement, err := db.Prepare(
		formattedQuery,
	)
	checkError(err)

	_, err := statement.Exec(values)
	checkError(err)

	return true
}

func main() {

	db, err := sql.Open(
		"sqlite3",
		"./phone.db",
	)
	checkError(err)

	success := initializeDatabase(
		db,
	)

	if success != nil {
		panic("there was a problem initializing the database")
	}

	phoneNumbers := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}

	for index, phoneNumber := range phoneNumbers {
		success := insertPhoneNumber(
			db,
			"phoneNumber",
			["number"],
			[phoneNumber],
		)
		if success != nil {
			fmt.Println("error while inserting data into the table")
		}
	}


	db.Close()
}
