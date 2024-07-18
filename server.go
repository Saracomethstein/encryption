package main

import (
	"database/sql"
	"encoding/json"
	connection "encryption/database"
	encrypt "encryption/encrypt"
	"log"
	"net/http"
)

type Request struct {
	Text string `json:"text"`
	Hash string `json:"hash"`
}

type Response struct {
	Result string `json:"result"`
}

var db *sql.DB

func main() {
	db = connection.SetupDB()
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("./ui")))

	http.HandleFunc("/api/encrypt", encryptHandler)

	http.HandleFunc("/api/decrypt", decryptHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var result string

	switch req.Hash {
	case "sha256":
		result = encrypt.Encrypt256(req.Text)
	case "sha384":
		result = encrypt.Encrypt384(req.Text)
	case "sha512":
		result = encrypt.Encrypt512(req.Text)
	default:
		result = "Unknown hash type"
	}

	AddDataInDB(req.Text, result)

	res := Response{Result: result}
	json.NewEncoder(w).Encode(res)
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := FindKeyByHash(req.Text)
	if key == "" {
		key = "Key not found"
	}

	res := Response{Result: key}
	json.NewEncoder(w).Encode(res)
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
