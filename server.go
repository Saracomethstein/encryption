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

	result := encrypt(req.Text, req.Hash)

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

func encrypt(promt string, hashType string) string {
	var result string
	switch hashType {
	case "sha256":
		hash := sha256.Sum256([]byte(promt))
		result = fmt.Sprintf("%x", hash)
	case "sha384":
		hash := sha512.Sum384([]byte(promt))
		result = fmt.Sprintf("%x", hash)
	case "sha512":
		hash := sha512.Sum512([]byte(promt))
		result = fmt.Sprintf("%x", hash)
	default:
		result = "Unknown hash type"
	}
	return result
}
