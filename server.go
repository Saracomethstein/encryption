package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
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
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	http.HandleFunc("/api/encrypt", encryptHandler)

	http.HandleFunc("/api/decrypt", decryptHandler)

	log.Println("Starting server on :8000")
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
		result = encrypt256(req.Text)
	case "sha384":
		result = encrypt384(req.Text)
	case "sha512":
		result = encrypt512(req.Text)
	default:
		result = "Unknown hash type"
	}

	res := Response{Result: result}
	json.NewEncoder(w).Encode(res)
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Здесь должна быть логика дешифрования
	res := Response{Result: "decrypted-" + req.Text}
	json.NewEncoder(w).Encode(res)
}

func encrypt256(promt string) string {
	hash := sha256.Sum256([]byte(promt))
	return fmt.Sprintf("%x", hash)
}

func encrypt384(promt string) string {
	hash := sha512.Sum384([]byte(promt))
	return fmt.Sprintf("%x", hash)
}

func encrypt512(promt string) string {
	hash := sha512.Sum512([]byte(promt))
	return fmt.Sprintf("%x", hash)
}
