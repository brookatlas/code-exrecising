package urlShortenerApi

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
)

func createUrlHash(Url string) string {
	hash := md5.Sum([]byte(Url))

	return hex.EncodeToString(hash[:])
}

func createUrlShortenerRecord(Url string) (UrlShortenerRecord, error) {
	urlHash := createUrlHash(Url)

	db, err := sql.Open(
		"sqlite3",
		"url-shortener.db",
	)
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
		return UrlShortenerRecord{}, err
	}

	recordInsertCommand := fmt.Sprintf(`
		INSERT INTO urlShortenedRecords (DestinationUrl, hash)
		VALUES ("%s", "%s")
	`, Url, urlHash)

	if _, err := db.Exec(recordInsertCommand); err != nil {
		fmt.Println(err.Error())
		return UrlShortenerRecord{}, err
	}

	return UrlShortenerRecord{
			DestinationUrl: Url,
			hash:           urlHash,
		},
		nil

}

func getDestinationUrlByHash(hash string) (string, error) {
	var err error
	db, err := sql.Open(
		"sqlite3",
		"url-shortener.db",
	)
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	recordQueryCommand := "SELECT * FROM urlShortenedRecords WHERE hash = ?"
	row := db.QueryRow(recordQueryCommand, hash)
	record := UrlShortenerDatabaseRecord{}
	err = row.Scan(&record.id, &record.DestinationUrl, &record.hash)
	if err != nil {
		return "", err
	}

	return record.DestinationUrl, nil
}
