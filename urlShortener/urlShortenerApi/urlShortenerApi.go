package urlShortenerApi

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type createShortURLRequest struct {
	Url string `json:"url"`
}

type UrlShortenerRecord struct {
	DestinationUrl string
	hash           string
}

type UrlShortenerDatabaseRecord struct {
	id             int
	DestinationUrl string
	hash           string
}

func validateURL(Url string) error {
	_, err := url.ParseRequestURI(Url)
	if err != nil {
		return err
	}

	_, domain, hasDot := strings.Cut(Url, ".")
	if hasDot == false {
		err := errors.New("bad formatted url.")
		return err
	}
	if len(domain) == 0 {
		err := errors.New("domain is missing.")
		return err
	}

	return nil
}

func init() {
	db, err := sql.Open(
		"sqlite3",
		"url-shortener.db",
	)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	const createTableCommand string = `
		CREATE TABLE IF NOT EXISTS urlShortenedRecords (
			id INTEGER NOT NULL PRIMARY KEY,
			DestinationUrl varchar(100),
			hash varchar(100)
		);
	`

	if _, err := db.Exec(createTableCommand); err != nil {
		fmt.Println(err.Error())
		return
	}
}
