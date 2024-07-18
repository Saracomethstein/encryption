package main

import (
	"bufio"
	"encoding/json"
	encrypt "encryption/encrypt"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	if !existsInList(req.Text, result) {
		addInList(req.Text, result)
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

	key := findKeyByHash(req.Text)
	if key == "" {
		key = "Key not found"
	}

	res := Response{Result: key}
	json.NewEncoder(w).Encode(res)
}

func addInList(key string, hash string) {
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

func existsInList(key string, hash string) bool {
	file, err := os.Open("list/list.txt")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entry := fmt.Sprintf("%s  %s", key, hash)

	for scanner.Scan() {
		if scanner.Text() == entry {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return false
}

func findKeyByHash(hash string) string {
	file, err := os.Open("list/list.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "  ")
		if len(parts) == 2 && parts[1] == hash {
			return parts[0]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return ""
}
