package database

import (
	"database/sql"
	"log"
)

func AddDataInDB(key string, hash string, db sql.DB) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM encryption_history WHERE key=$1 AND value=$2)`
	err := db.QueryRow(query, key, hash).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		query = `INSERT INTO encryption_history (key, value) VALUES ($1, $2)`
		_, err := db.Exec(query, key, hash)
		if err != nil {
			log.Fatal(err)
		}
	}
}
