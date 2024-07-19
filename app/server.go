package main

import (
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

func main() {
	connection.GetConnection()
	defer connection.CloseConnection()

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

	connection.AddDataInDB(req.Text, result)

	res := Response{Result: result}
	json.NewEncoder(w).Encode(res)
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := connection.FindKeyByHash(req.Text)
	if key == "" {
		key = "Key not found"
	}

	res := Response{Result: key}
	json.NewEncoder(w).Encode(res)
}
