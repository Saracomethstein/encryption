package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "encryption"
)

var db *sql.DB

func GetConnection() {
	db = SetupDB()
	fmt.Println("Successfully connected!")
}

func CloseConnection() {
	db.Close()
	fmt.Println("Connection is closed!")
}

func SetupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqldb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}
	return sqldb
}

func AddDataInDB(key string, hash string) {
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

func FindKeyByHash(hash string) string {
	var key string
	err := db.QueryRow("SELECT key FROM encryption_history WHERE value = $1", hash).Scan(&key)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		return ""
	}
	return key
}
