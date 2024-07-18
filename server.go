package main

import (
	"encoding/json"
	decrypt "encryption/decrypt"
	encrypt "encryption/encrypt"
	"fmt"
	"log"
	"net/http"
	"os"
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

	if !decrypt.ExistsInList(req.Text, result) {
		addDataInList(req.Text, result)
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

	key := decrypt.FindKeyByHash(req.Text)
	if key == "" {
		key = "Key not found"
	}

	res := Response{Result: key}
	json.NewEncoder(w).Encode(res)
}

func addDataInList(key string, hash string) {
	var str string = fmt.Sprintf("%s  %s\n", key, hash)
	file, err := os.OpenFile("list/list.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.WriteString(str)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	err = file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}
