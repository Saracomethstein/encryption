package main

import (
	"encoding/json"
	"html/template"
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

type ContactDetails struct {
	Input  string
	Output string
}

var (
	tmpl = template.Must(template.ParseFiles("../ui/index.html"))
)

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Response{Result: req.Text}
	json.NewEncoder(w).Encode(res)
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Response{Result: req.Text}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/api/encrypt", encryptHandler)
	http.HandleFunc("/api/decrypt", decryptHandler)
	log.Println("Starting server on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
